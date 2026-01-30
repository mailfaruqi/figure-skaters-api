package handlers

import (
	"encoding/json"
	"figure-skaters-api/models"
	"figure-skaters-api/services"
	"net/http"
	"strconv"
	"strings"
)

type ElementHandler struct {
service *services.ElementService
}

func NewElementHandler(service *services.ElementService) *ElementHandler {
return &ElementHandler{service: service}
}

// HandleElements - GET /api/elements
func (h *ElementHandler) HandleElements(w http.ResponseWriter, r *http.Request) {
switch r.Method {
case http.MethodGet:
h.GetAll(w, r)
case http.MethodPost:
h.Create(w, r)
default:
http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}
}

func (h *ElementHandler) GetAll(w http.ResponseWriter, r *http.Request) {
elements, err := h.service.GetAll()
if err != nil {
http.Error(w, err.Error(), http.StatusInternalServerError)
return
}
w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(elements)
}

func (h *ElementHandler) Create(w http.ResponseWriter, r *http.Request) {
var element models.Element
err := json.NewDecoder(r.Body).Decode(&element)
if err != nil {
http.Error(w, "Invalid request body", http.StatusBadRequest)
return
}
err = h.service.Create(&element)
if err != nil {
http.Error(w, err.Error(), http.StatusBadRequest)
return
}
w.Header().Set("Content-Type", "application/json")
w.WriteHeader(http.StatusCreated)
json.NewEncoder(w).Encode(element)
}

// HandleElementByID - GET/PUT/DELETE /api/elements/{id}
func (h *ElementHandler) HandleElementByID(w http.ResponseWriter, r *http.Request) {
switch r.Method {
case http.MethodGet:
h.GetByID(w, r)
case http.MethodPut:
h.Update(w, r)
case http.MethodDelete:
h.Delete(w, r)
default:
http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}
}

func (h *ElementHandler) GetByID(w http.ResponseWriter, r *http.Request) {
idStr := strings.TrimPrefix(r.URL.Path, "/api/elements/")
id, err := strconv.Atoi(idStr)
if err != nil {
http.Error(w, "Invalid element ID", http.StatusBadRequest)
return
}
element, err := h.service.GetByID(id)
if err != nil {
http.Error(w, err.Error(), http.StatusNotFound)
return
}
w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(element)
}

func (h *ElementHandler) Update(w http.ResponseWriter, r *http.Request) {
idStr := strings.TrimPrefix(r.URL.Path, "/api/elements/")
id, err := strconv.Atoi(idStr)
if err != nil {
http.Error(w, "Invalid element ID", http.StatusBadRequest)
return
}
var element models.Element
err = json.NewDecoder(r.Body).Decode(&element)
if err != nil {
http.Error(w, "Invalid request body", http.StatusBadRequest)
return
}
element.ID = id
err = h.service.Update(&element)
if err != nil {
http.Error(w, err.Error(), http.StatusBadRequest)
return
}
w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(element)
}

func (h *ElementHandler) Delete(w http.ResponseWriter, r *http.Request) {
idStr := strings.TrimPrefix(r.URL.Path, "/api/elements/")
id, err := strconv.Atoi(idStr)
if err != nil {
http.Error(w, "Invalid element ID", http.StatusBadRequest)
return
}
err = h.service.Delete(id)
if err != nil {
http.Error(w, err.Error(), http.StatusInternalServerError)
return
}
w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(map[string]string{
"message": "Element deleted successfully",
})
}