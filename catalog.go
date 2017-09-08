package catalog

type Flags interface {
	ToIndexes() ([]Index, error)
}

type Index interface {
	GetById(int64) (Record, error)
}

type Record interface {
	Id() int64
	Source() string
	URI() string
	Body() interface{}
}
