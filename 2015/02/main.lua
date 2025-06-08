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

-- formatInput formats the string input into L, W, H numbers
-- @param data string
-- @return number[]
local function formatInput(data)
    local array = {}
    for num in data:gmatch("%d+") do
        table.insert(array, tonumber(num))
    end
    return array
end


-- smallestArea
-- @param data string
-- @return number
local function smallestArea(data)
    local dims = formatInput(data)
    table.sort(dims)
    return dims[1] * dims[2]
end

-- boxSurface
-- @param data string
-- @return number
local function boxSurface(data)
    local l, w, h = table.unpack(formatInput(data))
    return 2 * l * w + 2 * w * h + 2 * h * l
end

-- ribbonLength
-- @param data string
-- @return number
local function ribbonLength(data)
    local dims = formatInput(data)
    table.sort(dims)
    local perimeter = 2 * (dims[1] + dims[2])
    local volume = dims[1] * dims[2] * dims[3]
    return perimeter + volume
end

local function main()
    local data = getData()
    local wrappingPaper = 0
    local ribbon = 0

    for line in data:gmatch("[^\r\n]+") do
        wrappingPaper = wrappingPaper + (boxSurface(line) + smallestArea(line))
        ribbon = ribbon + ribbonLength(line)
    end

    print("The amount of square feet of wrapping paper: " .. wrappingPaper)
    print("The amount of ribbon to order: " .. ribbon)
end

main()
