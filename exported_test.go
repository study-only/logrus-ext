package logrus_ext

import (
	"errors"
	"testing"
)

var log = Get("usecase", &Option{WithFunc: true})

func TestLogger_Errorf(t *testing.T) {
	repo := db{}
	repo.QueryTask(1)
}

type db struct {
}

func (s *db) QueryTask(id int) {
	err := errors.New("too many connection")
	log.WithError(err).Errorf("query task error: id=%d", id)
}
