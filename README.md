favdb: User Favorite DB handle in go (memory, github issue, postgesDB...)
======================

[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/kkdai/fav/master/LICENSE) [![GoDoc](https://godoc.org/github.com/kkdai/fav?status.svg)](https://godoc.org/github.com/kkdai/fav)[![Go](https://github.com/kkdai/photomgr/actions/workflows/go.yml/badge.svg)](https://github.com/kkdai/fav/actions/workflows/go.yml)

fav is a package to help you setup a single entry fo MemoryDB, Github Issue, PostgresSQL DB with user favorite.

The database must be simple and ease to use.

```
User1 -> Fav1, Fav2 ...
User2 -> Fav3
```

Install
--------------

    go get github.com/kkdai/favdb

Usage
---------------------

refer `example/github-issue`for more detail.
`GITHUB_URL` is a string combine with "USE/REPO/GITHU_TOKEN".

```go
var DB favdb.UserFavData

func main() {
    // PostgresSQL connect string
    url := os.Getenv("DATABASE_URL")
    // `GITHUB_URL` is a string combine with "USE/REPO/GITHU_TOKEN".
    gitUrl := os.Getenv("GITHUB_URL")
    if url != "" {
        // Use PostgresSQL as DB.
        DB = favdb.NewPGSql(url)
    } else if gitUrl != "" {
        // Use Github Issue as DB.
        DB = favdb.NewGithubDB(gitUrl)
    } else {
        //Use memory as DB
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
        //User data is not created, create a new one
        DB.Add(newUser)
    } else {
        //Record found, update it
        oldRecords := record.Favorites
        if exist, idx := favdb.InArray(newFavoriteArticle, oldRecords); exist == true {
            oldRecords = favdb.RemoveStringItem(oldRecords, idx)
        } else {
            oldRecords = append(oldRecords, newFavoriteArticle)
        }
        record.Favorites = oldRecords
        DB.Update(record)
    }
}
```

If you want to run it directly, just run

### Github Issue

```
    go install github.com/kkdai/photomgr/example/github_issue
```

Contribute
---------------

Please open up an issue on GitHub before you put a lot efforts on pull request.
The code submitting to PR must be filtered with `gofmt`

License
---------------

This package is licensed under MIT license. See LICENSE for details.
