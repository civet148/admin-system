package middleware

import (
	"admin-system/pkg/types"
	"encoding/json"
	"fmt"
	"github.com/civet148/log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const (
	CLAIM_EXPIRE       = "claim_expire"
	CLAIM_ISSUE_AT     = "claim_iat"
	CLAIM_USER_SESSION = "user_session"
)

const (
	DEFAULT_TOKEN_DURATION = 8760 * time.Hour // default one year
)

const (
	jwtTokenSecret = "7bdf27cffd5fd105af4efb20b1090bbe"
)

type JwtCode int

const (
	JWT_CODE_SUCCESS             JwtCode = 0
	JWT_CODE_ERROR_CHECK_TOKEN   JwtCode = -1
	JWT_CODE_ERROR_PARSE_TOKEN   JwtCode = -2
	JWT_CODE_ERROR_INVALID_TOKEN JwtCode = -3
	JWT_CODE_ERROR_TOKEN_EXPIRED JwtCode = -4
)

var codeMessages = map[JwtCode]string{
	JWT_CODE_SUCCESS:             "JWT_CODE_SUCCESS",
	JWT_CODE_ERROR_CHECK_TOKEN:   "JWT_CODE_ERROR_CHECK_TOKEN",
	JWT_CODE_ERROR_PARSE_TOKEN:   "JWT_CODE_ERROR_PARSE_TOKEN",
	JWT_CODE_ERROR_INVALID_TOKEN: "JWT_CODE_ERROR_INVALID_TOKEN",
	JWT_CODE_ERROR_TOKEN_EXPIRED: "JWT_CODE_ERROR_TOKEN_EXPIRED",
}

func (j JwtCode) GoString() string {
	return j.String()
}

func (j JwtCode) String() string {
	strMessage, ok := codeMessages[j]
	if ok {
		return strMessage
	}
	return fmt.Sprintf("JWT_CODE_UNKONWN<%d>", j)
}

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data interface{}
		if err := ParseToken(c); err != nil {
			c.JSON(http.StatusUnauthorized, types.HttpResponse{
				Header: types.HttpHeader{
					Code:    types.CODE_UNAUTHORIZED,
					Message: "unauthorized",
					Count:   0,
				},
				Data: data,
			})
			log.Errorf("[JWT] token parse failed, error [%s]", err.Error())
			c.Abort()
			return
		}

		c.Next()
	}
}

// generate JWT token
func GenerateToken(session interface{}, duration ...interface{}) (token string, err error) {

	var d time.Duration
	var claims = make(jwt.MapClaims)

	if len(duration) == 0 {
		d = DEFAULT_TOKEN_DURATION
	} else {
		var ok bool
		if d, ok = duration[0].(time.Duration); !ok {
			d = DEFAULT_TOKEN_DURATION
		}
	}
	var data []byte
	data, err = json.Marshal(session)
	if err != nil {
		return token, log.Errorf(err.Error())
	}
	sign := jwt.New(jwt.SigningMethodHS256)
	claims[CLAIM_EXPIRE] = time.Now().Add(d).Unix()
	claims[CLAIM_ISSUE_AT] = time.Now().Unix()
	claims[CLAIM_USER_SESSION] = string(data)
	sign.Claims = claims

	token, err = sign.SignedString([]byte(jwtTokenSecret))
	return token, err
}

// parse JWT token claims
func ParseToken(c *gin.Context) error {
	strAuthToken := GetAuthToken(c)
	if strAuthToken == "" {
		return log.Errorf("[JWT] request header have no any key '%s'", types.HEADER_AUTH_TOKEN)
	}
	claims, err := ParseTokenClaims(strAuthToken)
	if err != nil {
		return log.Errorf(err.Error())
	}
	c.Keys = make(map[string]interface{})
	c.Keys[CLAIM_EXPIRE] = int64(claims[CLAIM_EXPIRE].(float64))
	if c.Keys[CLAIM_EXPIRE].(int64) < time.Now().Unix() {
		return log.Errorf("[JWT] token [%s] expired at %v\n", strAuthToken, c.Keys[CLAIM_EXPIRE])
	}

	c.Keys[CLAIM_EXPIRE] = int64(claims[CLAIM_EXPIRE].(float64))
	c.Keys[CLAIM_ISSUE_AT] = int64(claims[CLAIM_ISSUE_AT].(float64))
	c.Keys[CLAIM_USER_SESSION] = claims[CLAIM_USER_SESSION].(string)
	return nil
}

func ParseTokenClaims(strAuthToken string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(strAuthToken, func(*jwt.Token) (interface{}, error) {
		return []byte(jwtTokenSecret), nil
	})
	if err != nil {
		return jwt.MapClaims{}, log.Errorf("[JWT] parse token error [%s]", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return jwt.MapClaims{}, log.Errorf("[JWT] parse token error: no claims found")
	}
	return claims, nil
}

func GetAuthToken(c *gin.Context) string {
	return c.Request.Header.Get(types.HEADER_AUTH_TOKEN)
}

func GetAuthSession(strAuthToken string, session interface{}) error {
	claims, err := ParseTokenClaims(strAuthToken)
	if err != nil {
		return log.Errorf(err.Error())
	}
	strSessionJson := claims[CLAIM_USER_SESSION].(string)
	err = json.Unmarshal([]byte(strSessionJson), session)
	if err != nil {
		return log.Errorf(err.Error())
	}
	return nil
}
