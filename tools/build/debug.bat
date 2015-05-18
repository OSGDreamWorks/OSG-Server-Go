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
go get code.google.com/p/goprotobuf/proto
go get code.google.com/p/snappy-go/snappy
cd %PRODIR%

set GOPATH=%PRODIR%;%PKGDIR%;

start go run %PRODIR%\src\runtime\dbserver\main.go &

Sleep 3
start go run %PRODIR%\src\runtime\authserver\main.go &

Sleep 3
start go run %PRODIR%\src\runtime\gateserver\main.go &

Sleep 3
start go run %PRODIR%\src\runtime\fightserver\main.go &

Sleep 3
start go run %PRODIR%\src\runtime\gameserver\main.go &

set GOPATH=%OLDGOPATH%

cd %PRODIR%tools\build\

:end
echo "debug successfully"