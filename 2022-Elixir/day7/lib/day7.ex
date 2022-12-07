defmodule Day7 do
  def readInput(path), do: File.read!(path) |> String.split("\n")

  def adjust_cwd(cwd, ".."), do: Tuple.delete_at(cwd, tuple_size(cwd) - 1)
  def adjust_cwd(cwd, dir), do: Tuple.append(cwd, String.to_atom(dir))

  def add_dir(cwd, fs, dir),
    do:
      put_in(
        fs,
        Tuple.append(cwd, String.to_atom(dir)) |> Tuple.to_list(),
        %{}
      )

  def add_file(cwd, fs, fName, fSize),
    do:
      put_in(
        fs,
        Tuple.append(cwd, String.to_atom(fName)) |> Tuple.to_list(),
        fSize
      )

  def handle_cmd(["$", "cd", "/"], cwd, fs), do: {cwd, fs}
  def handle_cmd(["$", "cd", dir], cwd, fs), do: {adjust_cwd(cwd, dir), fs}
  def handle_cmd(["$", "ls"], cwd, fs), do: {cwd, fs}
  def handle_cmd(["dir", dir], cwd, fs), do: {cwd, add_dir(cwd, fs, dir)}

  def handle_cmd([f_size, f_name], cwd, fs),
    do: {cwd, add_file(cwd, fs, f_name, String.to_integer(f_size))}

  def build_fs(cmd, {cwd, fs}), do: handle_cmd(String.split(cmd), cwd, fs)

  def size_of(v) when is_map(v) do
    Enum.reduce(v, 0, fn {_, v}, acc ->
      acc + size_of(v)
    end)
  end

  def size_of(v), do: v

  def compute_sizes(map), do: [{"/", size_of(map)} | compute_sizes(map, "/", [])]

  def compute_sizes(map, parent, sizes) do
    map
    |> Enum.reduce(sizes, fn {k, v}, acc ->
      key = parent <> to_string(k)

      cond do
        is_map(v) ->
          [{key, size_of(v)} | compute_sizes(v, key, acc)]

        true ->
          acc
      end
    end)
  end

  def part1(path \\ "input.test.txt") do
    readInput(path)
    |> Enum.reduce({{}, %{}}, &build_fs/2)
    |> elem(1)
    |> compute_sizes()
    |> Enum.filter(fn {_, s} -> s <= 100_000 end)
    |> Enum.reduce(0, fn {_, s}, acc -> acc + s end)
  end

  def part2(path \\ "input.test.txt") do
    sizes =
      readInput(path)
      |> Enum.reduce({{}, %{}}, &build_fs/2)
      |> elem(1)
      |> compute_sizes()
      |> Enum.map(fn {_, v} -> v end)
      |> Enum.sort(&(&2 < &1))

    available = 70_000_000 - hd(sizes)
    needed = max(30_000_000 - available, 0)
    IO.puts("available: #{available} needed: #{needed}, sizes: #{hd(sizes)}")

    Enum.filter(sizes, &(&1 >= needed))
    |> Enum.reverse()
    |> hd()
  end
end
