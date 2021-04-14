package main

import (
	"fmt"
	"io/ioutil"
	"plugin"
	"strings"
	"time"
)

func main() {
	//for {
	f, err := ioutil.ReadFile("plugin.conf")
	if err != nil {
		fmt.Println("open error", err)
		return
	} else if len(f) == 0 {
		fmt.Println("0 read")
		time.Sleep(time.Second * 5)
	} else {
		pluginName := strings.Trim(string(f), "\n")
		fmt.Println("plugin name", pluginName, len(pluginName)) // pluginName == "p.so"
		p, err := plugin.Open(pluginName)
		if err != nil {
			fmt.Println("import plugin error", err)
		}
		f, err := p.Lookup("PrintHello")
		if err != nil {
			fmt.Println("lookup func error", err)
		}
		f.(func())()
	}
	//}
}
