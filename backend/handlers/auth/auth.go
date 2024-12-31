package auth

import (
	"backend/config"
	"backend/models"
	"backend/serializers/users"
	"backend/utils/auth"
	"backend/utils/commonutils"
	"backend/utils/dbService"
	"backend/utils/validation"
	"github.com/sirupsen/logrus"
	"net/http"
)

func AuthLoginHandler(w http.ResponseWriter, r *http.Request) {
	// validate login request
	var request users.UserLoginRequest
	err := validation.ValidateRequest(w, r, &request)
	if err != nil {
		logrus.Error(err)
		return
	}

	// ensure that either username or email is provided
	if request.Email == "" && request.Username == "" {
		logrus.Error("Email or Username is required")
		http.Error(w, "Email or Username is required", http.StatusBadRequest)
	}

	// check if user exists
	sv := dbService.DBService{DB: config.DB}

	var condition models.User
	if request.Username != "" {
		condition.Username = request.Username

	} else if request.Email != "" {
		condition.Email = request.Email

	} else {
		logrus.Error("Email or Username is required")
		http.Error(w, "Email or Username is required", http.StatusBadRequest)
	}

	user, err := sv.Get(r.Context(), &condition)
	if err != nil {
		logrus.Error(err)
		http.Error(w, "User not found", http.StatusInternalServerError)
	}

	if user != nil {
		userObj := user.(*models.User)

		// compare password with that of retrieved record
		isValid, err := auth.ComparePasswords(userObj.Password, []byte(request.Password))
		if err != nil {
			logrus.Error(err)
			http.Error(w, "User not found", http.StatusInternalServerError)
		}

		if !isValid {
			logrus.Error("Username or Password is invalid")
			http.Error(w, "Error validating user", http.StatusBadRequest)
		}

		// generate auth token
		token, err := auth.GenerateAuthToken(*userObj)
		if err != nil {
			logrus.Error(err)
			http.Error(w, "Error validating user", http.StatusInternalServerError)
		}

		// return success response
		responseData := users.UserLoginResponse{
			Token:    token,
			UserData: user,
		}

		response := commonutils.Response{
			Message: "User login successful",
			Data:    responseData,
		}

		if err = commonutils.HTTPResponse(w, response, http.StatusOK); err != nil {
			logrus.Error(err)
			return
		}

	} else {
		logrus.Error("error retrieving user data")
		http.Error(w, "Error retrieving user data", http.StatusInternalServerError)
	}
}
