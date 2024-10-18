package main

import (
	"flag"

	"github.com/cloudwego/kitex/pkg/klog"
)

var (
	host         = flag.String("host", "127.0.0.1", "host")
	port         = flag.String("port", "8000", "port")
	clientType   = flag.String("clientType", "arbiter", "clientType")
	requestId    = flag.String("requestId", "", "requestId")
	organization = flag.String("organization", "", "organization")
	data         = flag.String("data", "{}", "data")
)

func main() {
	flag.Parse()
	flag.VisitAll(func(f *flag.Flag) {
		klog.Info(f.Name, ": ", f.Value)
	})
	switch *clientType {
	case "audio":
		audio()
	case "event":
		event()
	case "image":
		image()
	case "text":
		text()
	case "arbiter":
		fallthrough
	default:
		arbiter()
	}
}
