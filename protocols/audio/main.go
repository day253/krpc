package main

import (
	"log"

	re "github.com/ishumei/krpc/protocols/audio/kitex_gen/shumei/strategy/re/audiopredictor"
)

func main() {
	svr := re.NewServer(new(AudioPredictorImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
