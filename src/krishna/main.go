package main

import (
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/labstack/echo"
	"net/http"
)

type Message struct{
	Message string `json:"message"`
	Status string `json:"message"`
}

type Service struct{
	Name string `json:"name"`
	Args string `json:"args"`
}


func main(){
	client, _ := mongo.NewClient("mongodb://localhost:27017")
	e := echo.New()
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
    collection := client.Database("zoo_users").Collection("users")
	e.GET("/", func(c echo.Context) error {
		m := &Message{
			Message: "Ok",
			Status: "200",
		}
		return c.JSON(http.StatusOK, m);
	})

	e.GET("/serviceList", func(c echo.Context) error {
		serviceList := []string{}
		return c.JSON(http.StatusOK, serviceList)
	})

	e.GET("/service", func(c echo.Context) error {
		return c.JSON(http.StatusOK, &Message{})
	})

	e.GET("/register", func(c echo.Context) error {
		collection.InsertOne(ctx,bson.M{"key": c.FormValue("key")});
		return c.JSON(http.StatusOK, &Message{})
	})
}