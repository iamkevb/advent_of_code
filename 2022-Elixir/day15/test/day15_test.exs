defmodule Day15Test do
  use ExUnit.Case
  doctest Day15

  test "range at y 8,7" do
    s = Sensor.scan("Sensor at x=8, y=7: closest beacon is at x=2, y=10")
    assert 2..14 === Day15.range_at_y(s, 10)
  end

  test "range at y, y is on target" do
    s = Sensor.scan("Sensor at x=8, y=10: closest beacon is at x=2, y=10")
    assert 2..14 === Day15.range_at_y(s, 10)
  end

  test "range at y pyramid" do
    s = Sensor.scan("Sensor at x=5, y=5: closest beacon is at x=5, y=15")
    assert -5..15 === Day15.range_at_y(s, 5)
    assert -4..14 === Day15.range_at_y(s, 6)
    assert -3..13 === Day15.range_at_y(s, 7)
    assert -2..12 === Day15.range_at_y(s, 8)
  end

  test "range way outside" do
    s = Sensor.scan("Sensor at x=5, y=500000: closest beacon is at x=5, y=500001")
    assert :ignore === Day15.range_at_y(s, 5)
  end

  test "combine ranges" do
    assert [1..2, 3..5] === Day15.combine_overlapping_ranges([1..2, 3..5])
    assert [1..2, 3..5] === Day15.combine_overlapping_ranges([3..5, 1..2])
    assert [1..5] === Day15.combine_overlapping_ranges([3..5, 1..3])
    assert [1..10] === Day15.combine_overlapping_ranges([1..4, 6..10, 3..7])
    assert [-2..24] === Day15.combine_overlapping_ranges([-2..18, 16..24])
    assert [1..10] === Day15.combine_overlapping_ranges([1..10, 3..5])
  end
end
