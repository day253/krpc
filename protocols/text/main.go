package main

import (
	"log"

	re "github.com/ishumei/krpc/protocols/text/kitex_gen/shumei/strategy/re/textpredictor"
)

func main() {
	svr := re.NewServer(new(TextPredictorImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
