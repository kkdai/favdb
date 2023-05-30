package main

import (
	"log"
	"os"

	"github.com/kkdai/favdb"
)

var DB favdb.UserFavData

func main() {
	gitUrl := os.Getenv("GITHUB_URL")
	if gitUrl != "" {
		// Use Github Issue as DB.
		DB = favdb.NewGithubDB(gitUrl)
	} else {
		DB = favdb.NewMemDB()
	}
	addBookmarkArticle("title1", "Fav1")
	addBookmarkArticle("title1", "Fav2")
}

func addBookmarkArticle(user, fav string) {
	newFavoriteArticle := fav
	newUser := favdb.UserFavorite{
		UserId:    user,
		Favorites: []string{newFavoriteArticle},
	}
	if record, err := DB.Get(user); err != nil {
		log.Println("User data is not created, create a new one")
		DB.Add(newUser)
		log.Println(newFavoriteArticle, "Add user/fav")
	} else {
		log.Println("Record found, update it", record)
		oldRecords := record.Favorites

		if exist, idx := favdb.InArray(newFavoriteArticle, oldRecords); exist == true {
			log.Println(newFavoriteArticle, "Del fav")
			oldRecords = favdb.RemoveStringItem(oldRecords, idx)
		} else {
			log.Println(newFavoriteArticle, "Add fav")
			oldRecords = append(oldRecords, newFavoriteArticle)
		}
		record.Favorites = oldRecords
		DB.Update(record)
	}
}
