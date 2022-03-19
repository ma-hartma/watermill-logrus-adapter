/*
 * watermill-logrus-adapter
 *     Copyright (c) 2022, Matthias Hartmann <mahartma@mahartma.com>
 *
 *   For license see LICENSE
 */

package wla

import (
	"bytes"
	"testing"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func setup() (*logrus.Logger, *bytes.Buffer) {
	buf := bytes.NewBuffer([]byte{})

	log := logrus.New()
	log.Formatter = &logrus.TextFormatter{}
	log.Out = buf
	log.Level = logrus.TraceLevel

	return log, buf
}

func TestLogrusLoggerAdapter_no_fields(t *testing.T) {
	log, buf := setup()

	logger := NewLogrusLogger(log)

	logger.Info("foo", watermill.LogFields{})

	out := buf.String()
	assert.Contains(t, out, `level=info msg=foo`)
}

func TestLogrusLoggerAdapter_field_with_space(t *testing.T) {
	log, buf := setup()

	logger := NewLogrusLogger(log)

	logger.Info("foo", watermill.LogFields{"foo": `bar baz`})

	out := buf.String()
	assert.Contains(t, out, `level=info msg=foo foo="bar baz"`)
}

func TestLogrusLoggerAdapter_With_with_fields(t *testing.T) {
	log, buf := setup()

	baseLogger := NewLogrusLogger(log)
	fieldsLogger := baseLogger.With(watermill.LogFields{"foo": "1"})

	for msg, logger := range map[string]watermill.LoggerAdapter{"base": baseLogger, "with": fieldsLogger} {
		logger.Error(msg, nil, watermill.LogFields{"bar": "2"})
		logger.Info(msg, watermill.LogFields{"bar": "3"})
		logger.Debug(msg, watermill.LogFields{"bar": "4"})
		logger.Trace(msg, watermill.LogFields{"bar": "5"})
	}

	out := buf.String()

	assert.Contains(t, out, `level=error msg=base bar=2 err="<nil>"`)
	assert.Contains(t, out, `level=info msg=base bar=3`)
	assert.Contains(t, out, `level=debug msg=base bar=4`)
	assert.Contains(t, out, `level=trace msg=base bar=5`)

	assert.Contains(t, out, `level=error msg=with bar=2 err="<nil>" foo=1`)
	assert.Contains(t, out, `level=info msg=with bar=3 foo=1`)
	assert.Contains(t, out, `level=debug msg=with bar=4 foo=1`)
	assert.Contains(t, out, `level=trace msg=with bar=5 foo=1`)
}
