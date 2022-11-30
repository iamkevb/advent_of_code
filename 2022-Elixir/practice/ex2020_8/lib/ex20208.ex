defmodule Ex20208 do
  def readInput(path) do
    File.read!(path)
    |> String.split("\n")
  end

  def performOp(operations, idx, acc, completed) do
    [[_, op, d]] = Regex.scan(~r/(\S+) ([+-]\d+)/, Enum.at(operations, idx))

    case op do
      "acc" ->
        runProgram(operations, idx + 1, acc + String.to_integer(d), [idx | completed])

      "nop" ->
        runProgram(operations, idx + 1, acc, [idx | completed])

      "jmp" ->
        runProgram(operations, idx + String.to_integer(d), acc, [idx | completed])
    end
  end

  def runProgram(operations, idx, acc, completed) do
    if Enum.find(completed, fn e -> e === idx end) === idx do
      acc
    else
      performOp(operations, idx, acc, completed)
    end
  end

  def part1(path \\ "input.test.txt") do
    readInput(path) |> runProgram(0, 0, [])
  end
end
