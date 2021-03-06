package charts

import (
	"io"

	"github.com/chenjiandongx/go-echarts/common"
	"github.com/chenjiandongx/go-echarts/datasets"
)

type Map struct {
	BaseOpts
	Series

	mapType string
}

func (Map) chartType() string { return common.ChartType.Map }

func NewMap(mapType string, routers ...RouterOpts) *Map {
	chart := new(Map)
	chart.mapType = mapType
	chart.initBaseOpts(false, routers...)
	chart.JSAssets.Add("maps/" + datasets.MapFileNames[mapType] + ".js")
	return chart
}

func (c *Map) Add(name string, data map[string]float32, options ...seriesOptser) *Map {
	nvs := make([]common.NameValueItem, 0)
	for k, v := range data {
		nvs = append(nvs, common.NameValueItem{Name: k, Value: v})
	}
	series := singleSeries{Name: name, Type: common.ChartType.Map, MapType: c.mapType, Data: nvs}
	series.setSingleSeriesOpts(options...)
	c.Series = append(c.Series, series)
	c.setColor(options...)
	return c
}

func (c *Map) SetGlobalOptions(options ...globalOptser) *Map {
	c.BaseOpts.setBaseGlobalOptions(options...)
	return c
}

func (c *Map) validateOpts() {
	c.validateAssets(c.AssetsHost)
}

func (c *Map) Render(w ...io.Writer) error {
	c.insertSeriesColors(c.appendColor)
	c.validateOpts()
	return renderToWriter(c, "chart", []string{}, w...)
}
