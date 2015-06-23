-- logger

local ipairs = ipairs
local type = type
local table = table
local pairs = pairs
local pairs = pairs
local tostring = tostring
local next = next

local print = print
local os = os
local string = string
local debug = debug

module('logger')

Dump = function(...)
	local arg = {...}
	local root = {}
	for k,v in ipairs(arg) do
		if type(v) == "table" then
			root[k] = v
			if table.getn(arg) == 1 then
				root = v
			end
		else
			root[k] = v
		end
	end
	local cache = {  [root] = "." }
	local function _dump(t,space,name)
		local temp = {}
		for k,v in pairs(t) do
			local key = tostring(k)
			if cache[v] then
				table.insert(temp,"+" .. key .. " {" .. cache[v].."}")
			elseif type(v) == "table" then
				local new_key = name .. "." .. key
				cache[v] = new_key
				table.insert(temp,"+" .. key .. _dump(v,space .. (next(t,k) and "|" or " " ).. string.rep(" ",#key),new_key))
			else
				table.insert(temp,"+" .. key .. " [" .. tostring(v).."]")
			end
		end
		return table.concat(temp,"\n"..space)
	end
	print(_dump(root, "",""))
end

DumpString = function(v)

	if type(v) == "string" then
		local info = ""
		for i = 1, string.len(v) do
			local charcode = tonumber(string.byte(v, i, i));
			info = info .. string.format("%02X ", charcode);
		end
		print(os.date("%Y/%m/%d %H:%M:%S").." Info: [Lua] "..info)
	end

end

Debug = function(...)
    print(os.date("%Y/%m/%d %H:%M:%S").." Debug: [Lua] "..string.format(...))
end
Info = function(...)
    print(os.date("%Y/%m/%d %H:%M:%S").." Info: [Lua] "..string.format(...))
end
Warning = function(...)
    print(os.date("%Y/%m/%d %H:%M:%S").." Warning: [Lua] "..string.format(...))
end
Error = function(...)
    print(os.date("%Y/%m/%d %H:%M:%S").." Error: [Lua] "..string.format(...))
end
Fatal = function(...)
    print(os.date("%Y/%m/%d %H:%M:%S").." Fatal: [Lua] "..string.format(...))
	debug.traceback()
	os.exit(1)
end