defmodule Day9 do
  @start {1000, 1000}

  def readInput(path) do
    File.read!(path)
    |> String.split("\n")
  end

  def next_tail_loc(n1, n2) when n1 === n2, do: n2
  def next_tail_loc(n1, n2) when n1 > n2 + 1, do: n2 + 1
  def next_tail_loc(n1, n2) when n1 < n2 - 1, do: n2 - 1
  def next_tail_loc(_, n2), do: n2

  def move_tail({hx, hy}, {tx, ty}) do
    {ntx, nty} = {next_tail_loc(hx, tx), next_tail_loc(hy, ty)}

    cond do
      ntx === tx && nty === ty -> {ntx, nty}
      ntx !== tx && nty !== ty -> {ntx, nty}
      ntx !== tx && nty !== hy -> {ntx, hy}
      nty !== ty && ntx !== hx -> {hx, nty}
      true -> {ntx, nty}
    end
  end

  def handleInstruction(_, 0, visited), do: visited

  def handleInstruction("U", num, [{x, y} | visited]),
    do: handleInstruction("U", num - 1, [{x, y + 1} | [{x, y} | visited]])

  def handleInstruction("D", num, [{x, y} | visited]),
    do: handleInstruction("D", num - 1, [{x, y - 1} | [{x, y} | visited]])

  def handleInstruction("L", num, [{x, y} | visited]),
    do: handleInstruction("L", num - 1, [{x - 1, y} | [{x, y} | visited]])

  def handleInstruction("R", num, [{x, y} | visited]),
    do: handleInstruction("R", num - 1, [{x + 1, y} | [{x, y} | visited]])

  def mapKnot(prev, 0), do: prev

  def mapKnot(prev, knotid) do
    IO.puts("Next knot #{knotid} #{length(Enum.uniq(prev))}")

    Enum.reduce(prev, [@start], fn head, acc ->
      [tail | _] = acc
      next = move_tail(head, tail)
      [next | acc]
    end)
    |> Enum.reverse()
    |> mapKnot(knotid - 1)
  end

  def part1(path \\ "input.test.txt") do
    readInput(path)
    |> Enum.reduce([@start], fn op, visited ->
      [[_, direction, num]] = Regex.scan(~r{(\S) (\d+)}, op)
      handleInstruction(direction, String.to_integer(num), visited)
    end)
    |> Enum.reverse()
    |> mapKnot(1)
    |> Enum.uniq()
    |> length()
  end

  def part2(path \\ "input.test.2.txt") do
    readInput(path)
    |> Enum.reduce([@start], fn op, visited ->
      [[_, direction, num]] = Regex.scan(~r{(\S) (\d+)}, op)
      handleInstruction(direction, String.to_integer(num), visited)
    end)
    |> Enum.reverse()
    |> mapKnot(9)
    |> Enum.uniq()
    |> length()
  end
end
