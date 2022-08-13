package users

import "time"

var UserList *[]User

func GetUsers() {

	Time1, _ := time.Parse("2006-01-02", "2022-11-22")
	Time2, _ := time.Parse("2006-01-02", "2023-05-22")
	Time3, _ := time.Parse("2006-01-02", "2023-10-22")

	users := []User{
		{
			UserId:      1,
			LocationIds: []int{5010},
			BeforeDate:  Time1,
		},
		{
			UserId:      2,
			LocationIds: []int{5007},
			BeforeDate:  Time2,
		},
		{
			UserId:      3,
			LocationIds: []int{5001, 5007},
			BeforeDate:  Time3,
		},
		{
			UserId:      4,
			LocationIds: []int{5001},
			BeforeDate:  Time3,
		},
	}

	UserList = &users

}
