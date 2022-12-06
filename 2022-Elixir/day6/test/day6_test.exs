defmodule Day6Test do
  use ExUnit.Case
  doctest Day6

  test "Part 1, 1" do
    assert 5 == Day6.firstMarker("bvwbjplbgvbhsrlpgdmjqwftvncz")
  end

  test "Part 1, 2" do
    assert 6 == Day6.firstMarker("nppdvjthqldpwncqszvftbrmjlhg")
  end

  test "Part 1, 3" do
    assert 10 == Day6.firstMarker("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg")
  end

  test "Part 1, 4" do
    assert 11 == Day6.firstMarker("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw")
  end

  test "Part 2, 1" do
    assert 19 == Day6.firstMessage("mjqjpqmgbljsphdztnvjfqwrcgsmlb")
  end

  test "Part 2, 2" do
    assert 23 == Day6.firstMessage("bvwbjplbgvbhsrlpgdmjqwftvncz")
  end

  test "Part 2, 3" do
    assert 23 == Day6.firstMessage("nppdvjthqldpwncqszvftbrmjlhg")
  end

  test "Part 2, 4" do
    assert 29 == Day6.firstMessage("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg")
  end

  test "Part 2, 5" do
    assert 26 == Day6.firstMessage("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw")
  end
end
