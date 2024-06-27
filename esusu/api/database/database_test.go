package database

import (
	"fmt"
	"os"
	"testing"
)

func TestGetMemes(t *testing.T) {
	recs, index, err := loadMemes()
	if len(recs) == 0 || len(index) == 0 {
		t.Error("empty records")
	}
	if err != nil {
		t.Errorf("%v", err)
	}
}

func TestMemDbFind(t *testing.T) {
	db, _ := NewMemesDatabase()
	recs, _ := db.Find("dummy", "destroyed")
	if len(recs) == 0 {
		t.Error(`"destroyed" had to be found but wasn't`)
	}
	for idx := range recs {
		fmt.Printf("%v\n", recs[idx])
	}
}

func TestGetUsers(t *testing.T) {
	users, err := loadDefaultUsers()
	if err != nil {
		t.Errorf("failed to load default users: %v\n", err)
	}
	fmt.Printf("%#v %v\n", users, err)
}

func TestUpdateAndSaveUsers(t *testing.T) {
	db, err := NewUsersDb(``)
	if err != nil {
		t.Errorf("failed to load default users: %v\n", err)
	}
	if err := db.Update("10", TOKEN_CNT_FLD, "200"); err != nil {
		t.Errorf("failed to update %v", err)
	}
	if db.Flush() != nil {
		t.Error("failed to flush users db")
	}
	os.Remove(`./` + USERS_DB_FILE)
}
