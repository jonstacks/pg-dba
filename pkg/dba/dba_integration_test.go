// +build integration

package dba

import (
	"testing"

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
