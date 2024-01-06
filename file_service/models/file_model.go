package models

import (
	"context"

	db "github.com/flutterninja9/shoppie/image_service/database"
	"go.mongodb.org/mongo-driver/bson"
)

type EntityType string

const (
	EntityTypeProduct EntityType = "product"
	EntityTypeReview  EntityType = "review" // wont be used as of now
)

type FileModel struct {
	ID         string     `bson:"id" gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	EntityId   string     `bson:"entity_id" gorm:"not null" json:"entity_id"`
	EntityType EntityType `bson:"entity_type" gorm:"not null" json:"entity_type"`
	Name       string     `bson:"name" gorm:"not null" json:"name"`
	Path       string     `bson:"path" gorm:"not null" json:"path"`
	Type       string     `bson:"type" gorm:"not null" json:"type"`
	Metadata   string     `bson:"metadata" json:"metadata,omitempty"`
}

func (f *FileModel) Save() error {
	ctx := context.TODO()
	fileBson, err := bson.Marshal(f)
	if err != nil {
		return err
	}

	_, saveErr := db.DBClient.Database("files").Collection("files").InsertOne(ctx, fileBson)
	if saveErr != nil {
		return saveErr
	}

	return nil
}
