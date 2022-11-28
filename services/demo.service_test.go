package services_test

import (
	"testing"

	"github.com/chencc1988/cmd-template/models"
	"github.com/chencc1988/cmd-template/services"
)

func TestPrintConfiguration(t *testing.T) {

	models.Configuration.Option1 = 100
	models.Configuration.Option2 = "1"
	models.Configuration.Option3 = "2"
	models.Configuration.Option4 = "3"
	output := services.Demo.GetConfiguration()
	if output != "Option1 :100, Option2 :1, Option3 :2, Option4 :3" {
		t.Error("Configuration verification failed")
	}
}
