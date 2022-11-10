package main

import (
	"github.com/bxcodec/faker/v4"
)

func InsertFakeDate(db *DBManager) error {
	var user User
	for i := 1; i <= 10; i++ {
		user.Name = faker.FirstName()
		user.Lastname = faker.LastName()
		_, err := db.Create(&user)
		if err != nil {
			return err
		}
		//if i%100 == 0 {
		//	fmt.Println(i)
		//}
	}
	return nil
}
