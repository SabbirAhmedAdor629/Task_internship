package main

import (
	"database/sql"
	"fmt"
//	"time"
//	"log"
	"time"
	_ "github.com/lib/pq"
)
var db  *sql.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "messaging_development"
)


// func DeleteFromCampaign(db *sql.DB){
// 	var id int
// 	fmt.Scan(&id)
// 	del := fmt.Sprintf("DELETE FROM version_associations WHERE id = '%d'", id)
// 	_, err := db.Exec(del)
// 	CheckError(err)
// }

// func UpdateCampaign(db *sql .DB) {
// 	var event_name string
// 	var id int
// 	fmt.Println("Enter new event name : ")
// 	fmt.Scan(&event_name)
// 	fmt.Println("Enter id ")
// 	fmt.Scan(&id)
// 	update	:=	`UPDATE version_associations SET "foreign_key_name" = $1 WHERE "id" = $2`
// 	_, err :=	db .Exec(update, event_name, id)
// 	CheckError(err)
// }


func InsertIntoCategories(db *sql.DB){
	// var Id int
	// var Name string
	// var Description string
	// var Parent_category_id int
	// var Priority int
	// var Can_opt_out bool
	
	StartAt := time.Now()
	EndAt := time.Now()
	CreatedAt := time.Now()
	UpdatedAt := time.Now()
	// fmt.Println("Enter id : ")
	// fmt.Scan(&Id)
	// fmt.Println("Enter name : ")
	// fmt.Scan(&Name)
	// fmt.Println("Enter description : ")
	// fmt.Scan(&Description)
	// fmt.Println("Enter parent_catagory_id : ")
	// fmt.Scan(&Parent_category_id)
	// fmt.Println("Enter priority: ")
	// fmt.Scan(&Priority)	
	// fmt.Println("Enter can_opt_out: ")
	// fmt.Scan(&Can_opt_out)


// 	email, loginTime := "human@example.com", time.Now()
// result, err := db.Exec("INSERT INTO UserAccount VALUES ($1, $2)", email, loginTime)
// if err != nil {
//   panic(err)
// }

// insertTableAdId :=
// 				`INSERT INTO ` + env.DbTable +
// 					` (hashed_ad_id, app_ids, created_at, updated_at) VALUES ($1, $2, $3, $4) ` +
// 					` ON CONFLICT DO NOTHING`


	_, err := db.Exec("INSERT INTO campaigns (id, event_name, message_template_id, start_at, end_at, created_at, updated_at)  VALUES ($1,$2,$3,$4,$5,$6,$7)", 1, "class", 12,  StartAt, EndAt, CreatedAt,
	 UpdatedAt )
	CheckError(err)
	
	// insert := fmt.Sprintf("INSERT INTO categories VALUES ('%s');", event_name )
	// _, err := db.Exec(insert)
	
}


func main() {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	// close database
	defer db.Close()

	InsertIntoCategories(db)
	//UpdateCampaign(db)
	// DeleteFromCampaign(db)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
