package tax

import (
	"context"
	"os"
	"time"

	fiber "github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/RodolfoMurguia/beat-invoice/database"
)

type Tax struct {
	Id           primitive.ObjectID `json:"id"`
	Name         string             `json:"name"`
	UserId       string             `json:"userId"`
	CompanyName  string             `json:"companyName"`
	CompanyEmail string             `json:"companyEmail"`
	TaxIdNumber  string             `json:"taxIdNumber"`
	CreatedAt    primitive.DateTime `json:"createdAt"`
	UpdatedAt    primitive.DateTime `json:"updatedAt"`
}

func GetTaxProfiles(c *fiber.Ctx) error {

	//reciver id from params
	userId := c.Params("id")

	//convert id into mongo object id

	userIdObjectId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error converting id to object id"})
	}

	//build query
	query := bson.M{"userId": userIdObjectId}

	//get all tax profiles
	var result []Tax

	mongo := database.ConnectDB()
	ctx := context.Background()

	conn := mongo.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("DB_TAX_COLLECTION"))
	cursor, err := conn.Find(ctx, query)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error getting tax profiles"})
	}

	//iterate over cursor for parse data
	for cursor.Next(ctx) {
		var tax Tax
		cursor.Decode(&tax)
		result = append(result, tax)
	}

	//return result
	if len(result) > 0 {
		return c.JSON(result)
	} else {
		return c.Status(404).JSON(fiber.Map{"message": "No tax profiles found"})
	}

}

func AddTaxProfile(c *fiber.Ctx) error {

	taxprofile := new(Tax)
	c.BodyParser(&taxprofile)

	//manage the date
	taxprofile.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	taxprofile.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())

	//we genarate the connection and insert the tax profile
	mongo := database.ConnectDB()
	ctx := context.TODO()

	conn := mongo.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("DB_TAX_COLLECTION"))
	cursor, err := conn.InsertOne(ctx, taxprofile)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error adding tax profile"})
	}

	//return result
	return c.Status(200).JSON(fiber.Map{"ok": true, "data": cursor.InsertedID})
}

func GetTaxProfilesByName(c *fiber.Ctx) error {

	//reciver id from params
	name := c.Params("name")

	//convert id into mongo object id

	//build query
	query := bson.M{"name": name}

	//get all tax profiles
	var result []Tax

	mongo := database.ConnectDB()
	ctx := context.Background()

	conn := mongo.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("DB_TAX_COLLECTION"))
	cursor, err := conn.Find(ctx, query)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error getting tax profiles"})
	}

	//iterate over cursor for parse data
	for cursor.Next(ctx) {
		var tax Tax
		cursor.Decode(&tax)
		result = append(result, tax)
	}

	//return result
	if len(result) > 0 {
		return c.JSON(result)
	} else {
		return c.Status(404).JSON(fiber.Map{"message": "No tax profiles found"})
	}

}
