-- common
local common = {}

common.WatchSystemSignal = function()
    if WatchSystemSignal ~= nil then
        WatchSystemSignal()
    end
end

common.WriteObj = function(conn, buf)
    if WriteObj ~= nil then
        WriteObj(conn, buf)
    end
end

return common