defmodule Day12 do
  def parse_input(path) do
    File.read!(path)
    |> String.split("\n")
    |> Enum.with_index()
    |> Enum.reduce(%{}, fn {line, y}, acc ->
      String.graphemes(line)
      |> Enum.with_index()
      |> Enum.reduce(acc, fn {g, x}, acc ->
        Map.put(acc, {x, y}, g)
      end)
    end)
  end

  def find_value(grid, value) do
    Enum.find(grid, fn {_, v} -> v === value end)
    |> elem(0)
  end

  def elevation("S"), do: elevation("a")
  def elevation("E"), do: elevation("z")
  def elevation(<<v>>), do: v - ?a

  def neighbours({x, y}, heights) do
    h = heights[{x, y}]

    [{x + 1, y}, {x - 1, y}, {x, y + 1}, {x, y - 1}]
    |> Enum.reject(&(heights[&1] == nil))
    |> Enum.filter(&(heights[&1] <= h + 1))
  end

  def dijkstra(distances, heights, destination) do
    next = Enum.min_by(distances, &elem(&1, 1))

    case next do
      {^destination, dist} ->
        dist

      {cur, dist} ->
        cur
        |> neighbours(heights)
        |> Enum.reduce(distances, fn n, distances ->
          Map.replace_lazy(distances, n, &min(&1, dist + 1))
        end)
        |> Map.delete(cur)
        |> dijkstra(heights, destination)
    end
  end

  def part1(path \\ "input.test.txt") do
    grid = parse_input(path)
    start = find_value(grid, "S")
    destination = find_value(grid, "E")

    heights =
      grid
      |> Enum.map(fn {k, v} -> {k, elevation(v)} end)
      |> Map.new()

    distances =
      Map.keys(heights)
      |> Map.new(fn k -> {k, :infinity} end)
      |> Map.put(start, 0)

    dijkstra(distances, heights, destination)
  end

  def part2(path \\ "input.test.txt") do
    grid = parse_input(path)
    destination = find_value(grid, "E")

    starts = Enum.filter(grid, fn {_, v} -> v == "a" end) |> Enum.map(fn {s, _} -> s end)

    heights =
      grid
      |> Enum.map(fn {k, v} -> {k, elevation(v)} end)
      |> Map.new()

    distances =
      Map.keys(heights)
      |> Map.new(fn k -> {k, :infinity} end)

    distances =
      Enum.reduce(starts, distances, fn start, distances ->
        Map.put(distances, start, 0)
      end)

    dijkstra(distances, heights, destination)
  end
end
