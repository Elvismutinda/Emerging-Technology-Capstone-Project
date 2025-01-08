package transactions

import (
	"backend/config"
	"backend/models"
	transactionsSerializer "backend/serializers/transactions"
	"backend/utils/commonutils"
	"backend/utils/dbService"
	baseModels "backend/utils/models"
	"backend/utils/validation"
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
	"strings"
	"time"
)

func CreateTransactionHandler(w http.ResponseWriter, r *http.Request) {
	userId, ok := commonutils.GetUserIdFromContext(r.Context())
	if !ok {
		logrus.Error("userId is empty")
		return
	}

	// get transaction type from url
	transactionType := strings.ToUpper(mux.Vars(r)["transactionType"])

	if models.TransactionType(transactionType) != models.Income && models.TransactionType(transactionType) != models.Expense {
		logrus.Error("transactionType is invalid")
		http.Error(w, "transactionType is invalid", http.StatusBadRequest)
		return
	}

	// validate request
	var request transactionsSerializer.CreateTransactionRequest
	err := validation.ValidateRequest(w, r, &request)
	if err != nil {
		logrus.Error(err)
		return
	}

	sv := dbService.DBService{DB: config.DB}

	// get or create category
	conditionCategory := &models.Category{
		Name: request.Category,
		Type: models.TransactionType(transactionType),
	}
	category, err := sv.GetOrCreate(r.Context(), conditionCategory, conditionCategory)
	if err != nil {
		logrus.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// convert transaction date
	layout := "2006-01-02"
	transactionDate, err := time.Parse(layout, request.TransactionDate)
	if err != nil {
		logrus.Error(err)
		http.Error(w, "Invalid Transaction Date", http.StatusBadRequest)
		return
	}

	// create transaction
	condition := &models.Transaction{
		Description:     request.Description,
		UserId:          userId,
		CategoryId:      category.(*models.Category).Id,
		Amount:          request.Amount,
		TransactionDate: transactionDate,
		Type:            models.TransactionType(transactionType),
	}

	transaction, err := sv.GetOrCreate(r.Context(), condition, condition)
	if err != nil {
		logrus.Error(err)
		http.Error(w, "error creating transaction", http.StatusInternalServerError)
		return
	}

	// return response
	response := commonutils.Response{
		Message: "transaction created successfully",
		Data:    transaction,
	}

	if err = commonutils.HTTPResponse(w, response, http.StatusOK); err != nil {
		logrus.Error(err)
		http.Error(w, "Error writing http response", http.StatusInternalServerError)
		return
	}
}

func GetTransactionHandler(w http.ResponseWriter, r *http.Request) {
	userId, ok := commonutils.GetUserIdFromContext(r.Context())
	if !ok {
		logrus.Error("userId is empty")
		return
	}

	// get transaction id and type from url
	transactionId := mux.Vars(r)["transactionId"]

	if transactionId == "" {
		logrus.Error("transactionId is empty")
		http.Error(w, "transaction id is empty", http.StatusBadRequest)
	}

	// get transaction
	condition := &models.Transaction{
		Model:  baseModels.Model{Id: transactionId},
		UserId: userId,
	}

	sv := dbService.DBService{DB: config.DB}
	transaction, err := sv.Get(r.Context(), condition)
	if err != nil {
		logrus.Error(err)

		// check if its a record not found error
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			http.Error(w, "transaction not found", http.StatusNotFound)
			return
		}

		http.Error(w, "error retrieving transaction", http.StatusInternalServerError)
		return
	}

	// return response
	response := commonutils.Response{
		Message: "success",
		Data:    transaction,
	}

	if err = commonutils.HTTPResponse(w, response, http.StatusOK); err != nil {
		logrus.Error(err)
		http.Error(w, "Error writing http response", http.StatusInternalServerError)
		return
	}
}

func GetPaginatedTransactionsHandler(w http.ResponseWriter, r *http.Request) {
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

	// get transactions
	condition := &models.Transaction{
		UserId: userId,
	}

	sv := dbService.DBService{DB: config.DB}
	transactions, err := sv.GetPaginated(r.Context(), condition, page, pageSize)
	if err != nil {
		logrus.Error(err)

		// check if its a record not found error
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			http.Error(w, "transaction not found", http.StatusNotFound)
			return
		}

		http.Error(w, "error retrieving transaction", http.StatusInternalServerError)
		return
	}

	// return response
	response := commonutils.Response{
		Message: "success",
		Data:    transactions,
	}

	if err = commonutils.HTTPResponse(w, response, http.StatusOK); err != nil {
		logrus.Error(err)
		http.Error(w, "Error writing http response", http.StatusInternalServerError)
		return
	}
}

func UpdateTransactionHandler(w http.ResponseWriter, r *http.Request) {
	userId, ok := commonutils.GetUserIdFromContext(r.Context())
	if !ok {
		logrus.Error("userId is empty")
		return
	}

	// get transaction id from url
	transactionId := mux.Vars(r)["transactionId"]

	if transactionId == "" {
		logrus.Error("No transaction id provided")
		http.Error(w, "invalid transaction id", http.StatusBadRequest)
		return
	}

	// validate request
	var request transactionsSerializer.UpdateTransactionRequest
	err := validation.ValidateRequest(w, r, &request)
	if err != nil {
		logrus.Infoln("issue is here", request)
		logrus.Error(err)
		return
	}

	sv := dbService.DBService{DB: config.DB}

	// if transaction date is not empty, convert date
	layout := "2006-01-02"
	transactionDate, err := time.Parse(layout, request.TransactionDate)
	if err != nil {
		logrus.Error(err)
		http.Error(w, "Invalid Transaction Date", http.StatusBadRequest)
		return
	}

	// if category is not empty, get or create a category
	var categoryModel *models.Category

	model := &models.Transaction{
		Model:           baseModels.Model{Id: transactionId},
		UserId:          userId,
		Amount:          request.Amount,
		TransactionDate: transactionDate,
		Type:            models.TransactionType(request.Type),
		Description:     request.Description,
	}

	condition := &models.Transaction{
		Model:  baseModels.Model{Id: transactionId},
		UserId: userId,
	}

	// get category id if category is provided
	if request.Category != "" {
		// get or create category
		conditionCategory := &models.Category{
			Name: request.Category,
			Type: models.TransactionType(request.Type),
		}

		category, err := sv.GetOrCreate(r.Context(), conditionCategory, conditionCategory)
		if err != nil {
			logrus.Error(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		categoryModel = category.(*models.Category)
		model.CategoryId = categoryModel.Id
	}

	// update transaction
	transaction, err := sv.UpdateAndGet(r.Context(), model, condition)
	if err != nil {
		logrus.Error(err)
		http.Error(w, "error updating transaction", http.StatusInternalServerError)
		return
	}

	response := commonutils.Response{
		Message: "success",
		Data:    transaction,
	}

	if err = commonutils.HTTPResponse(w, response, http.StatusOK); err != nil {
		logrus.Error(err)
		http.Error(w, "Error writing http response", http.StatusInternalServerError)
	}
}

func DeleteTransactionHandler(w http.ResponseWriter, r *http.Request) {
	userId, ok := commonutils.GetUserIdFromContext(r.Context())
	if !ok {
		logrus.Error("userId is empty")
		return
	}

	// get transaction id from url
	transactionId := mux.Vars(r)["transactionId"]

	if transactionId == "" {
		logrus.Error("No transaction id provided")
		http.Error(w, "invalid transaction id", http.StatusBadRequest)
		return
	}

	// delete transaction
	condition := &models.Transaction{
		Model:  baseModels.Model{Id: transactionId},
		UserId: userId,
	}

	sv := dbService.DBService{DB: config.DB}

	transaction, err := sv.Get(r.Context(), condition)
	if err != nil {
		logrus.Error(err)
		http.Error(w, "transaction does not exist", http.StatusInternalServerError)
		return
	}

	if transaction == nil {
		logrus.Error("transaction does not exist")
		http.Error(w, "transaction does not exist", http.StatusNotFound)
		return
	}

	err = sv.SoftDelete(r.Context(), condition)
	if err != nil {
		logrus.Error(err)
		http.Error(w, "error deleting transaction", http.StatusInternalServerError)
		return
	}

	response := commonutils.Response{
		Message: "transaction deleted successfully",
	}

	if err = commonutils.HTTPResponse(w, response, http.StatusOK); err != nil {
		logrus.Error(err)
		http.Error(w, "Error writing http response", http.StatusInternalServerError)
	}
}

func GetTransactionsOverviewHandler(w http.ResponseWriter, r *http.Request) {
	userId, ok := commonutils.GetUserIdFromContext(r.Context())
	if !ok {
		logrus.Error("userId is empty")
		return
	}

	// get start and end date variables from url
	startDate, endDate, err := commonutils.ExtractDatesFromUrl(r)
	if err != nil {
		logrus.Error("error retrieving start and end dates")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var response transactionsSerializer.FinancesOverviewResponse
	var income, expense sql.NullFloat64

	// get all transaction that are incomes for the specified time range
	sv := dbService.DBService{DB: config.DB}

	tx := sv.DB.WithContext(r.Context()).Begin()
	if tx.Error != nil {
		logrus.Error(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	query := `SELECT SUM(amount) 
FROM transactions
WHERE type = ?
AND user_id = ?
AND transaction_date BETWEEN ? AND ?`

	err = tx.Debug().Raw(query, models.Income, userId, *startDate, *endDate).Scan(&income).Error
	if err != nil {
		logrus.Error(err)

		// find out the error type
		if err.Error() != gorm.ErrRecordNotFound.Error() {
			tx.Rollback()
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		} else {
			tx.Rollback()
		}
	}

	if income.Valid {
		response.Income = income.Float64
	} else {
		response.Income = 0
	}

	// get all transaction that are expenses for the specified time range
	err = tx.Debug().Raw(query, models.Expense, userId, startDate, endDate).Scan(&expense).Error
	if err != nil {
		logrus.Error(err)
		// find out the error type
		if err.Error() != gorm.ErrRecordNotFound.Error() {
			tx.Rollback()
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		} else {
			tx.Rollback()
		}
	}

	if expense.Valid {
		response.Expense = expense.Float64
	} else {
		response.Expense = 0
	}

	defer tx.Commit()

	// calculate the balance
	if response.Income == 0 && response.Expense == 0 {
		response.Balance = 0
	} else {
		response.Balance = response.Income - response.Expense
	}

	// send response
	responseData := commonutils.Response{
		Message: "success",
		Data:    response,
	}

	if err = commonutils.HTTPResponse(w, responseData, http.StatusOK); err != nil {
		logrus.Error(err)
		http.Error(w, "Error writing http response", http.StatusInternalServerError)
	}
}

func GetTransactionStatsByCategoryHandler(w http.ResponseWriter, r *http.Request) {
	userId, ok := commonutils.GetUserIdFromContext(r.Context())
	if !ok {
		logrus.Error("userId is empty")
		return
	}

	// get the category name from url
	categoryName := mux.Vars(r)["categoryName"]

	if categoryName == "" {
		logrus.Error("category is empty")
		http.Error(w, "invalid category", http.StatusBadRequest)
		return
	}

	// get the category id
	categoryCondition := &models.Category{
		Name:   categoryName,
		UserId: userId,
	}
	sv := dbService.DBService{DB: config.DB}

	category, err := sv.Get(r.Context(), categoryCondition)
	if err != nil {
		logrus.Error(err)
		http.Error(w, "category does not exist", http.StatusNotFound)
		return
	}
	categoryId := category.(*models.Category).Id

	// get start and end date variables from url
	startDate, endDate, err := commonutils.ExtractDatesFromUrl(r)
	if err != nil {
		logrus.Error("error retrieving start and end dates")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// initialize response struct
	var response transactionsSerializer.FinancesStatsByCategoryResponse

	// initialize transaction
	tx := sv.DB.WithContext(r.Context()).Begin()
	if tx.Error != nil {
		logrus.Error(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// define time series query
	query := `SELECT
        COALESCE(SUM(t.amount), 0) AS transactions
    FROM
        (SELECT generate_series(?::date, ?::date, '1 day'::interval)::date AS date) AS date_sequence
    LEFT JOIN
        transactions t ON date_sequence.date = DATE(t.transaction_date)
        AND t.type = ? 
        AND t.category_id = ?
        AND t.user_id = ?
    WHERE
        t.transaction_date IS NOT NULL
    GROUP BY
        date_sequence.date
    ORDER BY
        date_sequence.date`

	// get incomes
	err = tx.Debug().Raw(query, startDate, endDate, models.Income, categoryId, userId).Scan(&response.DailyIncomes).Error
	if err != nil {
		logrus.Error(err)
		tx.Rollback()
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// get expenses
	err = tx.Debug().Raw(query, startDate, endDate, models.Expense, categoryId, userId).Scan(&response.DailyExpenses).Error
	if err != nil {
		logrus.Error(err)
		tx.Rollback()
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// send response
	responseData := commonutils.Response{
		Message: "success",
		Data:    response,
	}

	if err = commonutils.HTTPResponse(w, responseData, http.StatusOK); err != nil {
		logrus.Error(err)
		http.Error(w, "Error writing http response", http.StatusInternalServerError)
	}
}
