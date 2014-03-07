
package network

import (
                "fmt"
                "net"
                "strings"
)

func GetMyIP() string {
        allIPs, err := net.InterfaceAddrs()
        if err != nil {
                fmt.Println("network.GetMyIP()--> Error receiving IPs. IP set to localhost. Consider setting IP manually")
                return "localhost"
        }
        
        IPString := make([]string, len(allIPs))
        for i := range allIPs {
                temp := allIPs[i].String()
                ip := strings.Split(temp, "/")
                IPString[i] = ip[0]
        }
        var myIP string
        for i:=range IPString {
                if IPString[i][0:3] == "129" {
                        myIP = IPString[i]
                }
        }
        return myIP
}
