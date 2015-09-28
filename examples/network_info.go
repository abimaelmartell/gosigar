package main

import (
	"fmt"
	"github.com/abimaelmartell/gosigar"
)

func main() {
	concreteSigar := sigar.ConcreteSigar{}
	networkInfo, _ := concreteSigar.GetNetworkInfo()

	fmt.Println("DefaultGateway:", networkInfo.DefaultGateway)
	fmt.Println("DefaultGatewayInterface:", networkInfo.DefaultGatewayInterface)
	fmt.Println("HostName:", networkInfo.HostName)
	fmt.Println("DomainName:", networkInfo.DomainName)
	fmt.Println("PrimaryDns:", networkInfo.PrimaryDns)
	fmt.Println("SecondaryDns:", networkInfo.SecondaryDns)

}
