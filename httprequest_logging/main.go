package main

import (
	"fmt"
	"time"

	"github.com/valyala/fasthttp"
)

var ch chan RequestInfo

func main() {
	ch = make(chan RequestInfo)

	go logDataJob()
	if err := fasthttp.ListenAndServe(":8080", handler); err != nil {
		fmt.Println("Error in ListenAndService")
	}
}

func handler(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json; charset=utf8")
	ctx.SetBodyString("{\"example\":\"ok\"}")

	ch <- RequestInfo{
		Method:   "GET",
		Body:     "",
		Response: "{\"example\":\"ok\"}",
	}
}

func logDataJob() {
	for {
		data := <-ch
		toPrint := "Request Method =" + data.Method + "; Body ="
		toPrint = toPrint + data.Body
		toPrint = toPrint + "; Response =" + data.Response

		time.Sleep(10 * time.Second)
		fmt.Println(toPrint)
		go func() {

		}()
	}
}

type RequestInfo struct {
	Method   string
	Body     string
	Response string
}
