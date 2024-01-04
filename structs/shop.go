package structs

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username string             `json:"username"`
	Email    string             `json:"email"`
	Password string             `json:"password"`
}

type Order struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID      primitive.ObjectID `json:"user_id,omitempty" bson:"user_id,omitempty"`
	TotalAmount float64            `json:"total_amount"`
	Status      string             `json:"status"`
}

type Product struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name"`
	Price       float64            `json:"price"`
	Description string             `json:"description"`
}

type Payment struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	OrderID primitive.ObjectID `json:"order_id,omitempty" bson:"order_id,omitempty"`
	Amount  float64            `json:"amount"`
	Status  string             `json:"status"`
}
