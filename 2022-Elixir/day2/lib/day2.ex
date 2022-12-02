defmodule Day2 do
  def readInput(path) do
    File.read!(path)
    |> String.split()
    |> Enum.chunk_every(2)
  end

  #  A X: Rock 1 point
  #  B Y: Paper 2 points
  #  C Z: Scissors 3 points

  # Score Points for your move + (6 for win|3 for draw|0 for loss)

  def score([opp, me]) do
    case {opp, me} do
      {"A", "X"} -> 1 + 3
      {"A", "Y"} -> 2 + 6
      {"A", "Z"} -> 3 + 0
      {"B", "X"} -> 1 + 0
      {"B", "Y"} -> 2 + 3
      {"B", "Z"} -> 3 + 6
      {"C", "X"} -> 1 + 6
      {"C", "Y"} -> 2 + 0
      {"C", "Z"} -> 3 + 3
    end
  end

  # A,B,C: Rock, Paper, Scissors
  # X,Y,Z: Lose, Draw, Win
  def moveAndScore([opp, result]) do
    score2(result) + move({opp, result})
  end

  def score2(result) do
    case result do
      "Y" -> 3
      "Z" -> 6
      _ -> 0
    end
  end

  def move({opp, result}) do
    case {opp, result} do
      # LOSE
      {"A", "X"} -> 3
      {"B", "X"} -> 1
      {"C", "X"} -> 2
      # DRAW
      {"A", "Y"} -> 1
      {"B", "Y"} -> 2
      {"C", "Y"} -> 3
      # WIN
      {"A", "Z"} -> 2
      {"B", "Z"} -> 3
      {"C", "Z"} -> 1
    end
  end

  def part1(path \\ "input.test.txt") do
    readInput(path)
    |> Enum.map(&score/1)
    |> Enum.sum()
  end

  def part2(path \\ "input.test.txt") do
    readInput(path)
    |> Enum.map(&moveAndScore/1)
    |> Enum.sum()
  end
end
