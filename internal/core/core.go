package core

import (
	"fmt"

	"github.com/ale8k/hexagonal-architecture-in-go-example/internal/ports"
)

type postalService struct {
	persistence ports.PersistencePort
}

func NewMailService(persistence ports.PersistencePort) *postalService {
	return &postalService{
		persistence: persistence,
	}
}

func (c *postalService) ReceiveMail(mail any) {
	fmt.Println("Received mail!", mail)
	c.persistence.Save(mail)
}
