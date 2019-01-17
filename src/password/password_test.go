package password

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPassword(t *testing.T) {
	require := require.New(t)

	t.Run("Generate", func(t *testing.T) {
		password := "123456"
		hash, err := Generate(password)
		t.Log(hash)
		require.Nil(err)

		require.NotEqual(hash, password)
	})

	t.Run("Compare", func(t *testing.T) {
		password := "123456"
		hash, err := Generate(password)
		t.Log(hash)
		require.Nil(err)

		require.False(Compare(hash, "abcde"))
		require.True(Compare(hash, password))
	})
}
