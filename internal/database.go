package database

import (
	"encoding/json"
	"log"
	"os"
	"time"
)


type Client struct {
    Dbpath string
}


type databaseSchema struct {
	Users map[string]User `json:"users"`
	Posts map[string]Post `json:"posts"`
}

type User struct {
	CreatedAt time.Time `json:"createdAt"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Name      string    `json:"name"`
	Age       int       `json:"age"`
}

type Post struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UserEmail string    `json:"userEmail"`
	Text      string    `json:"text"`
}

func NewClient(path string) Client {
    c := Client{path}
    return c
}
func (c Client) createDB() error {
    file, err := os.ReadFile(c.Dbpath)
    json.Marshal(file)
    return err
}
func (c Client) EnsureDB() error {
    file, _ := os.ReadFile(c.Dbpath)
    var err error
    if (file == nil) {
        err = c.createDB()
    }
    return err
}
