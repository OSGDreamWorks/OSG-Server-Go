
local logger = import("logger")
local mvc = import("mvc")

local FsConnMgr = class("FsConnMgr", mvc.AppBase)

function FsConnMgr:ctor(appName)
    FsConnMgr.super.ctor(self, appName)
end

function FsConnMgr:init(cfg)
    self.fspool = {}
    for key, value in pairs(cfg.FsHost) do
        logger.Debug("key %v, value:%v", key, value)
        self.fspool[key] = RpcClient.new(value)
    end

    self.poolsize = #cfg.FsHost
    self.workindex = 1

end

function FsConnMgr:GetWorkConn()

    self.workindex = self.workindex + 1
    if self.workindex > self.poolsize then
        self.workindex = 1
    end

    return self.fspool[self.workindex]
end

function FsConnMgr:Call(serviceMethod, req, rst)
    return self:GetWorkConn():Call(serviceMethod, req, rst)
end

return FsConnMgr