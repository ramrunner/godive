package util

import (
	"encoding/json"
	"github.com/ramrunner/godive/planner"
	"os"
)

func ReadJSONConf(fname string) (dc planner.DiveContext, err error) {
	var (
		fd   *os.File
		jdec *json.Decoder
	)

	if fd, err = os.Open(fname); err != nil {
		return
	}
	jdec = json.NewDecoder(fd)
	err = jdec.Decode(&dc)
	return
}
