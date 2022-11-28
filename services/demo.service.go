package services

import (
	"fmt"

	"github.com/chencc1988/cmd-template/models"
)

type DemoStruct struct {
}

func (DemoStruct) GetConfiguration() string {
	return fmt.Sprint(models.Configuration)
}
