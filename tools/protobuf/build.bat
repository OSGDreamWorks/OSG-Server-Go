@echo off

echo. 
echo. 
echo =============================================================
echo start build golang Protobuf

mkdir protos\golang\protobuf

protoc --plugin=protoc-gen-go="./protoc-gen-go.exe" --proto_path=./protos --go_out=./protos/golang/protobuf ./protos/PB_PacketCommon.proto
protoc --plugin=protoc-gen-go="./protoc-gen-go.exe" --proto_path=./protos --go_out=./protos/golang/protobuf ./protos/PB_PacketDefine.proto
protoc --plugin=protoc-gen-go="./protoc-gen-go.exe" --proto_path=./protos --go_out=./protos/golang/protobuf ./protos/PB_PacketServerDefine.proto
protoc --plugin=protoc-gen-go="./protoc-gen-go.exe" --proto_path=./protos --go_out=./protos/golang/protobuf ./protos/XShare_Logic.proto
protoc --plugin=protoc-gen-go="./protoc-gen-go.exe" --proto_path=./protos --go_out=./protos/golang/protobuf ./protos/XShare_Server.proto
protoc --plugin=protoc-gen-go="./protoc-gen-go.exe" --proto_path=./protos --go_out=./protos/golang/protobuf ./protos/CLPacket.proto
protoc --plugin=protoc-gen-go="./protoc-gen-go.exe" --proto_path=./protos --go_out=./protos/golang/protobuf ./protos/LCPacket.proto
protoc --plugin=protoc-gen-go="./protoc-gen-go.exe" --proto_path=./protos --go_out=./protos/golang/protobuf ./protos/SCPacket.proto
protoc --plugin=protoc-gen-go="./protoc-gen-go.exe" --proto_path=./protos --go_out=./protos/golang/protobuf ./protos/CSPacket.proto
protoc --plugin=protoc-gen-go="./protoc-gen-go.exe" --proto_path=./protos --go_out=./protos/golang/protobuf ./protos/ALPacket.proto
protoc --plugin=protoc-gen-go="./protoc-gen-go.exe" --proto_path=./protos --go_out=./protos/golang/protobuf ./protos/LAPacket.proto
protoc --plugin=protoc-gen-go="./protoc-gen-go.exe" --proto_path=./protos --go_out=./protos/golang/protobuf ./protos/SLPacket.proto
protoc --plugin=protoc-gen-go="./protoc-gen-go.exe" --proto_path=./protos --go_out=./protos/golang/protobuf ./protos/LSPacket.proto
protoc --plugin=protoc-gen-go="./protoc-gen-go.exe" --proto_path=./protos --go_out=./protos/golang/protobuf ./protos/SFPacket.proto
protoc --plugin=protoc-gen-go="./protoc-gen-go.exe" --proto_path=./protos --go_out=./protos/golang/protobuf ./protos/FSPacket.proto

move protos\golang\protobuf\* ..\..\src\protobuf

echo OK
echo =============================================================
echo. 
echo. 

echo =============================================================
echo "start build CSharp Protobuf"

mkdir protos\csharp\protobuf

cd protos
for %%i in (*.proto) do (  
echo %%i
"..\protogen.exe" -i:%%i -o:csharp\protobuf\%%i.cs

)
cd ..\

::move protos\csharp\protobuf\* ..\..\client\Assets\OSGClient\Scripts\protobuf

echo OK
echo =============================================================
echo. 
echo. 

echo =============================================================
echo start build javascript Protobuf

mkdir protos\javascript\protobuf

if not exist node_modules\protobufjs goto installProtobuf
echo get Protobuf
goto doneProtobuf
:installProtobuf
npm install protobufjs
:doneProtobuf
node  "node_modules\protobufjs\bin\pbjs" .\protos\PB_PacketCommon.proto -path=.\protos -target=js > .\protos\javascript\protobuf\PB_PacketCommon.js
node  "node_modules\protobufjs\bin\pbjs" .\protos\PB_PacketDefine.proto -path=.\protos -target=js > .\protos\javascript\protobuf\PB_PacketDefine.js
node  "node_modules\protobufjs\bin\pbjs" .\protos\PB_PacketServerDefine.proto -path=.\protos -target=js > .\protos\javascript\protobuf\PB_PacketServerDefine.js
node  "node_modules\protobufjs\bin\pbjs" .\protos\XShare_Logic.proto -path=.\protos -target=js > .\protos\javascript\protobuf\XShare_Logic.js
node  "node_modules\protobufjs\bin\pbjs" .\protos\XShare_Server.proto -path=.\protos -target=js > .\protos\javascript\protobuf\XShare_Server.js
node  "node_modules\protobufjs\bin\pbjs" .\protos\CLPacket.proto -path=.\protos -target=js > .\protos\javascript\protobuf\CLPacket.js
node  "node_modules\protobufjs\bin\pbjs" .\protos\LCPacket.proto -path=.\protos -target=js > .\protos\javascript\protobuf\LCPacket.js
node  "node_modules\protobufjs\bin\pbjs" .\protos\SCPacket.proto -path=.\protos -target=js > .\protos\javascript\protobuf\SCPacket.js
node  "node_modules\protobufjs\bin\pbjs" .\protos\CSPacket.proto -path=.\protos -target=js > .\protos\javascript\protobuf\CSPacket.js
node  "node_modules\protobufjs\bin\pbjs" .\protos\ALPacket.proto -path=.\protos -target=js > .\protos\javascript\protobuf\ALPacket.js
node  "node_modules\protobufjs\bin\pbjs" .\protos\LAPacket.proto -path=.\protos -target=js > .\protos\javascript\protobuf\LAPacket.js
node  "node_modules\protobufjs\bin\pbjs" .\protos\SLPacket.proto -path=.\protos -target=js > .\protos\javascript\protobuf\SLPacket.js
node  "node_modules\protobufjs\bin\pbjs" .\protos\LAPacket.proto -path=.\protos -target=js > .\protos\javascript\protobuf\LAPacket.js
node  "node_modules\protobufjs\bin\pbjs" .\protos\SFPacket.proto -path=.\protos -target=js > .\protos\javascript\protobuf\SFPacket.js
node  "node_modules\protobufjs\bin\pbjs" .\protos\FSPacket.proto -path=.\protos -target=js > .\protos\javascript\protobuf\FSPacket.js

::move protos\javascript\protobuf\* ..\..\client\SRPG\Test\SRPG\src\protobuf\

echo OK
echo =============================================================
echo. 
echo. 

echo =============================================================
echo begin build lua Protobuf

mkdir protos\lua\protobuf

cd protos
for %%i in (*.proto) do (  
echo %%i
"..\protoc.exe" --plugin=protoc-gen-lua="..\lua_protobuf\protoc-gen-lua.bat" --lua_out=lua\protobuf %%i

)
cd ..\

move protos\lua\protobuf\* ..\..\script\protobuf

echo OK
echo =============================================================
echo. 
echo. 