package handlers

import (
	"net/http"
	"go-training/internal/cohorts"
	"go-training/internal/upsert"

	"github.com/labstack/echo/v4"
)

func Home(c echo.Context) error {
	batch := cohorts.Batch{}

	err := c.Bind(&batch)

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	insertUser := upsert.GetInsertUsers(batch)
	deleteUser := upsert.GetDeleteUsers(batch)

	insertRequest := upsert.SendUpsertRequest("upsert", insertUser)
	deleteRequest := upsert.SendUpsertRequest("attribute/delete", deleteUser)

	return c.JSON(http.StatusOK, upsert.CombinedResult{
		Insert: insertRequest,
		Delete: deleteRequest,
	})
}
