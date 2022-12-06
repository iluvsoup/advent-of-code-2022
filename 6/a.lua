local file = io.open("input.txt", "r")
local input = file:read("*a")

local cursorPos = 1

while cursorPos <= #input - 3 do
  local characters = string.sub(input, cursorPos, cursorPos + 3)

  local duplicate = false
  for charIndex = 1, 4 do
    if duplicate then break end
    for i = 1, 4 do
      if i ~= charIndex and string.sub(characters, i, i) == string.sub(characters, charIndex, charIndex) then
        cursorPos = cursorPos + charIndex
        duplicate = true
        break
      end
    end
  end

  if duplicate == false then
    print(cursorPos + 3)
    break
  end
end