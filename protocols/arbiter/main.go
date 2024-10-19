package main

import (
	service "github.com/ishumei/krpc/protocols/arbiter/kitex_gen/com/shumei/service/predictor"
	"log"
)

func main() {
	svr := service.NewServer(new(PredictorImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
