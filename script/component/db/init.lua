-- db
local DBInit = DBInit
local DBQuery = DBQuery
local DBWrite = DBWrite
local DBDelete = DBDelete

module('db')

Init = function()
    if DBInit ~= nil then
        DBInit()
    end
end

--return value exist err
Query = function(tname, key, value)
    if DBQuery ~= nil then
        return DBQuery(tname, key, value)
    end
end

--return result err
Write = function(tname, key, value)
    if DBWrite ~= nil then
        return DBWrite(tname, key, value)
    end
end

--return result err
Delete = function(tname, key)
    if DBDelete ~= nil then
        return DBDelete(tname, key)
    end
end