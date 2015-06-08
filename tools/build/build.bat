@echo off

setlocal

if exist debug.bat goto ok
echo debug.bat must be run from its folder
goto end

:ok

set PRODIR=%~dp0..\..\
set PKGDIR=%~dp0..\..\3rdpkg\

set OLDGOPATH=%GOPATH%

if not exist %PKGDIR% mkdir %PKGDIR%

set GOPATH=%PKGDIR%
cd %PKGDIR%
go get github.com/googollee/go-socket.io
go get github.com/garyburd/redigo/redis
go get github.com/go-sql-driver/mysql
go get github.com/Shopify/go-lua
go get code.google.com/p/goprotobuf/proto
go get code.google.com/p/snappy-go/snappy
cd %PRODIR%

set GOPATH=%PRODIR%;%PKGDIR%;

go build -o %PRODIR%\bin\dbserver.exe		%PRODIR%\src\runtime\dbserver\main.go

go build -o %PRODIR%\bin\authserver.exe		%PRODIR%\src\runtime\authserver\main.go

go build -o %PRODIR%\bin\gateserver.exe		%PRODIR%\src\runtime\gateserver\main.go

go build -o %PRODIR%\bin\fightserver.exe		%PRODIR%\src\runtime\fightserver\main.go

go build -o %PRODIR%\bin\gameserver.exe		%PRODIR%\src\runtime\gameserver\main.go

set GOPATH=%OLDGOPATH%

cd %PRODIR%tools\build\

xcopy %PRODIR%etc %PRODIR%bin\etc /d/e/i

:end
echo "build successfully"