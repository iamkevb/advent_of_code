defmodule Day11 do
  def setOperation(monkey, op, "old") do
    %Monkey{monkey | operation: {op, -1}}
  end

  def setOperation(monkey, op, val) do
    %Monkey{monkey | operation: {op, String.to_integer(val)}}
  end

  def parseMonkeyLine(monkey, ["Starting items", items]) do
    String.split(items, ", ")
    |> Enum.map(&String.to_integer(String.trim(&1)))
    |> List.to_tuple()
    |> then(&%Monkey{monkey | items: &1})
  end

  def parseMonkeyLine(monkey, ["Operation", operation]) do
    [_, _, _, op, val] = String.split(operation)
    setOperation(monkey, op, val)
  end

  def parseMonkeyLine(monkey, ["Test", test]) do
    [_, _, val] = String.split(test)
    %Monkey{monkey | testValue: String.to_integer(val)}
  end

  def parseMonkeyLine(monkey, ["If true", action]) do
    [_, _, _, trueAction] = String.split(action)
    %Monkey{monkey | actionTrue: String.to_integer(trueAction)}
  end

  def parseMonkeyLine(monkey, ["If false", action]) do
    [_, _, _, falseAction] = String.split(action)
    %Monkey{monkey | actionFalse: String.to_integer(falseAction)}
  end

  def parseMonkeyLine(m, _), do: m

  def monkeyPlay(monkeys, key) do
    monkey =
      Map.get(monkeys, key)
      |> Monkey.inspect()

    tsize = tuple_size(monkey.items)

    if(tsize > 0) do
      nextMonkeyIndex = Monkey.nextMonkeyIndex(monkey)
      nextMonkey = Map.get(monkeys, nextMonkeyIndex)

      {item, monkey} = Monkey.throwItem(monkey)
      monkeys = Map.put(monkeys, key, monkey)

      nextMonkey = Monkey.catchItem(nextMonkey, item)
      monkeys = Map.put(monkeys, nextMonkeyIndex, nextMonkey)

      monkeyPlay(monkeys, key)
    else
      monkeys
    end
  end

  def playSingleRound(monkeys, key) do
    if map_size(monkeys) > key do
      monkeyPlay(monkeys, key)
      |> playSingleRound(key + 1)
    else
      monkeys
    end
  end

  def playRounds(monkeys, 0), do: monkeys

  def playRounds(monkeys, rounds) do
    IO.puts("#{rounds} to go")

    playSingleRound(monkeys, 0)
    |> playRounds(rounds - 1)
  end

  def parseMonkeys(path) do
    File.read!(path)
    |> String.split("\n\n")
    |> Enum.map(fn md ->
      String.split(md, "\n")
      |> Enum.reduce(%Monkey{inspected: 0, product: 0}, fn line, monkey ->
        parseMonkeyLine(monkey, String.split(String.trim(line), ":"))
      end)
    end)
  end

  def part1(path \\ "input.test.txt") do
    monkeys = parseMonkeys(path)
    product = Enum.reduce(monkeys, 1, fn m, acc -> acc * m.testValue end)
    IO.puts(product)

    Enum.map(monkeys, fn m -> %Monkey{m | product: product} end)
    |> Enum.with_index()
    |> Map.new(fn {v, k} -> {k, v} end)
    |> playRounds(20)
    |> Enum.map(fn {_, v} -> v.inspected end)
    |> Enum.sort(&(&2 < &1))
    |> Enum.take(2)
    |> then(fn [l, r] -> l * r end)
  end

  def part2(path \\ "input.test.txt") do
    monkeys = parseMonkeys(path)
    product = Enum.reduce(monkeys, 1, fn m, acc -> acc * m.testValue end)
    IO.puts(product)

    Enum.map(monkeys, fn m -> %Monkey{m | product: product} end)
    |> Enum.with_index()
    |> Map.new(fn {v, k} -> {k, v} end)
    |> playRounds(10000)
    |> Enum.map(fn {_, v} -> v.inspected end)
    |> Enum.sort(&(&2 < &1))
    |> Enum.take(2)
    |> then(fn [l, r] -> l * r end)
  end
end
