package favdb

import (
	"log"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

type PGSql struct {
	Db   *pg.DB
	data UserFavorite
}

func NewPGSql(url string) *PGSql {
	options, _ := pg.ParseURL(url)
	db := pg.Connect(options)

	err := createSchema(db)
	if err != nil {
		panic(err)
	}

	return &PGSql{
		Db: db,
	}
}

func (u *PGSql) Add(user UserFavorite) {
	_, err := u.Db.Model(user).Insert()
	if err != nil {
		log.Println(err)
	}
}

func (u *PGSql) Get(uid string) (result *UserFavorite, err error) {
	log.Println("***Get Fav uUID=", uid)
	err = u.Db.Model(&u.data).
		Where("user_id = ?", uid).
		Select()
	if err != nil {
		log.Println("Get err:", err)
		return nil, err
	}
	log.Println("UserFavorite DB result= ", u.data)
	return &u.data, nil
}

// ShowAll: Print all result.
func (u *PGSql) ShowAll() (result []UserFavorite, err error) {
	log.Println("***Get All DB")
	ret := []UserFavorite{}
	err = u.Db.Model(&ret).Select()
	if err != nil {
		log.Println(err)
	}
	log.Println("***Start server all users =", ret)
	if err != nil {
		log.Println("open file error:", err)
	}

	return ret, nil
}

func (u *PGSql) Update(user *UserFavorite) (err error) {
	log.Println("***Update Fav User=", u)
	_, err = u.Db.Model(user).
		Set("favorites = ?", user.Favorites).
		Where("user_id = ?", user.UserId).
		Update()
	if err != nil {
		log.Println("update err:", err)
	}
	// Update success, replace to memory.
	u.data = *user
	return nil
}

func createSchema(db *pg.DB) error {
	models := []interface{}{
		(*UserFavorite)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true})
		if err != nil {
			log.Println("createSchema err:", err)
			return err
		}
	}
	return nil
}
