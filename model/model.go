package model
type Movie struct{
	Id string `bson:"id"`
	Name string `bson:"name"`
	Watched bool `bson:"watched"`
}