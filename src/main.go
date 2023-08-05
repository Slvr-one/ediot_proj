package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

/*
	The Item struct represents each item in the tracker.

It contains fields:

	Name, Amount,
	Expiration, OrderArrival, and PreviousDates.
*/
type Item struct {
	Name          string
	Amount        int
	Expiration    time.Time
	OrderArrival  time.Time
	PreviousDates []time.Time
}

// The Tracker struct is responsible for holding a list of items.
type Tracker struct {
	Items []Item
}

// The AddItem method is used to create a new item and add it to the tracker.
func (t *Tracker) AddItem(name string, amount int, expiration time.Time, orderArrival time.Time, client *mongo.Client) error {
	item := Item{
		Name:          name,
		Amount:        amount,
		Expiration:    expiration,
		OrderArrival:  orderArrival,
		PreviousDates: []time.Time{},
	}
	t.Items = append(t.Items, item)

	// Save the item to the MongoDB database
	collection := client.Database("your_database_name").Collection("your_collection_name")
	_, err := collection.InsertOne(context.TODO(), item)
	if err != nil {
		return err
	}

	return nil
}

// The SavePreviousDate method is used to save the previous dates for a specific item.
func (t *Tracker) SavePreviousDate(name string, date time.Time, client *mongo.Client) error {
	for i := range t.Items {
		if t.Items[i].Name == name {
			t.Items[i].PreviousDates = append(t.Items[i].PreviousDates, date)

			// Update the item in the MongoDB database
			collection := client.Database("your_database_name").Collection("your_collection_name")
			_, err := collection.UpdateOne(
				context.TODO(),
				Item{Name: name},
				bson.M{"$set": bson.M{"previousdates": t.Items[i].PreviousDates}},
			)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func main() {
	// create a new instance of the Tracker and add some items to it.
	t := Tracker{}

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())

	// Add new items
	err = t.AddItem("Doritos Classic 40g X 50 X 3", 5, time.Now().AddDate(0, 0, 3), time.Now(), client)
	if err != nil {
		log.Fatal(err)
	}
	err = t.AddItem("Coce-cola 1.5L X 6 X 20", 10, time.Now().AddDate(0, 0, 5), time.Now(), client)
	if err != nil {
		log.Fatal(err)
	}

	// Save previous dates for an item
	err = t.SavePreviousDate("Item 1", time.Now().AddDate(0, 0, -1), client)
	if err != nil {
		log.Fatal(err)
	}
	err = t.SavePreviousDate("Item 1", time.Now().AddDate(0, 0, -3), client)
	if err != nil {
		log.Fatal(err)
	}
	err = t.SavePreviousDate("Item 2", time.Now().AddDate(0, 0, -2), client)
	if err != nil {
		log.Fatal(err)
	}
}
