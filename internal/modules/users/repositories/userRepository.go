package user_repositories

import (
	"fmt"

	user_model "mehrang.ir/school/internal/modules/users/models"
	"mehrang.ir/school/pkg/database"
)

func GetUsers(page int, pageSize int) (u []user_model.User, err error) {
	query := "SELECT id,name,created_at FROM users"
	query += " ORDER BY id DESC"
	if pageSize >= 0 {
		query += " LIMIT ?"
	}
	if page >= 1 {
		query += " OFFSET ?"
	}
	db, err := database.ConnectDB()
	if err != nil {
		return u, err
	}
	defer db.Close()
	var offset = ((page - 1) * pageSize)
	fmt.Printf("\n Q: %s ) - O: %d\n", query, offset)
	rows, err := db.Query(query, pageSize, offset)
	if err != nil {
		return u, err
	}
	var users []user_model.User
	defer rows.Close()
	for rows.Next() {
		var user user_model.User
		err = rows.Scan(&user.Id, &user.Name, &user.Created_at)
		if err != nil {
			return u, err
		}
		users = append(users, user)
	}
	return users, nil
}
