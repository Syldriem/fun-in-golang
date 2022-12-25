package database

import (
	"encoding/json"
	"os"
)


type Client struct {
    Dbpath string
}


type databaseSchema struct {
	Users map[string]User `json:"users"`
	Posts map[string]Post `json:"posts"`
}

func NewClient(path string) Client {
    c := Client{path}
    return c
}
func (c Client) createDB() error {
    data, err := json.Marshal(databaseSchema{
        Users: make(map[string]User),
        Posts: make(map[string]Post),
    })
    if err != nil {
        return err
    }
    err = os.WriteFile(c.Dbpath, data, 0600)
    if err != nil {
        return err
    }
    return nil
}
func (c Client) EnsureDB() error {
    file, _ := os.ReadFile(c.Dbpath)
    var err error
    if (file == nil) {
        err = c.createDB()
    }
    return err
}

func (c Client) updateDB(db databaseSchema) error {
    data, err := json.Marshal(db)
    if err != nil {
        return err
    }
    err = os.WriteFile(c.Dbpath, data, 0600)
    if err != nil {
       return err
    }
    return nil
}
func (c Client) readDB() (databaseSchema, error) {
    data, err := os.ReadFile(c.Dbpath)
    if err != nil {
        return databaseSchema{}, err
    }
    db := databaseSchema{}
    err = json.Unmarshal(data, &db)
    return db, err
}
