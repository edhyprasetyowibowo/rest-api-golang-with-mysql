package users

import (
	"context"
	"fmt"
	"log"
	"time"
	"github.com/repo/rest_api/config"
	"github.com/repo/rest_api/models"
)

const (
	table          = "user"
	layoutDateTime = "2006-01-02 15:04:05"
)


func GetAll(ctx context.Context) ([]models.Users, error) {

	var Users []models.Users

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v Order By id DESC", table)

	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var users models.Users
		var createdAt, updatedAt string

		if err = rowQuery.Scan(
			&users.ID,
			&users.Name,
			&users.Address,
			&users.Phone,
			&createdAt,
			&updatedAt); err != nil {
			return nil, err
		}

		//  Change format string to datetime for created_at and updated_at
		users.CreatedAt, err = time.Parse(layoutDateTime, createdAt)

		if err != nil {
			log.Fatal(err)
		}

		users.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)

		if err != nil {
			log.Fatal(err)
		}

		Users = append(Users, users)
	}

	return Users, nil
}
