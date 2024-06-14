package files

import (
	"database/sql"

	"github.com/go-chi/chi"
)

type handler struct {
	db *sql.DB
}

func setRoutes(r chi.Router, db *sql.DB) {
	h := handler{db}

	r.Put("/{id}", h.Modify)
}
