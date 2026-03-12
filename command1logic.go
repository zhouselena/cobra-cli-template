package cobraclitemplate

import (
	"fmt"
	"log"
)

// SampleStruct description.
type SampleStruct struct {
	Name	string
	SampleId int64
	Toggle bool
	ExtraInfo any
	Args map[string]any
}

// Command1 will print information.
func Command1(info1 string, info2 string) (string, error) {

	// validate input
	if (info1 == "") {
		return "", fmt.Errorf("invalid configuration")
	}

	// logic
	str := "Hello " + info1
	if (info2 != "") {
		str += " " + info2
	}

	log.Printf("%s", str)

	return str, nil

}