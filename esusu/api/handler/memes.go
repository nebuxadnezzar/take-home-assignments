package handler

import (
	"encoding/json"
	"fmt"
	db "maas/api/database"
	"maas/api/model"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func SendMeme(c echo.Context) error {
	objmap := c.Get("custom").(map[string]interface{})
	memdb := objmap[db.MEMESDB_KEY].(db.Database)
	usersdb := objmap[db.USERSDB_KEY].(db.Database)

	mr := model.MemeRequest{}
	err := c.Bind(&mr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.NewErrorMessage(
			http.StatusBadRequest,
			"bind: bad request data",
			err,
		))
	}
	c.Logger().Printf("PARAMS %#v", mr)
	if len(mr.Query) == 0 {
		return c.JSON(http.StatusBadRequest, model.NewErrorMessage(
			http.StatusBadRequest,
			"query parameter is missing",
			fmt.Errorf("query is a mandatory parameter"),
		))
	}

	authToken := extractToken(c.Request().Header)
	uid, ok := usersdb.(*db.UsersDatabase).FindUserIdByToken(authToken)
	if !ok {
		return c.JSON(http.StatusUnauthorized, model.NewErrorMessage(
			http.StatusUnauthorized,
			"invalid user token",
			fmt.Errorf("invalid user token used: %s", authToken),
		))
	}

	fld := `__`
	if mr.Fuzzy {
		fld = db.BY_FUZZY
	}
	recs, _ := memdb.Find(fld, strings.ToLower(mr.Query))
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	c.Response().WriteHeader(http.StatusOK)
	enc := json.NewEncoder(c.Response())
	for idx := range recs {
		mm := populateMeme(recs[idx])

		if err := enc.Encode(mm); err != nil {
			return c.JSON(http.StatusInternalServerError, model.NewErrorMessage(
				http.StatusInternalServerError,
				"meme conversion failed",
				err,
			))
		}
		c.Response().Flush()
	}

	usersdb.Update(uid, db.USER_STATS_FLD, c.Path())
	/*
		if recs, err := usersdb.Find("token", authToken); err == nil {
			for i, r := range recs {
				fmt.Printf("%d --> %v\n", i, r)
			}
		} else {
			fmt.Printf("FAILED TO GET USER: %v\n", err)
		}
	*/
	return nil
}

func populateMeme(rec []string) model.Meme {
	//ID,URL,Meme_Name,Meme_Page_URL,MD5,File_Size,Alternate_Text
	return model.Meme{
		ID:       rec[0],
		URL:      rec[1],
		Name:     rec[2],
		Page:     rec[3],
		MD5:      rec[4],
		FileSize: rec[5],
		Text:     rec[6],
	}
}

func extractToken(hdr http.Header) string {
	ss := hdr.Values("X-Token")
	if ss == nil || len(ss) == 0 {
		return ``
	}
	return ss[0]
}
