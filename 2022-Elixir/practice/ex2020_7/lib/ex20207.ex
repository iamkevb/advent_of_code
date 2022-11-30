defmodule Ex20207 do
  def readInput(path) do
    File.read!(path)
    |> String.split("\n")
  end

  def parseRule(ruleDescriptor) do
    [bag, valueStr] =
      String.split(ruleDescriptor, "bags contain")
      |> Enum.map(&String.trim/1)

    case valueStr do
      "no other bags." ->
        %{}

      _ ->
        contains =
          valueStr
          |> String.split(",")
          |> Enum.reduce(%{}, fn s, acc ->
            [[_, count, bag]] = Regex.scan(~r/(\d+)\s(.*)\sbag/, s)
            Map.put(acc, bag, String.to_integer(count))
          end)

        %{bag => contains}
    end
  end

  def createRules(rulesDescriptors) do
    rulesDescriptors
    |> Enum.reduce(%{}, fn r, acc -> Map.merge(acc, parseRule(r)) end)
  end

  def holds(rule, bag) do
    case Map.fetch(rule, bag) do
      {:ok, count} -> count
      _ -> 0
    end
  end

  def containers(ruleMap, bag, containers \\ []) do
    Enum.reduce(ruleMap, containers, fn {k, v}, acc ->
      unless holds(v, bag) === 0 do
        containers(ruleMap, k, [v | acc])
      else
        acc
      end
    end)
  end

  def countBagsInside(ruleMap, bag, count \\ 0) do
    case Map.fetch(ruleMap, bag) do
      {:ok, val} ->
        Enum.reduce(val, count, fn {k, v}, acc ->
          inside = countBagsInside(ruleMap, k, 0)
          IO.puts("#{k}: #{v} + #{inside} * #{v} inside")
          acc + v + inside * v
        end)

      _ ->
        0
    end
  end

  @spec part1(
          binary
          | maybe_improper_list(
              binary | maybe_improper_list(any, binary | []) | char,
              binary | []
            )
        ) :: non_neg_integer
  def part1(path \\ "input.prod.txt") do
    readInput(path) |> createRules |> containers("shiny gold") |> Enum.uniq() |> length()
  end

  def part2(path \\ "input.prod.txt") do
    readInput(path) |> createRules |> countBagsInside("shiny gold")
  end
end
