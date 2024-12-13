package main

import (
	"log"

	re "github.com/day253/krpc/protocols/image/kitex_gen/shumei/strategy/re/imagepredictor"
)

func main() {
	svr := re.NewServer(new(ImagePredictorImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
