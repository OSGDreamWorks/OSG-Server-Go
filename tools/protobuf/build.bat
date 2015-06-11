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
::node  "node_modules\protobufjs\bin\pbjs" msg.proto -target=js > msg.js

::move msg.js ..\..\client\SRPG\Test\SRPG\src\msg.js
echo OK
echo =============================================================
echo. 
echo. 