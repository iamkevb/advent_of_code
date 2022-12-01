defmodule Day1 do
  def part1(path \\ "input.prod.txt") do
    File.read!(path)
    |> String.split("\n\n")
    |> Enum.reduce(0, fn e, acc ->
      String.split(e, "\n")
      |> Enum.reduce(0, fn f, sum -> sum + String.to_integer(f) end)
      |> max(acc)
    end)
  end

  def part2(path \\ "input.prod.txt") do
    File.read!(path)
    |> String.split("\n\n")
    |> Enum.map(fn e ->
      String.split(e, "\n")
      |> Enum.reduce(0, fn f, sum -> sum + String.to_integer(f) end)
    end)
    |> Enum.sort(&(&1 >= &2))
    |> Enum.slice(0, 3)
    |> Enum.reduce(0, fn e, acc -> acc + e end)
  end
end
