package initialize

import (
	"fmt"
	"net"
	"stellar/common"
	"stellar/pkg/utils"
	"strings"
)

// 生成客户端 ID，格式为：projectName-ip-port-randomString
func ClientId() {
	var ip net.IP
	addrs, _ := net.InterfaceAddrs()
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip = ipnet.IP
				break
			}
		}
	}

	var ipstr string
	if ip == nil {
		fmt.Println("No valid IP address found")
		ipstr = "127.0.0.1"
	} else {
		ipstr = ip.String()
	}

	clientId := fmt.Sprintf("%s-%s-%s-%s", strings.ToUpper(common.SystemProjectName), ipstr, common.Config.System.Port, utils.RandString(8, common.Numbers+common.UppercaseLetters))
	common.ClientId = &clientId
}
