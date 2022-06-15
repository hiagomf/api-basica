package main

import (
	"log"
	"net/http"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

func main() {

	group := errgroup.Group{}

	// define a porta que o servidor deve ficar escutando com suas respectivas rotas
	group.Go(
		func() error {
			return endless.ListenAndServe(":8081", externalAuthRouter())
		})

	// define que as go routinas não devem encerrar
	if err := group.Wait(); err != nil {
		log.Println("error on init")
	}

}

// externalAuthRouter - responsável por iniciar as rotas através do gin
func externalAuthRouter() http.Handler {
	ginInstance := gin.New()
	return ginInstance
}
