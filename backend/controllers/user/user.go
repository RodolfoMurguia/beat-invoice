package user

import (
	"context"
	"os"
	"time"

	"github.com/RodolfoMurguia/beat-invoice/database"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	//"github.com/aliereno/go-pagination"
	//"github.com/aliereno/go-pagination/frameworks"
)

type User struct {
	Id          primitive.ObjectID `json:"id"`
	FirstName   string             `json:"firstName"`
	LastName    string             `json:"lastName"`
	Email       string             `json:"email"`
	Password    string             `json:"password"`
	UserProfile int                `json:"userProfile"`
	CreatedAt   primitive.DateTime `json:"createdAt"`
	UpdatedAt   primitive.DateTime `json:"updatedAt"`
	IsActive    bool               `json:"isActive"`
}

func GetUsers(c *fiber.Ctx) error {

	//define query and response variables
	var result []User
	query := bson.M{"isActive": bson.M{"$exists": true}}

	mongo := database.ConnectDB()
	ctx := context.Background()

	conn := mongo.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("DB_USER_COLLECTION"))
	cursor, err := conn.Find(ctx, query)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error getting users"})
	}

	//iterate over cursor for parse data
	for cursor.Next(ctx) {
		var user User
		cursor.Decode(&user)
		result = append(result, user)
	}

	//return result, validating if is empty
	if len(result) == 0 {
		return c.Status(200).JSON(fiber.Map{"message": "Users not found"})
	} else {
		return c.Status(200).JSON(fiber.Map{"data": result})
	}

}

func AddUser(c *fiber.Ctx) error {

	userData := new(User)
	c.BodyParser(userData)

	//validate user data
	if userData.FirstName == "" || userData.LastName == "" || userData.Email == "" || userData.Password == "" {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid data"})
	}

	//we add createdAt and updatedAt fields
	userData.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	userData.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())

	//mongo connection

	mongo := database.ConnectDB()
	ctx := context.Background()

	conn := mongo.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("DB_USER_COLLECTION"))
	_, err := conn.InsertOne(ctx, userData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error adding user"})
	}

	//return result
	return c.Status(200).JSON(fiber.Map{"ok": true, "data": userData})

}

func GetUser(c *fiber.Ctx) error {

	//get user id from url
	userId := c.Params("id")
	ObjId, err := primitive.ObjectIDFromHex(userId)

	//validate user id
	if userId == "" || err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Something went wrong"})
	}

	//mongo connection
	mongo := database.ConnectDB()
	ctx := context.Background()

	conn := mongo.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("DB_USER_COLLECTION"))
	user := conn.FindOne(ctx, bson.M{"_id": ObjId})
	if user.Err() != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error getting user"})
	}

	//parse data
	var result User
	user.Decode(&result)

	//return result
	return c.Status(200).JSON(fiber.Map{"data": result})
}

func UpdateUser(c *fiber.Ctx) error {

	userData := new(User)
	c.BodyParser(userData)

	//validate user data
	if userData.FirstName == "" || userData.LastName == "" || userData.Email == "" || userData.Password == "" {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid data"})
	}

	//get user id from url
	userId := c.Params("id")

	//validate user id
	if userId == "" {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid data"})
	}

	//mongo connection
	mongo := database.ConnectDB()
	ctx := context.Background()

	conn := mongo.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("DB_USER_COLLECTION"))
	_, err := conn.UpdateOne(ctx, bson.M{"_id": userId}, bson.M{"$set": userData})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error updating user"})
	}

	//return result
	return c.Status(200).JSON(fiber.Map{"ok": true, "data": userData})
}

func DeactivateUser(c *fiber.Ctx) error {

	//get user id from url
	userId := c.Params("id")
	ObjId, err := primitive.ObjectIDFromHex(userId)
	if userId == "" || err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Something went wrong"})
	}

	//validate user id
	if userId == "" {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid data"})
	}

	//mongo connection
	mongo := database.ConnectDB()
	ctx := context.Background()

	conn := mongo.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("DB_USER_COLLECTION"))
	_, err = conn.UpdateOne(ctx, bson.M{"_id": ObjId}, bson.M{"$set": bson.M{"isActive": false}})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error deactivating user"})
	}

	//return result
	return c.Status(200).JSON(fiber.Map{"ok": true})
}

func LoginUser(c *fiber.Ctx) error {

	//get user data from body
	userData := new(User)
	c.BodyParser(userData)

	//validate user data
	if userData.Email == "" || userData.Password == "" {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid data"})
	}

	//mongo connection
	mongo := database.ConnectDB()
	ctx := context.Background()

	conn := mongo.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("DB_USER_COLLECTION"))
	user := conn.FindOne(ctx, bson.M{"email": userData.Email, "password": userData.Password})
	if user.Err() != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error logging user"})
	}

	//parse data
	var result User
	user.Decode(&result)

	//return result
	return c.Status(200).JSON(fiber.Map{"data": result})
}
