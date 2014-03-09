package goauth

import (
    //"os"
    //"errors"
    "labix.org/v2/mgo"
    "labix.org/v2/mgo/bson"
    "fmt"
    "errors"
)

// GobFileAuthBackend stores user data and the location of the gob file.
type MongodbAuthBackend struct {
    mongoUrl string
    database string
}

// NewGobFileAuthBackend initializes a new backend by loading a map of users
// from a file.
func NewMongodbBackend(mongoUrl string, database string) (b MongodbAuthBackend, err error) {
    b.mongoUrl = mongoUrl
    b.database = database
    _, err = mgo.Dial(b.mongoUrl)
    if err != nil {
        return b, errors.New("Can't connect to mongodb: " + err.Error())
    }
    return b, nil
}

func (b MongodbAuthBackend) connect() (c *mgo.Collection, err error) {
    sesh, err := mgo.Dial(b.mongoUrl)
    if err != nil {
        return c, errors.New("Can't connect to mongodb: " + err.Error())
    }
    return sesh.DB(b.database).C("goauth"), nil
}

// User returns the user with the given username.
func (b MongodbAuthBackend) User(username string) (user UserData, ok bool) {
    c, err := b.connect()
    if err != nil {
        panic(err)
    }
    u := c.Find(bson.M{"Username":"username"})
    fmt.Println(u.Count())
    fmt.Println("U")
    fmt.Println(u)
    return user, false
}

// Users returns a slice of all users.
func (b MongodbAuthBackend) Users() (us []UserData) {
    c, err := b.connect()
    if err != nil {
        panic(err)
    }
    results := c.Find(nil).All
    fmt.Println(results)
    return
}

// SaveUser adds a new user, replacing one with the same username, and saves a
// gob file.
func (b MongodbAuthBackend) SaveUser(user UserData) error {
    c, err := b.connect()
    if err != nil {
        panic(err)
    }
    err = c.Insert(bson.M{"Username":user.Username, "Hash":user.Hash, "Email":user.Email})
    return err
}

/*
func (b GobFileAuthBackend) save() error {
    f, err := os.Create(b.filepath)
    defer f.Close()
    if err != nil {
        return errors.New("auth file can't be edited. Is the data folder there?")
    }
    enc := gob.NewEncoder(f)
    err = enc.Encode(b.users)
    return err
}

// DeleteUser removes a user.
func (b GobFileAuthBackend) DeleteUser(username string) error {
    delete(b.users, username)
    err := b.save()
    return err
}
*/
