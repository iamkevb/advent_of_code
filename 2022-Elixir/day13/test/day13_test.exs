defmodule Day13Test do
  use ExUnit.Case
  doctest Day13

  test "compare two integers" do
    assert Day13.compare([1, 2, 3], [3, 2, 1])
    refute Day13.compare([3, 2, 1], [1, 2, 3])
    assert Day13.compare([1, 1], [1, 1])
    assert Day13.compare([1, 1], [1, 1, 1])
    refute Day13.compare([1, 1, 1], [1, 1])
    assert Day13.compare([1, 1, 3, 1, 1], [1, 1, 5, 1, 1])
  end

  test "compare integer and list" do
    refute Day13.compare([9], [[8, 7, 6]])
    assert Day13.compare([[8, 7, 6]], [9])
  end

  test "compare list and list" do
    refute Day13.compare([1, [2, [3, [4, [5, 6, 7]]]], 8, 9], [1, [2, [3, [4, [5, 6, 0]]]], 8, 9])
    assert Day13.compare([1, [2, [3, [4, [5, 6, 0]]]], 8, 9], [1, [2, [3, [4, [5, 6, 7]]]], 8, 9])
  end

  test "example case failing" do
    assert Day13.compare([[4, 4], 4, 4], [[4, 4], 4, 4, 4])
  end

  test "another failed test case" do
    refute Day13.compare([[[]]], [[]])
  end
end
