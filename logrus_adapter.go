/*
 * watermill-logrus-adapter
 *     Copyright (c) 2022, Matthias Hartmann <mahartma@mahartma.com>
 *
 *   For license see LICENSE
 */

package wla

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/sirupsen/logrus"
)

// LogrusLoggerAdapter is a watermill logger adapter for logrus.
type LogrusLoggerAdapter struct {
	log    *logrus.Logger
	fields watermill.LogFields
}

// NewLogrusLogger returns a LogrusLoggerAdapter that sends all logs to
// the passed logrus instance.
func NewLogrusLogger(log *logrus.Logger) watermill.LoggerAdapter {
	return &LogrusLoggerAdapter{log: log}
}

// Error logs on level error with err as field and optional fields.
func (l *LogrusLoggerAdapter) Error(msg string, err error, fields watermill.LogFields) {
	l.createEntry(fields.Add(watermill.LogFields{"err": err})).Error(msg)
}

// Info logs on level info with optional fields.
func (l *LogrusLoggerAdapter) Info(msg string, fields watermill.LogFields) {
	l.createEntry(fields).Info(msg)
}

// Debug logs on level debug with optional fields.
func (l *LogrusLoggerAdapter) Debug(msg string, fields watermill.LogFields) {
	l.createEntry(fields).Debug(msg)
}

// Trace logs on level trace with optional fields.
func (l *LogrusLoggerAdapter) Trace(msg string, fields watermill.LogFields) {
	l.createEntry(fields).Trace(msg)
}

// With returns a new LogrusLoggerAdapter that includes fields
// to be re-used between logging statements.
func (l *LogrusLoggerAdapter) With(fields watermill.LogFields) watermill.LoggerAdapter {
	return &LogrusLoggerAdapter{
		log:    l.log,
		fields: l.fields.Add(fields),
	}
}

// createEntry is a helper to add fields to a logrus entry if necessary.
func (l *LogrusLoggerAdapter) createEntry(fields watermill.LogFields) *logrus.Entry {
	entry := logrus.NewEntry(l.log)

	allFields := fields.Add(l.fields)

	if len(allFields) > 0 {
		entry = entry.WithFields(logrus.Fields(allFields))
	}

	return entry
}
