-- config
local config = {}

config.ReadConfig = function(file_name)
    local f = assert(io.open(file_name, 'r'))
    local jsonString = f:read("*all")
    if Json ~= nil then
        jsonString = Json.Decode(jsonString)
    end
    f:close()
    return jsonString
end

return config