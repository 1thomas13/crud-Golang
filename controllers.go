package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type task struct {
	ID      int
	Name    string `json:"Name"`
	Content string `json:"Content"`
}

type allTasks []task

var tasks = allTasks{
	{
		ID:      1,
		Name:    "clean",
		Content: "clean bedroom",
	},
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func getTask(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	taskID, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprint(w, "Invalid Id")
		return
	}

	for _, task := range tasks {
		if task.ID == taskID {
			w.Header().Set("content-type", "application/json")
			json.NewEncoder(w).Encode(task)
		}
	}
}

func deleteTask(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	taskID, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprint(w, "Invalid Id")
		return
	}

	for i, task := range tasks {
		if task.ID == taskID {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Fprintf(w, "The task with ID %v has been remove successfully.", taskID)
		}
	}
}

func editTask(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])
	var updatedTask task

	if err != nil {
		fmt.Fprint(w, "Invalid Id")
		return
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprint(w, "Please entrer a valid data")
		return
	}

	json.Unmarshal(reqBody, &updatedTask)

	for i, task := range tasks {
		if task.ID == taskID {
			tasks = append(tasks[:i], tasks[i+1:]...)
			updatedTask.ID = taskID
			tasks = append(tasks, updatedTask)
			fmt.Fprintf(w, "The task with ID %v has been updated successfully.", taskID)
		}
	}
}

func createTask(w http.ResponseWriter, r *http.Request) {
	var newTask task
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Insert a valid task")
	}
	json.Unmarshal(reqBody, &newTask)

	newTask.ID = len(tasks) + 1

	tasks = append(tasks, newTask)

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}
