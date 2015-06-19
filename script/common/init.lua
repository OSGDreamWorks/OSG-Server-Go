-- common
local cfg = import(".config")
local logger = import(".logger")
local watchSystemSignal = WatchSystemSignal
local writeObj = WriteObj
local setInterval = SetInterval
local clearInterval = ClearInterval
local print = print

module('common')

config = cfg
coroutine_map = {}

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

Sleep = function(seconds)
    if sleep ~= nil then
        sleep(seconds)
    end
end

SetInterval = function (identifier_, duration_, func_)
    if setInterval ~= nil then
        local func = function()
            if func_ ~= nil then func_() end
        end
        setInterval(identifier_, 0, duration_, func)
    end
end

ClearInterval = function (identifier_)
    if clearInterval ~= nil then
        clearInterval(identifier_)
    end
end

SetTimeout = function (identifier_, delay_, func_)
    if setInterval ~= nil then
        local func = function()
            func_()
            clearInterval(identifier_)
        end
        setInterval(identifier_, delay_, 1, func)
    end
end

ClearTimeout = function (identifier_)
    if clearInterval ~= nil then
        clearInterval(identifier_)
    end
end

return common