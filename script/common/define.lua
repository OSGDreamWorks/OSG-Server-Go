-- common

local common = {}

common.WatchSystemSignal = function()
    if WatchSystemSignal ~= nil then
        WatchSystemSignal()
    end
end

return common