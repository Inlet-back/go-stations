package handler

import (
	"context"
	"encoding/json"
	"strconv"

	"net/http"

	"github.com/TechBowl-japan/go-stations/model"
	"github.com/TechBowl-japan/go-stations/service"
)

// A TODOHandler implements handling REST endpoints.
type TODOHandler struct {
	svc *service.TODOService
}

// NewTODOHandler returns TODOHandler based http.Handler.
func NewTODOHandler(svc *service.TODOService) *TODOHandler {
	return &TODOHandler{
		svc: svc,
	}
}
func (t *TODOHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	
	if r.Method == http.MethodPost{
		var todoRequest model.CreateTODORequest
		if err := json.NewDecoder(r.Body).Decode(&todoRequest);
		 err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if(todoRequest.Subject == "" ){
		 http.Error(w, "subject is required", http.StatusBadRequest)
		 return
		}
		todo,err := t.Create(r.Context(),&todoRequest)
		if err!=nil{
			http.Error(w,err.Error(),http.StatusBadRequest)
			return
		}
		if err :=json.NewEncoder(w).Encode(todo) ;err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
			
		}
		
	}
	if r.Method == http.MethodPut{
		var todoRequest model.UpdateTODORequest
		if err := json.NewDecoder(r.Body).Decode(&todoRequest);
		err !=nil {
			http.Error(w,err.Error(),http.StatusBadRequest)
			return
		}
		if todoRequest.ID ==0 || todoRequest.Subject =="" {
			http.Error(w, "request is invalid ", http.StatusBadRequest)
			return
		}
		todo,err := t.Update(r.Context(),&todoRequest)
		if err != nil {
	
			http.Error(w,err.Error(),http.StatusNotFound)
			return
		}
		if err := json.NewEncoder(w).Encode(todo); err != nil{
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		
	}
	if r.Method == http.MethodGet{
		var todoRequest model.ReadTODORequest
		var err error
		qyeryParams := r.URL.Query()
		
		if qyeryParams.Get("size") !=""{
			todoRequest.PrevID,_ = strconv.ParseInt(qyeryParams.Get("prev_id"),10,64)

			todoRequest.Size,err  = strconv.ParseInt(qyeryParams.Get("size"),10,64)
			if err != nil  {
				http.Error(w,err.Error(),http.StatusBadRequest)
				return
			}
		}else{
			todoRequest.Size = 5
		}
		todos,err := t.Read(r.Context(),&todoRequest)
		if err!=nil{
			http.Error(w,err.Error(),http.StatusBadRequest)
			return
		}
		if err :=json.NewEncoder(w).Encode(todos) ;err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
			
		}	


	}

}

// Create handles the endpoint that creates the TODO.
func (h *TODOHandler) Create(ctx context.Context, req *model.CreateTODORequest) (*model.CreateTODOResponse, error) {
	
	todo, err := h.svc.CreateTODO(ctx, req.Subject, req.Description)
	return &model.CreateTODOResponse{TODO: *todo}, err
}

// Read handles the endpoint that reads the TODOs.
func (h *TODOHandler) Read(ctx context.Context, req *model.ReadTODORequest) (*model.ReadTODOResponse, error) {
	todos, err := h.svc.ReadTODO(ctx, req.PrevID, req.Size)
	return &model.ReadTODOResponse{TODOs: todos}, err
}

// Update handles the endpoint that updates the TODO.
func (h *TODOHandler) Update(ctx context.Context, req *model.UpdateTODORequest) (*model.UpdateTODOResponse, error) {
	todo, err := h.svc.UpdateTODO(ctx, req.ID,req.Subject , req.Description)
	return &model.UpdateTODOResponse{TODO: *todo}, err
}

// Delete handles the endpoint that deletes the TODOs.
func (h *TODOHandler) Delete(ctx context.Context, req *model.DeleteTODORequest) (*model.DeleteTODOResponse, error) {
	_ = h.svc.DeleteTODO(ctx, nil)
	return &model.DeleteTODOResponse{}, nil
}
