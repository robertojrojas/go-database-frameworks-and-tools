package main
import (
	"log"
	"time"
	"database/sql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Todo struct {
	Id int
	Subject sql.NullString
	Description sql.NullString
	Completed sql.NullBool
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`

}

func insertSelect(db *sql.DB) {

	log.Println("Insert Select...")
	log.Println("================")
	now := time.Now()
	res, err := db.Exec("Insert into todos(subject, description, created_at, updated_at) values ($1, $2, $3, $4)", "Mow the lawn", "", now, now)

	if err != nil {
		log.Fatal(err)
	}
	affected, _ := res.RowsAffected()
	log.Printf("Rows affected %d\n", affected)

	var subject string

	rows, err := db.Query("select subject from todos")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		if err = rows.Scan(&subject); err != nil {
			log.Fatal(err)
		}
		log.Printf("Subject is %s\n", subject)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
}

func queryWhere(db *sqlx.DB) {

	log.Println("Query with Where...")
	log.Println("================")
	var subject string

	rows, err := db.Query("select subject from todos where id = $1", 1)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		if err = rows.Scan(&subject); err != nil {
			log.Fatal(err)
		}
		log.Printf("Subject is %s\n", subject)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
}

func queryWhereToStruct(db *sqlx.DB) {

	log.Println("Query with Where Map to Struct...")
	log.Println("=================================")

	todos := []Todo{}
	err := db.Select(&todos, "select * from todos")

	if err != nil {
		log.Fatal(err)
	}

	for _, todo := range todos {
		log.Printf("Subject is %s\n", todo.CreatedAt)
	}

}

func queryWhereToStructTX(db *sqlx.DB) {

	log.Println("Query with Where Map to Struct TX...")
	log.Println("====================================")

	tx := db.MustBegin()
    now := time.Now()
	t := Todo{
		Subject: sql.NullString{String:"Mow Lawn"},
		Description: sql.NullString{String:"Yuck!"},
		CreatedAt: now,
		UpdatedAt: now,
	}

	tx.Exec("Insert into todos(subject, description, created_at, updated_at) values ($1, $2, $3, $4)", t.Subject, t.Description, t.CreatedAt, t.UpdatedAt)

	tx.Commit()

	todos := []Todo{}
	err := db.Select(&todos, "select * from todos")

	if err != nil {
		log.Fatal(err)
	}

	for _, todo := range todos {
		log.Printf("Subject is %s\n", todo.CreatedAt)
	}

}


func main() {

	db, err := sqlx.Open("postgres",
		"postgres://go_user:password1@localhost/godos_development?sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//insertSelect(db)

	queryWhere(db)

	queryWhereToStruct(db)

	queryWhereToStructTX(db)


}
