// +build integration

package dba

import (
	"testing"

	"github.com/jonstacks/pg-dba/pkg/config"
	"github.com/stretchr/testify/assert"
)

// Tests The automatic DBA
func TestDBA(t *testing.T) {
	dba, err := New(config.DBConnectionString())
	assert.Nil(err)

	assert.Nil(dba.Run())
}
