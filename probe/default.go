package probe

import (
	"github.com/whosonfirst/go-whosonfirst-inspector"
	_ "log"
	"sort"
	"time"
)

type DefaultProbe struct {
	catalog.Probe
	indexes []catalog.Index
}

type DefaultProbeResults struct {
	catalog.ProbeResults `json:",omitempty"`
	ProbeRecordSet       catalog.RecordSet `json:"recordset"`
	ProbeTiming          time.Duration     `json:"timings"`
}

func (r *DefaultProbeResults) RecordSet() catalog.RecordSet {
	return r.ProbeRecordSet
}

func (r *DefaultProbeResults) Timing() time.Duration {
	return r.ProbeTiming
}

type DefaultRecordSet struct {
	catalog.RecordSet `json:",omitempty"`
	DefaultRecords    []catalog.Record `json:"records"`
}

func (r *DefaultRecordSet) Records() []catalog.Record {
	return r.DefaultRecords
}

func NewDefaultProbe(indexes ...catalog.Index) (catalog.Probe, error) {

	p := DefaultProbe{
		indexes: indexes,
	}

	return &p, nil
}

func (p *DefaultProbe) GetById(id int64) (catalog.ProbeResults, error) {

	records := make([]catalog.Record, 0)
	pending := len(p.indexes)

	err_ch := make(chan error)
	done_ch := make(chan bool)
	record_ch := make(chan catalog.Record)

	t1 := time.Now()

	for _, idx := range p.indexes {

		go func(idx catalog.Index, id int64) {

			defer func() {
				done_ch <- true
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

	lookup := make(map[int64][]catalog.Record)
	timestamps := make([]int64, 0)

	for _, r := range records {

		lastmod := r.LastModified()

		lastmod_records, ok := lookup[lastmod]

		if !ok {
			lastmod_records = make([]catalog.Record, 0)
			timestamps = append(timestamps, lastmod)
		}

		lastmod_records = append(lastmod_records, r)
		lookup[lastmod] = lastmod_records
	}

	sort.Slice(timestamps, func(i, j int) bool { return i > j })

	sorted_records := make([]catalog.Record, 0)

	for _, ts := range timestamps {

		for _, r := range lookup[ts] {
			sorted_records = append(sorted_records, r)
		}
	}

	t2 := time.Since(t1)

	rs := DefaultRecordSet{
		DefaultRecords: sorted_records,
	}

	pr := DefaultProbeResults{
		ProbeRecordSet: &rs,
		ProbeTiming:    t2,
	}

	return &pr, nil
}
