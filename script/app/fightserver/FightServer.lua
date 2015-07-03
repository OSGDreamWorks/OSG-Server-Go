
local logger = import("logger")
local mvc = import("mvc")

local FightServer = class("FightServer", mvc.AppBase)

-- ��������
FightServer["battles"]       = {}          -- ���ս������

function FightServer:ctor(appName)
    FightServer.super.ctor(self, appName)
end

function FightServer:CreateServices(cfg, n)

    self.rpcServer = RpcServer.new()
    self.rpcServer:Register(self, self.class)

    self.rpcServer:ListenAndServe(cfg.FsHost[n])

end

function FightServer:SF_StartBattle(req, rst)
    logger.Debug(req)
    return "rst"
end

return FightServer