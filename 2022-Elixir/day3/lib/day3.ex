defmodule Day3 do
  defp readInput(path) do
    File.read!(path)
    |> String.split("\n")
  end

  def computePriority(e) do
    Enum.reduce(e, 0, fn v, acc ->
      [c | _] = String.to_charlist(v)

      if c >= ?a do
        acc + c - ?a + 1
      else
        acc + c - ?A + 27
      end
    end)
  end

  def findBadge([a, b, c]) do
    commonGlyphs(a, commonGlyphs(b, c))
  end

  def commonGlyphs(a, b) do
    bg =
      String.graphemes(b)
      |> Enum.uniq()
      |> Enum.dedup()

    String.graphemes(a)
    |> Enum.uniq()
    |> Enum.dedup()
    |> Enum.filter(fn c -> Enum.find(bg, fn v -> v === c end) end)
    |> Enum.reduce("", fn c, acc -> acc <> c end)
  end

  def part1(path \\ "input.test.txt") do
    readInput(path)
    |> Enum.map(fn v ->
      {a, b} = String.split_at(v, div(String.length(v), 2))
      commonGlyphs(a, b)
    end)
    |> computePriority()
  end

  def part2(path \\ "input.test.txt") do
    readInput(path)
    |> Enum.chunk_every(3)
    |> Enum.map(&findBadge/1)
    |> computePriority()
  end
end
