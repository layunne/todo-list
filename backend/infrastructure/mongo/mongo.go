package mongo

import (
	"context"
	"github.com/gobeam/mongo-go-pagination"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"reflect"
)

type Client interface {
	GetAll(collection string, result interface{}) error
	Insert(collection string, document interface{}) error
	InsertOrUpdate(collection string, document interface{}, id string) error
	InsertMany(collection string, document []interface{}) error
	Remove(collection string, id string) error
	RemoveAll(collection string) error
	GetById(collection string, id string, result interface{}) error
	Find(collection string, keyName string, value string, result interface{}) error
	FindOne(collection string, keyName string, value string, result interface{}) error
	GetAllWithPagination(collection string, limit int64, page int64, result interface{}) (*mongopagination.PaginationData, error)
}

func NewMongoClient(mongoHost string, databaseName string) Client {

	log.Println(mongoHost, databaseName)
	clientOptions := options.Client().ApplyURI(mongoHost)

	// Connect to MongoDB
	conn, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatalf("connect error: %v", err)
	}

	// Check the connection
	err = conn.Ping(context.Background(), nil)

	if err != nil {
		log.Fatalf("Check connection: %v", err)
	}

	log.Println("Connected to MongoDB!")

	return &client{
		mongo: conn.Database(databaseName),
	}
}

type client struct {
	mongo *mongo.Database
}

func (c *client) GetAll(collection string, result interface{}) error {

	mongoResult, err := c.mongo.Collection(collection).Find(context.TODO(), bson.D{})

	if err != nil {
		return err
	}

	rVal := reflect.ValueOf(result).Elem()

	rType := rVal.Type().Elem()

	for mongoResult.Next(context.TODO()) {

		elemVal := reflect.New(rType)
		elem := elemVal.Interface()

		// create a value into which the single document can be decoded
		err := mongoResult.Decode(elem)
		if err != nil {
			return err
		}

		rVal.Set(reflect.Append(rVal, elemVal.Elem()))
	}
	return nil
}

func (c *client) Find(collection string, keyName string, value string, result interface{}) error {

	mongoResult, err := c.mongo.Collection(collection).Find(context.TODO(), bson.M{keyName: value})

	if err != nil {
		return err
	}

	rVal := reflect.ValueOf(result).Elem()

	rType := rVal.Type().Elem()

	for mongoResult.Next(context.TODO()) {

		elemVal := reflect.New(rType)
		elem := elemVal.Interface()

		// create a value into which the single document can be decoded
		err := mongoResult.Decode(elem)
		if err != nil {
			return err
		}

		rVal.Set(reflect.Append(rVal, elemVal.Elem()))
	}
	return nil
}

func (c *client) Insert(collection string, document interface{}) error {
	_, err := c.mongo.Collection(collection).InsertOne(context.TODO(), document)
	return err
}

func (c *client) InsertOrUpdate(collection string, document interface{}, id string) error {
	_, err := c.mongo.Collection(collection).UpdateOne(context.TODO(), bson.M{"_id": id},  bson.M{"$set": document}, options.Update().SetUpsert(true))
	return err
}


func (c *client) InsertMany(collection string, document []interface{}) error {

	_, err := c.mongo.Collection(collection).InsertMany(context.TODO(), document)

	return err
}

func (c *client) Remove(collection string, id string) error {
	_, err := c.mongo.Collection(collection).DeleteOne(context.TODO(), bson.M{"_id": id})
	return err
}

func (c *client) RemoveAll(collection string) error {
	_, err := c.mongo.Collection(collection).DeleteMany(context.TODO(), bson.D{{}})
	return err
}

func (c *client) GetById(collection string, id string, result interface{}) error {

	return c.FindOne(collection, "_id", id, result)
}
func (c *client) FindOne(collection string, keyName string, value string, result interface{}) error {

	res := c.mongo.Collection(collection).FindOne(context.TODO(), bson.M{keyName: value})

	if res.Err() != nil {
		return res.Err()
	}
	return res.Decode(result)
}

func (c *client) GetAllWithPagination(collection string, limit int64, page int64, result interface{}) (*mongopagination.PaginationData, error) {
	coll := c.mongo.Collection(collection)

	paginatedData, err := mongopagination.New(coll).Limit(limit).Page(page).Filter(bson.D{}).Find()
	if err != nil {
		return nil, err
	}

	rVal := reflect.ValueOf(result).Elem()

	rType := rVal.Type().Elem()

	for _, raw := range paginatedData.Data {
		elemVal := reflect.New(rType)
		elem := elemVal.Interface()
		err := bson.Unmarshal(raw, elem)
		if err != nil {
			return nil, err
		}

		rVal.Set(reflect.Append(rVal, elemVal.Elem()))
	}
	return &paginatedData.Pagination, nil
}
