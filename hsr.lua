local hsr = {}

function hsr.printf(s, ...) return io.write(s:format(...)) end

function hsr.cut(array, start, finish)
    -- printf("cut called with %s and %s\n", start, finish)
    local result = {}
    local z = 1

    for i = start, finish, 1 do
        result[z] = array[i]
        z = z + 1
    end

    if type(result) == nil then result = {} end

    return result
end

function hsr.pop(array)
    local a = array

    if hsr.len(a) > 1 then
        a = hsr.cut(a, 2, hsr.len(a))
    else
        a = {}
    end
    return a
end

function hsr.len(array)
    local result
    result = table.getn(array)
    return result
end

return hsr
