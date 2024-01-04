package controllers

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"backend-go/database"
	"backend-go/structs"
)

func CreateProduct(c *fiber.Ctx) error {
	_, db := database.ConnectDB()
	productsCollection := db.Collection("products")
	var product structs.Product
	if err := c.BodyParser(&product); err != nil {
		return err
	}
	product.ID = primitive.NewObjectID()
	_, err := productsCollection.InsertOne(context.Background(), product)
	if err != nil {
		return err
	}
	return c.JSON(product)
}

func GetAllProducts(c *fiber.Ctx) error {
	_, db := database.ConnectDB()
	productsCollection := db.Collection("products")

	var products []structs.Product
	cursor, err := productsCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return err
	}
	defer cursor.Close(context.Background())
	if err := cursor.All(context.Background(), &products); err != nil {
		return err
	}
	return c.JSON(products)
}

func UpdateProduct(c *fiber.Ctx) error {
	_, db := database.ConnectDB()
	var updatedProduct structs.Product

	productsCollection := db.Collection("products")

	id := c.Params("id")
	productID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	// var updatedProduct structs.Product
	if err := c.BodyParser(&updatedProduct); err != nil {
		return err
	}
	update := bson.M{
		"$set": updatedProduct,
	}
	_, err = productsCollection.UpdateOne(context.Background(), bson.M{"_id": productID}, update)
	if err != nil {
		return err
	}
	updatedProduct.ID = productID
	return c.JSON(updatedProduct)
}

func DeleteProduct(c *fiber.Ctx) error {
	_, db := database.ConnectDB()
	productsCollection := db.Collection("products")
	id := c.Params("id")
	productID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = productsCollection.DeleteOne(context.Background(), bson.M{"_id": productID})
	if err != nil {
		return err
	}
	return c.SendString("Product deleted successfully")
}
