// package main

// import (
// 	"context"
// 	"log"

// 	"github.com/gofiber/fiber/v2"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// type User struct {
// 	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
// 	Username string             `json:"username"`
// 	Email    string             `json:"email"`
// 	Password string             `json:"password"`
// }

// type Order struct {
// 	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
// 	UserID      primitive.ObjectID `json:"user_id,omitempty" bson:"user_id,omitempty"`
// 	TotalAmount float64            `json:"total_amount"`
// 	Status      string             `json:"status"`
// }

// type Product struct {
// 	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
// 	Name        string             `json:"name"`
// 	Price       float64            `json:"price"`
// 	Description string             `json:"description"`
// }

// type Payment struct {
// 	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
// 	OrderID primitive.ObjectID `json:"order_id,omitempty" bson:"order_id,omitempty"`
// 	Amount  float64            `json:"amount"`
// 	Status  string             `json:"status"`
// }

// func main() {
// 	app := fiber.New()

// 	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://pixelcoin121:F0E8SWSiE5yhtAf0@ac-ls9abwa-shard-00-00.jc4c8s7.mongodb.net:27017,ac-ls9abwa-shard-00-01.jc4c8s7.mongodb.net:27017,ac-ls9abwa-shard-00-02.jc4c8s7.mongodb.net:27017/?replicaSet=atlas-o4u1d1-shard-0&ssl=true&authSource=admin"))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer client.Disconnect(context.Background())

// 	db := client.Database("test0")
// 	usersCollection := db.Collection("users")
// 	ordersCollection := db.Collection("orders")
// 	productsCollection := db.Collection("products")
// 	paymentsCollection := db.Collection("payments")

// 	app.Post("/api/users", func(c *fiber.Ctx) error {
// 		var user User
// 		if err := c.BodyParser(&user); err != nil {
// 			return err
// 		}
// 		user.ID = primitive.NewObjectID()
// 		_, err := usersCollection.InsertOne(context.Background(), user)
// 		if err != nil {
// 			return err
// 		}
// 		return c.JSON(user)
// 	})

// 	app.Get("/api/users", func(c *fiber.Ctx) error {
// 		var users []User
// 		cursor, err := usersCollection.Find(context.Background(), bson.M{})
// 		if err != nil {
// 			return err
// 		}
// 		defer cursor.Close(context.Background())
// 		if err := cursor.All(context.Background(), &users); err != nil {
// 			return err
// 		}
// 		return c.JSON(users)
// 	})

// 	app.Put("/api/users/:id", func(c *fiber.Ctx) error {
// 		id := c.Params("id")
// 		userID, err := primitive.ObjectIDFromHex(id)
// 		if err != nil {
// 			return err
// 		}
// 		var updatedUser User
// 		if err := c.BodyParser(&updatedUser); err != nil {
// 			return err
// 		}
// 		update := bson.M{
// 			"$set": updatedUser,
// 		}
// 		_, err = usersCollection.UpdateOne(context.Background(), bson.M{"_id": userID}, update)
// 		if err != nil {
// 			return err
// 		}
// 		updatedUser.ID = userID
// 		return c.JSON(updatedUser)
// 	})

// 	app.Delete("/api/users/:id", func(c *fiber.Ctx) error {
// 		id := c.Params("id")
// 		userID, err := primitive.ObjectIDFromHex(id)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = usersCollection.DeleteOne(context.Background(), bson.M{"_id": userID})
// 		if err != nil {
// 			return err
// 		}
// 		return c.SendString("User deleted successfully")
// 	})

// 	// Similar CRUD operations for Orders, Products, Payments...
// 	// คำสั่ง CRUD สำหรับ Orders
// 	app.Post("/api/orders", func(c *fiber.Ctx) error {
// 		var order Order
// 		if err := c.BodyParser(&order); err != nil {
// 			return err
// 		}
// 		order.ID = primitive.NewObjectID()
// 		_, err := ordersCollection.InsertOne(context.Background(), order)
// 		if err != nil {
// 			return err
// 		}
// 		return c.JSON(order)
// 	})

// 	app.Get("/api/orders", func(c *fiber.Ctx) error {
// 		var orders []Order
// 		cursor, err := ordersCollection.Find(context.Background(), bson.M{})
// 		if err != nil {
// 			return err
// 		}
// 		defer cursor.Close(context.Background())
// 		if err := cursor.All(context.Background(), &orders); err != nil {
// 			return err
// 		}
// 		return c.JSON(orders)
// 	})

// 	app.Put("/api/orders/:id", func(c *fiber.Ctx) error {
// 		id := c.Params("id")
// 		orderID, err := primitive.ObjectIDFromHex(id)
// 		if err != nil {
// 			return err
// 		}
// 		var updatedOrder Order
// 		if err := c.BodyParser(&updatedOrder); err != nil {
// 			return err
// 		}
// 		update := bson.M{
// 			"$set": updatedOrder,
// 		}
// 		_, err = ordersCollection.UpdateOne(context.Background(), bson.M{"_id": orderID}, update)
// 		if err != nil {
// 			return err
// 		}
// 		updatedOrder.ID = orderID
// 		return c.JSON(updatedOrder)
// 	})

// 	app.Delete("/api/orders/:id", func(c *fiber.Ctx) error {
// 		id := c.Params("id")
// 		orderID, err := primitive.ObjectIDFromHex(id)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = ordersCollection.DeleteOne(context.Background(), bson.M{"_id": orderID})
// 		if err != nil {
// 			return err
// 		}
// 		return c.SendString("Order deleted successfully")
// 	})

// 	// คำสั่ง CRUD สำหรับ Products
// 	app.Post("/api/products", func(c *fiber.Ctx) error {
// 		var product Product
// 		if err := c.BodyParser(&product); err != nil {
// 			return err
// 		}
// 		product.ID = primitive.NewObjectID()
// 		_, err := productsCollection.InsertOne(context.Background(), product)
// 		if err != nil {
// 			return err
// 		}
// 		return c.JSON(product)
// 	})

// 	app.Get("/api/products", func(c *fiber.Ctx) error {
// 		var products []Product
// 		cursor, err := productsCollection.Find(context.Background(), bson.M{})
// 		if err != nil {
// 			return err
// 		}
// 		defer cursor.Close(context.Background())
// 		if err := cursor.All(context.Background(), &products); err != nil {
// 			return err
// 		}
// 		return c.JSON(products)
// 	})

// 	app.Put("/api/products/:id", func(c *fiber.Ctx) error {
// 		id := c.Params("id")
// 		productID, err := primitive.ObjectIDFromHex(id)
// 		if err != nil {
// 			return err
// 		}
// 		var updatedProduct Product
// 		if err := c.BodyParser(&updatedProduct); err != nil {
// 			return err
// 		}
// 		update := bson.M{
// 			"$set": updatedProduct,
// 		}
// 		_, err = productsCollection.UpdateOne(context.Background(), bson.M{"_id": productID}, update)
// 		if err != nil {
// 			return err
// 		}
// 		updatedProduct.ID = productID
// 		return c.JSON(updatedProduct)
// 	})

// 	app.Delete("/api/products/:id", func(c *fiber.Ctx) error {
// 		id := c.Params("id")
// 		productID, err := primitive.ObjectIDFromHex(id)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = productsCollection.DeleteOne(context.Background(), bson.M{"_id": productID})
// 		if err != nil {
// 			return err
// 		}
// 		return c.SendString("Product deleted successfully")
// 	})

// 	// คำสั่ง CRUD สำหรับ Payments
// 	app.Post("/api/payments", func(c *fiber.Ctx) error {
// 		var payment Payment
// 		if err := c.BodyParser(&payment); err != nil {
// 			return err
// 		}
// 		payment.ID = primitive.NewObjectID()
// 		_, err := paymentsCollection.InsertOne(context.Background(), payment)
// 		if err != nil {
// 			return err
// 		}
// 		return c.JSON(payment)
// 	})

// 	app.Get("/api/payments", func(c *fiber.Ctx) error {
// 		var payments []Payment
// 		cursor, err := paymentsCollection.Find(context.Background(), bson.M{})
// 		if err != nil {
// 			return err
// 		}
// 		defer cursor.Close(context.Background())
// 		if err := cursor.All(context.Background(), &payments); err != nil {
// 			return err
// 		}
// 		return c.JSON(payments)
// 	})

// 	app.Put("/api/payments/:id", func(c *fiber.Ctx) error {
// 		id := c.Params("id")
// 		paymentID, err := primitive.ObjectIDFromHex(id)
// 		if err != nil {
// 			return err
// 		}
// 		var updatedPayment Payment
// 		if err := c.BodyParser(&updatedPayment); err != nil {
// 			return err
// 		}
// 		update := bson.M{
// 			"$set": updatedPayment,
// 		}
// 		_, err = paymentsCollection.UpdateOne(context.Background(), bson.M{"_id": paymentID}, update)
// 		if err != nil {
// 			return err
// 		}
// 		updatedPayment.ID = paymentID
// 		return c.JSON(updatedPayment)
// 	})

// 	app.Delete("/api/payments/:id", func(c *fiber.Ctx) error {
// 		id := c.Params("id")
// 		paymentID, err := primitive.ObjectIDFromHex(id)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = paymentsCollection.DeleteOne(context.Background(), bson.M{"_id": paymentID})
// 		if err != nil {
// 			return err
// 		}
// 		return c.SendString("Payment deleted successfully")
// 	})

// 	log.Fatal(app.Listen(":3000"))

// }
