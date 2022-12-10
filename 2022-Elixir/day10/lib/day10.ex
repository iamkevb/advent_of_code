defmodule Day10 do
  def readInput(path) do
    File.read!(path)
    |> String.split("\n")
  end

  def part1(path \\ "input.test.txt") do
    readInput(path)
    |> Enum.reduce([1], fn i, acc ->
      handleInstruction(i, acc)
    end)
    |> Enum.reverse()
    |> List.to_tuple()
    |> then(fn l ->
      #
      elem(l, 19) * 20 + elem(l, 59) * 60 + elem(l, 99) * 100 + elem(l, 139) * 140 +
        elem(l, 179) * 180 +
        elem(l, 219) * 220
    end)
  end

  def handleInstruction(instruction, xValues) do
    [x | xValues] = xValues
    {num, val} = cycles(instruction)
    # IO.puts("x:#{x} num:#{num} val:#{val} len xValues:#{length(xValues)}")
    [x + val | List.duplicate(x, num) ++ xValues]
  end

  def cycles("noop"), do: {1, 0}

  def cycles(addx) do
    [_ | [c | _]] = String.split(addx)
    {2, String.to_integer(c)}
  end
end
