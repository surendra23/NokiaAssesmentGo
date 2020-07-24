package controller

import (
	"NokiaAssesmentGo/db"
	"NokiaAssesmentGo/utils"
	"encoding/json"
	"net/http"
	"strconv"
)

// GetFromDBAndStoreInCache get the persisted data from DB and stores in cache
// @Summary gets data from DB and stores in Redis cache
// @Produce json
// @Success 200 {strig}
// @Router getFromDBAndStoreInCache [get]
func GetFromDBAndStoreInCache(res http.ResponseWriter, req *http.Request) {
	db.GetPersonFromDB()
	json.NewEncoder(res).Encode("Data stored into cache")
}

// ReadFromCache get the data from cache
// @Summary gets data from from Redis cache
// @Produce json
// @Success 200 {strig}
// @Router ReadFromCache [get]
func ReadFromCache(res http.ResponseWriter, req *http.Request) {
	//people :=[]db.Person{}
	id := req.FormValue("ID")
	if id == "" {
		id = "*"
	}
	offsetStr := req.FormValue("offset")
	if offsetStr == "" {
		http.Error(res, "offset is required", http.StatusBadRequest)
		return
	}
	offset, _ := strconv.Atoi(offsetStr)

	limitStr := req.FormValue("limit")
	if limitStr == "" {
		http.Error(res, "limit is required", http.StatusBadRequest)
		return
	}

	limit, _ := strconv.Atoi(limitStr)
	entries := utils.FetchFromCache(id, offset, limit)
	// person := &db.Person{}
	// json.Unmarshal([]byte(entry), &person)
	json.NewEncoder(res).Encode(entries)
}
