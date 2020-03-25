package users

import (
	"context"
	"fmt"
	"log"
	"time"
	"errors"
	"database/sql"
	"github.com/repo/rest_api/config"
	"github.com/repo/rest_api/models"
)

const (
	table          = "user"
	layoutDateTime = "2006-01-02 15:04:05"
)

//GetAll
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

// Insert
func Insert(ctx context.Context, user models.Users) error {

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("INSERT INTO %v (name, address, phone, created_at, updated_at) values('%v','%v','%v','%v','%v')", table,
		user.Name,
		user.Address,
		user.Phone,
		time.Now().Format(layoutDateTime),
		time.Now().Format(layoutDateTime))

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}

	return nil
}

// Update
func Update(ctx context.Context, user models.Users) error {
 
    db, err := config.MySQL()
 
    if err != nil {
        log.Fatal("Can't connect to MySQL", err)
    }
 
    queryText := fmt.Sprintf("UPDATE %v set name ='%s', address = '%s', phone = '%s', updated_at = '%v' where id = '%d'",
        table,
        user.Name,
		user.Address,
		user.Phone,
        time.Now().Format(layoutDateTime),
        user.ID,
    )
    fmt.Println(queryText)
 
    _, err = db.ExecContext(ctx, queryText)
 
    if err != nil {
        return err
    }
 
    return nil
}



// Delete
func Delete(ctx context.Context, user models.Users) error {
 
    db, err := config.MySQL()
 
    if err != nil {
        log.Fatal("Can't connect to MySQL", err)
    }
 
    queryText := fmt.Sprintf("DELETE FROM %v where id = '%d'", table, user.ID)
 
    s, err := db.ExecContext(ctx, queryText)
 
    if err != nil && err != sql.ErrNoRows {
        return err
    }
 
    check, err := s.RowsAffected()
    fmt.Println(check)
    if check == 0 {
        return errors.New("id tidak ada")
    }
 
    return nil
}
