package peda

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/aiteung/atdb"
	"github.com/whatsauth/watoken"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetConnection(mongoenv, dbname string) *mongo.Database {
	var DBmongoinfo = atdb.DBInfo{
		DBString: os.Getenv(mongoenv),
		DBName:   dbname,
	}
	return atdb.MongoConnect(DBmongoinfo)
}

func CompareUsername(mongoenv *mongo.Database, collname, username string) bool {
	filter := bson.M{"username": username}
	err := atdb.GetOneDoc[User](mongoenv, collname, filter)
	users := err.Username
	if users == "" {
		return false
	}
	return true
}

func GetAllBangunanLineString(mongoenv *mongo.Database, collname string) []GeoJson {
	lokasi := atdb.GetAllDoc[[]GeoJson](mongoenv, collname)
	return lokasi
}

func PostPoint(mongoconn *mongo.Database, collection string, pointdata GeoJsonPoint) interface{} {
	return atdb.InsertOneDoc(mongoconn, collection, pointdata)
}

func PostLinestring(mongoconn *mongo.Database, collection string, linestringdata GeoJsonLineString) interface{} {
	return atdb.InsertOneDoc(mongoconn, collection, linestringdata)
}

func PostPolygon(mongoconn *mongo.Database, collection string, polygondata GeoJsonPolygon) interface{} {
	return atdb.InsertOneDoc(mongoconn, collection, polygondata)
}

func MemasukkanKoordinat(MongoConn *mongo.Database, colname string, coordinate []float64, name, volume, tipe string) (InsertedID interface{}) {
	req := new(Coordinate)
	req.Type = tipe
	req.Coordinates = coordinate
	req.Name = name

	ins := atdb.InsertOneDoc(MongoConn, colname, req)
	return ins
}

func GetNameAndPassowrd(mongoenv *mongo.Database, collname string) []User {
	user := atdb.GetAllDoc[[]User](mongoenv, collname)
	return user
}

func GetAllUser(mongoenv *mongo.Database, collname string) []User {
	user := atdb.GetAllDoc[[]User](mongoenv, collname)
	return user
}
func CreateNewUserRole(mongoenv *mongo.Database, collname string, userdata User) interface{} {
	// Hash the password before storing it
	hashedPassword, err := HashPassword(userdata.Password)
	if err != nil {
		return err
	}
	userdata.Password = hashedPassword

	// Insert the user data into the database
	return atdb.InsertOneDoc(mongoenv, collname, userdata)
}
func CreateUserAndAddedToeken(privatekey string, mongoenv *mongo.Database, collname string, userdata User) interface{} {
	// Hash the password before storing it
	hashedPassword, err := HashPassword(userdata.Password)
	if err != nil {
		return err
	}
	userdata.Password = hashedPassword

	// Insert the user data into the database
	atdb.InsertOneDoc(mongoenv, collname, userdata)

	// Create a token for the user
	tokenstring, err := watoken.Encode(userdata.Username, os.Getenv(privatekey))
	if err != nil {
		return err
	}
	userdata.Token = tokenstring

	// Update the user data in the database
	return atdb.ReplaceOneDoc(mongoenv, collname, bson.M{"username": userdata.Username}, userdata)
}

func DeleteUser(mongoenv *mongo.Database, collname string, userdata User) interface{} {
	filter := bson.M{"username": userdata.Username}
	return atdb.DeleteOneDoc(mongoenv, collname, filter)
}
func ReplaceOneDoc(mongoenv *mongo.Database, collname string, filter bson.M, userdata User) interface{} {
	return atdb.ReplaceOneDoc(mongoenv, collname, filter, userdata)
}
func FindUser(mongoenv *mongo.Database, collname string, userdata User) User {
	filter := bson.M{"username": userdata.Username}
	return atdb.GetOneDoc[User](mongoenv, collname, filter)
}

func FindUserUser(mongoenv *mongo.Database, collname string, userdata User) User {
	filter := bson.M{
		"username": userdata.Username,
	}
	return atdb.GetOneDoc[User](mongoenv, collname, filter)
}

func IsPasswordValid(mongoenv *mongo.Database, collname string, userdata User) bool {
	filter := bson.M{"username": userdata.Username}
	res := atdb.GetOneDoc[User](mongoenv, collname, filter)
	hashChecker := CheckPasswordHash(userdata.Password, res.Password)
	return hashChecker
}

func InsertUserdata(mongoenv *mongo.Database, collname, username, role, password string) (InsertedID interface{}) {
	req := new(User)
	req.Username = username
	req.Password = password
	req.Role = role
	return atdb.InsertOneDoc(mongoenv, collname, req)
}

func usernameExists(mongoenv, dbname string, userdata User) bool {
	mconn := SetConnection(mongoenv, dbname).Collection("user")
	filter := bson.M{"username": userdata.Username}

	var user User
	err := mconn.FindOne(context.Background(), filter).Decode(&user)
	return err == nil
}

func PostStructWithToken[T any](tokenkey string, tokenvalue string, structname interface{}, urltarget string) (result T, errormessage string) {
	client := http.Client{}
	mJson, _ := json.Marshal(structname)
	req, err := http.NewRequest("POST", urltarget, bytes.NewBuffer(mJson))
	if err != nil {
		errormessage = "http.NewRequest Got error :" + err.Error()
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add(tokenkey, tokenvalue)
	resp, err := client.Do(req)
	if err != nil {
		errormessage = "client.Do(req) Error occured. Error is :" + err.Error()
		return
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		errormessage = "Error Read Data data from request." + err.Error()
		return
	}
	if er := json.Unmarshal(respBody, &result); er != nil {
		errormessage = string(respBody) + "Error Unmarshal from Response : " + er.Error()
	}
	return
}

func CreateResponse(status bool, message string, data interface{}) Jaja {
	response := Jaja{
		Status:  status,
		Message: message,
		Data:    data,
	}
	return response
}
