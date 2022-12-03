defmodule Day3 do
  @spec readInput(String.t()) :: [String.t()]
  defp readInput(path) do
    File.read!(path)
    |> String.split("\n")
  end

  @spec computePriority([String.t()]) :: integer()
  def computePriority(e) do
    Enum.reduce(e, 0, fn v, acc ->
      c = String.to_charlist(v) |> hd

      if c >= ?a do
        acc + c - ?a + 1
      else
        acc + c - ?A + 27
      end
    end)
  end

  @spec findBadge([String.t()]) :: String.t()
  def findBadge([a, b, c]) do
    commonGlyphs(a, commonGlyphs(b, c))
  end

  @spec commonGlyphs(String.t(), String.t()) :: String.t()
  def commonGlyphs(a, b) do
    bg =
      String.graphemes(b)
      |> Enum.uniq()

    String.graphemes(a)
    |> Enum.uniq()
    |> Enum.filter(fn c -> Enum.find(bg, fn v -> v === c end) end)
    |> Enum.reduce("", fn c, acc -> acc <> c end)
  end

  @spec part1(String.t()) :: integer
  def part1(path \\ "input.test.txt") do
    readInput(path)
    |> Enum.map(fn v ->
      {a, b} = String.split_at(v, div(String.length(v), 2))
      commonGlyphs(a, b)
    end)
    |> computePriority()
  end

  @spec part2(String.t()) :: integer
  def part2(path \\ "input.test.txt") do
    readInput(path)
    |> Enum.chunk_every(3)
    |> Enum.map(&findBadge/1)
    |> computePriority()
  end
end
