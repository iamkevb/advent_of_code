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
      elem(l, 19) * 20 + elem(l, 59) * 60 + elem(l, 99) * 100 + elem(l, 139) * 140 +
        elem(l, 179) * 180 +
        elem(l, 219) * 220
    end)
  end

  def handleInstruction(instruction, xValues) do
    [x | xValues] = xValues
    {num, val} = cycles(instruction)
    [x + val | List.duplicate(x, num) ++ xValues]
  end

  def cycles("noop"), do: {1, 0}

  def cycles(addx) do
    [_ | [c | _]] = String.split(addx)
    {2, String.to_integer(c)}
  end

  def part2(path \\ "input.test.txt") do
    readInput(path)
    |> Enum.reduce([1], fn i, acc ->
      handleInstruction(i, acc)
    end)
    |> Enum.reverse()
    |> Enum.with_index()
    |> Enum.map(fn {x, idx} ->
      pos = rem(idx, 40)

      cond do
        abs(x - pos) <= 1 -> "#"
        true -> " "
      end
    end)
    |> Enum.chunk_every(40)
    |> Enum.map(&Enum.join/1)
    |> Enum.each(&IO.puts/1)
  end
end
