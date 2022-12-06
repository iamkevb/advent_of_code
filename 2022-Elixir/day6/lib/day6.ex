defmodule Day6 do
  def isMarker(g) do
    Enum.take(g, 4)
    |> Enum.uniq()
    |> length()
    |> then(&(&1 == 4))
  end

  def findMarker(g, idx) do
    cond do
      isMarker(g) -> idx
      true -> tl(g) |> findMarker(idx + 1)
    end
  end

  def firstMarker(s) do
    String.graphemes(s)
    |> findMarker(3)
    |> then(&(&1 + 1))
  end

  def part1(path \\ "input.prod.txt") do
    File.read!(path)
    |> firstMarker()
  end

  def isMessage(g) do
    Enum.take(g, 14)
    |> Enum.uniq()
    |> length()
    |> then(&(&1 == 14))
  end

  def findMessage(g, idx) do
    cond do
      isMessage(g) -> idx
      true -> tl(g) |> findMessage(idx + 1)
    end
  end

  def firstMessage(s) do
    String.graphemes(s)
    |> findMessage(13)
    |> then(&(&1 + 1))
  end

  def part2(path \\ "input.prod.txt") do
    File.read!(path)
    |> firstMessage()
  end
end
