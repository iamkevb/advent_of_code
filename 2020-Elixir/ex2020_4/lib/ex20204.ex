defmodule Ex20204 do
  def readInput(path) do
    File.read!(path)
    |> String.split("\n\n")
    |> Enum.map(fn s -> String.replace(s, "\n", " ") end)
  end

  # byr (Birth Year)
  # iyr (Issue Year)
  # eyr (Expiration Year)
  # hgt (Height)
  # hcl (Hair Color)
  # ecl (Eye Color)
  # pid (Passport ID)
  # cid (Country ID)
  def parse(entry) do
    String.split(entry)
    |> Enum.reduce(%{}, fn e, acc ->
      [k, v] = String.split(e, ":")
      Map.put(acc, k, v)
    end)
  end

  # valid if all 8 keys are present, or only cid is missing
  # so if first 7 keys are present?
  def validate(%{
        "byr" => _,
        "iyr" => _,
        "eyr" => _,
        "hgt" => _,
        "hcl" => _,
        "ecl" => _,
        "pid" => _
      }) do
    true
  end

  def validate(_) do
    false
  end

  # byr (Birth Year) - four digits; at least 1920 and at most 2002.
  def validByr(byr) do
    year = String.to_integer(byr)
    1920 <= year && year <= 2002
  end

  # iyr (Issue Year) - four digits; at least 2010 and at most 2020.
  def validIyr(iyr) do
    year = String.to_integer(iyr)
    2010 <= year && year <= 2020
  end

  # eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
  def validEyr(eyr) do
    year = String.to_integer(eyr)
    2020 <= year && year <= 2030
  end

  # hgt (Height) - a number followed by either cm or in:
  # If cm, the number must be at least 150 and at most 193.
  # If in, the number must be at least 59 and at most 76.
  def validHgt(hgt) do
    try do
      [[_, h, u]] = Regex.scan(~r/(\d+)(cm|in)/, hgt)

      hi = String.to_integer(h)

      case u do
        "cm" -> 150 <= hi && hi <= 193
        "in" -> 59 <= hi && hi <= 76
        _ -> false
      end
    rescue
      _ -> false
    end
  end

  # hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
  def validHcl(hcl) do
    Regex.match?(~r/#[a-zA-Z0-9]{6}/, hcl)
  end

  # ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
  def validEcl(ecl) do
    ecl === "amb" ||
      ecl === "blu" ||
      ecl === "brn" ||
      ecl === "gry" ||
      ecl === "grn" ||
      ecl === "hzl" ||
      ecl === "oth"
  end

  # pid (Passport ID) - a nine-digit number, including leading zeroes.
  def validPid(pid) do
    p = Integer.parse(pid)

    if p === :error do
      false
    else
      {_, r} = p
      r === "" && String.length(pid) === 9
    end
  end

  def validate2(%{
        "byr" => byr,
        "iyr" => iyr,
        "eyr" => eyr,
        "hgt" => hgt,
        "hcl" => hcl,
        "ecl" => ecl,
        "pid" => pid
      }) do
    validByr(byr) && validIyr(iyr) && validEyr(eyr) && validHgt(hgt) && validHcl(hcl) &&
      validEcl(ecl) && validPid(pid)
  end

  def validate2(_) do
    false
  end

  def part1(path \\ "input.test.txt") do
    readInput(path)
    |> Enum.map(&parse(&1))
    |> Enum.filter(&validate/1)
    |> length()
  end

  def part2(path \\ "input.test.txt") do
    readInput(path)
    |> Enum.map(&parse(&1))
    |> Enum.filter(&validate2/1)
    |> length()
  end
end
