package exec

import (
	"github.com/influxdata/influxdb/models"
	"github.com/influxdata/influxdb/tsdb/engine/tsm1"
	"github.com/influxdata/platform/pkg/data/gen"
)

type seriesGenerator struct {
	name  []byte
	tags  gen.TagsSequence
	field string
	vg    gen.ValuesSequence
	buf   []byte
}

func NewSeriesGenerator(name []byte, field string, vg gen.ValuesSequence, tags gen.TagsSequence) gen.SeriesGenerator {
	return &seriesGenerator{
		name:  name,
		field: field,
		vg:    vg,
		tags:  tags,
	}
}

func (g *seriesGenerator) Next() bool {
	if !g.tags.Next() {
		return false
	}

	g.vg.Reset()
	g.buf = models.AppendMakeKey(g.buf[:0], g.name, g.tags.Value())

	return true
}

func (g *seriesGenerator) Key() []byte {
	return tsm1.SeriesFieldKeyBytes(string(g.buf), g.field)
}

func (g *seriesGenerator) ValuesGenerator() gen.ValuesSequence { return g.vg }
