defmodule Day10Test do
  use ExUnit.Case
  doctest Day10

  # At the start of the first cycle, the noop instruction begins execution. During the first cycle, X is 1. After the first cycle, the noop instruction finishes execution, doing nothing.
  # At the start of the second cycle, the addx 3 instruction begins execution. During the second cycle, X is still 1.
  # During the third cycle, X is still 1. After the third cycle, the addx 3 instruction finishes execution, setting X to 4.
  # At the start of the fourth cycle, the addx -5 instruction begins execution. During the fourth cycle, X is still 4.
  # During the fifth cycle, X is still 4. After the fifth cycle, the addx -5 instruction finishes execution, setting X to -1.

  test "handle noop" do
    x = Day10.handleInstruction("noop", [1])
    assert [1, 1] == x
    x = Day10.handleInstruction("addx 3", x)
    assert [4, 1, 1, 1] == x
    x = Day10.handleInstruction("addx -5", x)
    assert [-1, 4, 4, 1, 1, 1] == x
  end
end
