package jwt

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSignVerify(t *testing.T) {
	require := require.New(t)

	d := map[string]interface{}{
		"abcd": "123",
	}

	token, err := Sign(d)
	t.Log(token)
	require.Nil(err)

	claims, err := Verify(token)
	require.Nil(err)
	require.Equal(claims.Get("abcd").(string), "123")
}
