package catalog

import (
	"time"
)

type Flags interface {
	ToIndexes() ([]Index, error)
}

type Index interface {
	GetById(int64) (Record, error)
}

type Probe interface {
	GetById(int64) (ProbeResults, error)
}

type ProbeResults interface {
	RecordSet() RecordSet
	Timings() time.Duration
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
