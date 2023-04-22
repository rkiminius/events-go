package events

import (
	"errors"
	"events/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

const collectionName = "event"

type Event struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name         string             `json:"name" bson:"name"`
	Date         time.Time          `json:"date" bson:"date"`
	Languages    []string           `json:"languages" bson:"languages"`
	VideoQuality []string           `json:"videoQuality" bson:"videoQuality"`
	AudioQuality []string           `json:"audioQuality" bson:"audioQuality"`
	Invitees     []string           `json:"invitees" bson:"invitees"`
	Description  string             `json:"description,omitempty" bson:"description"`
	Options      EventOptions       `json:"options,omitempty" bson:"options"`
	UpdatedAt    *time.Time         `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
	DeletedAt    *time.Time         `json:"deletedAt,omitempty" bson:"deletedAt,omitempty"`
}

type EventOptions struct {
	DefaultVideoQuality string `json:"default_video_quality" bson:"default_video_quality"`
	DefaultAudioQuality string `json:"default_audio_quality" bson:"default_audio_quality"`
}

var defaultMaxInvitees = 100

// insert function used to insert a new event into a database collection.
func insert(event *Event) (*Event, error) {
	collection := getCollection()
	ctx, cancel := db.GetTimeoutContext()
	defer cancel()

	if event.ID == primitive.NilObjectID {
		event.ID = primitive.NewObjectID()
	}

	insertOneResult, err := collection.InsertOne(ctx, event)
	if err != nil {
		return nil, err
	}

	eventFromDb, err := getById(insertOneResult.InsertedID.(primitive.ObjectID))
	if err != nil {
		return nil, err
	}

	return eventFromDb, nil
}

// getById used to retrieve an event from the database by its ID.
func getById(id primitive.ObjectID) (*Event, error) {
	var event Event
	ctx, cancel := db.GetTimeoutContext()
	defer cancel()
	filter := bson.M{"_id": id, "deletedAt": nil}
	singleResult := getCollection().FindOne(ctx, filter)
	if err := singleResult.Decode(&event); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		}
		log.Fatal(err)
		return nil, err
	}

	return &event, nil
}

// getAll used to retrieve an events list from the database.
func getAll() (*[]Event, error) {
	events := make([]Event, 0)
	ctx, cancel := db.GetTimeoutContext()
	defer cancel()
	cur, err := getCollection().Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var event Event
		err := cur.Decode(&event)
		if err != nil {
			return nil, err
		}
		if event.DeletedAt == nil {
			events = append(events, event)
		}
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	return &events, nil
}

func updateEvent(objId primitive.ObjectID, event *Event) error {
	ctx, cancel := db.GetTimeoutContext()
	defer cancel()

	filter := bson.M{"_id": objId}
	update := bson.M{"$set": bson.M{
		"name":         event.Name,
		"date":         event.Date,
		"languages":    event.Languages,
		"videoQuality": event.VideoQuality,
		"audioQuality": event.AudioQuality,
		"invitees":     event.Invitees,
		"description":  event.Description,
		"options": bson.M{
			"default_video_quality": event.Options.DefaultVideoQuality,
			"default_audio_quality": event.Options.DefaultAudioQuality,
		},
		"updatedAt": time.Now(),
	}}

	result, err := getCollection().UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.ModifiedCount <= 0 {
		return errors.New("event was not updated")
	}

	return nil
}

func deleteEvent(objId primitive.ObjectID) (int, error) {
	ctx, cancel := db.GetTimeoutContext()
	defer cancel()

	filter := bson.M{"_id": objId}
	update := bson.M{"$set": bson.M{"deletedAt": time.Now()}}
	result, err := getCollection().UpdateOne(ctx, filter, update)
	if err != nil {
		return 0, err
	}
	return int(result.ModifiedCount), nil
}

func getCollection() *mongo.Collection {
	return db.GetMongoCollection(collectionName)
}
