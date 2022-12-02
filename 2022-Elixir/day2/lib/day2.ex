defmodule Day2 do
  def readInput(path) do
    File.read!(path)
    |> String.split()
    |> Enum.chunk_every(2)
  end

  #  A X: Rock 1 point
  #  B Y: Paper 2 points
  #  C Z: Scissors 3 points

  # Score Points for your move + (6 for win|3 for draw|0 for loss)
  defp score(["A", "X"]), do: 1 + 3
  defp score(["A", "Y"]), do: 2 + 6
  defp score(["A", "Z"]), do: 3 + 0
  defp score(["B", "X"]), do: 1 + 0
  defp score(["B", "Y"]), do: 2 + 3
  defp score(["B", "Z"]), do: 3 + 6
  defp score(["C", "X"]), do: 1 + 6
  defp score(["C", "Y"]), do: 2 + 0
  defp score(["C", "Z"]), do: 3 + 3
  defp score([elf, player]), do: raise("oops #{elf} #{player}")

  # A,B,C: Rock, Paper, Scissors
  # X,Y,Z: Lose, Draw, Win
  defp gameScore([_, "Y"]), do: 3
  defp gameScore([_, "Z"]), do: 6
  defp gameScore(_), do: 0

  defp moveScore(["A", "X"]), do: 3
  defp moveScore(["A", "Y"]), do: 1
  defp moveScore(["A", "Z"]), do: 2
  defp moveScore(["B", "X"]), do: 1
  defp moveScore(["B", "Y"]), do: 2
  defp moveScore(["B", "Z"]), do: 3
  defp moveScore(["C", "X"]), do: 2
  defp moveScore(["C", "Y"]), do: 3
  defp moveScore(["C", "Z"]), do: 1
  defp moveScore([elf, player]), do: raise("oops #{elf} #{player}")

  def part1(path \\ "input.test.txt") do
    readInput(path)
    |> Enum.reduce(0, fn e, acc -> acc + score(e) end)
  end

  def part2(path \\ "input.test.txt") do
    readInput(path)
    |> Enum.reduce(0, fn e, acc ->
      acc + moveScore(e) + gameScore(e)
    end)
  end
end
