package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	PostgesHost      = "localhost"
	PotsgresUser     = "postgres"
	PostgresPassword = "7"
	PostgresPort     = 5432
	PostgresDatabase = "golang"
)

func main() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s database=%s sslmode=disable",
		PostgesHost,
		PostgresPort,
		PotsgresUser,
		PostgresPassword,
		PostgresDatabase,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("failed to open connection: %v", err)
	}

	dbManager := NewDBManager(db)

	//err = InsertFakeDate(dbManager)
	//if err != nil {
	//	log.Fatalf("failed to insert fake date: %v", err)
	//}

	//users, err := dbManager.Get(55)
	//if err != nil {
	//	log.Fatalf("failed to get date:%v", err)
	//}

	sendUser := User{
		Id:       6070,
		Name:     "Muhammadyusuf",
		Lastname: "Adhamov",
	}
	users, err := dbManager.Update(&sendUser)
	if err != nil {
		log.Fatalf("failed to update date:%v", err)
	}

	PrintUser(users)

	//user, err := dbManager.GetAll()
	//
	//if err != nil {
	//	log.Fatalf("failed to get all date:%v", err)
	//}
	//
	//for _, users := range user {
	//	PrintUser(users)
	//}

	////
	//err = dbManager.Delete(6063)
	//if err != nil {
	//	log.Fatalf("failed to get date:%v", err)
	//}

}

func PrintUser(user *User) {
	fmt.Println("-----------User-----------")
	fmt.Println("ID:", user.Id)
	fmt.Println("Name:", user.Name)
	fmt.Println("Age:", user.Lastname)
	fmt.Println("Created_at:", user.CreatedAt.Format("2006-01-02 15:04"))
	fmt.Println("---------------------")
}
