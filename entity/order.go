package entity

import (
	"github.com/aaydin-tr/e-commerce/valueobject"
	"github.com/google/uuid"
)

type Order struct {
	ID          uuid.UUID
	ProductCode valueobject.Code
	Quantity    valueobject.Quantity
}
