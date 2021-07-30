package repositories

import (
	"context"
	"ittech24/rest/apidemo/entities"
	"log"
	"time"

	"github.com/rs/xid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetPersonByID Gets a record by ID
func (a *Repository) GetPersonByID(id string) entities.Person {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var Person entities.Person

	filter := bson.D{{Key: "_id", Value: id}}
	collection := a.Factory.Database.Collection(PeopleCollectionName)

	collection.FindOne(ctx, filter).Decode(&Person)

	return Person
}

// GetAllPersons Gets all the records from the Person collection
func (a *Repository) GetAllPersons() []entities.Person {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var Persons []entities.Person

	filter := bson.D{{}}
	collection := a.Factory.Database.Collection(PeopleCollectionName)

	cur, err := collection.Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(ctx) {
		var Person entities.Person
		err := cur.Decode(&Person)
		if err != nil {
			log.Fatal(err)
		}

		Persons = append(Persons, Person)
	}

	return Persons
}

// UpsertPerson Update/Insert and Person in the database
func (a *Repository) UpsertPerson(Person entities.Person) *mongo.UpdateResult {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opts := options.Update().SetUpsert(true)
	if len(Person.ID) == 0 {
		Person.ID = xid.New().String()
	}

	filter := bson.D{{Key: "_id", Value: Person.ID}}
	collection := a.Factory.Database.Collection(PeopleCollectionName)
	update := bson.D{{
		Key: "$set",
		Value: bson.D{
			{Key: "name", Value: Person.Name},
			{Key: "height", Value: Person.Height},
			{Key: "mass", Value: Person.Mass},
			{Key: "hair_color", Value: Person.HairColor},
			{Key: "skin_color", Value: Person.SkinColor},
			{Key: "eye_color", Value: Person.EyeColor},
			{Key: "birth_year", Value: Person.BirthYear},
			{Key: "gender", Value: Person.Gender},
			{Key: "homeworld", Value: Person.Homeworld},
			{Key: "films", Value: Person.Films},
			{Key: "species", Value: Person.Species},
			{Key: "vehicles", Value: Person.Vehicles},
			{Key: "starships", Value: Person.Starships},
			{Key: "created", Value: Person.Created},
			{Key: "edited", Value: Person.Edited},
			{Key: "url", Value: Person.URL},
		}}}

	upsertResult, err := collection.UpdateOne(ctx, filter, update, opts)

	if err != nil {
		log.Fatal(err)
	}

	return upsertResult
}

// UpdatePerson Updates an Person record
func (a *Repository) UpdatePerson(Person entities.Person) *mongo.UpdateResult {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opts := options.Update().SetUpsert(false)
	if len(Person.ID) == 0 {
		Person.ID = xid.New().String()
	}
	filter := bson.D{{Key: "_id", Value: Person.ID}}
	collection := a.Factory.Database.Collection(PeopleCollectionName)
	update := bson.D{{
		Key: "$set",
		Value: bson.D{
			{Key: "name", Value: Person.Name},
			{Key: "height", Value: Person.Height},
			{Key: "mass", Value: Person.Mass},
			{Key: "hair_color", Value: Person.HairColor},
			{Key: "skin_color", Value: Person.SkinColor},
			{Key: "eye_color", Value: Person.EyeColor},
			{Key: "birth_year", Value: Person.BirthYear},
			{Key: "gender", Value: Person.Gender},
			{Key: "homeworld", Value: Person.Homeworld},
			{Key: "films", Value: Person.Films},
			{Key: "species", Value: Person.Species},
			{Key: "vehicles", Value: Person.Vehicles},
			{Key: "starships", Value: Person.Starships},
			{Key: "created", Value: Person.Created},
			{Key: "edited", Value: Person.Edited},
			{Key: "url", Value: Person.URL},
		}}}

	upsertResult, err := collection.UpdateOne(ctx, filter, update, opts)

	if err != nil {
		log.Fatal(err)
	}

	return upsertResult
}

// UpsertManyPersons Bulk Inserts/Updates into the Persons collection
func (a *Repository) UpsertManyPersons(Persons []entities.Person) *mongo.BulkWriteResult {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var operations []mongo.WriteModel
	for i := 0; i < len(Persons); i++ {
		if len(Persons[i].ID) == 0 {
			Persons[i].ID = xid.New().String()
		}
		filter := bson.D{{Key: "_id", Value: Persons[i].ID}}
		update := bson.D{{
			Key: "$set",
			Value: bson.D{
				{Key: "name", Value: Persons[i].Name},
				{Key: "height", Value: Persons[i].Height},
				{Key: "mass", Value: Persons[i].Mass},
				{Key: "hair_color", Value: Persons[i].HairColor},
				{Key: "skin_color", Value: Persons[i].SkinColor},
				{Key: "eye_color", Value: Persons[i].EyeColor},
				{Key: "birth_year", Value: Persons[i].BirthYear},
				{Key: "gender", Value: Persons[i].Gender},
				{Key: "homeworld", Value: Persons[i].Homeworld},
				{Key: "films", Value: Persons[i].Films},
				{Key: "species", Value: Persons[i].Species},
				{Key: "vehicles", Value: Persons[i].Vehicles},
				{Key: "starships", Value: Persons[i].Starships},
				{Key: "created", Value: Persons[i].Created},
				{Key: "edited", Value: Persons[i].Edited},
				{Key: "url", Value: Persons[i].URL},
			}}}
		op := mongo.NewUpdateOneModel()
		op.SetUpsert(true)
		op.SetFilter(filter)
		op.SetUpdate(update)
		operations = append(operations, op)
	}

	collection := a.Factory.Database.Collection(PeopleCollectionName)
	bulkOptions := options.BulkWriteOptions{}

	upsertResult, err := collection.BulkWrite(ctx, operations, &bulkOptions)

	if err != nil {
		log.Fatal(err)
	}

	return upsertResult
}

// DeletePerson Deletes a record from the Person collection
func (a *Repository) DeletePerson(id string) *mongo.DeleteResult {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.D{{Key: "_id", Value: id}}

	collection := a.Factory.Database.Collection(PeopleCollectionName)

	deleteResult, err := collection.DeleteOne(ctx, filter)

	if err != nil {
		log.Fatal(err)
	}

	return deleteResult
}
