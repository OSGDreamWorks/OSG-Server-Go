@echo off

echo. 
echo. 
echo =============================================================
echo start build golang Protobuf

mkdir protos\golang\protobuf

protoc --plugin=protoc-gen-go="./protoc-gen-go.exe" --proto_path=./protos --go_out=./protos/golang/protobuf ./protos/CLPacket.proto
protoc --plugin=protoc-gen-go="./protoc-gen-go.exe" --proto_path=./protos --go_out=./protos/golang/protobuf ./protos/CSPacket.proto
protoc --plugin=protoc-gen-go="./protoc-gen-go.exe" --proto_path=./protos --go_out=./protos/golang/protobuf ./protos/LAPacket.proto
protoc --plugin=protoc-gen-go="./protoc-gen-go.exe" --proto_path=./protos --go_out=./protos/golang/protobuf ./protos/LCPacket.proto
protoc --plugin=protoc-gen-go="./protoc-gen-go.exe" --proto_path=./protos --go_out=./protos/golang/protobuf ./protos/PB_PacketCommon.proto
protoc --plugin=protoc-gen-go="./protoc-gen-go.exe" --proto_path=./protos --go_out=./protos/golang/protobuf ./protos/PB_PacketDefine.proto
protoc --plugin=protoc-gen-go="./protoc-gen-go.exe" --proto_path=./protos --go_out=./protos/golang/protobuf ./protos/PB_PacketServerDefine.proto
protoc --plugin=protoc-gen-go="./protoc-gen-go.exe" --proto_path=./protos --go_out=./protos/golang/protobuf ./protos/SCPacket.proto
protoc --plugin=protoc-gen-go="./protoc-gen-go.exe" --proto_path=./protos --go_out=./protos/golang/protobuf ./protos/SLPacket.proto
protoc --plugin=protoc-gen-go="./protoc-gen-go.exe" --proto_path=./protos --go_out=./protos/golang/protobuf ./protos/XShare_Logic.proto
protoc --plugin=protoc-gen-go="./protoc-gen-go.exe" --proto_path=./protos --go_out=./protos/golang/protobuf ./protos/XShare_Server.proto

move protos\golang\protobuf\* ..\..\src\protobuf

echo OK
echo =============================================================
echo. 
echo. 

echo =============================================================
echo "start build CSharp Protobuf"

mkdir protos\csharp\protobuf

::protogen -i:msg.proto -o:Msg.cs

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
node  "node_modules\protobufjs\bin\pbjs" .\protos\CLPacket.proto -path=.\protos -target=js > .\protos\javascript\protobuf\CLPacket.js
node  "node_modules\protobufjs\bin\pbjs" .\protos\CSPacket.proto -path=.\protos -target=js > .\protos\javascript\protobuf\CSPacket.js
node  "node_modules\protobufjs\bin\pbjs" .\protos\LAPacket.proto -path=.\protos -target=js > .\protos\javascript\protobuf\LAPacket.js
node  "node_modules\protobufjs\bin\pbjs" .\protos\LCPacket.proto -path=.\protos -target=js > .\protos\javascript\protobuf\LCPacket.js
node  "node_modules\protobufjs\bin\pbjs" .\protos\PB_PacketCommon.proto -path=.\protos -target=js > .\protos\javascript\protobuf\PB_PacketCommon.js
node  "node_modules\protobufjs\bin\pbjs" .\protos\PB_PacketDefine.proto -path=.\protos -target=js > .\protos\javascript\protobuf\PB_PacketDefine.js
node  "node_modules\protobufjs\bin\pbjs" .\protos\PB_PacketServerDefine.proto -path=.\protos -target=js > .\protos\javascript\protobuf\PB_PacketServerDefine.js
node  "node_modules\protobufjs\bin\pbjs" .\protos\SCPacket.proto -path=.\protos -target=js > .\protos\javascript\protobuf\SCPacket.js
node  "node_modules\protobufjs\bin\pbjs" .\protos\SLPacket.proto -path=.\protos -target=js > .\protos\javascript\protobuf\SLPacket.js
node  "node_modules\protobufjs\bin\pbjs" .\protos\XShare_Logic.proto -path=.\protos -target=js > .\protos\javascript\protobuf\XShare_Logic.js
node  "node_modules\protobufjs\bin\pbjs" .\protos\XShare_Server.proto -path=.\protos -target=js > .\protos\javascript\protobuf\XShare_Server.js

::move msg.js ..\..\client\SRPG\Test\SRPG\src\msg.js
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

echo OK
echo =============================================================
echo. 
echo. 