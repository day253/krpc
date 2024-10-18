package main

import (
	re "github.com/ishumei/krpc/protocols/text/kitex_gen/shumei/strategy/re/textpredictor"
	"log"
)

func main() {
	svr := re.NewServer(new(TextPredictorImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
