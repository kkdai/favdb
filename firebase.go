package favdb

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"google.golang.org/api/option"
)

const (
	dbPath = "favDB"
)

type FireDB struct {
	*db.Client
}

type FirebaseDB struct {
	CTX    context.Context
	UID    string
	Client FireDB
}

func NewDB(url string, json string, uid string) *FirebaseDB {
	ctx := context.Background()
	opt := option.WithCredentialsJSON([]byte(json))
	config := &firebase.Config{DatabaseURL: url}
	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v", err)
		return nil
	}
	client, err := app.Database(ctx)
	if err != nil {
		log.Fatalf("error initializing database: %v", err)
		return nil
	}

	return &FirebaseDB{
		CTX:    ctx,
		UID:    uid,
		Client: FireDB{client},
	}
}

func (u *FirebaseDB) Add(user UserFavorite) {

}

func (u *FirebaseDB) Get(uid string) (result *UserFavorite, err error) {
	return nil, nil
}

// ShowAll: Print all result.
func (u *FirebaseDB) ShowAll() (result []UserFavorite, err error) {
	log.Println("***Get All DB- Not support now.")
	err = u.Client.NewRef(dbPath).Get(u.CTX, &result)
	if err != nil {
		fmt.Println("load memory failed, ", err)
	}

	return nil, nil
}

func (u *FirebaseDB) Update(user *UserFavorite) (err error) {
	return nil
}

func (u *FirebaseDB) saveIssue(title, body string) error {
	return nil
}

func (u *FirebaseDB) getIssue(title string) (string, int, error) {
	return "", 0, nil
}

func (u *FirebaseDB) updateIssue(number int, title string, updatedCnt string) error {
	return nil
}
