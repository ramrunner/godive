package main

import (
	"flag"
	"fmt"
	//"github.com/ramrunner/godive/planner"
	"github.com/ramrunner/godive/util"
)

func main() {
	flag.Parse()
	if len(flag.Args()) != 1 {
		fmt.Printf("json config file required\n")
		return
	}
	dc, err := util.ReadJSONConf(flag.Arg(0))
	if err != nil {
		fmt.Printf("error in reading config:%s\n", err)
		return
	}
	fmt.Printf("got dive context:%+v\n", dc)
	if err := dc.Assert(); err != nil {
		fmt.Printf("error in dive context: %s\n", err)
	} else {
		fmt.Printf("dive context loaded sucessfully\n")
	}
	return
}
