package protobuf

const (
	Ok      = 0
	NoExist = 404
)

type Network_Protocol uint32

const (
	DB_Protocol_eDB_PacketBegin Network_Protocol = 0x01000000
	// ----------------------------
	DB_Protocol_eDB_Query  Network_Protocol = 0x01000000
	DB_Protocol_eDB_Write  Network_Protocol = 0x01000001
	DB_Protocol_eDB_Delete Network_Protocol = 0x01000002
	// ----------------------------
	DB_Protocol_eDB_PacketEnd Network_Protocol = 0x01100000
)

var DB_Protocol_value = map[string]int32{
	"eDB_PacketBegin": 0x01000000,
	"eLC_eDB_Query":   0x01000000,
	"eLC_eDB_Write":   0x01000001,
	"eLC_eDB_Delete":  0x01000002,
	"eDB_PacketEnd":   0x01100000,
}

type DBQuery struct {
	Table string
	Key   string
}

type DBQueryResult struct {
	Code  uint32
	Value []byte
}

type DBDel struct {
	Table string
	Key   string
}

type DBDelResult struct {
	Code uint32
}

type DBWrite struct {
	Table string
	Key   string
	Value []byte
}

type DBWriteResult struct {
	Code uint32
}

//账号服务器转用
type AuthDbWrite struct {
	Table string
	Key   string
	Value string
}

type AuthDbWriteResult struct {
	Code uint32
}

type AuthDbDel struct {
	Table string
	Key   string
}

type AuthDbDelResult struct {
	Code uint32
}

type AuthDbQuery struct {
	Table string
	Key   string
}

type AuthDbQueryResult struct {
	Code  uint32
	Value string
}
