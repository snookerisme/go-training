package flight

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func NewGetAllHandler(db *sql.DB) handler {
	return handler{db: db}
}

func (h handler) GetAll(c *gin.Context) {

	stmt, err := h.db.Prepare("SELECT id, number, airlineCode, destination, arrival FROM flight") // HL
	if err != nil {
		log.Fatal("can't prepare query all todos statment", err)
	}

	f := []Flight{}

	rows, err := stmt.Query() // HL
	if err != nil {
		log.Fatal("can't query all todos", err)
	}
	for rows.Next() { // HL
		var id, number int
		var airlineCode, destination, arrival string
		err := rows.Scan(&id, &number, &airlineCode, &destination, &arrival) // HL
		if err != nil {
			log.Fatal("can't Scan row into variable", err)
		}
		f = append(f, Flight{id, number, airlineCode, destination, arrival})
	} // HL

	c.JSON(http.StatusOK, gin.H{
		"data": f,
	})
}
