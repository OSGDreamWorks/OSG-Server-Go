package protobuf

const (
	Ok      = 0
	NoExist = 404
)

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
