package darknet

import (
	"fmt"
	"testing"
)

func TestReadSectionsFromCfg(t *testing.T) {
	sections, err := readCfg("./cmd/examples/yolov7-tiny.cfg")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, s := range sections {
		fmt.Println(s.Type)
		for _, o := range s.Options {
			fmt.Println("\t", o)
		}
	}
}
