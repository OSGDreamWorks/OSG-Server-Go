-- common
local cfg = import(".config")
local logger = import(".logger")
local watchSystemSignal = WatchSystemSignal
local writeObj = WriteObj

module('common')

config = cfg

WatchSystemSignal = function()
    if watchSystemSignal ~= nil then
        watchSystemSignal()
    end
end

WriteObj = function(conn, buf)
    if writeObj ~= nil then
        writeObj(conn, buf)
    end
end

return common