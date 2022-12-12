defmodule Day11Test do
  use ExUnit.Case
  doctest Day11

  @m1 %Monkey{
    items: {1, 2, 3},
    operation: {"*", "100"},
    testValue: 2,
    actionTrue: 1,
    actionFalse: 2,
    inspected: 0
  }

  @m2 %Monkey{
    items: {4, 5, 6},
    operation: {"*", "1000"},
    testValue: 2,
    actionTrue: 2,
    actionFalse: 0,
    inspected: 0
  }

  @m3 %Monkey{
    items: {7, 8, 9},
    operation: {"+", "100"},
    testValue: 2,
    actionTrue: 0,
    actionFalse: 1,
    inspected: 0
  }

  test "monkeyPlay" do
    monkeys = %{0 => @m1, 1 => @m2, 2 => @m3}
    played = Day11.monkeyPlay(monkeys, 0)
    assert {} == Map.get(played, 0).items
    assert {4, 5, 6, 66, 100} == Map.get(played, 1).items
    assert {7, 8, 9, 33} == Map.get(played, 2).items
  end

  test "playSingleRound" do
    IO.puts("PLAYSINGLE")
    monkeys = %{0 => @m1, 1 => @m2, 2 => @m3}
    played = Day11.playSingleRound(monkeys, 0)
    assert {} === Map.get(played, 0).items
  end
end
