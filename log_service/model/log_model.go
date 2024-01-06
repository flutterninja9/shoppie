package model

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type LogModel struct {
	gorm.Model
	Log string `gorm:"not null" json:"log"`
}

type Logs []LogModel

func (l *LogModel) Save(c *mongo.Client) error {
	coll := c.Database("logs").Collection("logs")
	_, err := coll.InsertOne(context.TODO(), l)

	if err != nil {
		fmt.Println(err)
	}

	return nil
}

func (l *Logs) GetLogs(c *mongo.Client) {
	coll := c.Database("logs").Collection("logs")
	cur, err := coll.Find(context.TODO(), bson.D{})

	if err != nil {
		fmt.Println("Error finding logs:", err)
		return
	}

	defer cur.Close(context.TODO())

	// Clearing existing logs
	*l = Logs{}

	// Iterating through the cursor
	for cur.Next(context.TODO()) {
		var log LogModel
		err := cur.Decode(&log)
		if err != nil {
			fmt.Println("Error decoding log:", err)
			continue
		}
		*l = append(*l, log)
	}

	if err := cur.Err(); err != nil {
		fmt.Println("Cursor error:", err)
	}
}
