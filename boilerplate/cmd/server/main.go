package main

import (
	"flag"

	_ "github.com/ishumei/krpc/autolimit"
	"github.com/ishumei/krpc/boilerplate/internal/handler"
	"github.com/ishumei/krpc/kframe/arbiter"
	"github.com/ishumei/krpc/kframe/sconfig"
	"github.com/samber/do"
)

func main() {
	flag.Parse()
	handler.BackgroundTask()
	injector := sconfig.Injector
	arbiterService := do.MustInvoke[*arbiter.ArbiterService](injector)
	defer func() { _ = arbiterService.Shutdown() }()
	defer func() { _ = injector.Shutdown() }()
	arbiterService.Start()
}
