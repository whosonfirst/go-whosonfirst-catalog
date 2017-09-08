package probe

import (
	"github.com/whosonfirst/go-whosonfirst-catalog"
)

type Probe struct {
	indexes []catalog.Index
}

type RecordSet struct {
	Records []catalog.Record
}

func NewProbe(indexes ...catalog.Index) (*Probe, error) {

	p := Probe{
		indexes: indexes,
	}

	return &p, nil
}

func (p *Probe) GetById(id int64) (ResultSet, error) {

	records := make([]catalog.Record, 0)
	pending := len(p.indexes)

	err_ch := make(chan error)
	done_ch := make(chan bool)
	record_ch := make(chan catalog.Record)

	for _, idx := range p.indexes {

		go func(idx catalog.Index, id int64) {

			defer func() {
				done_ch <- bool
			}()

			r, err := idx.GetById(id)

			if err != nil {
				err_ch <- err
				return
			}

			record_ch <- r

		}(idx, id)

	}

	for pending > 0 {

		select {
		case err := <-err_ch:
			return nil, err
		case r := <-record_ch:
			records = append(records, r)
		case <-done_ch:
			pending -= 1
		}
	}

	rs := RecordSet{
		Records: records,
	}

	return *rs, nil
}
