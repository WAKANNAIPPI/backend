package main

import "src/database"

func main() {

	db := database.DBconnect()

	u := database.User{}

	u.Id = "testUser"
	u.Name = "テストユーザー"
	u.Password = "password"

	database.CreateUser(u)

	item := make([]database.User_item, 4)

	item[0].Iid = "red"
	item[0].Uid = "testUser"
	item[1].Iid = "blue"
	item[1].Uid = "testUser"
	item[2].Iid = "yellow"
	item[2].Uid = "testUser"
	item[3].Iid = "rare"
	item[3].Uid = "testUser"

	for i := 0; i < len(item); i++ {
		db.Debug().Create(&item[i])
	}
}
