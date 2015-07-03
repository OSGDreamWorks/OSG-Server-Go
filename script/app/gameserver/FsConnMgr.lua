
local logger = import("logger")
local mvc = import("mvc")

local FsConnMgr = class("FsConnMgr", mvc.AppBase)

function FsConnMgr:ctor(appName)
    FsConnMgr.super.ctor(self, appName)
end

function FsConnMgr:init(cfg)

    self.fspool = {}
    for key, value in pairs(cfg.FsHost) do
        local fs = self
        self.fspool[key] = RpcClient.new(value)
        self.fspool[key]:AddDisCallback(
            function(err)
                fs.fspool[key]:ReConnect(value)
            end
        )
    end

    self.poolsize = #cfg.FsHost
    self.workindex = 1
    self.l = RWMutex.new()

end

function FsConnMgr:GetWorkConn()

    self.l:Lock()

    self.workindex = self.workindex + 1
    if self.workindex > self.poolsize then
        self.workindex = 1
    end

    self.l:Unlock()

    return self.fspool[self.workindex]

end

function FsConnMgr:Call(serviceMethod, req, rst)
    return self:GetWorkConn():Call(serviceMethod, req, rst)
end

return FsConnMgr