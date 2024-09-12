package flight

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type updateHandler struct {
	db *sql.DB
}

func NewUpdateHandler(db *sql.DB) updateHandler {
	return updateHandler{db: db}
}

func (h updateHandler) Update(c *gin.Context) {
	idS, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	var f Flight
	err = c.ShouldBindJSON(&f)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	stmt, err := h.db.Prepare("UPDATE flight SET airlineCode=$2,destination=$3,arrival=$4 WHERE id=$1;") // HL

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "can't prepare statment update" + err.Error(),
		})
		return
	}

	if _, err := stmt.Exec(idS, f.AirlineCode, f.Destination, f.Arrival); err != nil { // HL
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error execute update " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "update success",
	})
}
