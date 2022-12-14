defmodule Day14 do
  def parse_input(path) do
    File.read!(path)
    |> String.split("\n")
  end

  def parse_line([_ | []], map), do: map

  def parse_line([{x, y1} | [{x, y2} | tail]], map) do
    m = Enum.reduce(y1..y2, map, fn y, acc -> [{x, y} | acc] end)
    parse_line([{x, y2} | tail], m)
  end

  def parse_line([{x1, y} | [{x2, y} | tail]], map) do
    m = Enum.reduce(x1..x2, map, fn x, acc -> [{x, y} | acc] end)
    parse_line([{x2, y} | tail], m)
  end

  def parse_lines(lines) do
    Enum.reduce(lines, [], fn line, map ->
      String.split(line, " -> ")
      |> Enum.map(fn e ->
        [[_, x, y]] = Regex.scan(~r{(\d+),(\d+)}, e)
        {String.to_integer(x), String.to_integer(y)}
      end)
      |> parse_line(map)
    end)
    |> Map.new(fn key -> {key, true} end)
  end

  def drop(grid, {_x, y}, max_y) when y > max_y, do: {:halt, grid}

  def drop(grid, {x, y}, max_y) do
    moves = [
      {x, y + 1},
      {x - 1, y + 1},
      {x + 1, y + 1}
    ]

    next =
      Enum.reject(moves, fn m -> grid[m] end)
      |> Enum.take(1)

    case next do
      [] -> {:cont, Map.put(grid, {x, y}, true)}
      _ -> drop(grid, hd(next), max_y)
    end
  end

  def drop_sand(grid) do
    max_y = Map.keys(grid) |> Enum.reduce(0, fn {_, v}, acc -> max(acc, v) end)

    Enum.reduce_while(Stream.cycle([{500, 0}]), {grid, 0}, fn e, {map, count} ->
      {c, m} = drop(map, e, max_y)
      # this will be off by one because I'm adding the :halt
      {c, {m, count + 1}}
    end)
  end

  def part1(path \\ "input.test.txt") do
    parse_input(path)
    |> parse_lines()
    |> drop_sand
    |> elem(1)
    |> then(&(&1 - 1))
  end
end

# Day14.part1()
