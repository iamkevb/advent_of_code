defmodule Day5Test do
  use ExUnit.Case
  doctest Day5

  test "move" do
    stacks = %{1 => ["N", "Z"], 2 => ["D", "C", "M"], 3 => ["P"]}
    newStacks = Day5.move(1, 1, 2, stacks)
    assert %{1 => ["Z"], 2 => ["N", "D", "C", "M"], 3 => ["P"]} === newStacks
  end
end
