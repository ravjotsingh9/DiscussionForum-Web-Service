package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/ravjotsingh9/discussionForum-Web-Service/db"
	"github.com/ravjotsingh9/discussionForum-Web-Service/schema"
	"github.com/ravjotsingh9/discussionForum-Web-Service/util"
	"github.com/segmentio/ksuid"
)

func getWelcomeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	util.ResponseOk(w, "Welcome "+vars["name"])
}

func getTopicHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	ctx := r.Context()

	comment := schema.Comment{
		ID:      vars["id"],
		Content: "",
		PID:     "",
		TID:     vars["id"],
	}

	comments, err := db.GetComment(ctx, comment)
	if err != nil {
		log.Println(err)
		util.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	util.ResponseOk(w, comments)
}

func createCommentHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	// create a ID
	createdAt := time.Now().UTC()
	id, err := ksuid.NewRandomWithTime(createdAt)
	if err != nil {
		util.ResponseError(w, http.StatusInternalServerError, "Failed to create ID")
		return
	}

	decoder := json.NewDecoder(r.Body)
	var com schema.Comment
	err = decoder.Decode(&com)
	if err != nil {
		panic(err)
	}

	if com.PID == "" && com.TID == "" { // consider it a new topic
		// create a TID
		createdAt = time.Now().UTC()
		tid, err := ksuid.NewRandomWithTime(createdAt)
		if err != nil {
			util.ResponseError(w, http.StatusInternalServerError, "Failed to create ID")
			return
		}
		com.PID = tid.String()
		com.TID = com.PID
	}

	if com.TID != "" && com.PID == "" { // use tid as pid
		com.PID = com.TID
	}

	if com.PID != "" && com.TID == "" {
		comment := schema.Comment{
			ID:      com.PID,
			Content: "",
			PID:     "",
			TID:     "",
		}

		comments, err := db.GetComment(ctx, comment)
		if err != nil {
			log.Println(err)
			util.ResponseError(w, http.StatusInternalServerError, err.Error())
			return
		}
		if len(comments) < 1 {
			com.TID = comments[0].TID
		} else {
			util.ResponseError(w, http.StatusNotAcceptable, "Topic id is not specified, if its a new topic, do not specify the parent id either.")
		}

	}

	comment := schema.Comment{
		ID:      id.String(),
		Content: com.Content,
		PID:     com.PID,
		TID:     com.TID,
	}

	if err := db.InsertComment(ctx, comment); err != nil {
		log.Println(err)
		util.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.ResponseOk(w, comment)
}
