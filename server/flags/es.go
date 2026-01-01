package flags

import (
	"bufio"
	"fmt"
	"os"
	"server/model/elasticsearch"
	"server/service"
)

func Elasticsearch() error {
	esService := service.ServiceGroupApp.EsService

	indexExists, err := esService.IndexExists(elasticsearch.ArticleIndex())
	if err != nil {
		return err
	}

	if indexExists {
		fmt.Println("The index already exists. Do you want to delete it? (y/n)")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()

		switch input {
		case "y":
			fmt.Println("Deleting the index...")
			if err := esService.IndexDelete(elasticsearch.ArticleIndex()); err != nil {
				return err
			}
		case "n":
			fmt.Println("Skipping the index deletion.")
			os.Exit(0)
		default:
			fmt.Println("Invalid input. Please enter y to delete the index or n to skip the deletion.")
			return Elasticsearch()
		}
	}

	return esService.IndexCreate(elasticsearch.ArticleIndex(), elasticsearch.ArticleMapping())
}
