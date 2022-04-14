package handle_run_request

import (
	. "fmt"
	"math/rand"
	"time"
)

func randomGenerateInfo() (runtime, runstep string) {
	rand.Seed(time.Now().UnixNano())
	var runTime = 640 + rand.Intn(390)
	//var runDistance = 2400 + rand.Intn(6)
	var runStep = 1024 + rand.Intn(512)
	return Sprintf("%d", runTime), Sprintf("%d", runStep)
}
