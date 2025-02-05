package billing

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ger9000/memes-as-a-service/internal/delivery/rest"
	"github.com/ger9000/memes-as-a-service/internal/infrastructure"
	"github.com/ger9000/memes-as-a-service/internal/usecases/billing"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Test_RechargeController(t *testing.T) {
	t.Run("OK response", func(t *testing.T) {
		// Given
		controller, db := Init()

		token := "test"
		body, err := json.Marshal(RechargeAvailableCallRequest{
			Token:            token,
			AmountToRecharge: 10,
		})
		if err != nil {
			assert.Fail(t, "cannot decode body")
		}

		r := httptest.NewRequest("POST", "/billing/recharge", bytes.NewBuffer(body))
		w := httptest.NewRecorder()

		db.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "tracks" WHERE token = $1 ORDER BY "tracks"."token" LIMIT $2`)).WithArgs(token, 1).
			WillReturnRows(sqlmock.NewRows([]string{"token", "available_calls"}).AddRow(token, 1))
		db.ExpectBegin()
		db.ExpectExec("UPDATE").
			WillReturnResult(sqlmock.NewResult(1, 1))
		db.ExpectCommit()

		// When
		controller.Invoke(w, r)

		// Then
		assert.Equal(t, http.StatusNoContent, w.Code)
	})

	t.Run("Error response given invalid body", func(t *testing.T) {
		// Given
		controller, _ := Init()

		r := httptest.NewRequest("POST", "/billing/recharge", bytes.NewBuffer([]byte("body")))
		w := httptest.NewRecorder()

		// When
		controller.Invoke(w, r)

		// Then
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Ok response when token has never been tracked", func(t *testing.T) {
		// Given
		controller, db := Init()

		token := "test"
		body, err := json.Marshal(RechargeAvailableCallRequest{
			Token:            token,
			AmountToRecharge: 10,
		})
		if err != nil {
			assert.Fail(t, "cannot decode body")
		}

		r := httptest.NewRequest("POST", "/billing/recharge", bytes.NewBuffer(body))
		w := httptest.NewRecorder()

		db.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "tracks" WHERE token = $1 ORDER BY "tracks"."token" LIMIT $2`)).WithArgs(token, 1).
			WillReturnRows(&sqlmock.Rows{})
		db.ExpectBegin()
		db.ExpectExec("UPDATE").WillReturnResult(sqlmock.NewResult(1, 1))
		db.ExpectCommit()

		// When
		controller.Invoke(w, r)

		// Then
		assert.Equal(t, http.StatusNoContent, w.Code)
	})

	t.Run("Error response when an error is throwed finding track", func(t *testing.T) {
		// Given
		controller, db := Init()

		token := "test"
		body, err := json.Marshal(RechargeAvailableCallRequest{
			Token:            token,
			AmountToRecharge: 10,
		})
		if err != nil {
			assert.Fail(t, "cannot decode body")
		}

		r := httptest.NewRequest("POST", "/billing/recharge", bytes.NewBuffer(body))
		w := httptest.NewRecorder()

		db.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "tracks" WHERE token = $1 ORDER BY "tracks"."token" LIMIT $2`)).WithArgs(token, 1).
			WillReturnError(errors.New("error on db instance"))
		db.ExpectBegin()
		db.ExpectExec("UPDATE").WillReturnResult(sqlmock.NewResult(1, 1))
		db.ExpectCommit()

		// When
		controller.Invoke(w, r)

		// Then
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("Error response when an error is throwed updating track", func(t *testing.T) {
		// Given
		controller, db := Init()

		token := "test"
		body, err := json.Marshal(RechargeAvailableCallRequest{
			Token:            token,
			AmountToRecharge: 10,
		})
		if err != nil {
			assert.Fail(t, "cannot decode body")
		}

		r := httptest.NewRequest("POST", "/billing/recharge", bytes.NewBuffer(body))
		w := httptest.NewRecorder()

		db.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "tracks" WHERE token = $1 ORDER BY "tracks"."token" LIMIT $2`)).WithArgs(token, 1).
			WillReturnRows(sqlmock.NewRows([]string{"token", "available_calls"}).AddRow(token, 1))
		db.ExpectBegin()
		db.ExpectExec("UPDATE").WillReturnError(errors.New("error on db instance"))
		db.ExpectCommit()

		// When
		controller.Invoke(w, r)

		// Then
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

func Init() (rest.IController, sqlmock.Sqlmock) {
	mockDb, mock, _ := sqlmock.New()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})

	service := infrastructure.NewRepository(db)
	action := billing.NewRechargeAvailableCallAction(service)

	return NewRechargeAvailableCallController(action), mock
}
