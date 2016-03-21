@echo off

echo. 
echo. 
echo =============================================================
echo start build Protobuf

mkdir build
del /a /f /q /s ..\..\src\protobuf

cd proto3
for /R %%i in (*.proto) do (
echo. 
echo -----------copy----------
echo %%i
xcopy %%i ..\build
echo Done!
echo -------------------------
echo. 
)
cd ..

cd build
..\tools\protoc.exe --plugin=protoc-gen-go="..\tools\protoc-gen-go.exe" --proto_path=.\ -I=.\ --go_out=..\..\..\src\protobuf\ *.proto
cd ..

xcopy proto3\server\db_proto.go ..\..\src\protobuf\

rd /q /s build

echo Done!
echo =============================================================
echo. 

pause