const fs = require("fs")

const data = fs.readFileSync("input.txt", { encoding: "utf8" })

const elves = data.split("\n\n")

let highest = 0

for (const elf of elves) {
  let calories = 0
  for (const line of elf.split("\n")) {
    calories += parseInt(line)
  }
  if (calories > highest) {
    highest = calories
  }
}

console.log(highest)