package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hdmorales87/go-gorm-restapi/dto"
	"github.com/hdmorales87/go-gorm-restapi/services"
)

// GetTasksHandler godoc
// @Summary Get all tasks
// @Description Get a list of all tasks
// @Tags tasks
// @Accept  json
// @Produce  json
// @Success 200 {array} dto.TaskDTO
// @Failure 500 {object} map[string]string
// @Router /tasks [get]
func GetTasksHandler(taskService *services.TaskService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tasks, err := taskService.GetAllTasks()
		w.Header().Set("Content-Type", "application/json")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(tasks)
	}
}

// GetTaskHandler godoc
// @Summary Get task by ID
// @Description Get a single task by ID
// @Tags tasks
// @Accept  json
// @Produce  json
// @Param id path string true "Task ID"
// @Success 200 {object} dto.TaskDTO
// @Failure 404 {object} map[string]string
// @Router /tasks/{id} [get]
func GetTaskHandler(taskService *services.TaskService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		task, err := taskService.GetTaskByID(params["id"])
		w.Header().Set("Content-Type", "application/json")
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(task)
	}
}

// CreateTaskHandler godoc
// @Summary Create a new task
// @Description Create a new task with the provided data
// @Tags tasks
// @Accept  json
// @Produce  json
// @Param task body dto.CreateTaskDTO true "Task data"
// @Success 201 {object} dto.TaskDTO
// @Failure 500 {object} map[string]string
// @Router /tasks [post]
func CreateTaskHandler(taskService *services.TaskService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var taskDTO dto.CreateTaskDTO
		json.NewDecoder(r.Body).Decode(&taskDTO)
		createdTask, err := taskService.CreateTask(taskDTO)
		w.Header().Set("Content-Type", "application/json")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(createdTask)
	}
}

// UpdateTaskHandler godoc
// @Summary Update a task
// @Description Update an existing task by ID
// @Tags tasks
// @Accept  json
// @Produce  json
// @Param id path string true "Task ID"
// @Param task body dto.UpdateTaskDTO true "Task data"
// @Success 200 {object} dto.TaskDTO
// @Failure 404 {object} map[string]string
// @Router /tasks/{id} [put]
func UpdateTaskHandler(taskService *services.TaskService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		var taskDTO dto.UpdateTaskDTO
		json.NewDecoder(r.Body).Decode(&taskDTO)
		updatedTask, err := taskService.UpdateTask(params["id"], taskDTO)
		w.Header().Set("Content-Type", "application/json")
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(updatedTask)
	}
}

// DeleteTaskHandler godoc
// @Summary Delete a task
// @Description Delete a task by ID (soft delete)
// @Tags tasks
// @Accept  json
// @Produce  json
// @Param id path string true "Task ID"
// @Success 204
// @Failure 404 {object} map[string]string
// @Router /tasks/{id} [delete]
func DeleteTaskHandler(taskService *services.TaskService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		err := taskService.DeleteTask(params["id"])
		w.Header().Set("Content-Type", "application/json")
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
