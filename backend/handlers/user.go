package handlers

import (
	"backend/config"
	"backend/models"
	"backend/serializers"
	"backend/utils/commonutils"
	"backend/utils/dbService"
	baseModels "backend/utils/models"
	"backend/utils/validation"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	// request serializer
	var request serializers.UserCreateSerializer
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

	createdUser, err := sv.Create(r.Context(), condition, condition)
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

}
