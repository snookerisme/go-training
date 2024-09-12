package flight

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type createHandler struct {
	db *sql.DB
}

func NewCreateHandler(db *sql.DB) createHandler {
	return createHandler{db: db}
}

func (h createHandler) Create(c *gin.Context) {
	var f Flight
	err := c.ShouldBindJSON(&f)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	q := "INSERT INTO flight (id,number, airlineCode, destination, arrival) values ($1, $2, $3, $4,$5) RETURNING id"
	row := h.db.QueryRow(q, f.ID, f.Number, f.AirlineCode, f.Destination, f.Arrival)
	var id int
	err = row.Scan(&id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "add success",
	})
}
