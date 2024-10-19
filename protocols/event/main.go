package main

import (
	"log"

	re "github.com/ishumei/krpc/protocols/event/kitex_gen/shumei/strategy/re/eventpredictor"
)

func main() {
	svr := re.NewServer(new(EventPredictorImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
