defmodule Ex20203 do
  def readInput(path) do
    File.read!(path)
    |> String.split("\n")
  end

  def isTree(nil, _) do
    false
  end

  def isTree(rowString, col) do
    len = String.length(rowString)

    if col == 0 || col < len do
      String.at(rowString, col) === "#"
    else
      isTree(rowString, col - len)
    end
  end

  def isTree(rows, row, col) do
    Enum.at(rows, row)
    |> isTree(col)
  end

  def countTrees(rows, visits) do
    Enum.filter(visits, fn v -> isTree(rows, v.row, v.col) end)
    |> length()
  end

  def visited(maxRow, col, row, colStep, rowStep) do
    nextCol = col + colStep
    nextRow = row + rowStep

    if nextRow > maxRow do
      []
    else
      [%{row: nextRow, col: nextCol} | visited(maxRow, nextCol, nextRow, colStep, rowStep)]
    end
  end

  def part1(path \\ "input.prod.txt") do
    grid = readInput(path)
    visits = visited(length(grid), 0, 0, 3, 1)
    countTrees(grid, visits)
  end

  def part2(path \\ "input.prod.txt") do
    grid = readInput(path)

    countTrees(grid, visited(length(grid), 0, 0, 1, 1)) *
      countTrees(grid, visited(length(grid), 0, 0, 3, 1)) *
      countTrees(grid, visited(length(grid), 0, 0, 5, 1)) *
      countTrees(grid, visited(length(grid), 0, 0, 7, 1)) *
      countTrees(grid, visited(length(grid), 0, 0, 1, 2))
  end
end
