package middleware

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

type ctxKey uint

const (
	ctxKeyLog ctxKey = iota
	//	ctxKeyDB
)

func GetLogger(r *http.Request) *logrus.Entry {
	return r.Context().Value(ctxKeyLog).(*logrus.Entry)
}

//func GetDB(r *http.Request) data.DBConnector {
//	return r.Context().Value(ctxKeyDB).(data.DBConnector)
//}
