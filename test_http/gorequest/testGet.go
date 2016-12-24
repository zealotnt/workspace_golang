package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/parnurzeal/gorequest"
)

var PR_INFO = fmt.Println
var PR_DUMP = spew.Dump
var STR_SEPERATOR = "**********************************************"

func main() {
	request := gorequest.New()
	resp, body, errs := request.Get("http://example.com/").End()
	if errs != nil {
		PR_DUMP(errs)
		panic(errs)
	}
	PR_DUMP(resp)
	PR_INFO(STR_SEPERATOR)
	PR_DUMP(body)
}
