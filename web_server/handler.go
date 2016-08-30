package web_server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.ibm.com/almaden-containers/spectrum-common.git/models"
	"github.ibm.com/almaden-containers/spectrum-container-plugin.git/core"
	Db "github.ibm.com/almaden-containers/spectrum-common.git/core"
)

type Handler struct {
	Controller *core.Controller
	log        *log.Logger
}

func NewHandler(logger *log.Logger, filesystem, mountpath string, Dbclient *Db.DatabaseClient) *Handler {
	return &Handler{log: logger, Controller: core.NewController(logger, filesystem, mountpath, Dbclient)}
}

func (c *Handler) Activate(w http.ResponseWriter, r *http.Request) {
	c.log.Println("Handler: activate start")
	defer c.log.Println("Handler: activate end")
	activateResponse := c.Controller.Activate()
	activateResponse.WriteResponse(w)
}
func (c *Handler) Create(w http.ResponseWriter, r *http.Request) {
	c.log.Println("Handler: create start")
	defer c.log.Println("Handler: create end")
	var createRequest models.CreateRequest
	err := extractRequestObject(r, &createRequest)
	if err != nil {
		genericResponse := &models.GenericResponse{Err: err.Error()}
		genericResponse.WriteResponse(w)
		return
	}
	createResponse := c.Controller.Create(&createRequest)
	createResponse.WriteResponse(w)

}

func (c *Handler) Remove(w http.ResponseWriter, r *http.Request) {
	c.log.Println("Handler: remove start")
	defer c.log.Println("Handler: remove end")
	var removeRequest models.GenericRequest
	err := extractRequestObject(r, &removeRequest)
	if err != nil {
		genericResponse := &models.GenericResponse{Err: err.Error()}
		genericResponse.WriteResponse(w)
		return
	}
	removeResponse := c.Controller.Remove(&removeRequest)
	removeResponse.WriteResponse(w)
}

func (c *Handler) Mount(w http.ResponseWriter, r *http.Request) {
	c.log.Println("Handler: mount start")
	defer c.log.Println("Handler: mount end")
	var mountRequest models.GenericRequest
	err := extractRequestObject(r, &mountRequest)
	if err != nil {
		mountResponse := &models.MountResponse{Err: err.Error()}
		mountResponse.WriteResponse(w)
		return
	}
	mountResponse := c.Controller.Mount(&mountRequest)
	mountResponse.WriteResponse(w)
}

func (c *Handler) Unmount(w http.ResponseWriter, r *http.Request) {
	c.log.Println("Handler: unmount start")
	defer c.log.Println("Handler: unmount end")
	var unmountRequest models.GenericRequest
	err := extractRequestObject(r, &unmountRequest)
	if err != nil {
		genericResponse := &models.GenericResponse{Err: err.Error()}
		genericResponse.WriteResponse(w)
		return
	}
	unmountResponse := c.Controller.Unmount(&unmountRequest)
	unmountResponse.WriteResponse(w)
}

func (c *Handler) Path(w http.ResponseWriter, r *http.Request) {
	c.log.Println("Handler: path start")
	defer c.log.Println("Handler: path end")
	var pathRequest models.GenericRequest
	err := extractRequestObject(r, &pathRequest)
	if err != nil {
		mountResponse := &models.MountResponse{Err: err.Error()}
		mountResponse.WriteResponse(w)
		return
	}
	pathResponse := c.Controller.Path(&pathRequest)
	pathResponse.WriteResponse(w)
}

func (c *Handler) Get(w http.ResponseWriter, r *http.Request) {
	c.log.Println("Handler: get start")
	defer c.log.Println("Handler: get end")
	var getRequest models.GenericRequest
	err := extractRequestObject(r, &getRequest)
	if err != nil {
		errorResponse := &models.GetResponse{Err: err.Error()}
		errorResponse.WriteResponse(w)
		return
	}
	getResponse := c.Controller.Get(&getRequest)
	getResponse.WriteResponse(w)
}

func (c *Handler) List(w http.ResponseWriter, r *http.Request) {
	c.log.Println("Handler: list start")
	defer c.log.Println("Handler: list end")
	listResponse := c.Controller.List()
	listResponse.WriteResponse(w)
}

func extractRequestObject(r *http.Request, request interface{}) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("Error reading request body: %s", err.Error())
	}
	err = json.Unmarshal(body, &request)
	if err != nil {
		return fmt.Errorf("Error unmarshalling request: %s", err.Error())
	}
	return nil
}
