@echo off

setlocal

if exist install-require.bat goto ok
echo install-require.bat must be run from its folder
exit 1
goto end

:ok

npm install string-format

:end
echo "install successfully"