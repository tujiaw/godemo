package search
import (
	"log"
	"fmt"
)

type Result struct {
	Field string
	Content string
}

type Matcher interface {
	Search(feed *Feed, searchTerm string)([]*Result, error)
}

func Match(matcher Matcher, feed *Feed, searchTerm string, results chan<- *Result) {
	searchResults, err := matcher.Search(feed, searchTerm)
	if err != nil {
		log.Println(err)
		return
	}

	for _, result := range searchResults {
		results <- result
	}
}

func Display(results chan *Result) {	
	i := 0
	for result := range results {
		fmt.Printf("display[%d] %s:\n%s\n\n", i, result.Field, result.Content)
		i += 1
	}
}

