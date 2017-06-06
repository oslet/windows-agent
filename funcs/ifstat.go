package funcs

import (
	"log"
	"strings"
	//	"github.com/freedomkk-qfeng/windows-agent/g"
	"github.com/open-falcon/common/model"
	"github.com/shirou/gopsutil/net"
)

func net_status(ifacePrefix []string) ([]net.IOCountersStat, error) {
	net_iocounter, err := net.IOCounters(true)
	netIfs := []net.IOCountersStat{}
	for _, iface := range ifacePrefix {
		for _, netIf := range net_iocounter {
			if strings.Contains(netIf.Name, iface) {
				netIfs = append(netIfs, netIf)
			}
		}
	}
	return netIfs, err
}

func NetMetrics() []*model.MetricValue {
	return CoreNetMetrics()
}

func CoreNetMetrics() []*model.MetricValue {

	netIfs, err := NetIOCounters(true)
	if err != nil {
		log.Println("Get netInfo fail: ", err)
		return []*model.MetricValue{}
	}

	cnt := len(netIfs)
	ret := make([]*model.MetricValue, cnt*8)

	for idx, netIf := range netIfs {
		iface := "iface=" + netIf.Name
		ret[idx*8+0] = CounterValue("net.if.in.bytes", netIf.BytesRecv, iface)
		ret[idx*8+1] = CounterValue("net.if.in.packets", netIf.PacketsRecv, iface)
		ret[idx*8+2] = CounterValue("net.if.in.errors", netIf.Errin, iface)
		ret[idx*8+3] = CounterValue("net.if.in.dropped", netIf.Dropin, iface)
		ret[idx*8+4] = CounterValue("net.if.out.bytes", netIf.BytesSent, iface)
		ret[idx*8+5] = CounterValue("net.if.out.packets", netIf.PacketsSent, iface)
		ret[idx*8+6] = CounterValue("net.if.out.errors", netIf.Errout, iface)
		ret[idx*8+7] = CounterValue("net.if.out.dropped", netIf.Dropout, iface)

	}
	return ret
}
