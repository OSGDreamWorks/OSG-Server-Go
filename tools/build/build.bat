@echo off

setlocal

if exist build.bat goto ok
echo build.bat must be run from its folder
goto end

:ok

set PRODIR=%~dp0..\..\
set PKGDIR=%~dp0..\..\3rdpkg\

set OLDGOPATH=%GOPATH%

if not exist %PKGDIR% goto req

cd %PRODIR%

set GOPATH=%PRODIR%;%PKGDIR%

go build -o %PRODIR%\bin\dbserver.exe		%PRODIR%\src\runtime\dbserver\main.go

go build -o %PRODIR%\bin\authserver.exe		%PRODIR%\src\runtime\authserver\main.go

go build -o %PRODIR%\bin\loginserver.exe		%PRODIR%\src\runtime\loginserver\main.go

::go build -o %PRODIR%\bin\fightserver.exe		%PRODIR%\src\runtime\fightserver\main.go

go build -o %PRODIR%\bin\gameserver.exe		%PRODIR%\src\runtime\gameserver\main.go

set GOPATH=%OLDGOPATH%

cd %PRODIR%tools\build\

xcopy %PRODIR%etc %PRODIR%bin\etc /d/e/i

goto end

:req
echo "please run install-require.bat first."

:end
echo "build finish"