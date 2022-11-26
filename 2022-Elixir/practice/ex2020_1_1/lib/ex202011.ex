defmodule Ex202011 do
  def readInput(path) do
    File.read!(path)
    |> String.split()
    |> Enum.map(&String.to_integer/1)
    |> Enum.uniq()
    |> Enum.sort()
  end

  def findProduct2(_, []) do
    0
  end

  # find 2 values that sum to 2020, return their product
  def findProduct2(sum, [head | tail]) do
    v2 = Enum.find(tail, fn x -> sum - x === head end)

    if v2 != nil do
      v2 * head
    else
      findProduct2(sum, tail)
    end
  end

  def findProduct3(_, []) do
    0
  end

  def findProduct3(sum, [head | tail]) do
    prod2 = findProduct2(sum - head, tail)

    if prod2 > 0 do
      head * prod2
    else
      findProduct3(sum, tail)
    end
  end

  def part1 do
    values = readInput("input.txt")
    findProduct2(2020, values)
  end

  def part2 do
    values = readInput("input.txt")
    findProduct3(2020, values)
  end
end
