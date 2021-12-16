package ride

import (
	"context"
	"os"

	"github.com/RodolfoMurguia/beat-invoice/database"
	//"github.com/RodolfoMurguia/beat-invoice/controller/tax"
	//"github.com/RodolfoMurguia/beat-invoice/controller/user"
	//"github.com/RodolfoMurguia/beat-invoice/controller/invoice"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Ride struct {
	Id          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserId      string             `json:"userId" bson:"userId,omitempty"`
	RideDate    primitive.DateTime `json:"rideDate" bson:"rideDate,omitempty"`
	PickupZone  string             `json:"pickupZone" bson:"pickupZone,omitempty"`
	DropoffZone string             `json:"dropoffZone" bson:"dropoffZone,omitempty"`
	RideCost    float64            `json:"rideCost" bson:"rideCost,omitempty"`
	IsInvoiced  bool               `json:"isInvoiced" bson:"isInvoiced,omitempty"`
}

func GetRidesbyUser(c *fiber.Ctx) error {

	//get userId from url
	userId := c.Params("userId")

	//define query and response variables
	var result []Ride
	query := bson.M{"userId": userId}
	mongo := database.ConnectDB()
	ctx := context.Background()

	//define connection and excecute query
	conn := mongo.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("DB_RIDE_COLLECTION"))
	cursor, err := conn.Find(ctx, query)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error getting rides"})
	}

	//iterate over cursor for parse data

	for cursor.Next(ctx) {
		var ride Ride
		cursor.Decode(&ride)
		result = append(result, ride)
	}

	//return result, validating if is empty
	if len(result) == 0 {
		return c.Status(200).JSON(fiber.Map{"message": "Rides not found"})
	}

	return c.Status(200).JSON(fiber.Map{"data": result})
}

func GetRideById(c *fiber.Ctx) error {
	// get rideId from url
	rideId := c.Params("rideId")

	//convert rideId to ObjectId
	oid, err := primitive.ObjectIDFromHex(rideId)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error getting ride"})
	}

	//define query and response variables
	var result Ride
	query := bson.M{"_id": oid}
	mongo := database.ConnectDB()
	ctx := context.Background()

	//define connection and excecute query

	conn := mongo.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("DB_RIDE_COLLECTION"))
	err = conn.FindOne(ctx, query).Decode(&result)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error getting ride"})
	}

	return c.Status(200).JSON(fiber.Map{"data": result})
}

func AddRide(c *fiber.Ctx) error {

	//get data from body
	var ride Ride
	c.BodyParser(&ride)

	//we validate if the data is valid
	if ride.PickupZone == "" || ride.DropoffZone == "" {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid data"})
	}

	//define query and response variables
	mongo := database.ConnectDB()
	ctx := context.Background()

	//define connection and excecute query
	conn := mongo.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("DB_RIDE_COLLECTION"))
	_, err := conn.InsertOne(ctx, ride)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error adding ride"})
	}

	return c.Status(200).JSON(fiber.Map{"message": "Ride added successfully"})
}

func GenerateInvoiceByRide(c *fiber.Ctx) {

	type GenerateInvoice struct {
		RideId       string `json:"rideId" bson:"rideId,omitempty"`
		UserId       string `json:"userId" bson:"userId,omitempty"`
		TaxProfileId string `json:"taxProfileId" bson:"taxProfileId,omitempty"`
	}

	//get data from body

	var generateInvoice GenerateInvoice
	c.BodyParser(&generateInvoice)

	//define query and response variables

	//retrieve the data from our collections to generate the invoice

	//build the body of or unvoice

	//we call the addInvoice function to add the invoice to our database

	//return the invoice

}
