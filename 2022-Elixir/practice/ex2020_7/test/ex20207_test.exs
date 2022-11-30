defmodule Ex20207Test do
  use ExUnit.Case
  doctest Ex20207

  test "parseRule" do
    r = "light red bags contain 1 bright white bag, 2 muted yellow bags."
    m = Ex20207.parseRule(r)
    assert Map.keys(m) === ["light red"]

    assert m["light red"]["bright white"] === 1
    assert m["light red"]["muted yellow"] === 2
  end

  test "holds" do
    rule = %{"faded blue" => 9, "shiny gold" => 2}
    assert Ex20207.holds(rule, "shiny gold") === 2
  end

  test "countBagsInside dotted black" do
    rules = ["dotted black bags contain no other bags."]
    ruleMap = Ex20207.createRules(rules)
    assert Ex20207.countBagsInside(ruleMap, "dotted black") === 0
  end

  test "countBagsInside holds 1 black" do
    rules = [
      "dotted black bags contain no other bags.",
      "polka dot bags contain 1 dotted black bag."
    ]

    ruleMap = Ex20207.createRules(rules)
    assert Ex20207.countBagsInside(ruleMap, "polka dot") === 1
  end

  test "countBagsInside holds 2 black" do
    rules = [
      "dotted black bags contain no other bags.",
      "polka dot bags contain 2 dotted black bags."
    ]

    ruleMap = Ex20207.createRules(rules)
    inPolkaDot = Ex20207.countBagsInside(ruleMap, "polka dot")

    assert inPolkaDot === 2
  end

  test "purple dinosaur hold 1 polka dot holds 2 black" do
    rules = [
      "dotted black bags contain no other bags.",
      "polka dot bags contain 2 dotted black bags.",
      "purple dinosaur bags contain 1 polka dot bag."
    ]

    ruleMap = Ex20207.createRules(rules)
    assert Ex20207.countBagsInside(ruleMap, "purple dinosaur") === 3
  end

  test "dark olive has 7 inside" do
    rules = [
      "dark olive bags contain 3 faded blue bags, 4 dotted black bags.",
      "faded blue bags contain no other bags.",
      "dotted black bags contain no other bags."
    ]

    ruleMap = Ex20207.createRules(rules)
    assert Ex20207.countBagsInside(ruleMap, "dark olive") === 7
  end

  @tag runnable: true
  test "shiny gold has 24" do
    rules = [
      "shiny gold bags contain 2 vibrant plum bags.",
      "vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.",
      "faded blue bags contain no other bags.",
      "dotted black bags contain no other bags."
    ]

    ruleMap = Ex20207.createRules(rules)
    assert Ex20207.countBagsInside(ruleMap, "shiny gold") === 24
  end
end
