defmodule Day9Test do
  use ExUnit.Case
  doctest Day9

  test "do tail doesn't move when next or on head" do
    # tail doesn't move when within 1 of head
    assert 0 === Day9.next_tail_loc(0, 0)
    assert 0 === Day9.next_tail_loc(1, 0)
    assert 1 === Day9.next_tail_loc(0, 1)
  end

  test "tail moves 1 closer when not next to head" do
    assert 1 === Day9.next_tail_loc(2, 0)
    assert 1 === Day9.next_tail_loc(0, 2)
  end

  test "move tail none" do
    assert {0, 0} === Day9.move_tail({0, 0}, {0, 0})
    assert {0, 0} === Day9.move_tail({1, 0}, {0, 0})
    assert {0, 0} === Day9.move_tail({0, 1}, {0, 0})
    assert {0, 0} === Day9.move_tail({1, 1}, {0, 0})
    assert {3, 0} === Day9.move_tail({4, 1}, {3, 0})
  end

  test "move tail right" do
    assert {1, 0} === Day9.move_tail({2, 0}, {0, 0})
  end

  test "move tail left" do
    assert {1, 0} === Day9.move_tail({0, 0}, {2, 0})
  end

  test "move tail up" do
    assert {0, 1} === Day9.move_tail({0, 0}, {0, 2})
  end

  test "move tail down" do
    assert {0, 1} === Day9.move_tail({0, 2}, {0, 0})
  end

  test "move tail diagonal" do
    assert {1, 1} === Day9.move_tail({0, 0}, {2, 2})
    assert {1, 1} === Day9.move_tail({2, 2}, {0, 0})
    assert {1, 1} === Day9.move_tail({0, 2}, {2, 0})
    assert {1, 1} === Day9.move_tail({2, 0}, {0, 2})
    assert {4, 1} === Day9.move_tail({4, 2}, {3, 0})
  end

  test "handleInstruction up" do
    visited = Day9.handleInstruction("U", 1, [{0, 0}])
    assert [{0, 1}, {0, 0}] === visited

    visited = Day9.handleInstruction("U", 2, [{0, 0}])
    assert [{0, 2}, {0, 1}, {0, 0}] === visited
  end
end
