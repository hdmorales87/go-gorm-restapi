package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hdmorales87/go-gorm-restapi/dto"
	"github.com/hdmorales87/go-gorm-restapi/services"
)

// GetUsersHandler godoc
// @Summary Get all users
// @Description Get a list of all users
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {array} dto.UserDTO
// @Failure 500 {object} map[string]string
// @Router /users [get]
func GetUsersHandler(userService services.IUserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := userService.GetAllUsers()
		w.Header().Set("Content-Type", "application/json")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}
		json.NewEncoder(w).Encode(users)
	}
}

// GetUserHandler godoc
// @Summary Get user by ID
// @Description Get a single user by ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {object} dto.UserDTO
// @Failure 404 {object} map[string]string
// @Router /users/{id} [get]
func GetUserHandler(userService services.IUserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		user, err := userService.GetUserByID(params["id"])
		w.Header().Set("Content-Type", "application/json")
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}
		json.NewEncoder(w).Encode(user)
	}
}

// CreateUserHandler godoc
// @Summary Create a new user
// @Description Create a new user with the provided data
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body dto.CreateUserDTO true "User data"
// @Success 200 {object} dto.UserDTO
// @Failure 500 {object} map[string]string
// @Router /users [post]
func CreateUserHandler(userService services.IUserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var userDTO dto.CreateUserDTO
		json.NewDecoder(r.Body).Decode(&userDTO)
		createdUser, err := userService.CreateUser(userDTO)
		w.Header().Set("Content-Type", "application/json")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}
		json.NewEncoder(w).Encode(createdUser)
	}
}

// UpdateUserHandler godoc
// @Summary Update a user
// @Description Update an existing user by ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Param user body dto.UpdateUserDTO true "User data"
// @Success 200 {object} dto.UserDTO
// @Failure 404 {object} map[string]string
// @Router /users/{id} [put]
func UpdateUserHandler(userService services.IUserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		var userDTO dto.UpdateUserDTO
		json.NewDecoder(r.Body).Decode(&userDTO)
		updatedUser, err := userService.UpdateUser(params["id"], userDTO)
		w.Header().Set("Content-Type", "application/json")
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(updatedUser)
	}
}

// DeleteUserHandler godoc
// @Summary Delete a user
// @Description Delete a user by ID (soft delete)
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 204
// @Failure 404 {object} map[string]string
// @Router /users/{id} [delete]
func DeleteUserHandler(userService services.IUserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		err := userService.DeleteUser(params["id"])
		w.Header().Set("Content-Type", "application/json")
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
