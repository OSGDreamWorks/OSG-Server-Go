@echo off

setlocal

if exist debug.bat goto ok
echo debug.bat must be run from its folder
goto end

:ok

set PRODIR=%~dp0..\..\
set PKGDIR=%~dp0..\..\3rdpkg\

set OLDGOPATH=%GOPATH%

if not exist %PKGDIR% goto req

cd %PRODIR%

set GOPATH=%PRODIR%;%PKGDIR%

start go run %PRODIR%\src\runtime\dbserver\main.go &

Sleep 3
start go run %PRODIR%\src\runtime\authserver\main.go &

Sleep 3
start go run %PRODIR%\src\runtime\loginserver\main.go &

Sleep 3
::start go run %PRODIR%\src\runtime\fightserver\main.go &

Sleep 3
start go run %PRODIR%\src\runtime\gameserver\main.go &

set GOPATH=%OLDGOPATH%

cd %PRODIR%tools\build\

goto end

:req
echo "please run install-require.bat first."

:end
echo "debug successfully"