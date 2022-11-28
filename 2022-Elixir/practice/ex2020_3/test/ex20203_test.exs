defmodule Ex20203Test do
  use ExUnit.Case
  doctest Ex20203

  test "isTree/2" do
    rowString = ".#."
    refute Ex20203.isTree(rowString, 0)
    assert Ex20203.isTree(rowString, 1)
    refute Ex20203.isTree(rowString, 2)
    refute Ex20203.isTree(rowString, 3)
    assert Ex20203.isTree(rowString, 4)
    refute Ex20203.isTree(rowString, 5)

    assert Ex20203.isTree("#.#", 3)
  end

  test "isTree/3" do
    rows = [".#.", "#.#"]
    refute Ex20203.isTree(rows, 0, 0)
    assert Ex20203.isTree(rows, 0, 1)
    refute Ex20203.isTree(rows, 0, 2)
    refute Ex20203.isTree(rows, 0, 3)
    assert Ex20203.isTree(rows, 0, 4)
    refute Ex20203.isTree(rows, 0, 5)
    assert Ex20203.isTree(rows, 1, 0)
    refute Ex20203.isTree(rows, 1, 1)
    assert Ex20203.isTree(rows, 1, 2)
    assert Ex20203.isTree(rows, 1, 3)
    refute Ex20203.isTree(rows, 1, 4)
    assert Ex20203.isTree(rows, 1, 5)
  end

  test "countTreesEmpty" do
    rows = ["..."]
    visits = [%{row: 0, col: 0}]
    assert Ex20203.countTrees(rows, visits) === 0
  end

  test "countTreesForest" do
    rows = ["#.#", ".#."]

    visits = [
      %{row: 0, col: 0},
      %{row: 0, col: 1},
      %{row: 1, col: 1}
    ]

    assert Ex20203.countTrees(rows, visits) === 2
  end

  test "visited" do
    assert Ex20203.visited(3, 0, 0, 2, 1) === [
             %{row: 1, col: 2},
             %{row: 2, col: 4},
             %{row: 3, col: 6}
           ]
  end
end
