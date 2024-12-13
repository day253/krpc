package main

import (
	re "github.com/day253/krpc/protocols/image/kitex_gen/shumei/strategy/re/imagepredictor"
	"log"
)

func main() {
	svr := re.NewServer(new(ImagePredictorImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
