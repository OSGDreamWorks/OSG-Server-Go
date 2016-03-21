@echo off

setlocal

if exist install-require.bat goto ok
echo install-require.bat must be run from its folder
goto end

:ok

set PRODIR=%~dp0..\
set PKGDIR=%~dp0..\3rdpkg\

set OLDGOPATH=%GOPATH%

if not exist %PKGDIR% mkdir %PKGDIR%

set GOPATH=%PKGDIR%
cd %PKGDIR%
go get github.com/googollee/go-socket.io
go get github.com/garyburd/redigo/redis
go get github.com/go-sql-driver/mysql
go get github.com/yuin/gopher-lua
go get -u github.com/golang/protobuf
go get github.com/golang/snappy

::go build -o protoc-gen-go.exe github.com/golang/protobuf\protoc-gen-go

cd %PRODIR%

set GOPATH=%OLDGOPATH%

cd %PRODIR%tools\


:end
echo "install successfully"