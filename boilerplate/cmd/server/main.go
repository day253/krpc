package main

import (
	"flag"

	"github.com/ishumei/krpc/boilerplate/pkg/handler"
	"github.com/ishumei/krpc/kserver"
	"github.com/ishumei/krpc/kserver/arbiter"
	"github.com/samber/do"
)

func main() {
	flag.Parse()
	handler.BackgroundTask()
	injector := kserver.Injector
	arbiterService := do.MustInvoke[*arbiter.ArbiterService](injector)
	defer func() { _ = arbiterService.Shutdown() }()
	defer func() { _ = injector.Shutdown() }()
	arbiterService.Start()
}
