package categories

import (
	"backend/config"
	"backend/models"
	categorySerializer "backend/serializers/categories"
	"backend/utils/commonutils"
	"backend/utils/dbService"
	baseModels "backend/utils/models"
	"backend/utils/validation"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

func CreateCategoryHandler(w http.ResponseWriter, r *http.Request) {
	// get user id
	userId, ok := commonutils.GetUserIdFromContext(r.Context())
	if !ok {
		logrus.Error("userId is empty")
		return
	}

	// validate request
	var request categorySerializer.CreateCategoryRequest
	err := validation.ValidateRequest(w, r, &request)
	if err != nil {
		logrus.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	categoryType := strings.ToUpper(request.CategoryType)

	// create category
	condition := &models.Category{
		UserId: userId,
	}

	model := &models.Category{
		UserId: userId,
		Name:   request.Name,
		Type:   models.TransactionType(categoryType),
	}

	sv := dbService.DBService{DB: config.DB}

	// check if category exists first
	var category interface{}
	categoryExisting, err := sv.Get(r.Context(), model)
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			category, err = sv.Create(r.Context(), model, condition)
			if err != nil {
				logrus.Error(err)
				http.Error(w, "error creating category", http.StatusInternalServerError)
				return
			}
		} else {
			logrus.Error(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}

	response := commonutils.Response{
		Message: "success",
		Data:    category,
	}

	if categoryExisting != nil {
		response.Data = categoryExisting
	} else if category != nil {
		response.Data = category
	} else {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if err = commonutils.HTTPResponse(w, response, http.StatusOK); err != nil {
		logrus.Error(err)
		http.Error(w, "Error writing http response", http.StatusInternalServerError)
		return
	}
}

func GetCategoryByIdHandler(w http.ResponseWriter, r *http.Request) {
	// get user id
	userId, ok := commonutils.GetUserIdFromContext(r.Context())
	if !ok {
		logrus.Error("userId is empty")
		return
	}

	// get category Id
	categoryId := mux.Vars(r)["categoryId"]

	if categoryId == "" {
		logrus.Error("categoryId is empty")
		http.Error(w, "categoryId is empty", http.StatusBadRequest)
		return
	}

	// create category
	condition := &models.Category{
		Model:  baseModels.Model{Id: categoryId},
		UserId: userId,
	}

	sv := dbService.DBService{DB: config.DB}

	category, err := sv.Get(r.Context(), condition)
	if err != nil {
		logrus.Error(err)
		http.Error(w, "error retrieving category", http.StatusInternalServerError)
		return
	}

	response := commonutils.Response{
		Message: "success",
		Data:    category,
	}

	if err = commonutils.HTTPResponse(w, response, http.StatusOK); err != nil {
		logrus.Error(err)
		http.Error(w, "Error writing http response", http.StatusInternalServerError)
		return
	}

}

func GetCategoryByNameHandler(w http.ResponseWriter, r *http.Request) {
	// get user id
	userId, ok := commonutils.GetUserIdFromContext(r.Context())
	if !ok {
		logrus.Error("userId is empty")
		return
	}

	// get category Id
	categoryName := mux.Vars(r)["categoryName"]

	if categoryName == "" {
		logrus.Error("category name is empty")
		http.Error(w, "category name is empty", http.StatusBadRequest)
		return
	}

	// create category
	condition := &models.Category{
		UserId: userId,
		Name:   categoryName,
	}

	sv := dbService.DBService{DB: config.DB}

	category, err := sv.Get(r.Context(), condition)
	if err != nil {
		logrus.Error(err)
		http.Error(w, "error retrieving category", http.StatusInternalServerError)
		return
	}

	response := commonutils.Response{
		Message: "success",
		Data:    category,
	}

	if err = commonutils.HTTPResponse(w, response, http.StatusOK); err != nil {
		logrus.Error(err)
		http.Error(w, "Error writing http response", http.StatusInternalServerError)
		return
	}
}

func GetPaginatedCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	userId, ok := commonutils.GetUserIdFromContext(r.Context())
	if !ok {
		logrus.Error("userId is empty")
		return
	}

	// get pagination params
	pageParam := r.URL.Query().Get("page")
	pageSizeParam := r.URL.Query().Get("page_size")

	// assign default pagination param values
	var page, pageSize int
	if pageParam == "" && pageSizeParam == "" {
		page = 1
		pageSize = 10
	} else {
		page, pageSize = commonutils.GetPaginationParams(pageParam, pageSizeParam)
	}

	// get categories
	condition := &models.Category{
		UserId: userId,
	}

	sv := dbService.DBService{DB: config.DB}
	categories, err := sv.GetPaginated(r.Context(), condition, page, pageSize)
	if err != nil {
		logrus.Error(err)

		// check if its a record not found error
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			http.Error(w, "category not found", http.StatusNotFound)
			return
		}

		http.Error(w, "error retrieving category", http.StatusInternalServerError)
		return
	}

	// return response
	response := commonutils.Response{
		Message: "success",
		Data:    categories,
	}

	if err = commonutils.HTTPResponse(w, response, http.StatusOK); err != nil {
		logrus.Error(err)
		http.Error(w, "Error writing http response", http.StatusInternalServerError)
		return
	}
}

func GetPaginatedCategoriesByCategoryTypeHandler(w http.ResponseWriter, r *http.Request) {
	userId, ok := commonutils.GetUserIdFromContext(r.Context())
	if !ok {
		logrus.Error("userId is empty")
		return
	}

	// get category Type
	categoryType := mux.Vars(r)["categoryType"]

	if categoryType == "" {
		logrus.Error("category type is empty")
		http.Error(w, "category type is empty", http.StatusBadRequest)
		return
	}

	// get pagination params
	pageParam := r.URL.Query().Get("page")
	pageSizeParam := r.URL.Query().Get("page_size")

	// assign default pagination param values
	var page, pageSize int
	if pageParam == "" && pageSizeParam == "" {
		page = 1
		pageSize = 10
	} else {
		page, pageSize = commonutils.GetPaginationParams(pageParam, pageSizeParam)
	}

	// get categories
	condition := &models.Category{
		UserId: userId,
		Type:   models.TransactionType(categoryType),
	}

	sv := dbService.DBService{DB: config.DB}
	categories, err := sv.GetPaginated(r.Context(), condition, page, pageSize)
	if err != nil {
		logrus.Error(err)

		// check if its a record not found error
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			http.Error(w, "category not found", http.StatusNotFound)
			return
		}

		http.Error(w, "error retrieving category", http.StatusInternalServerError)
		return
	}

	// return response
	response := commonutils.Response{
		Message: "success",
		Data:    categories,
	}

	if err = commonutils.HTTPResponse(w, response, http.StatusOK); err != nil {
		logrus.Error(err)
		http.Error(w, "Error writing http response", http.StatusInternalServerError)
		return
	}
}

func UpdateCategoryHandler(w http.ResponseWriter, r *http.Request) {
	// get user id
	userId, ok := commonutils.GetUserIdFromContext(r.Context())
	if !ok {
		logrus.Error("userId is empty")
		return
	}

	// get category Id
	categoryId := mux.Vars(r)["categoryId"]

	if categoryId == "" {
		logrus.Error("categoryId is empty")
		http.Error(w, "categoryId is empty", http.StatusBadRequest)
		return
	}

	// validate request
	var request categorySerializer.UpdateCategoryRequest
	err := validation.ValidateRequest(w, r, &request)
	if err != nil {
		logrus.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// update category
	condition := &models.Category{
		Model:  baseModels.Model{Id: categoryId},
		UserId: userId,
	}

	var model models.Category
	if request.Name == "" && request.CategoryType == "" {
		logrus.Error("all request fields are empty")
		http.Error(w, "all request fields are empty", http.StatusBadRequest)
		return
	}

	if request.Name != "" {
		model.Name = request.Name
	}

	if request.CategoryType != "" {
		model.Type = models.TransactionType(request.CategoryType)
	}

	sv := dbService.DBService{DB: config.DB}

	category, err := sv.UpdateAndGet(r.Context(), model, condition)
	if err != nil {
		logrus.Error(err)
		http.Error(w, "error updating category", http.StatusInternalServerError)
		return
	}

	// return response
	response := commonutils.Response{
		Message: "success",
		Data:    category,
	}

	if err = commonutils.HTTPResponse(w, response, http.StatusOK); err != nil {
		logrus.Error(err)
		http.Error(w, "Error writing http response", http.StatusInternalServerError)
		return
	}
}

func DeleteCategoryHandler(w http.ResponseWriter, r *http.Request) {
	userId, ok := commonutils.GetUserIdFromContext(r.Context())
	if !ok {
		logrus.Error("userId is empty")
		return
	}

	// get category id from url
	categoryId := mux.Vars(r)["categoryId"]

	if categoryId == "" {
		logrus.Error("No category id provided")
		http.Error(w, "invalid category id", http.StatusBadRequest)
		return
	}

	// delete category
	condition := &models.Category{
		Model:  baseModels.Model{Id: categoryId},
		UserId: userId,
	}

	sv := dbService.DBService{DB: config.DB}

	category, err := sv.Get(r.Context(), condition)
	if err != nil {
		logrus.Error(err)
		http.Error(w, "category does not exist", http.StatusInternalServerError)
		return
	}

	if category == nil {
		logrus.Error("category does not exist")
		http.Error(w, "category does not exist", http.StatusNotFound)
		return
	}

	err = sv.SoftDelete(r.Context(), condition)
	if err != nil {
		logrus.Error(err)
		http.Error(w, "error deleting category", http.StatusInternalServerError)
		return
	}

	response := commonutils.Response{
		Message: "category deleted successfully",
	}

	if err = commonutils.HTTPResponse(w, response, http.StatusOK); err != nil {
		logrus.Error(err)
		http.Error(w, "Error writing http response", http.StatusInternalServerError)
	}

}
