package autolimit

import (
	"log"

	"github.com/KimMachineGun/automemlimit/memlimit"
	"go.uber.org/automaxprocs/maxprocs"
)

func init() {
	_, _ = memlimit.SetGoMemLimit(0.9)
}

func init() {
	_, _ = maxprocs.Set(maxprocs.Logger(log.Printf))
}
