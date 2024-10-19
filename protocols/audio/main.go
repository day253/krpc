package main

import (
	re "github.com/ishumei/krpc/protocols/audio/kitex_gen/shumei/strategy/re/audiopredictor"
	"log"
)

func main() {
	svr := re.NewServer(new(AudioPredictorImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
