package main
import (
	"log"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"database/sql"
	"encoding/json"
	"os"
	"database/sql/driver"
	"strings"
)

type User struct {
	Id int               `db:"id" json:"id"`
	Email string         `db:"email" json:"email"`
	Name MyNulString  `db:"name" json:"name"`
}

type MyNulString sql.NullString


// ###### Row.Scanner Interface ######
// Called when reading from the DB
func (ns *MyNulString) Scan(value interface{}) error {
	n := sql.NullString{String:ns.String}
	err := n.Scan(value)
	ns.String, ns.Valid = n.String, n.Valid
	return err
}

// ##### Valuer.Valuer Interface #####
// Called when writing to the DB
func (ns *MyNulString) Value() (driver.Value, error) {
	n := sql.NullString{String:ns.String}
	return n.Value()
}



func readFromDB() {
	db, err := sqlx.Open("postgres",
		"postgres://go_user:password1@localhost/godos_development?sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	user := &User{}
	err  = db.Get(user, "select * from users")

	if err != nil {
		log.Fatal(err)
	}

	log.Println(user)
	log.Println(user.Email)
	log.Println(user.Name)

	json.NewEncoder(os.Stdout).Encode(user)
}

// ##### Marshaler Interface #####
// Called when generating JSON output
func (ns *MyNulString) MarshalJSON() ([]byte, error) {
	if ns.Valid {
		return json.Marshal(ns.String)
	}

	return json.Marshal(nil)

}

// ##### UnMarshaler Interface #####
// Called when reading JSON input
func (ns *MyNulString) UnmarshalJSON(text []byte) error {
	ns.Valid = false

	if string(text) == "null" {
		return nil
	}

	s := ""
	err := json.Unmarshal(text, &s)
	if err == nil {
		ns.Valid = true
		ns.String = s
	}
    return err
}

func readJSONWithNull() {
	log.Println("+++ Parsing JSON with null")
	user := User{}
	jsonData := `{"id":1, "email":"roberto@example.com", "name":null}`
	json.NewDecoder(strings.NewReader(jsonData)).Decode(&user)
	printUser(user)

}

func readJSON() {
	log.Println("+++ Parsing normal JSON")
	user := User{}
	jsonData := `{"id":1, "email":"roberto@example.com", "name":"Roberto"}`
	json.NewDecoder(strings.NewReader(jsonData)).Decode(&user)
	printUser(user)
}

func printUser(user User) {
	log.Println(user)
	log.Println(user.Email)
	log.Printf("user.Name.String [%s]\n",user.Name.String)
}

func main() {

	//readFromDB()

	readJSON()

	readJSONWithNull()

	// Cool package that handles nulls my Mark: github.com/markbates/going/nulls

}
