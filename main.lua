#!/usr/bin/luajit

local numbers = {"Mac", "GNU/Linux", "Windows", "BSD", "plan9", "Solaris"}

local hsr = require "hsr"

local function printr(array)
    if type(array) == "table" then
        local size = hsr.len(array)

        for i = 1, size do hsr.printf("%s. %s\n", i, array[i]) end
    end
end

local function sort(array)
    local al = array
    local size = hsr.len(al) -- 2
    local left, right
    local sorted = {}

    if size > 1 then
        local middle = math.floor(size / 2) -- 1

        left = sort(hsr.cut(al, 1, middle))

        right = sort(hsr.cut(al, middle + 1, size))

        while (type(left) == "table" and hsr.len(left) > 0) or
            (type(right) == "table" and hsr.len(right) ~= 0) do

            if type(left) ~= "table" or hsr.len(left) == 0 then
                table.insert(sorted, right[1])
                right = hsr.pop(right)
            elseif type(right) ~= "table" or hsr.len(right) == 0 then
                table.insert(sorted, left[1])
                left = hsr.pop(left)
            else
                io.write("What is better: (1) ", left[1], " or (2) ", right[1],
                         "? ")
                local response = io.read()

                if response == "2" then
                    table.insert(sorted, right[1])
                    right = hsr.pop(right)
                elseif response == "1" then
                    table.insert(sorted, left[1])
                    left = hsr.pop(left)
                else
                    print("????")
                    os.exit()
                end
            end
        end

        al = sorted
    end

    return al
end

local final = sort(numbers)
printr(final)
