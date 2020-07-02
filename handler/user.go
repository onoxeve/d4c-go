package handler

import (
	"d4c/main/model"
	"github.com/labstack/echo/v4"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

func (h *Handler) Sighup(c echo.Context) (err error) {
	// Bind
	u := &model.User{ID: bson.NewObjectId()}
	if err = c.Bind(u); err != nil {
		return
	}

	// Validate
	if u.Email == "" || u.Password == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid email or password"}
	}

	// Create
	db := h.DB.Clone()
	defer db.Close()
	if err = db.DB("d4c").C("users").Insert(u); err != nil {
		return
	}

	return c.JSON(http.StatusCreated, u)
}
