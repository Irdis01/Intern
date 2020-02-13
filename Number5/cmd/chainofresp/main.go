package main

import (
	"log"

	"github.com/Irdis01/Intern/Number5/pkg/chainofresp/meet_behavior"
	"github.com/Irdis01/Intern/Number5/pkg/chainofresp/attack_behavior"
	"github.com/Irdis01/Intern/Number5/pkg/chainofresp/customer"
	"github.com/Irdis01/Intern/Number5/pkg/chainofresp/warehouse"
)

func main() {
	baseMapper := customer.NewUserBase()
	userRegistrator := attack_behavior.NewRegistrator(baseMapper)
	respService := warehouse.NewService()
	userAuth := meet_behavior.NewAuthentificator(baseMapper, userRegistrator, respService)
	msg := make([]byte, 0)
	err := userAuth.ConnectUser("1", "1", &msg)
	log.Println(err)
	log.Println(string(msg))
	err = userAuth.ConnectUser("1", "2", &msg)
	log.Println(err)
}
