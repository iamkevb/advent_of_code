defmodule Day18 do
  def parseInput(path) do
    File.read!(path)
    |> String.split("\n")
    |> Enum.map(fn s ->
      String.split(s, ",")
      |> Enum.map(&String.to_integer/1)
      |> List.to_tuple()
    end)
    |> MapSet.new()
  end

  def neighbours({px, py, pz}) do
    [{1, 0, 0}, {-1, 0, 0}, {0, 1, 0}, {0, -1, 0}, {0, 0, 1}, {0, 0, -1}]
    |> Enum.map(fn {x, y, z} -> {x + px, y + py, z + pz} end)
  end

  def part1(path \\ "input.test.txt") do
    m = parseInput(path)

    Enum.map(m, fn p ->
      neighbours(p) |> Enum.reject(&MapSet.member?(m, &1))
    end)
    |> List.flatten()
    |> length()
  end
end
