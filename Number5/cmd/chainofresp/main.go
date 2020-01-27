package main

import (
	"log"

	"github.com/Irdis01/Intern/Number5/pkg/chainofresp/authentication"
	"github.com/Irdis01/Intern/Number5/pkg/chainofresp/registrator"
	"github.com/Irdis01/Intern/Number5/pkg/chainofresp/service"
	"github.com/Irdis01/Intern/Number5/pkg/chainofresp/userbase"
)

func main() {
	baseMapper := userbase.NewUserBase()
	userRegistrator := registrator.NewRegistrator(baseMapper)
	respService := service.NewService()
	userAuth := authentication.NewAuthentificator(baseMapper, userRegistrator, respService)
	msg := make([]byte, 0)
	userAuth.ConnectUser("1", "1", &msg)
	log.Println(string(msg))
}
