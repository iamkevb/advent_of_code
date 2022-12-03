defmodule Day3Test do
  use ExUnit.Case
  doctest Day3

  test "find badge" do
    badge =
      Day3.findBadge([
        "WjvRSdSQjvpjWzNlnZlNZqCCMzZZ",
        "nJtJsbctPBPwLNcDZNNGLClC",
        "tsFJHBgJwgJbnvSHHWVWHhVhpQ"
      ])

    assert badge === "n"
  end

  test "common glyphs" do
    a = "WjvRSdSQjvpjWzNlnZlNZqCCMzZZ"
    b = "nJtJsbctPBPwLNcDZNNGLClC"
    c = "tsFJHBgJwgJbnvSHHWVWHhVhpQ"
    d = Day3.commonGlyphs(a, b)

    assert d === "NlnZC"
    e = Day3.commonGlyphs(c, d)
    assert e === "n"
  end
end
