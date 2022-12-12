defmodule MonkeyTest do
  use ExUnit.Case
  doctest Monkey

  test "Monkey operate *" do
    m = %Monkey{operation: {"*", "2"}}
    assert 246 = Monkey.operate(123, m.operation)
  end

  test "Monkey operate +" do
    m = %Monkey{operation: {"+", "2"}}
    assert 125 = Monkey.operate(123, m.operation)
  end

  test "Monkey operate * old" do
    m = %Monkey{operation: {"*", "old"}}
    assert 16 = Monkey.operate(4, m.operation)
  end

  test "Monkey operate + old" do
    m = %Monkey{operation: {"+", "old"}}
    assert 8 = Monkey.operate(4, m.operation)
  end

  test "throw item" do
    m1 = %Monkey{items: {1, 2, 3}}
    {item, m} = Monkey.throwItem(m1)
    assert 1 === item
    assert {2, 3} === m.items
  end

  test "catch item" do
    m1 = %Monkey{items: {1, 2, 3}}
    m = Monkey.catchItem(m1, 9)
    assert {1, 2, 3, 9} === m.items
  end

  test "inspect" do
    m = %Monkey{
      items: {79, 98},
      operation: {"*", "19"},
      testValue: 23,
      actionTrue: 2,
      actionFalse: 3,
      inspected: 0
    }

    m2 = Monkey.inspect(m)
    # 79 * 19 / 3 = 500
    assert {500, 98} === m2.items
    assert 1 === m2.inspected
  end

  test "next monkey index" do
    m = %Monkey{
      items: {79, 98},
      operation: {"*", "19"},
      testValue: 23,
      actionTrue: 2,
      actionFalse: 3,
      inspected: 0
    }

    assert 3 = Monkey.nextMonkeyIndex(m)
  end
end
