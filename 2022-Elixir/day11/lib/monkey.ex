defmodule Monkey do
  @type t :: %__MODULE__{
          items: {Integer},
          operation: {String.t(), Integer},
          testValue: Integer,
          actionTrue: Integer,
          actionFalse: Integer,
          inspected: Integer
        }
  defstruct [:items, :operation, :testValue, :actionTrue, :actionFalse, :inspected, :product]

  def operate(level, {"*", -1}), do: level * level
  def operate(level, {"+", -1}), do: level + level
  def operate(level, {"*", value}), do: level * value
  def operate(level, {"+", value}), do: level + value

  def inspect(monkey) do
    if tuple_size(monkey.items) === 0 do
      monkey
    else
      # monkey inspects, then gets bored
      # part 2 removes |> div(3)
      item = operate(elem(monkey.items, 0), monkey.operation)
      item = rem(item, monkey.product)
      [_ | il] = Tuple.to_list(monkey.items)

      %Monkey{monkey | items: List.to_tuple([item | il]), inspected: monkey.inspected + 1}
    end
  end

  def action(monkey, 0), do: monkey.actionTrue
  def action(monkey, _), do: monkey.actionFalse

  def nextMonkeyIndex(monkey) do
    level = elem(monkey.items, 0)
    action(monkey, rem(level, monkey.testValue))
  end

  def throwItem(monkey),
    do: {elem(monkey.items, 0), %Monkey{monkey | items: Tuple.delete_at(monkey.items, 0)}}

  def catchItem(monkey, item), do: %Monkey{monkey | items: Tuple.append(monkey.items, item)}
end
