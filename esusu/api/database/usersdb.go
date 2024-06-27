package database

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"maas/api/model"
	"os"
	"strconv"
)

//go:embed data/users.json
var usersjson []byte

func NewUsersDb(usersfile string) (Database, error) {
	var users map[string]*model.User
	var err error
	if usersfile == `` {
		users, err = loadDefaultUsers()
	} else {
		users, err = loadUsers(usersfile)
	}

	if err != nil {
		return nil, err
	}

	return &UsersDatabase{
		filepath:    usersfile,
		users:       users,
		userbytoken: createUserByToken(users),
	}, nil
}

func (db *UsersDatabase) FindUserIdByToken(token string) (string, bool) {
	id, ok := db.userbytoken[token]
	return strconv.Itoa(id), ok
}

func (db *UsersDatabase) Find(fldname, fldval string) ([][]string, error) {
	retval := [][]string{}
	var rec *model.User
	var ok = false
	var id string
	switch fldname {
	case BY_TOKEN:
		id, ok = db.FindUserIdByToken(fldval)
		if ok {
			rec = db.users[id]
		}
	case BY_ID:
		rec, ok = db.users[fldval]
	default:
		return nil, fmt.Errorf("invalid field name: %s", fldname)
	}

	if ok {
		if u, err := json.Marshal(rec); err == nil {
			retval = append(retval, []string{string(u)})
		} else {
			return retval, err
		}
	}
	return retval, nil
}

func (db *UsersDatabase) Flush() error {
	b, err := json.MarshalIndent(db.users, "", " ")
	if err != nil {
		return err
	}
	path := `./` + USERS_DB_FILE
	fmt.Printf("FLUSHING DB to %s\n", path)
	if err := os.WriteFile(path, b, os.ModePerm); err != nil {
		return err
	}
	return nil
}

func (db *UsersDatabase) Update(id, fldname, fldval string) error {
	db.mutex.RLock()
	defer db.mutex.RUnlock()
	u, ok := db.users[id]
	if !ok {
		return fmt.Errorf("user id %s doesn't exist", id)
	}
	switch fldname {
	case TOKEN_CNT_FLD:
		return db.updateTokenCount(u, fldval)
	case USER_STATS_FLD:
		db.updateUserStats(u, fldval)
	default:
		return fmt.Errorf("there is no update procedure for %s field", fldname)
	}
	return nil
}

func (db *UsersDatabase) updateUserStats(user *model.User, uri string) {
	if user.Hits == nil {
		user.Hits = make(map[string]int)
	}
	user.Hits[uri] += 1
	user.TokenCount -= 1
}

func (db *UsersDatabase) updateTokenCount(user *model.User, val string) error {
	if n, err := strconv.Atoi(val); err == nil {
		user.TokenCount += n
	} else {
		return err
	}
	return nil
}

func createUserByToken(users map[string]*model.User) map[string]int {
	ubt := map[string]int{}
	for _, user := range users {
		ubt[user.AuthToken] = user.ID
	}
	return ubt
}

func loadUsers(path string) (map[string]*model.User, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return bytesToUserMap(b)
}

func loadDefaultUsers() (map[string]*model.User, error) {
	return bytesToUserMap(usersjson)
}

func bytesToUserMap(b []byte) (map[string]*model.User, error) {
	var users map[string]*model.User
	if err := json.Unmarshal(b, &users); err != nil {
		return nil, err
	}
	return users, nil
}
