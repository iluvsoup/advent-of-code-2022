local file = io.open("input.txt", "r")
local input = file:read("*a")

local cursorPos = 1

while cursorPos <= #input - 13 do
  local characters = string.sub(input, cursorPos, cursorPos + 13)

  local duplicate = false
  for charIndex = 1, 14 do
    if duplicate then break end
    for i = 1, 14 do
      if i ~= charIndex and string.sub(characters, i, i) == string.sub(characters, charIndex, charIndex) then
        cursorPos = cursorPos + charIndex
        duplicate = true
        break
      end
    end
  end

  if duplicate == false then
    print(cursorPos + 13)
    break
  end
end