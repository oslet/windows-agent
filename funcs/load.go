package funcs

import (
	"encoding/json"
	"log"

	"github.com/open-falcon/common/model"
	"github.com/oslet/agent/tools/load"
)

type LoadAvgStat struct {
	Load1  float64 `json:"load1"`
	Load5  float64 `json:"load5"`
	Load15 float64 `json:"load15"`
}

func LoadMetrics() (L []*model.MetricValue) {

	loadVal, err := load.LoadAvg()
	if err != nil {
		log.Println("Get load fail: ", err)
		return nil
	}

	L = append(L, CounterValue("load.load1", loadVal.Load1))
	L = append(L, CounterValue("load.load5", loadVal.Load5))
	L = append(L, CounterValue("load.load15", loadVal.Load15))

	return
}

func (l LoadAvgStat) String() string {
	s, _ := json.Marshal(l)
	return string(s)
}

func LoadAvg() (*LoadAvgStat, error) {
	ret := LoadAvgStat{}

	return &ret, nil
}
