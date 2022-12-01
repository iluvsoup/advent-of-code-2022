const fs = require("fs")

const data = fs.readFileSync("input.txt", { encoding: "utf8" })

const elves = data.split("\n\n")

let top1 = 0
let top2 = 0
let top3 = 0

for (const elf of elves) {
  let calories = 0
  for (const line of elf.split("\n")) {
    calories += parseInt(line)
  }
  if (calories > top1) {
    top3 = top2
    top2 = top1
    top1 = calories
  } else if (calories > top2) {
    top3 = top2
    top2 = calories
  } else if (calories > top3) {
    top3 = calories
  }
}

console.log(top1 + top2 + top3)