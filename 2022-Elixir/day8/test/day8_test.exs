defmodule Day8Test do
  use ExUnit.Case
  doctest Day8

  test "compute tree score 1" do
    t1 = %Tree{height: 1}
    t2 = %Tree{height: 2}
    t3 = %Tree{height: 3}
    t4 = %Tree{height: 4}
    row = Day8.computeTreeScore([t4, t3, t2, t1], [])
    scores = Enum.map(row, & &1.score)
    assert [3, 2, 1, 0] == scores
  end

  test "compute tree score 2" do
    t1 = %Tree{height: 1}
    t2 = %Tree{height: 2}
    t3 = %Tree{height: 3}
    t4 = %Tree{height: 4}
    row = Day8.computeTreeScore([t1, t2, t3, t4], [])
    scores = Enum.map(row, & &1.score)
    assert [1, 1, 1, 0] == scores
  end

  test "compute tree score 3" do
    t1 = %Tree{height: 1}
    t2 = %Tree{height: 2}
    t3 = %Tree{height: 3}
    t4 = %Tree{height: 4}
    row = Day8.computeTreeScore([t2, t1, t1, t4], [])
    scores = Enum.map(row, & &1.score)
    assert [3, 1, 1, 0] == scores
  end
end
