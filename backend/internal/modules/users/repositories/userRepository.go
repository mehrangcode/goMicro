package user_repositories

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	user_model "mehrang.ir/school/internal/modules/users/models"
	"mehrang.ir/school/pkg/database"
	"mehrang.ir/school/utils"
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
func GetUserById(userId string) (user user_model.User, err error) {
	query := "SELECT id,name,email,created_at FROM users WHERE id=$1"
	db, err := database.ConnectDB()
	if err != nil {
		return user, err
	}
	defer db.Close()
	err = db.QueryRow(query, userId).Scan(&user.Id, &user.Name, &user.Email, &user.Created_at)
	if err != nil {
		return user, errors.New("user not found")
	}
	return user, nil
}

func CreateUser(w http.ResponseWriter, r *http.Request) (id *int, err error) {
	db, err := database.ConnectDB()
	if err != nil {
		return id, err
	}
	defer db.Close()
	var query = `INSERT INTO users (name,email,password,created_at) VALUES($1,$2,$3,$4) RETURNING id`
	var user user_model.User
	utils.ReadJson(w, r, &user)
	fmt.Println("\n USER ", user)
	row := db.QueryRow(query, user.Name, user.Email, user.Password, time.Now())
	err = row.Scan(&id)
	if err != nil {
		return id, err
	}
	return id, nil

}

func UpdateUser(w http.ResponseWriter, r *http.Request) error {
	userID := chi.URLParam(r, "userId")
	db, err := database.ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()
	var exists bool
	err = db.QueryRow("SELECT EXISTS (SELECT 1 FROM users WHERE id = $1)", userID).Scan(&exists)
	if err != nil {
		return err
	} else if !exists {
		return errors.New("user does not exists")
	} else {
		var query = `UPDATE users SET name=$1, email=$2 WHERE id=$3`
		var user user_model.User
		err = utils.ReadJson(w, r, &user)
		if err != nil {
			return err
		}
		_, err := db.Exec(query, user.Name, user.Email, userID)
		if err != nil {
			return err
		}
	}
	return nil
}

func DeleteUser(userId string) error {
	db, err := database.ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()
	var exists bool
	err = db.QueryRow("SELECT EXISTS (SELECT 1 FROM users WHERE id = $1)", userId).Scan(&exists)
	if err != nil {
		return err
	} else if !exists {
		return errors.New("user does not exists")
	} else {
		var query = `DELETE FROM users WHERE id=$1;`
		_, err := db.Exec(query, userId)
		if err != nil {
			return err
		}
	}
	return nil
}
