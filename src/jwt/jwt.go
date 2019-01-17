package jwt

import "time"
import "github.com/SermoDigital/jose/jws"
import "github.com/SermoDigital/jose/jwt"
import "github.com/SermoDigital/jose/crypto"
import "github.com/isayme/go-user/src/conf"
import "github.com/isayme/go-user/src/httperror"

// Verify verify token and return claims
func Verify(token string) (jwt.Claims, error) {
	t, err := jws.ParseJWT([]byte(token))
	if err != nil {
		return nil, err
	}

	cfg := conf.Get()
	method := jws.GetSigningMethod(cfg.JWT.Method)

	for _, key := range cfg.JWT.Keys {
		if err = t.Validate([]byte(key), method); err != nil {
			// try next key only if signature invalid
			if err == crypto.ErrSignatureInvalid {
				continue
			}
			break
		} else {
			// success
			break
		}
	}

	if err != nil {
		return nil, httperror.AccessTokenInvalid.WithMsg(err.Error())
	}

	return t.Claims(), nil
}

// Sign generate token
func Sign(data map[string]interface{}) (string, error) {
	payload := jws.Claims{}
	for k, v := range data {
		payload.Set(k, v)
	}

	now := time.Now()
	payload.SetIssuedAt(now)

	cfg := conf.Get()

	expire := time.Duration(cfg.JWT.Expire)
	if expire.Nanoseconds() != 0 {
		payload.SetExpiration(now.Add(expire))
	}

	method := jws.GetSigningMethod(cfg.JWT.Method)
	jwt := jws.NewJWT(payload, method)

	tokenBytes, err := jwt.Serialize([]byte(cfg.JWT.Keys[0]))
	if err != nil {
		return "", err
	}
	return string(tokenBytes), nil
}
