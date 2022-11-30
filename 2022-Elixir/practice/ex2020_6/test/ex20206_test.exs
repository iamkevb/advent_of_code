defmodule Ex20206Test do
  use ExUnit.Case
  doctest Ex20206

  test "countGroup" do
    assert Ex20206.countGroup("123") === 3
  end
end
