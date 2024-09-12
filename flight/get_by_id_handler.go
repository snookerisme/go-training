package flight

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type getByIDHandler struct {
	db *sql.DB
}

func NewGetByIDHandler(db *sql.DB) getByIDHandler {
	return getByIDHandler{db: db}
}

func (h getByIDHandler) GetByID(c *gin.Context) {
	idS, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	stmt, err := h.db.Prepare("SELECT id, number, airlineCode, destination, arrival FROM flight where id=$1") // HL
	if err != nil {
		log.Fatal("can't prepare query one row statment", err)
	}

	row := stmt.QueryRow(idS) // HL
	var id, number int
	var airlineCode, destination, arrival string

	err = row.Scan(&id, &number, &airlineCode, &destination, &arrival) // HL
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	data := Flight{id, number, airlineCode, destination, arrival}

	c.JSON(http.StatusOK, data)
}
