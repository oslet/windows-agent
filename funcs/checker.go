package funcs

import (
	"fmt"
)

func CheckCollector() {

	output := make(map[string]bool)
	_, procStatErr := CPUTimes(false)

	output["df.bytes"] = len(DeviceMetrics()) > 0
	output["net.if  "] = len(CoreNetMetrics()) > 0
	output["loadavg "] = len(LoadMetrics()) > 0
	output["cpustat "] = procStatErr == nil
	output["disk.io "] = len(DiskIOMetrics()) > 0
	output["memory  "] = len(MemMetrics()) > 0
	output["tcpip   "] = len(TcpipMetrics()) > 0

	for k, v := range output {
		status := "fail"
		if v {
			status = "ok"
		}
		fmt.Println(k, "...", status)
	}
}
