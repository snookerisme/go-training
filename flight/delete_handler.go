package flight

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type deleteHandler struct {
	db *sql.DB
}

func NewDeleteHandler(db *sql.DB) deleteHandler {
	return deleteHandler{db: db}
}

func (h deleteHandler) Delete(c *gin.Context) {
	idS, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	stmt, err := h.db.Prepare("DELETE FROM flight WHERE id = $1")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "can't prepare statment delete" + err.Error(),
		})
		return
	}

	if _, err := stmt.Exec(idS); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error execute delete " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "delete success",
	})
}
