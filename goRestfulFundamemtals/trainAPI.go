package gorestfulfundamemtals

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/emicklei/go-restful"
	_ "github.com/mattn/go-sqlite3"

	dbutils "github.com/Dav16Akin/go-dictionary/railAPI/dbUtils"
)

var DB *sql.DB

type Train struct{}

type TrainResource struct {
	ID              int    `json:"id"`
	DriverName      string `json:"driver_name"`
	OperatingStatus bool   `json:"operating_status"`
}

type StationResource struct {
	ID          int
	Name        string
	OpeningTime time.Time
	ClosingTime time.Time
}

type ScheduleResource struct {
	ID          int
	TrainID     int
	StationID   int
	ArrivalTime time.Time
}

func (t *Train) Register(container *restful.Container) {
	ws := new(restful.WebService)

	/* with this we only entertain content-type application/json , if any other type is passed we will get a not supported media type error */
	ws.Path("/v1/trains").Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)

	ws.Route(ws.GET("/{train-id}").To(t.getTrain))
	ws.Route(ws.POST("").To(t.createTrain))
	ws.Route(ws.DELETE("/{train-id}").To(t.removeTrain))
	container.Add(ws)
}

// GET http://localhost:8000/v1/trains/1
func (t *Train) getTrain(req *restful.Request, resp *restful.Response) {
	id := req.PathParameter("train-id")

	var train TrainResource

	err := DB.QueryRow("select ID, DRIVER_NAME, OPERATING_STATUS FROM train where id=?", id).Scan(&train.ID, &train.DriverName, &train.OperatingStatus)

	if err != nil {
		if err == sql.ErrNoRows {
			resp.WriteErrorString(http.StatusNotFound, "Train could not be found")
		} else {
			log.Printf("Database error in getTrain : %v", err)
			resp.WriteErrorString(http.StatusInternalServerError, "Internal server error")
		}
		return
	}

	resp.WriteEntity(train)
}

// POST http://localhost:8000/v1/trains
func (t *Train) createTrain(req *restful.Request, resp *restful.Response) {
	var b TrainResource

	decoder := json.NewDecoder(req.Request.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&b); err != nil {
		log.Println("Invalid json", err)
		resp.WriteErrorString(http.StatusBadRequest, "Invalid JSON body")
		return
	}

	if b.DriverName == "" {
		resp.WriteErrorString(http.StatusBadRequest, "driver_name is required")
		return
	}

	statement, err := DB.Prepare("insert into train (DRIVER_NAME, OPERATING_STATUS) values (?,?)")
	if err != nil {
		log.Printf("Error preparing statement : %v", err)
		resp.WriteErrorString(http.StatusInternalServerError, "Database Error")
		return
	}

	defer statement.Close()

	result, err := statement.Exec(b.DriverName, b.OperatingStatus)
	if err != nil {
		log.Printf("Error executing insert : %v", err)
		resp.WriteErrorString(http.StatusInternalServerError, "Could not save train")
		return
	}

	newID, err := result.LastInsertId()
	if err != nil {
		log.Printf("Error getting last insert ID: %v", err)
		resp.WriteErrorString(http.StatusInternalServerError, "Database Error")
		return
	}

	b.ID = int(newID)

	resp.WriteHeaderAndEntity(http.StatusCreated, b)
}

func (t *Train) removeTrain(req *restful.Request, resp *restful.Response) {
	id := req.PathParameter("train-id")

	statement, err := DB.Prepare("delete from train where ID=?")
	if err != nil {
		log.Printf("Error removing train : %v", err)
		resp.WriteErrorString(http.StatusInternalServerError, "Database error")
		return
	}

	defer statement.Close()

	result, err := statement.Exec(id)
	if err != nil {
		log.Printf("delete exec error: %v", err)
		resp.WriteErrorString(http.StatusInternalServerError, "Could not delete train")
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("rowsAffected error: %v", err)
		resp.WriteErrorString(http.StatusInternalServerError, "Database error")
		return
	}

	if rowsAffected == 0 {
		resp.WriteErrorString(http.StatusNotFound, "Train not found")
		return
	}

	resp.WriteHeader(http.StatusNoContent)
}

func RunTrain() {
	var err error
	DB, err = sql.Open("sqlite3", "./railapi.db")
	if err != nil {
		log.Printf("Error Opening Database : %v", err)
	}

	// Always good practice to check if the DB is actually reachable
	if err = DB.Ping(); err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	dbutils.Initialize(DB)
	wsContainer := restful.NewContainer()
	wsContainer.Router(restful.CurlyRouter{})

	t := Train{}
	t.Register(wsContainer)

	fmt.Println("Server is running on PORT 8000...")

	server := &http.Server{Addr: ":8000", Handler: wsContainer}
	log.Fatal(server.ListenAndServe())
}
