package sentry

import "github.com/getsentry/sentry-go"

// NewInit initializes sentry in run time
func NewInit(dsn string) error {
	return sentry.Init(sentry.ClientOptions{Dsn: dsn, AttachStacktrace: true})
}
