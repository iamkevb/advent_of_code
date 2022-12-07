defmodule Day7Test do
  use ExUnit.Case
  doctest Day7

  test "adjust_cwd" do
    assert {} === Day7.adjust_cwd({:a}, "..")
    assert {:a} === Day7.adjust_cwd({}, "a")
    assert {:a, :b} === Day7.adjust_cwd({:a}, "b")
    assert {:a, :a} === Day7.adjust_cwd({:a}, "a")
  end

  test "add_dir" do
    assert %{a: %{}} === Day7.add_dir({}, %{}, "a")
    assert %{a: %{}, b: %{}} === Day7.add_dir({}, %{a: %{}}, "b")
    assert %{a: %{b: %{}}} === Day7.add_dir({:a}, %{a: %{}}, "b")
    assert %{a: %{b: %{c: %{}}}} === Day7.add_dir({:a, :b}, %{a: %{b: %{}}}, "c")
  end

  test "add_file" do
    assert %{a: %{f: 123}} === Day7.add_file({:a}, %{a: %{}}, "f", 123)
    assert %{a: %{f: 123, g: 234}} === Day7.add_file({:a}, %{a: %{f: 123}}, "g", 234)
  end
end
