defmodule Ex20202 do
  def readInput(path) do
    File.read!(path)
    |> String.split("\n")
  end

  def isValid(passwordEntry) do
    # 1-3 a: abcde
    [[_, min, max, letter, password]] = Regex.scan(~r/(\d+)-(\d+) (\S): (\S+)/, passwordEntry)

    count = Enum.filter(String.graphemes(password), fn c -> c == letter end) |> length()

    # IO.puts(
    #   "min #{min}, max #{max}, count #{count} min<=count #{String.to_integer(min, 10) <= count} count <= max #{count <= String.to_integer(max, 10)}"
    # )

    String.to_integer(min, 10) <= count && count <= String.to_integer(max, 10)
  end

  def countValid(passwords) do
    Enum.filter(passwords, &isValid(&1))
    |> length()
  end

  def countValid2(passwords) do
    Enum.filter(passwords, &isValid2(&1))
    |> length()
  end

  def isValid2(passwordEntry) do
    # 1-3 a: abcde
    [[_, idx1, idx2, letter, password]] = Regex.scan(~r/(\d+)-(\d+) (\S): (\S+)/, passwordEntry)

    at1 = String.at(password, String.to_integer(idx1) - 1)
    at2 = String.at(password, String.to_integer(idx2) - 1)
    (at1 === letter || at2 === letter) && at1 !== at2
  end

  def part1() do
    readInput("real_input.txt")
    |> countValid
  end

  def part2() do
    readInput("real_input.txt")
    |> countValid2
  end
end
