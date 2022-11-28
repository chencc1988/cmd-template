package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/chencc1988/cmd-template/models"
	"github.com/chencc1988/cmd-template/services"

	"github.com/jedib0t/go-pretty/v6/table"
)

func main() {

	var (
		help    bool
		Option1 uint   = 10000
		Option2 string = "http://127.0.0.1"
		Option3 string = "demo"
		Option4 string = "product"
	)

	if v, exist := os.LookupEnv("Option1"); exist {
		if p, err := strconv.Atoi(v); err != nil {
			log.Fatal("Option1 should be a number")
		} else {
			Option1 = uint(p)
		}
	}
	if v, exist := os.LookupEnv("Option2"); exist {
		Option2 = v
	}
	if v, exist := os.LookupEnv("Option3"); exist {
		Option3 = v
	}
	if v, exist := os.LookupEnv("Option4"); exist {
		Option4 = v
	}

	flag.UintVar(&Option1, "o1", Option1, "Option1 is an unit input")
	flag.StringVar(&Option2, "o2", Option2, "Option1 is a string input")
	flag.StringVar(&Option3, "o3", Option3, "Option1 is a string input")
	flag.StringVar(&Option4, "o4", Option4, "Option1 is a string input")

	flag.BoolVar(&help, "h", false, "Help")

	flag.Parse()

	if help {
		flag.Usage()
		return
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Arg", "Env Var", "Value"})
	t.AppendRow([]interface{}{"-o1", "Option1", Option1})
	t.AppendRow([]interface{}{"-o2", "Option2", Option2})
	t.AppendRow([]interface{}{"-o3", "Option3", Option3})
	t.AppendRow([]interface{}{"-o4", "Option4", Option4})
	t.SetAllowedRowLength(70)
	t.Render()

	models.Configuration.Option1 = Option1
	models.Configuration.Option2 = Option2
	models.Configuration.Option3 = Option3
	models.Configuration.Option4 = Option4
	fmt.Print(services.Demo.GetConfiguration())
}
