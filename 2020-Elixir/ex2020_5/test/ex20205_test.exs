defmodule Ex20205Test do
  use ExUnit.Case
  doctest Ex20205

  test "find row" do
    assert Ex20205.findRow("BFFFBBFRRR") == 70
    assert Ex20205.findRow("FFFBBBFRRR") == 14
    assert Ex20205.findRow("BBFFBBFRLL") == 102
    assert Ex20205.findRow("BBBBBBBRLL") == 127
    assert Ex20205.findRow("FFFFFFFRLL") == 0
  end

  test "find seat" do
    assert Ex20205.findSeat("BFFFBBFRRR") == 7
    assert Ex20205.findSeat("FFFBBBFRRR") == 7
    assert Ex20205.findSeat("BBFFBBFRLL") == 4
  end
end
