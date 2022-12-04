sum = 0

IO.foreach("input.txt") { |line|
  ranges = line.split(",")
  range1 = ranges[0].split("-")
  range2 = ranges[1].split("-")

  # For some reason you're allowed to compare strings in ruby
  # but it yields the wrong answer ???
  range1[0] = Integer(range1[0])
  range1[1] = Integer(range1[1])
  range2[0] = Integer(range2[0])
  range2[1] = Integer(range2[1])
  
  if range1[1] >= range2[0] and range1[0] <= range2[1] then
    sum += 1
  end
}

puts sum