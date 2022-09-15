package tag

import (
	"context"
	"fmt"
	"strconv"

	httperrors "github.com/myrachanto/erroring"
	"github.com/myrachanto/microservice/tag/src/db"
	"github.com/myrachanto/microservice/tag/src/support"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// tagrepository ...
var (
	Tagrepository TagRepoInterface = &tagrepository{}
	ctx                            = context.TODO()
)

type TagRepoInterface interface {
	Create(tag *Tag) (*Tag, httperrors.HttpErr)
	GetOne(id string) (*Tag, httperrors.HttpErr)
	GetAll() ([]*Tag, httperrors.HttpErr)
	Featured(code string, status bool) httperrors.HttpErr
	Feature(Bizname string) ([]Tag, httperrors.HttpErr)
	Update(id string, tag *Tag) (*Tag, httperrors.HttpErr)
	Delete(id string) (string, httperrors.HttpErr)
	GetunobyName(name, Bizname string) (result *Tag, err httperrors.HttpErr)
}

type tagrepository struct {
	Mongodb *mongo.Database
	Bizname string
	Cancel  context.CancelFunc
}

func NewtagRepo() TagRepoInterface {
	return &tagrepository{}
}

func (r *tagrepository) Create(tag *Tag) (*Tag, httperrors.HttpErr) {
	if err1 := tag.Validate(); err1 != nil {
		return nil, err1
	}

	code, err1 := r.genecode()
	if err1 != nil {
		return nil, err1
	}

	ok := r.CheckifExist(tag.Name, Bizname)
	if ok {
		return nil, httperrors.NewNotFoundError("That Major Category already exist!")
	}
	tag.Code = code
	tag.Shopalias = Bizname
	tag.Url = support.Joins(tag.Name)
	// fmt.Println("create tag-----------step2", tag)
	collection := db.Mongodb.Collection("tag")
	_, err := collection.InsertOne(ctx, tag)
	if err != nil {
		// fmt.Println("err -------", err)
		return nil, httperrors.NewBadRequestError(fmt.Sprintf("Create product Failed, %d", err))
	}
	tag, e := r.getuno(tag.Code)

	if e != nil {
		return nil, e
	}
	return tag, nil
}

func (r *tagrepository) GetOne(id string) (tag *Tag, errors httperrors.HttpErr) {

	stringresults := httperrors.ValidStringNotEmpty(id)
	if stringresults.Noerror() {
		return nil, stringresults
	}
	collection := db.Mongodb.Collection("tag")
	filter := bson.D{
		{"$and", bson.A{
			bson.D{{"code", id}},
			bson.D{{"shopalias", Bizname}},
		}},
	}
	err := collection.FindOne(ctx, filter).Decode(&tag)
	if err != nil {
		return nil, httperrors.NewBadRequestError(fmt.Sprintf("Could not find resource with this id, %d", err))
	}
	return tag, nil
}

func (r *tagrepository) GetAll() ([]*Tag, httperrors.HttpErr) {
	tags := []*Tag{}
	// fmt.Println("tag get all-------------", db.Mongodb)
	collection := db.Mongodb.Collection("tag")
	filter := bson.D{
		{"$and", bson.A{
			bson.D{{"shopalias", Bizname}},
		}},
	}
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, httperrors.NewNotFoundError("no results found")
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var tag Tag
		err := cur.Decode(&tag)
		if err != nil {
			return nil, httperrors.NewNotFoundError("Error while decoding results!")
		}
		tags = append(tags, &tag)
	}
	if err := cur.Err(); err != nil {
		return nil, httperrors.NewNotFoundError("Error with cursor!")
	}
	return tags, nil

}
func (r *tagrepository) Feature(Bizname string) ([]Tag, httperrors.HttpErr) {

	collection := db.Mongodb.Collection("tag")
	// filter := bson.M{"featured": true}

	filter := bson.D{
		{"$and", bson.A{
			bson.D{{"featured", true}},
			bson.D{{"shopalias", Bizname}},
		}},
	}
	tags := []Tag{}
	options := options.Find()
	// options.SetLimit(5)
	options.SetSort(bson.D{{"name", 1}})
	cursor, err := collection.Find(ctx, filter, options)
	// fmt.Println(cursor)
	if err != nil {
		return nil, httperrors.NewNotFoundError("No records found!")
	}
	if err = cursor.All(ctx, &tags); err != nil {
		return nil, httperrors.NewNotFoundError("Error decoding!")
	}
	// fmt.Println(products)
	return tags, nil

}
func (r *tagrepository) Featured(code string, status bool) httperrors.HttpErr {
	stringresults := httperrors.ValidStringNotEmpty(code)
	if stringresults.Noerror() {
		return stringresults
	}

	upay := &Tag{}

	collection := db.Mongodb.Collection("tag")

	filter := bson.D{
		{"$and", bson.A{
			bson.D{{"code", code}},
			bson.D{{"shopalias", Bizname}},
		}},
	}
	err := collection.FindOne(ctx, filter).Decode(&upay)
	if err != nil {
		return httperrors.NewBadRequestError(fmt.Sprintf("Could not find resource with this id, %d", err))
	}
	_, errs := collection.UpdateOne(
		ctx,
		filter,
		bson.D{
			{"$set", bson.D{{"featured", status}}},
		},
	)
	// update := bson.M{"$set": pay}
	// _, errs := collection.UpdateOne(ctx, filter, update)
	if errs != nil {
		return httperrors.NewNotFoundError("Error updating!")
	}
	return nil
}

func (r *tagrepository) Update(code string, tag *Tag) (*Tag, httperrors.HttpErr) {

	stringresults := httperrors.ValidStringNotEmpty(code)
	if stringresults.Noerror() {
		return nil, stringresults
	}
	// fmt.Println("result+++++++++++++++++++++++++++")
	result, err3 := r.GetunobyCode(code, Bizname)
	if err3 != nil {
		fmt.Println(err3)
	}
	// fmt.Println("result+++++++++++++++++++++++++++step1", result)
	if tag.Name == "" {
		tag.Name = result.Name
	}
	if tag.Title == "" {
		tag.Title = result.Title
	}
	if tag.Description == "" {
		tag.Description = result.Description
	}
	if tag.Description == "" {
		tag.Description = result.Description
	}
	if tag.Code == "" {
		tag.Code = result.Code
	}
	if tag.Shopalias == "" {
		tag.Shopalias = result.Shopalias
	}
	if !tag.Featured {
		tag.Featured = result.Featured
	}
	if tag.Picture == "" {
		tag.Picture = result.Picture
	}
	collection := db.Mongodb.Collection("tag")
	filter := bson.D{
		{"$and", bson.A{
			bson.D{{"code", code}},
			bson.D{{"shopalias", Bizname}},
		}},
	}
	// fmt.Println("result+++++++++++++++++++++++++++step2")
	update := bson.M{"$set": tag}
	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, httperrors.NewNotFoundError("Error updating!")
	}

	tag, e := r.getuno(tag.Code)

	if e != nil {
		return nil, e
	}
	return tag, nil
}
func (r *tagrepository) Delete(id string) (string, httperrors.HttpErr) {

	stringresults := httperrors.ValidStringNotEmpty(id)
	if stringresults.Noerror() {
		return "", stringresults
	}
	collection := db.Mongodb.Collection("tag")
	filter := bson.D{
		{"$and", bson.A{
			bson.D{{"code", id}},
			bson.D{{"shopalias", Bizname}},
		}},
	}
	ok, err := collection.DeleteOne(ctx, filter)
	if ok == nil {
		return "", httperrors.NewNotFoundError(fmt.Sprintf("deletion of %d failed", err))
	}
	return "deleted successfully", nil
}
func (r *tagrepository) genecode() (string, httperrors.HttpErr) {

	special := support.Stamper()
	special2 := Bizname[0:5]
	collection := db.Mongodb.Collection("tag")
	filter := bson.D{
		{"$and", bson.A{
			bson.D{{"shopalias", Bizname}},
		}},
	}
	count, err := collection.CountDocuments(ctx, filter)
	co := count + 1
	if err != nil {
		code := "tagCode" + special2 + strconv.FormatUint(uint64(co), 10) + "-" + special
		return code, nil
	}
	code := "tagCode" + special2 + strconv.FormatUint(uint64(co), 10) + "-" + special
	return code, nil
}
func (r *tagrepository) getuno(code string) (result *Tag, err httperrors.HttpErr) {
	stringresults := httperrors.ValidStringNotEmpty(code)
	if stringresults.Noerror() {
		return nil, stringresults
	}
	collection := db.Mongodb.Collection("tag")
	filter := bson.D{
		{"$and", bson.A{
			bson.D{{"code", code}},
			bson.D{{"shopalias", Bizname}},
		}},
	}
	err1 := collection.FindOne(ctx, filter).Decode(&result)
	if err1 != nil {
		return nil, httperrors.NewNotFoundError("no results found")
	}
	return result, nil
}

func (r *tagrepository) GetunobyCode(name, Bizname string) (result *Tag, err httperrors.HttpErr) {
	stringresults := httperrors.ValidStringNotEmpty(name)
	if stringresults.Noerror() {
		return nil, stringresults
	}
	collection := db.Mongodb.Collection("tag")
	filter := bson.D{
		{"$and", bson.A{
			bson.D{{"code", name}},
			bson.D{{"shopalias", Bizname}},
		}},
	}
	fmt.Println("---------------------------", filter)
	err1 := collection.FindOne(ctx, filter).Decode(&result)
	if err1 != nil {
		return nil, httperrors.NewNotFoundError("no results found")
	}
	return result, nil
}

func (r *tagrepository) GetunobyName(name, Bizname string) (result *Tag, err httperrors.HttpErr) {
	stringresults := httperrors.ValidStringNotEmpty(name)
	if stringresults.Noerror() {
		return nil, stringresults
	}
	collection := db.Mongodb.Collection("tag")
	filter := bson.D{
		{"$and", bson.A{
			bson.D{{"name", name}},
			bson.D{{"shopalias", Bizname}},
		}},
	}
	err1 := collection.FindOne(ctx, filter).Decode(&result)
	if err1 != nil {
		return nil, httperrors.NewNotFoundError("no results found")
	}
	return result, nil
}
func (r *tagrepository) CheckifExist(name, Bizname string) bool {

	stringresults := httperrors.ValidStringNotEmpty(name)
	if stringresults.Noerror() {
		return false
	}
	result := Tag{}
	collection := db.Mongodb.Collection("tag")
	filter := bson.D{
		{"$and", bson.A{
			bson.D{{"name", name}},
			bson.D{{"shopalias", Bizname}},
		}},
	}
	err1 := collection.FindOne(ctx, filter).Decode(&result)
	if err1 != nil {
		return false
	}

	return true
}
