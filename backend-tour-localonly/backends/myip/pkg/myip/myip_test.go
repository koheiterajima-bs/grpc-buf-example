// Package myip は、リソースの実装
package myip

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {

	var name string
	t.Run("create", func(t *testing.T) {
		got, err := Create(context.Background(), IPV4)
		assert.NoError(t, err)
		name = got.Name
	})

	t.Run("get", func(t *testing.T) {
		got, ok := Get(context.Background(), name)
		assert.True(t, ok)
		assert.Equal(t, name, got.Name)
	})

}
