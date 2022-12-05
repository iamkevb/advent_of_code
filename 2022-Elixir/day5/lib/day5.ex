defmodule Day5 do
  def readInput(path) do
    File.read!(path)
    |> String.split("\n\n")
  end

  def parseStacks(input) do
    input
    |> String.replace(["[", "]"], " ")
    |> String.split("\n")
    |> Enum.drop(-1)
    |> Enum.map(&String.graphemes/1)
    |> Enum.zip()
    |> Enum.map(fn stack -> stack |> Tuple.to_list() |> Enum.filter(&(&1 != " ")) end)
    |> Enum.reject(&Enum.empty?/1)
    |> Enum.with_index(&{&2 + 1, &1})
    |> Map.new()
  end

  def move(n, from, to, stacks) do
    dropped = Enum.drop(stacks[from], n)
    added = (stacks[from] |> Enum.take(n) |> Enum.reverse()) ++ stacks[to]
    Map.put(stacks, from, dropped) |> Map.put(to, added)
  end

  def move9001(n, from, to, stacks) do
    dropped = Enum.drop(stacks[from], n)
    added = (stacks[from] |> Enum.take(n)) ++ stacks[to]
    Map.put(stacks, from, dropped) |> Map.put(to, added)
  end

  def part1(path \\ "input.test.txt") do
    [stacks, operations] = readInput(path)
    startStacks = parseStacks(stacks)

    String.split(operations, "\n")
    |> Enum.reduce(startStacks, fn op, acc ->
      [[_, n, from, to]] = Regex.scan(~r/move (\d+) from (\d+) to (\d+)/, op)
      num = String.to_integer(n)
      from = String.to_integer(from)
      to = String.to_integer(to)
      move(num, from, to, acc)
    end)
    |> Enum.reduce("", fn {_, xs}, acc -> acc <> hd(xs) end)
  end

  def part2(path \\ "input.test.txt") do
    [stacks, operations] = readInput(path)
    startStacks = parseStacks(stacks)

    String.split(operations, "\n")
    |> Enum.reduce(startStacks, fn op, acc ->
      [[_, n, from, to]] = Regex.scan(~r/move (\d+) from (\d+) to (\d+)/, op)
      num = String.to_integer(n)
      from = String.to_integer(from)
      to = String.to_integer(to)
      move9001(num, from, to, acc)
    end)
    |> Enum.reduce("", fn {_, xs}, acc -> acc <> hd(xs) end)
  end
end
