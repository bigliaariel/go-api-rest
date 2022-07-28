package main

import (
	"api/entities"
	"api/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func route_notes_all(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var noteModel models.NoteModel
	var notes []entities.Note
	notes, err := noteModel.FindAll()
	if err != nil {
		fmt.Println(err)
		json.NewEncoder(w).Encode([]entities.Note{})
	} else {
		if notes == nil {
			notes = []entities.Note{}
		}
		json.NewEncoder(w).Encode(notes)
	}
}

func route_notes_create(w http.ResponseWriter, req *http.Request) {

	var note entities.Note
	var noteModel models.NoteModel
	_ = json.NewDecoder(req.Body).Decode(&note)
	noteModel.Add(note.Title, note.Body, 1)
	route_notes_all(w, req)

}

func route_notes_update(w http.ResponseWriter, req *http.Request) {

	var note entities.Note
	var noteModel models.NoteModel
	_ = json.NewDecoder(req.Body).Decode(&note)
	noteModel.Update(note.Id, note.Title, note.Body, note.Color)
	route_notes_all(w, req)

}

func route_notes_delete(w http.ResponseWriter, req *http.Request) {

	var noteModel models.NoteModel
	params := mux.Vars(req)
	p1 := params["id"]
	id, _ := strconv.ParseInt(p1, 10, 32)
	noteModel.Remove(int32(id))
	route_notes_all(w, req)

}

func route_notes_find(w http.ResponseWriter, req *http.Request) {

	var noteModel models.NoteModel
	params := mux.Vars(req)
	p1 := params["id"]
	id, _ := strconv.ParseInt(p1, 10, 32)
	notes, err := noteModel.Find(int32(id))
	if err != nil {
		fmt.Println(err)
		json.NewEncoder(w).Encode([]entities.Note{})
	} else {
		if notes == nil {
			notes = &entities.Note{}
		}
		json.NewEncoder(w).Encode(notes)
	}

}

func main() {

	port := os.Getenv("PORT") // Produccion
	//port := "3000" // Desarrollo
	router := mux.NewRouter()
	fs := http.FileServer(http.Dir("./public/"))
	router.PathPrefix("/files/").Handler(http.StripPrefix("/files/", fs))

	router.HandleFunc("/note", route_notes_all).Methods("GET")
	router.HandleFunc("/note/{id}", route_notes_find).Methods("GET")
	router.HandleFunc("/note", route_notes_create).Methods("POST")
	router.HandleFunc("/note/{id}", route_notes_delete).Methods("DELETE")
	router.HandleFunc("/note", route_notes_update).Methods("PUT")
	log.Print("Listening on :" + port)
	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "DELETE"}), handlers.AllowedOrigins([]string{"*"}))(router)))
	//log.Fatal(http.ListenAndServe(":"+port, router))

}
