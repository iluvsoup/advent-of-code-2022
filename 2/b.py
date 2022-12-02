file = open("input.txt", "r")
input = file.read()
rounds = input.split("\n")

score = 0

for round in rounds:
  moves = round.split(" ")
  opponentMove = moves[0]
  roundResult = moves[1]

  playerMove = None

  # lose
  if roundResult == "X":
    if opponentMove == "A":
      playerMove = "Z"
    elif opponentMove == "B":
      playerMove = "X"
    elif opponentMove == "C":
      playerMove = "Y"
  # tie
  elif roundResult == "Y":
    score += 3
    
    if opponentMove == "A":
      playerMove = "X"
    elif opponentMove == "B":
      playerMove = "Y"
    elif opponentMove == "C":
      playerMove = "Z"
  # win
  elif roundResult == "Z":
    score += 6
    
    if opponentMove == "A":
      playerMove = "Y"
    elif opponentMove == "B":
      playerMove = "Z"
    elif opponentMove == "C":
      playerMove = "X"
  
  if playerMove == "X":
    score += 1
  elif playerMove == "Y":
    score += 2
  elif playerMove == "Z":
    score += 3

print(score)