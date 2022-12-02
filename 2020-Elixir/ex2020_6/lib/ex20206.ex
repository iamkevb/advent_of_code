defmodule Ex20206 do
  def readInput(path) do
    File.read!(path)
    |> String.split("\n\n")
  end

  def part1(path \\ "input.prod.txt") do
    readInput(path)
    |> Enum.map(fn g ->
      String.replace(g, "\n", "")
      |> String.graphemes()
      |> Enum.uniq()
      |> length()
    end)
    |> Enum.reduce(0, fn e, acc -> acc + e end)
  end

  def countGroup(group) do
    people = String.split(group, "\n")
    len = length(people)

    Enum.reduce(people, [], fn p, acc -> [String.graphemes(p) | acc] end)
    |> Enum.flat_map(fn x -> x end)
    |> Enum.frequencies()
    |> Enum.reduce(0, fn {_, e}, acc ->
      if e === len do
        acc + 1
      else
        acc
      end
    end)
  end

  def part2(path \\ "input.prod.txt") do
    readInput(path)
    |> Enum.reduce(0, fn g, acc ->
      countGroup(g) + acc
    end)
  end
end
