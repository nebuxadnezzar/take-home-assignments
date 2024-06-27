package database

import (
	"embed"
	"encoding/csv"
	"fmt"
	"os"
	"regexp"
	"strings"
	"sync"
)

//go:embed data
var memfolder embed.FS

func NewMemesDatabase() (Database, error) {
	recs, index, err := loadMemes()
	if err != nil {
		return nil, err
	}
	return &MemesDatabase{recs: recs, index: index}, nil
}

func (db *MemesDatabase) Find(fldname, fldval string) ([][]string, error) {
	var (
		offsets []int
		ok      bool
	)
	if fldname == BY_FUZZY {
		offsets, ok = fuzzyFind(fldval, db.index)
	} else {
		offsets, ok = find(fldval, db.index)
	}
	if !ok {
		return [][]string{}, nil
	}
	retval := make([][]string, len(offsets))
	for idx := range offsets {
		retval[idx] = db.recs[offsets[idx]]
	}
	return retval, nil
}

func (db *MemesDatabase) Flush() error { return nil }

func find(term string, index map[string][]int) ([]int, bool) {
	offsets, ok := index[term]
	if !ok {
		return []int{}, false
	}
	return offsets, true
}

// Regex search of all index keys takes long to run since it's a linear scan
// plus matching. in real life there are better solutions for search such as
// prefix or ternary trees, but it'll suffice for now.
func fuzzyFind(term string, index map[string][]int) ([]int, bool) {
	retval := []int{}
	pattern := fmt.Sprintf("%s.*", term)
	ch := make(chan string, 1)

	wg := new(sync.WaitGroup)

	for k := range index {
		wg.Add(1)
		go func(k string) {
			if matches(pattern, k) {
				ch <- k
			}
			wg.Done()
		}(k)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	for key := range ch {
		retval = append(retval, index[key]...)
	}
	return retval, len(retval) > 0
}

func loadMemes() ([][]string, map[string][]int, error) {
	f, _ := memfolder.Open("data/memes.csv")
	defer f.Close()
	reader := csv.NewReader(f)
	index := map[string][]int{}
	recs, err := reader.ReadAll()
	if err != nil {
		return nil, nil, err
	}
	for idx, rec := range recs {
		addtoindex(index, strings.Split(strings.Trim(strings.ToLower(rec[2]), " \n\r"), ` `), idx)
		addtoindex(index, strings.Split(strings.Trim(strings.ToLower(rec[6]), " \n\r"), ` `), idx)
	}
	return recs, index, nil
}

func addtoindex(index map[string][]int, keys []string, id int) {
	for _, k := range keys {
		if k == "" {
			continue
		}

		if v, ok := index[k]; ok {
			index[k] = append(v, id)
		} else {
			index[k] = []int{id}
		}
	}
}

func matches(pattern string, data string) bool {
	matched, err := regexp.MatchString(pattern, data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Regex parsing error: %v\n", err)
		return false
	} else {
		return matched
	}
}
