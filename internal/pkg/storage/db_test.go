package storage

import (
	"github.com/Azatik1000/distsys-hw/internal/pkg/models"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
	"testing"
)

type env struct {
	storage Storage
}

func newEnv() (*env, func(), error) {
	storage, err := NewDB()
	if err != nil {
		return nil, nil, err
	}

	stop := func() {
		_ = storage.Close()
	}

	return &env{storage}, stop, nil
}

func TestSimpleAdd(t *testing.T) {
	env, stop, err := newEnv()
	require.NoError(t, err)
	defer stop()

	added, err := env.storage.AddProduct(&models.Product{
		Name:  "some-name",
		Code:  "some-code",
		Kind:  models.ProductKind(1),
	})

	require.NoError(t, err)
	found, err := env.storage.GetProduct(added.ID)

	require.NoError(t, err)
	require.True(t, cmp.Equal(added, found))
}