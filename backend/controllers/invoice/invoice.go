package invoice

import (
	"context"
	"os"

	"github.com/RodolfoMurguia/beat-invoice/database"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Invoice struct {
	Id              primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserId          primitive.ObjectID `json:"userId" bson:"userId,omitempty"`
	Status          int32              `json:"status" bson:"status,omitempty"`
	InvoiceType     int32              `json:"invoiceType" bson:"invoiceType,omitempty"`
	InvoiceRide     primitive.ObjectID `json:"invoiceRide" bson:"invoiceRide,omitempty"`
	InvoiceSubtotal float32            `json:"invoiceSubtotal" bson:"invoiceSubtotal,omitempty"`
	InvoiceTax      float32            `json:"invoiceTax" bson:"invoiceTax,omitempty"`
	InvoiceTotal    float32            `json:"invoiceTotal" bson:"invoiceTotal,omitempty"`
	Notified        bool               `json:"notified" bson:"notified,omitempty"`
	CreatedAt       primitive.DateTime `json:"createdAt" bson:"createdAt,omitempty"`
	UpdatedAt       primitive.DateTime `json:"updatedAt" bson:"updatedAt,omitempty"`
}

func GetInvoices(c *fiber.Ctx) error {

	var result []Invoice

	mongo := database.ConnectDB()
	ctx := context.Background()

	conn := mongo.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("DB_INVOICE_COLLECTION"))
	cursor, err := conn.Find(ctx, bson.M{})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error getting invoices"})
	}

	for cursor.Next(ctx) {
		var invoice Invoice
		cursor.Decode(&invoice)
		result = append(result, invoice)
	}

	if len(result) == 0 {
		return c.Status(200).JSON(fiber.Map{"message": "Invoices not found"})
	}

	return c.Status(200).JSON(fiber.Map{"data": result})
}

func GetInvoicesByUserId(c *fiber.Ctx) error {

	//get userId from url
	userId := c.Params("userId")
	//convert userId to ObjectId
	userIdObjectId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error getting invoices"})
	}

	var result []Invoice
	mongo := database.ConnectDB()
	ctx := context.Background()

	conn := mongo.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("DB_INVOICE_COLLECTION"))
	cursor, err := conn.Find(ctx, bson.M{"userId": userIdObjectId})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error getting invoices"})
	}

	for cursor.Next(ctx) {
		var invoice Invoice
		cursor.Decode(&invoice)
		result = append(result, invoice)
	}

	if len(result) == 0 {
		return c.Status(200).JSON(fiber.Map{"message": "Invoices not found"})
	}
	return c.Status(200).JSON(fiber.Map{"data": result})
}

func GetPendingInvoices(c *fiber.Ctx) error {

	var result []Invoice
	mongo := database.ConnectDB()
	ctx := context.Background()

	conn := mongo.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("DB_INVOICE_COLLECTION"))
	cursor, err := conn.Find(ctx, bson.M{"status": 0})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error getting invoices"})
	}

	for cursor.Next(ctx) {
		var invoice Invoice
		cursor.Decode(&invoice)
		result = append(result, invoice)
	}

	if len(result) == 0 {
		return c.Status(200).JSON(fiber.Map{"message": "Invoices not found"})
	}
	return c.Status(200).JSON(fiber.Map{"data": result})

}

//export invoices to csv
func ExportInvoices(c *fiber.Ctx) error {

	var result []Invoice
	mongo := database.ConnectDB()
	ctx := context.Background()

	conn := mongo.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("DB_INVOICE_COLLECTION"))
	cursor, err := conn.Find(ctx, bson.M{})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error getting invoices"})
	}

	for cursor.Next(ctx) {
		var invoice Invoice
		cursor.Decode(&invoice)
		result = append(result, invoice)
	}

	if len(result) == 0 {
		return c.Status(200).JSON(fiber.Map{"message": "Invoices not found"})
	}

	//generate csv, write to buffer and return

	//create csv file to attach to response

	//return csv file
	return c.Status(200).JSON(fiber.Map{"message": "Invoices exported"})
}

func AddInvoice(c *fiber.Ctx) error {

	var invoice Invoice
	c.BodyParser(&invoice)

	mongo := database.ConnectDB()
	ctx := context.Background()

	conn := mongo.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("DB_INVOICE_COLLECTION"))
	_, err := conn.InsertOne(ctx, invoice)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error adding invoice"})
	}

	return c.Status(200).JSON(fiber.Map{"message": "Invoice added successfully"})

}
