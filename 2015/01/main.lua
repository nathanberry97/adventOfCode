#!/usr/bin/lua

-- getData
-- @return string
local function getData()
    local puzzleInput = io.open("./input.txt", "r")

    if not puzzleInput then
        return ""
    end

    local data = puzzleInput:read("a")

    puzzleInput:close()

    return data
end

-- main
local function main()
    local data = getData()
    local answer, basement = 0, true

    for i = 1, #data do
        local char = string.sub(data, i, i)

        if char == "(" then
            answer = answer + 1
        elseif char == ")" then
            answer = answer - 1
        end

        if answer == -1 and basement then
            print("Entered the basement at: " .. i)
            basement = false
        end
    end

    print("Final floor: " .. answer)
end

main()
