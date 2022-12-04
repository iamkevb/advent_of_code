defmodule Day4 do
  defp readFile(path) do
    File.read!(path)
    |> String.split("\n")
  end

  defp parseLine(l) do
    Regex.scan(~r/(\d+)-(\d+),(\d+)-(\d+)/, l)
    |> Enum.map(fn x -> tl(x) end)
    |> List.flatten()
    |> Enum.map(fn x -> String.to_integer(x) end)
  end

  defp overlapFull([a, b, c, d]), do: (a <= c && b >= d) || (c <= a && d >= b)
  defp overlapFull(_), do: false

  defp overlapAny([a, b, c, d]), do: max(a, c) <= min(b, d)
  defp overlapAny(_), do: false

  def part1(path \\ "input.test.txt") do
    readFile(path)
    |> Enum.map(&parseLine/1)
    |> Enum.filter(&overlapFull/1)
    |> length()
  end

  def part2(path \\ "input.test.txt") do
    readFile(path)
    |> Enum.map(&parseLine/1)
    |> Enum.filter(&overlapAny/1)
    |> length()
  end
end
