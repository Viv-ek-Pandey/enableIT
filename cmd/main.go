package main

import (
	"enableit/api"
	"flag"
)

func rawParam(args *string) map[string]string {
	m := map[string]string{}
	if args != nil && *args != "" {
		m["raw_params"] = *args
	}
	return m
}

func main() {
	//parse arguments
	args := flag.String("p", "", "raw params to task")
	flag.Parse()

	params := rawParam(args)

	api.Process(params)

}
