local AppBase = class("AppBase")

function AppBase:ctor(appName, packageRoot)

    self.name = appName
    self.packageRoot = packageRoot or "game"
    self.snapshots_ = {}

    -- set global app
    app = self

end

return AppBase
