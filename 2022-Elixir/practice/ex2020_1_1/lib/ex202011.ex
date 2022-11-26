defmodule Ex202011 do
  def readInput(path) do
    File.read!(path)
    |> String.split()
    |> Enum.map(&String.to_integer/1)
    |> Enum.uniq()
    |> Enum.sort()
  end

  def findProduct([]) do
    0
  end

  # find 2 values that sum to 2020, return their product
  def findProduct([head | tail]) do
    v2 = Enum.find(tail, fn x -> 2020 - x === head end)

    if v2 != nil do
      v2 * head
    else
      findProduct(tail)
    end
  end

  def run do
    values = readInput("input.txt")
    findProduct(values)
  end
end
