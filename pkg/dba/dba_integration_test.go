// +build integration

package dba

import (
	"testing"
	"time"

	"github.com/jonstacks/pg-dba/pkg/config"
	"github.com/stretchr/testify/assert"
)

// Tests The automatic DBA
func TestDBACase1(t *testing.T) {
	opts := NewDefaultOptions()
	opts.Verbose = config.Verbose()
	dba := New(config.DBConnectionString(), opts)
	assert.Nil(t, dba.Run())
}

func TestDBANonVerbose(t *testing.T) {
	opts := NewDefaultOptions()
	opts.Verbose = false
	dba := New(config.DBConnectionString(), opts)
	assert.Nil(t, dba.Run())
}

func TestDBATimeout(t *testing.T) {
	opts := NewDefaultOptions()
	opts.AnalyzeTimeout = 50 * time.Millisecond
	opts.Verbose = config.Verbose()
	dba := New(config.DBConnectionString(), opts)
	assert.Error(t, dba.Run())
}
