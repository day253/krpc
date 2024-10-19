package grace

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/ishumei/krpc/kserver/sconfig"
	registry_zookeeper "github.com/ishumei/krpc/registry-zookeeper"
	"github.com/samber/do"
)

func DefaultUserExitSignal() <-chan error {
	errCh := make(chan error, 1)
	go func() {
		sig := UserExitSignal()
		defer signal.Stop(sig)
		<-sig
		errCh <- nil
	}()
	return errCh
}

func DefaultDeregisterSignal() {
	go func() {
		sig := DeregisterSignal()
		defer signal.Stop(sig)
		for range sig {
			err := do.MustInvoke[*registry_zookeeper.ZookeeperRegistry](sconfig.Injector).Deregister(nil)
			klog.Info("deregister service", err)
		}
		<-sig
	}()
}

func UserExitSignal() chan os.Signal {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGHUP)
	return signals
}

func DeregisterSignal() chan os.Signal {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM, syscall.SIGUSR1)
	return signals
}
