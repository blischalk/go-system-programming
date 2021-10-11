package main

import (
  "context"
	"fmt"
	"log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"html/template"
	"net/http"
	"os"
	"time"
)

var DatabaseName string
var collectionName string

type Document struct {
	P1 int
	P2 int
	P3 int
	P4 int
	P5 int
}

func content(w http.ResponseWriter, r *http.Request) {
	var Data []Document
	myT := template.Must(template.ParseGlob("mongoDB.gohtml")) 
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	collection := client.Database(DatabaseName).Collection(collectionName)

  cur, err := collection.Find(ctx, bson.D{})
  if err != nil {
		fmt.Println(err)
		return
	}
  defer cur.Close(ctx)

  err = cur.All(ctx, &Data)
  if err != nil {
    panic(err)
  }

	fmt.Println("Found:", len(Data), "results!")
	myT.ExecuteTemplate(w, "mongoDB.gohtml", Data)
}

func main() {
	arguments := os.Args

	if len(arguments) <= 2 {
		fmt.Println("Please provide a Database and a Collection!")
		os.Exit(100)
	} else {
		DatabaseName = arguments[1]
		collectionName = arguments[2]
	}

	http.HandleFunc("/", content)
	http.ListenAndServe(":8001", nil)
}
