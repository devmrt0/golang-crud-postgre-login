package api

import (
	"cloudgobackend/api/router"
	"cloudgobackend/auto"
	"fmt"

	"cloudgobackend/config"
	"log"
	"net/http"
)

func Run() {
	config.Load()
	auto.Load()
	fmt.Printf("running... at port %d", config.PORT)
	listen(config.PORT)

}

func listen(port int) {
	r := router.New()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router.LoadCORS(r)))
}
