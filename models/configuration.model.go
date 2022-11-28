package models

import (
	"fmt"
)

type ConfigurationStruct struct {
	Option1 uint
	Option2 string
	Option3 string
	Option4 string
}

func (ConfigurationStruct) String() string {
	cfg := fmt.Sprintf(
		"Option1 :%v, Option2 :%v, Option3 :%v, Option4 :%v",
		Configuration.Option1,
		Configuration.Option2,
		Configuration.Option3,
		Configuration.Option4,
	)
	return cfg
}
