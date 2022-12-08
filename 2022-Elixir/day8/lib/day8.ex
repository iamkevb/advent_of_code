defmodule Tree do
  defstruct height: 0, visible: false
end

defmodule Day8 do
  def readInput(path) do
    File.read!(path)
    |> String.split()
    |> Enum.map(fn s ->
      String.graphemes(s)
      |> Enum.map(fn e -> %Tree{height: e} end)
    end)
  end

  def markVisibleTreeRow(row) do
    row
    |> Enum.reduce({-1, {}}, fn tree, {maxH, row} ->
      visible = tree.visible || tree.height > maxH

      {
        max(maxH, tree.height),
        Tuple.append(row, %Tree{height: tree.height, visible: visible})
      }
    end)
    |> elem(1)
    |> Tuple.to_list()
  end

  def markVisibleTrees(grid) do
    grid
    |> Enum.map(fn row -> markVisibleTreeRow(row) end)
  end

  def part1(path \\ "input.test.txt") do
    readInput(path)

    # L->R
    |> markVisibleTrees()
    # R->L
    |> Enum.map(&Enum.reverse/1)
    |> markVisibleTrees()

    # Bottom->Top
    |> Enum.zip()
    |> Enum.map(&Tuple.to_list/1)
    |> markVisibleTrees()

    # T-> B
    |> Enum.map(&Enum.reverse/1)
    |> markVisibleTrees()

    # count trues
    |> Enum.reduce(0, fn row, acc ->
      Enum.reduce(row, acc, fn %{visible: v}, acc ->
        case v do
          true -> acc + 1
          false -> acc
        end
      end)
    end)
  end
end
