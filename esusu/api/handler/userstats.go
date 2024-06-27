package handler

import (
	_ "encoding/json"
	"fmt"
	db "maas/api/database"
	"maas/api/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func SendStats(c echo.Context) error {

	objmap := c.Get("custom").(map[string]interface{})
	usersdb := objmap[db.USERSDB_KEY].(db.Database)

	ur := model.UserRequest{}
	err := c.Bind(&ur)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.NewErrorMessage(
			http.StatusBadRequest,
			"bind: bad request data",
			err,
		))
	}
	c.Logger().Printf("PARAMS %#v", ur)
	id, authToken := ur.ID, ur.AuthToken
	if id < 1 && authToken == `` {
		return c.JSON(http.StatusBadRequest, model.NewErrorMessage(
			http.StatusBadRequest,
			"bad request data, no id or token provided",
			fmt.Errorf("bad request data, no id or token provided"),
		))
	}
	var fld, val string

	if id > 0 {
		fld = db.BY_ID
		val = strconv.Itoa(id)
	} else {
		fld = db.BY_TOKEN
		val = authToken
	}
	recs, err := usersdb.Find(fld, val)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.NewErrorMessage(
			http.StatusInternalServerError,
			fmt.Sprintf("failed to find user for field %s and value %s", fld, val),
			err,
		))
	}

	for _, r := range recs {
		c.Response().Write([]byte(r[0]))
	}

	return nil
}
