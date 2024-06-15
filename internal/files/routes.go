package files

import (
	"database/sql"

	"github.com/JoaoPauloFontana/drive/internal/bucket"
	"github.com/JoaoPauloFontana/drive/internal/queue"
	"github.com/go-chi/chi"
)

type handler struct {
	db     *sql.DB
	bucket *bucket.Bucket
	queue  *queue.Queue
}

func setRoutes(r chi.Router, db *sql.DB, b *bucket.Bucket, q *queue.Queue) {
	h := handler{db, b, q}

	r.Post("/", h.Create)
	r.Put("/{id}", h.Modify)
}
