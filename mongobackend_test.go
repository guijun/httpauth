package goauth

import (
    "testing"
    "labix.org/v2/mgo"
    "fmt"
    //"os"
    "bytes"
)

var (
    mb MongodbAuthBackend
)

func TestNewMongodbBackend(t *testing.T) {
    sesh, err := mgo.Dial("mongodb://localhost:27017")
    if err != nil {
        panic(err.Error())
    }
    c := sesh.DB("goauth").C("goauth")
    c.DropCollection()

    mb, err := NewMongodbBackend("mongodb://localhost:27017", "goauth")
    if err != nil {
        t.Fatal("Error creating mongodbbackend")
    }
    if mb.mongoUrl != "mongodb://localhost:27017" {
        t.Fatal("Database url not saved.")
    }
    if mb.database != "goauth" {
        t.Fatal("Database name not saved.")
    }
}

func TestSaveUser_mongo(t *testing.T) {
    user := UserData{"username", "email", []byte("passwordhash")}
    if err := mb.SaveUser(user); err != nil {
        t.Fatalf("Save user (1): %v", err)
    }

    user2 := UserData{"username2", "email2", []byte("passwordhash2")}
    if err := mb.SaveUser(user2); err != nil {
        t.Fatalf("Save user (1): %v", err)
    }

    fmt.Println("test")
}

func TestUser_existing_mongo(t *testing.T) {
    if user, ok := mb.User("username"); ok {
        if user.Username != "username" {
            t.Fatal("Username not correct.")
        }
        if user.Email != "email" {
            t.Fatal("User email not correct.")
        }
        if !bytes.Equal(user.Hash, []byte("passwordhash")) {
            t.Fatal("User password not correct.")
        }
    } else {
        t.Fatal("User not found")
    }
}


