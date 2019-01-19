package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	//Use _ because it is needed for mysql driver to be imported.
	_ "github.com/go-sql-driver/mysql"
	model "github.com/project-quiz/quiz-go-model"
)

type DatabaseService struct {
	database *sql.DB
}

//New connect with the questions database.
func New() (*DatabaseService, error) {
	ip, ipSuccess := os.LookupEnv("DATABASE_IP")
	port, portSuccess := os.LookupEnv("DATABASE_PORT")
	username, usernameSuccess := os.LookupEnv("DATABASE_USERNAME")
	password, passwordSuccess := os.LookupEnv("DATABASE_PASSWORD")

	success := ipSuccess && portSuccess && usernameSuccess && passwordSuccess

	if !success {
		return nil, errors.New("not enough database values available")
	}

	connection := username + ":" + password + "@tcp(" + ip + ":" + port + ")/quiz"
	db, err := sql.Open("mysql", connection)

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

//GetQuestions Get all questions from the database.
func (service *DatabaseService) GetQuestions() (questions []model.Question, err error) {
	Result, err := service.database.Query("SELECT guid,text FROM questions")

	if err != nil {
		fmt.Println("Question error: " + err.Error())
	}

	q := []model.Question{}

	for Result.Next() {
		var (
			guid string
			text string
		)
		if err := Result.Scan(&guid, &text); err != nil {
			log.Fatal(err)
		}

		a, answerErr := service.GetAnswers(guid)

		if answerErr != nil {
			fmt.Println("Answer err for question " + guid + ": " + answerErr.Error())
		}

		question := model.Question{Guid: guid, Question: text, Answers: a}
		q = append(q, question)
	}

	return q, err
}

func (service *DatabaseService) GetAnswers(questionGuid string) (answers []*model.Answer, err error) {
	Result, err := service.database.Query("SELECT guid,answer FROM answers WHERE question='" + questionGuid + "'")

	a := []*model.Answer{}

	for Result.Next() {
		var (
			guid   string
			answer string
		)
		if err := Result.Scan(&guid, &answer); err != nil {
			log.Fatal(err)
		}
		obj := model.Answer{Guid: guid, Text: answer}

		a = append(a, &obj)
	}
	return a, err
}
