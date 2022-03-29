package util

import (
	"database/sql"
	"fmt"
	"log"
	"newssync/src"
	"time"
)

type SourceEntry struct {
	postId     string
	postTitle  string
	postSource string
	postData   uint64
}

func CreateTable(db *sql.DB) {
	createSourceTable := `CREATE TABLE IF NOT EXISTS source (
		"postId" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"postTitle" TEXT,
		"postSource" TEXT,
		"postDate" REAL
	  );`

	log.Println("Creating source table...")
	statement, err := db.Prepare(createSourceTable)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
	log.Println("Source table created")
}

func CheckIfPostHasBeenPosted(db *sql.DB, postName string) bool {
	checkExistenceQuery := fmt.Sprintf(`SELECT * FROM source WHERE
		"postTitle" IS "%s"
	`, postName)

	result, err := db.Query(checkExistenceQuery)
	defer result.Close()

	for result.Next() {
		var postId int
		var postDate time.Time

		result.Scan(&postId, &postDate)
		log.Println(fmt.Sprintf("Error, post %d already exists %s", postId, postDate.String()))
		return true
	}

	if err != nil {
		log.Panic(err)
	}

	return false
}

func StoreNewPost(db *sql.DB, h *src.History) error {
	query := "INSERT INTO source (postTitle, postSource, postDate) VALUES ($1, $2, $3)"
	_, err := db.Exec(query, h.LastArticleTitle, h.Source, time.Now().Unix())
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
