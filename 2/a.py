file = open("input.txt", "r")
input = file.read()
rounds = input.split("\n")

score = 0

for round in rounds:
  moves = round.split(" ")
  opponentMove = moves[0]
  playerMove = moves[1]

  if playerMove == "X":
    score += 1
  elif playerMove == "Y":
    score += 2
  elif playerMove == "Z":
    score += 3

  if (opponentMove == "A" and playerMove == "X") or (opponentMove == "B" and playerMove == "Y") or (opponentMove == "C" and playerMove == "Z"):
    score += 3
  elif opponentMove == "A" and playerMove == "Y":
    score += 6
  elif opponentMove == "B" and playerMove == "Z":
    score += 6
  elif opponentMove == "C" and playerMove == "X":
    score += 6

print(score)