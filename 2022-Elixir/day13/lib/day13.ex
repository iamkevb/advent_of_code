defmodule Day13 do
  def compare({l, r}), do: compare(l, r)
  def compare(l, r) when is_integer(l) and is_integer(r) and l < r, do: true
  def compare(l, r) when is_integer(l) and is_integer(r) and l > r, do: false
  def compare(l, r) when is_integer(l) and is_integer(r), do: :continue

  def compare([lh | lt], [rh | rt]),
    do: if(is_boolean(r = compare(lh, rh)), do: r, else: compare(lt, rt))

  def compare([], [_ | _]), do: true
  def compare([_ | _], []), do: false
  def compare([], []), do: :continue
  def compare(l, r) when is_list(l) and is_integer(r), do: compare(l, [r])
  def compare(l, r) when is_integer(l) and is_list(r), do: compare([l], r)

  def part1(path \\ "input.test.txt") do
    File.read!(path)
    |> String.split("\n")
    |> Enum.reject(&(String.length(&1) === 0))
    |> Enum.map(&Code.eval_string/1)
    |> Enum.map(&elem(&1, 0))
    |> Enum.chunk_every(2)
    |> Enum.map(&List.to_tuple/1)
    |> Enum.map(fn {l, r} ->
      compare(l, r)
    end)
    |> Enum.with_index()
    |> Enum.filter(fn {v, _} -> v end)
    |> Enum.reduce(0, fn {_, i}, acc ->
      acc + i + 1
    end)
  end

  def part2(path \\ "input.test.txt") do
    ("[2]\n[6]\n" <> File.read!(path))
    |> String.split("\n")
    |> Enum.reject(&(String.length(&1) === 0))
    |> Enum.map(&Code.eval_string/1)
    |> Enum.map(&elem(&1, 0))
    |> Enum.sort(&compare/2)
    |> Enum.with_index()
    |> Enum.filter(fn {e, _} -> e === [2] || e === [6] end)
    |> Enum.reduce(1, fn {_, i}, acc -> acc * (i + 1) end)
  end
end
