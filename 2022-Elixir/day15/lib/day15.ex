defmodule Sensor do
  defstruct position: {0, 0}, beacon: {0, 0}, distance: 0

  def scan(s) do
    [sx, sy, bx, by] =
      Regex.scan(
        ~r{Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)},
        s
      )
      |> hd()
      |> tl()

    sx = String.to_integer(sx)
    sy = String.to_integer(sy)
    bx = String.to_integer(bx)
    by = String.to_integer(by)
    d = abs(sx - bx) + abs(sy - by)

    %Sensor{position: {sx, sy}, beacon: {bx, by}, distance: d}
  end
end

defmodule Day15 do
  alias Inspect.Range

  def parse(path) do
    File.read!(path)
    |> String.split("\n")
    |> Enum.map(&Sensor.scan/1)
  end

  def range_at_y(sensor, target_y) do
    {x, y} = sensor.position

    d_remain = sensor.distance - abs(y - target_y)

    cond do
      d_remain < 0 ->
        :ignore

      true ->
        (x - d_remain)..(x + d_remain)
    end
  end

  def combine_overlapping_ranges(ranges) do
    sorted_ranges = Enum.sort(ranges, fn _..s, _..e -> e < s end)

    combined_ranges =
      sorted_ranges
      |> Enum.reduce([], fn next, acc ->
        # acc_a..acc_b = acc
        a..b = next

        case acc do
          # just add first range
          [] ->
            [next]

          # does a..b intersect acc_a..acc_b
          [acc_a..acc_b | _] when b >= acc_a ->
            # IO.puts("1 #{a}..#{b} #{acc_b}..#{acc_b}")
            [min(a, acc_a)..max(b, acc_b) | Enum.drop(acc, 1)]

          # is a..b consecutive to acc_a..acc_b (1..2, 3..5)
          [acc_a..acc_b | _] when b + 1 == acc_a ->
            # IO.puts("2 a..b#{a}..#{b} acc_a..acc_b#{acc_a}..#{acc_b}")
            [a..acc_b | Enum.drop(acc, 1)]

          # no intersection, prepend new range
          _ ->
            [acc_a..acc_b | _] = acc
            # IO.puts("3 a..b #{a}..#{b} acc_a..acc_b #{acc_a}..#{acc_b}")
            [a..b | acc]
        end
      end)

    combined_ranges
  end

  def part1(path \\ "input.test.txt") do
    # sensor_row = 2_000_000
    sensor_row = 10
    sensors = parse(path)

    sensors_on_row =
      Enum.filter(sensors, fn s ->
        {_, sy} = s.position
        sy === sensor_row
      end)
      |> length()

    IO.puts("sensors_on_row = #{sensors_on_row}")

    beacons_on_row =
      Enum.reduce(sensors, [], fn s, acc ->
        {_, by} = s.beacon

        if by === sensor_row do
          [s.beacon | acc]
        else
          acc
        end
      end)
      |> Enum.uniq()
      |> length()

    IO.puts("beacons_on_row = #{beacons_on_row}")

    Enum.map(sensors, fn s ->
      range_at_y(s, sensor_row)
    end)
    |> Enum.reject(&(&1 == :ignore))
    |> combine_overlapping_ranges()
    |> Enum.reduce(0, fn s..e, acc -> acc + e - s + 1 end)
    |> then(&(&1 - sensors_on_row - beacons_on_row))
  end

  def search_row(_, 0), do: :you_failed_kevin

  def search_row(sensors, row) do
    ranges =
      Enum.map(sensors, fn s ->
        range_at_y(s, row)
      end)
      |> Enum.reject(&(&1 == :ignore))
      |> combine_overlapping_ranges()

    if length(ranges) == 1 do
      search_row(sensors, row - 1)
    else
      # ranges
      [_..s, e.._] = ranges
      {s, e, row}
    end
  end

  def part2(path \\ "input.test.txt") do
    max_row = 4_000_000

    parse(path)
    |> search_row(max_row)
  end
end
