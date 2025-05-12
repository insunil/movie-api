package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"module/model"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbname = "netflix"
const collname = "watchlist"

var collection *mongo.Collection

func init() {

	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientoption := options.Client().ApplyURI("mongodb+srv://sunil:sunil@cluster0.bzpjx.mongodb.net/")
	client, _ := mongo.Connect(context, clientoption)
	collection = client.Database(dbname).Collection(collname)
	fmt.Println("controller")
}

// create
func addmovie(m model.Movie) {
	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, _ := collection.InsertOne(context, m)
	fmt.Println(result.InsertedID)
}

// getSpecific
func getSpecificmovie(nid string) model.Movie {
	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var movie model.Movie
	collection.FindOne(context, bson.M{"id": nid}).Decode(&movie)

	return movie
}

// get all
func getallmovies() []model.Movie {
	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, _ := collection.Find(context, bson.M{})
	defer cursor.Close(context)
	var movies []model.Movie
	for cursor.Next(context) {
		var movie model.Movie
		cursor.Decode(&movie)

		movies = append(movies, movie)

	}
	return movies
}

// update
func updatemovie(nid string) {
	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	res, _ := collection.UpdateOne(context, bson.M{"id": nid}, bson.M{"$set": bson.M{"watched": true}})
	fmt.Println(res.ModifiedCount)
}

// deletespecific
func deletespecificmovie(nid string) {
	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	res, _ := collection.DeleteOne(context, bson.M{"id": nid})
	fmt.Println(res.DeletedCount)
}

// deleteall
func deleteall() {
	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	res, _ := collection.DeleteMany(context, bson.M{})
	fmt.Println(res.DeletedCount)
}

// controller
func Cadd(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	var movie model.Movie
	json.NewDecoder(r.Body).Decode(&movie)
	addmovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func Cgetspecificw(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	params := mux.Vars(r)
	var movie model.Movie
	movie = getSpecificmovie(params["id"])
	json.NewEncoder(w).Encode(&movie)
}

func Cgetall(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	allmovies := getallmovies()
	json.NewEncoder(w).Encode(allmovies)
}
func Cupdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	params := mux.Vars(r)
	updatemovie(params["id"])
}
