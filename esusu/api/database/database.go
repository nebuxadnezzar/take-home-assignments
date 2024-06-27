package database

import (
	"maas/api/model"
	"sync"
)

const (
	MEMESDB_KEY    = "memesdb"
	USERSDB_KEY    = "usersdb"
	BY_TOKEN       = "token"
	BY_NAME        = "name"
	BY_ID          = "id"
	BY_FUZZY       = "fuzzy"
	TOKEN_CNT_FLD  = "tokenCount"
	USER_STATS_FLD = "userStats"
	USERS_DB_FILE  = "users_flushed.json"
)

type (
	Database interface {
		Find(fldname, fldval string) ([][]string, error)
		Flush() error
		Update(id, fldname, fldval string) error
	}
	MemesDatabase struct {
		index map[string][]int
		recs  [][]string
		Database
	}
	UsersDatabase struct {
		mutex       sync.RWMutex
		users       map[string]*model.User
		userbytoken map[string]int
		filepath    string
	}
)
