package users

import (
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestList(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Error(err)
	}

	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "login", "password", "created_at", "modified_at", "deleted", "last_login"}).
		AddRow(1, "João", "jpfontana12", "123456", time.Now(), time.Now(), false, time.Now()).
		AddRow(2, "João Paulo", "jpfontana123", "123456", time.Now(), time.Now(), false, time.Now())

	mock.ExpectQuery(regexp.QuoteMeta(`select * from "users" where deleted = false`)).
		WillReturnRows(rows)

	_, err = SelectAll(db)

	if err != nil {
		t.Error(err)
	}

	err = mock.ExpectationsWereMet()

	if err != nil {
		t.Error(err)
	}
}
