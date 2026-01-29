package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Diya28-mehra/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionstring ="mongodb+srv://letsconnectdm_db_user:diya%40123@golang.zvmkbzc.mongodb.net/?appName=GoLang"
const dbName = "Netflix"
const colName = "watchlist"

//Most Important
var collection *mongo.Collection

func init(){
	//client option
	client_options := options.Client().ApplyURI(connectionstring)

	//connect with mongodb
	client,err:=mongo.Connect(context.TODO(),client_options)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("MongoDb Connection Success")

	collection=client.Database(dbName).Collection(colName)

	//collection instance ready for me
	fmt.Println("Collections Referance is Ready",collection)
}

//Mongodb helpers - file

//insert one record 
func insertOneMovie(movie models.Netflix){
	inserted,err:=collection.InsertOne(context.Background(),movie)
	if err!=nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted One movie in DB ",inserted.InsertedID)
}

//Update a record
func updateMovie(movieId string){
	id,_:=primitive.ObjectIDFromHex(movieId)
	filter:=bson.M{"_id":id}
	update:=bson.M{"$set":bson.M{"watched":true}}
	result,err := collection.UpdateOne(context.Background(),filter,update)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("Movie is updated: ",result.ModifiedCount)
}

//delete 1 record
func deleteMovie(movieId string){
	id,_:=primitive.ObjectIDFromHex(movieId)
	filter:=bson.M{"_id":id}
	result,err:=collection.DeleteOne(context.Background(),filter)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("Movie got deleted with delete count: ",result)
}

//delete all records
func deleteAllRecords() int64{
	result,err:=collection.DeleteMany(context.Background(),bson.D{{}})
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("Deleted all movies with delete count: ",result.DeletedCount)
	return result.DeletedCount
}

func getAllMovies()[]primitive.M{
	cur,err:=collection.Find(context.Background(),bson.D{{}})
	if err!=nil{
		log.Fatal(err)
	}

	var movies []primitive.M
	for cur.Next(context.Background()) {
		var movie bson.M
		err:=cur.Decode(&movie)
		if err!=nil{
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}

	defer cur.Close(context.Background())
	return movies
}


//Actual Controllers - file
func GetMyAllMovies(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Access-Control-Allow-Methods", "POST")


	var movie models.Netflix
	_=json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkedAsWatched(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Allow-Control-Allow-Methods","PUT")

	params:=mux.Vars(r)

	updateMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Allow-Control-Allow-Methods","POST")

	params:=mux.Vars(r)
	deleteMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovies(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Allow-Control-Allow-Methods","POST")

	count:=deleteAllRecords()
	json.NewEncoder(w).Encode(count)
}