package autolimit

import (
	"github.com/KimMachineGun/automemlimit/memlimit"
	"github.com/cloudwego/kitex/pkg/klog"
	"go.uber.org/automaxprocs/maxprocs"
)

func init() {
	_, _ = memlimit.SetGoMemLimit(0.9)
}

func init() {
	_, _ = maxprocs.Set(maxprocs.Logger(klog.Infof))
}
