// +build integration

package dba

import (
	"testing"

	"github.com/jonstacks/pg-dba/pkg/config"
	"github.com/stretchr/testify/assert"
)

// Tests The automatic DBA
func TestDBA(t *testing.T) {
	dba := New(config.DBConnectionString(), config.Verbose())
	assert.Nil(t, dba.Run())
}
