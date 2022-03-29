package boot

import (
	"fmt"
	"log"
	"newssync/src"
	"newssync/src/impl"
	"os"

	"github.com/joho/godotenv"
)

// RetrieveSources will yield all configured sources we want to source data from to fan out to our sinks
func RetrieveSources() []src.Source {
	hackernews := &impl.HackerNewsSource{
		Url:          src.HACKER_NEWS_URL,
		Filemetadata: "metadata/hackernews.json",
		DatabaseLoc:  "./databases/source.db",
	}

	return []src.Source{hackernews}
}

// RetrieveSinks retrieves all valid upload sinks to post new content to, such as Twitter or Reddit
func RetrieveSinks() []src.Sink {
	twitter := &impl.TwitterSink{
		Uri:  "https://api.twitter.com/2/tweets",
		Auth: src.Credentials{},
	}

	return []src.Sink{twitter}
}

func GenerateDatabases(pathPrefix string, databaseFiles ...string) []string {
	formatPath := func(prefixPath, file string) string {
		return fmt.Sprintf("%s/%s", prefixPath, file)
	}

	var createdDatabaseFiles []string
	for _, databaseFile := range databaseFiles {

		if _, err := os.Stat(formatPath(pathPrefix, databaseFile)); err != nil {
			log.Println(fmt.Sprintf("%s does not exist.. creating..", formatPath(pathPrefix, databaseFile)))
			createdDatabaseFiles = append(createdDatabaseFiles, formatPath(pathPrefix, databaseFile))
			_, err := generateDatabaseFile(pathPrefix, formatPath(pathPrefix, databaseFile))
			checkErr(err)
			continue
		}

		log.Println(fmt.Sprintf("%s exists already... skipping", databaseFile))
	}

	return createdDatabaseFiles
}

func generateDatabaseFile(directory, file string) (*os.File, error) {
	err := os.MkdirAll(directory, 0700)
	checkErr(err)

	log.Println("Writing " + file)
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE, 0700)
	f.Close()

	checkErr(err)
	log.Println(file + " database created")
	return f, nil
}

func LoadEnvironment() {
	godotenv.Load()
	log.Println("Environment variables have been loaded correctly...")
}

func checkErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}
