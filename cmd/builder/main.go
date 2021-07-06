package main

import (
	"builderTest/pkg/builder"
	"fmt"
)


func main() {
	builder := builder.New()
	server:=builder.PickClientsAndroid("1.1.1.1").PickClientsIOS("2.2.2.2").Build()
	//fmt.Println(server.ShowAndroidIP())
	//fmt.Println(server.ShowIOSIP())
}
