package users

import (
	"backend/config"
	"backend/models"
	"backend/serializers/users"
	"backend/utils/commonutils"
	"backend/utils/dbService"
	baseModels "backend/utils/models"
	"backend/utils/validation"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	// request serializer
	var request users.UserCreateSerializer
	err := validation.ValidateRequest(w, r, &request)
	if err != nil {
		logrus.Error(err)
		return
	}

	// initialize db
	sv := dbService.DBService{DB: config.DB}

	// create user
	userId := uuid.NewString()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		logrus.Error(err)
		return
	}

	condition := &models.User{
		Model:    baseModels.Model{Id: userId},
		Username: request.Username,
		Email:    request.Email,
		Password: hashedPassword,
	}

	createdUser, err := sv.GetOrCreate(r.Context(), condition, condition)
	if err != nil {
		logrus.Error(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	response := commonutils.Response{
		Message: "User created successfully",
		Data:    createdUser,
	}

	if err = commonutils.HTTPResponse(w, response, http.StatusOK); err != nil {
		logrus.Error(err)
		return
	}
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	// get user id
	userId, ok := commonutils.GetUserIdFromContext(r.Context())
	if !ok {
		logrus.Error("User id not found")
		return
	}

	// get user id from url param
	userIdUrl := mux.Vars(r)["userId"]

	if userIdUrl == "" {
		logrus.Error("User id cannot be empty")
		http.Error(w, "url param is empty", http.StatusBadRequest)
		return
	}

	// check if the url param userId and the one in the context are the same
	if userIdUrl != userId {
		logrus.Error("User id does not match")
		http.Error(w, "user id does not match", http.StatusBadRequest)
		return
	}

	// get user record using user Id
	condition := &models.User{Model: baseModels.Model{Id: userId}}

	sv := dbService.DBService{DB: config.DB}

	user, err := sv.Get(r.Context(), condition)
	if err != nil {
		logrus.Error(err)
		http.Error(w, "User not found", http.StatusInternalServerError)
		return
	}

	// return http response
	response := commonutils.Response{
		Message: "success",
		Data:    user,
	}

	if err = commonutils.HTTPResponse(w, response, http.StatusOK); err != nil {
		logrus.Error(err)
		http.Error(w, "Error writing http response", http.StatusInternalServerError)
		return
	}
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	userId, ok := commonutils.GetUserIdFromContext(r.Context())
	if !ok {
		logrus.Error("User id not found")
		return
	}

	// check if userId is valid
	userIdUrl := mux.Vars(r)["userId"]

	// validate request body
	var request users.UserUpdateRequest
	err := validation.ValidateRequest(w, r, &request)
	if err != nil {
		logrus.Error(err)
		http.Error(w, "invalid field(s)", http.StatusBadRequest)
		return
	}

	isValid, err := commonutils.CompareUserIds(userIdUrl, userId)
	if err != nil || !isValid {
		logrus.Error(err)
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		return
	}

	// check if username, email or password is provided
	var model models.User
	if request.Username != "" {
		model.Username = request.Username

	} else if request.Email != "" {
		model.Email = request.Email

	} else if request.Password != "" {
		// hash password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
		if err != nil {
			logrus.Error(err)
			return
		}
		model.Password = hashedPassword
	} else {
		logrus.Error("All fields are empty")
		http.Error(w, "Fill in at least one field", http.StatusBadRequest)
		return
	}

	// condition for update
	condition := &models.User{Model: baseModels.Model{Id: userIdUrl}}
	sv := dbService.DBService{DB: config.DB}

	// check if user exists then update
	user, err := sv.UpdateAndGet(r.Context(), model, condition)
	if err != nil {
		logrus.Error(err)
		http.Error(w, "error updating user", http.StatusInternalServerError)
		return
	}

	response := commonutils.Response{
		Message: "user updated successfully",
		Data:    user,
	}

	if err = commonutils.HTTPResponse(w, response, http.StatusOK); err != nil {
		logrus.Error(err)
		http.Error(w, "Error writing http response", http.StatusInternalServerError)
		return
	}
}
