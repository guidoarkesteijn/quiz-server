package database

import (
	"database/sql"
	"fmt"
	"log"

	//Use _ because it is needed for mysql driver to be imported.
	_ "github.com/go-sql-driver/mysql"
)

type DatabaseService struct {
	database *sql.DB
}

//Connect connect with the questions database.
func Connect(ip string, port string) (*DatabaseService, error) {
	db, err := sql.Open("mysql", "root:RrDLvT70IHAV@tcp("+ip+":"+port+")/quiz")

	if err != nil {
		fmt.Println("err" + err.Error())
		return nil, err
	}

	service := DatabaseService{db}

	return &service, err
}

func (service *DatabaseService) GetQuestion(guid string) {
	Result, errDB := service.database.Query("SELECT guid,text FROM questions WHERE guid='" + guid + "'")

	if errDB != nil {
		fmt.Println("Error" + errDB.Error())
	}

	for Result.Next() {
		var (
			guid string
			text string
		)
		if err := Result.Scan(&guid, &text); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("guid %s question is %s\n", guid, text)
	}
}

func (service *DatabaseService) GetQuestions() {
	Result, errDB := service.database.Query("SELECT guid,text FROM questions")

	if errDB != nil {
		fmt.Println("Error" + errDB.Error())
	}

	for Result.Next() {
		var (
			guid string
			text string
		)
		if err := Result.Scan(&guid, &text); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("guid %s question is %s\n", guid, text)
	}
}
