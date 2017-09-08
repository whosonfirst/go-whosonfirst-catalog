package catalog

type Flags interface {
	ToIndexes() ([]Index, error)
}

type Index interface {
	GetById(int64) (Record, error)
}

type Probe interface {
	GetById(int64) (RecordSet, error)
}

type Record interface {
	Id() int64
	Type() string
	Source() string
	URI() string
	Body() interface{}
	Hash() string
}

type RecordSet interface {
	Records() []Record
}
