defmodule Ex20205 do
  def part1(path \\ "input.prod.txt") do
    readInput(path)
    |> Enum.reduce(0, fn p, acc ->
      max(acc, findRow(p) * 8 + findSeat(p))
    end)
  end

  def part2(path \\ "input.prod.txt") do
    readInput(path)
    |> Enum.map(fn p ->
      row = findRow(p)

      case row do
        0 -> nil
        127 -> nil
        _ -> row * 8 + findSeat(p)
      end
    end)
    |> Enum.sort()
    |> findMissing()
  end

  def findMissing(sorted_list) do
    [left | tail] = sorted_list
    [curr | _] = tail

    if curr - left !== 1 do
      curr - 1
    else
      findMissing(tail)
    end
  end

  def readInput(path) do
    File.read!(path)
    |> String.split("\n")
  end

  def findRow(pass) do
    String.graphemes(pass)
    |> Enum.reduce({0, 127}, fn e, acc ->
      e0 = elem(acc, 0)
      e1 = elem(acc, 1)
      mid = e0 + div(e1 - e0, 2)

      case e do
        "F" ->
          {e0, mid}

        "B" ->
          {mid, e1}

        _ ->
          acc
      end
    end)
    |> elem(1)
  end

  def findSeat(pass) do
    String.graphemes(pass)
    |> Enum.reduce({0, 7}, fn e, acc ->
      e0 = elem(acc, 0)
      e1 = elem(acc, 1)
      mid = e0 + div(e1 - e0, 2)

      case e do
        "L" ->
          {e0, mid}

        "R" ->
          {mid, e1}

        _ ->
          acc
      end
    end)
    |> elem(1)
  end
end
