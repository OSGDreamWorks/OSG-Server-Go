@echo off

echo. 
echo. 
echo =============================================================
echo start build golang Protobuf

protoc --plugin=protoc-gen-go="./protoc-gen-go.exe" --go_out=./ *.proto

move msg.pb.go ..\..\src\protobuf\msg.pb.go
echo OK
echo =============================================================
echo. 
echo. 

echo =============================================================
echo "start build CSharp Protobuf"
protogen -i:msg.proto -o:Msg.cs
echo OK
echo =============================================================
echo. 
echo. 

echo =============================================================
echo start build javascript Protobuf
if not exist node_modules\protobufjs goto installProtobuf
echo get Protobuf
goto doneProtobuf
:installProtobuf
npm install protobufjs
:doneProtobuf
node  "node_modules\protobufjs\bin\pbjs" msg.proto -target=js > msg.js

::move msg.js ..\..\client\SRPG\Test\SRPG\src\msg.js
echo OK
echo =============================================================
echo. 
echo. 