package ginfundamentals

import (
	"database/sql"
	"log"
	"net/http"

	dbutils "github.com/Dav16Akin/go-dictionary/railAPI/dbUtils"
	_ "github.com/mattn/go-sqlite3"

	"github.com/gin-gonic/gin"
)

var DB *sql.DB

type StationResource struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	OpeningTime string `json:"opening_time"`
	ClosingTime string `json:"closing_time"`
}

func GetStations(c *gin.Context) {
	rows, err := DB.Query("SELECT ID, NAME, CAST(OPENING_TIME as CHAR), CAST(CLOSING_TIME as CHAR) FROM station")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error" : err.Error()})
		return
	}

	defer rows.Close()

	stations := []StationResource{}

	for rows.Next() {
		var station StationResource

		err := rows.Scan(
			&station.ID,
			&station.Name,
			&station.OpeningTime,
			&station.ClosingTime,
		)

		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				gin.H{"error": "Error scanning station"},
			)
			return
		}

		stations = append(stations, station)
	}

	if err := rows.Err(); err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": "Row iteration error"},
		)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"stations": stations,
	})
}
func GetStation(c *gin.Context) {
	var station StationResource

	id := c.Param("station_id")
	err := DB.QueryRow("select ID, NAME, CAST(OPENING_TIME as CHAR), CAST(CLOSING_TIME as CHAR) from station where id=?", id).Scan(
		&station.ID, &station.Name, &station.OpeningTime, &station.ClosingTime,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting data from the database"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"result": station})
		return
	}
}

func CreateStation(c *gin.Context) {
	var station StationResource

	if err := c.BindJSON(&station); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	statement, err := DB.Prepare("insert into station (NAME, OPENING_TIME, CLOSING_TIME) values (?, ?, ?)")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to prepare statement"})
		return
	}

	defer statement.Close()

	result, err := statement.Exec(station.Name, station.OpeningTime, station.ClosingTime)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to insert station",
		})
		return
	}

	newID, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get inserted ID",
		})
		return
	}

	station.ID = int(newID)

	c.JSON(http.StatusCreated, gin.H{
		"result": station,
	})
}

func RemoveStation(c *gin.Context) {
	id := c.Param("station_id")
	statement, err := DB.Prepare("delete from station where id=?")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	defer statement.Close()

	result, err := statement.Exec(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting station"})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "databbase error"})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "station not found"})
		return
	}

	c.Status(http.StatusNoContent)
}

func RunGinAPI() {
	var err error
	DB, err = sql.Open("sqlite3", "./railapi.db")
	if err != nil {
		log.Printf("Error Opening Database : %v", err)
	}

	dbutils.Initialize(DB)

	router := gin.Default()

	router.GET("/v1/stations", GetStations)
	router.GET("/v1/stations/:station_id", GetStation)
	router.POST("/v1/stations", CreateStation)
	router.DELETE("/v1/stations/:station_id", RemoveStation)

	router.Run(":8000")
}
