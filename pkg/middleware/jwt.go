package middleware

import (
	"admin-system/pkg/types"
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
	CLAIM_USER_ID      = "claim_id"
	CLAIM_USER_NAME    = "claim_username"
	CLAIM_USER_ALIAS   = "claim_alias"
	CLAIM_PHONE_NUMBER = "claim_phone_number"
	CLAIM_IS_ADMIN     = "claim_is_admin"
)

const (
	DEFAULT_TOKEN_DURATION = 12 * time.Hour
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
		var code JwtCode
		var data interface{}

		if code = ParseToken(c); code != JWT_CODE_SUCCESS {
			c.JSON(http.StatusUnauthorized, types.HttpResponse{
				Header: types.HttpHeader{
					Code:    types.CODE_UNAUTHORIZED,
					Message: "unauthorized",
					Count:   0,
				},
				Data: data,
			})
			log.Errorf("[JWT] token parse failed, error code [%s]", code.String())
			c.Abort()
			return
		}

		c.Next()
	}
}

// generate JWT token
func GenerateToken(s *types.Session, duration ...interface{}) (token string, claims jwt.MapClaims, err error) {

	var d time.Duration
	if len(duration) == 0 {
		d = DEFAULT_TOKEN_DURATION
	} else {
		var ok bool
		if d, ok = duration[0].(time.Duration); !ok {
			d = DEFAULT_TOKEN_DURATION
		}
	}
	sign := jwt.New(jwt.SigningMethodHS256)
	claims = make(jwt.MapClaims)
	claims[CLAIM_EXPIRE] = time.Now().Add(d).Unix()
	claims[CLAIM_ISSUE_AT] = time.Now().Unix()
	claims[CLAIM_USER_ID] = s.UserId
	claims[CLAIM_USER_NAME] = s.UserName
	claims[CLAIM_USER_ALIAS] = s.Alias
	claims[CLAIM_PHONE_NUMBER] = s.PhoneNumber
	claims[CLAIM_IS_ADMIN] = s.IsAdmin
	sign.Claims = claims

	token, err = sign.SignedString([]byte(jwtTokenSecret))
	return token, claims, err
}

// parse JWT token claims
func ParseToken(c *gin.Context) JwtCode {
	authToken := GetAuthToken(c)
	if authToken == "" {
		log.Errorf("[JWT] request header have no any key '%s'", types.HEADER_AUTH_TOKEN)
		return JWT_CODE_ERROR_CHECK_TOKEN
	}
	token, err := jwt.Parse(authToken, func(*jwt.Token) (interface{}, error) {
		return []byte(jwtTokenSecret), nil
	})
	if err != nil {
		log.Errorf("[JWT] parse token error [%s]", err)
		return JWT_CODE_ERROR_PARSE_TOKEN
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !CheckClaims(claims) {
		log.Errorf("[JWT] token [%s] have no claims or check failed", authToken)
		return JWT_CODE_ERROR_INVALID_TOKEN
	}

	c.Keys = make(map[string]interface{})
	c.Keys[CLAIM_EXPIRE] = int64(claims[CLAIM_EXPIRE].(float64))
	if c.Keys[CLAIM_EXPIRE].(int64) < time.Now().Unix() {
		log.Errorf("[JWT] token [%s] expired at %v\n", authToken, c.Keys[CLAIM_EXPIRE])
		return JWT_CODE_ERROR_TOKEN_EXPIRED
	}

	c.Keys[CLAIM_EXPIRE] = int64(claims[CLAIM_EXPIRE].(float64))
	c.Keys[CLAIM_ISSUE_AT] = int64(claims[CLAIM_ISSUE_AT].(float64))
	c.Keys[CLAIM_USER_ID] = int32(claims[CLAIM_USER_ID].(float64))
	c.Keys[CLAIM_USER_NAME] = claims[CLAIM_USER_NAME].(string)
	c.Keys[CLAIM_USER_ALIAS] = claims[CLAIM_USER_ALIAS].(string)
	c.Keys[CLAIM_PHONE_NUMBER] = claims[CLAIM_PHONE_NUMBER].(string)
	c.Keys[CLAIM_IS_ADMIN] = claims[CLAIM_IS_ADMIN].(bool)
	return JWT_CODE_SUCCESS
}

func CheckClaims(claims jwt.MapClaims) bool {

	if _, ok := claims[CLAIM_EXPIRE]; !ok {
		log.Errorf("[JWT] claims of CLAIM_EXPIRE is nil")
		return false
	}
	if _, ok := claims[CLAIM_ISSUE_AT]; !ok {
		log.Errorf("[JWT] claims of CLAIM_ISSUE_AT is nil")
		return false
	}
	if _, ok := claims[CLAIM_USER_ID]; !ok {
		log.Errorf("[JWT] claims of CLAIM_USER_ID is nil")
		return false
	}
	if _, ok := claims[CLAIM_USER_NAME]; !ok {
		log.Errorf("[JWT] claims of CLAIM_USER_NAME is nil")
		return false
	}
	if _, ok := claims[CLAIM_USER_ALIAS]; !ok {
		log.Errorf("[JWT] claims of CLAIM_USER_ALIAS is nil")
		return false
	}
	if _, ok := claims[CLAIM_PHONE_NUMBER]; !ok {
		log.Errorf("[JWT] claims of CLAIM_PHONE_NUMBER is nil")
		return false
	}
	if _, ok := claims[CLAIM_IS_ADMIN]; !ok {
		log.Errorf("[JWT] claims of CLAIM_IS_ADMIN is nil")
		return false
	}
	return true
}

func GetAuthToken(c *gin.Context) string {
	return c.Request.Header.Get(types.HEADER_AUTH_TOKEN)
}

func GetClaimExpire(c *gin.Context) int64 {
	if _, ok := c.Keys[CLAIM_EXPIRE]; !ok {
		log.Errorf("[JWT] gin context keys of CLAIM_EXPIRE is nil")
		return 0
	}
	return int64(c.Keys[CLAIM_EXPIRE].(float64))
}

func GetClaimIssueAt(c *gin.Context) int64 {
	if _, ok := c.Keys[CLAIM_ISSUE_AT]; !ok {
		log.Errorf("[JWT] gin context keys of CLAIM_ISSUE_AT is nil")
		return 0
	}
	return int64(c.Keys[CLAIM_ISSUE_AT].(float64))
}

func GetClaimUserId(c *gin.Context) int32 {
	if _, ok := c.Keys[CLAIM_USER_ID]; !ok {
		log.Errorf("[JWT] gin context keys of CLAIM_USER_ID is nil")
		return 0
	}
	return c.Keys[CLAIM_USER_ID].(int32)
}

func GetClaimUserName(c *gin.Context) string {
	if _, ok := c.Keys[CLAIM_USER_NAME]; !ok {
		log.Errorf("[JWT] gin context keys of CLAIM_USER_NAME is nil")
		return ""
	}
	return c.Keys[CLAIM_USER_NAME].(string)
}
func GetClaimUserAlias(c *gin.Context) string {
	if _, ok := c.Keys[CLAIM_USER_ALIAS]; !ok {
		log.Errorf("[JWT] gin context keys of CLAIM_USER_ALIAS is nil")
		return ""
	}
	return c.Keys[CLAIM_USER_ALIAS].(string)
}
func GetClaimPhoneNumber(c *gin.Context) string {
	if _, ok := c.Keys[CLAIM_PHONE_NUMBER]; !ok {
		log.Errorf("[JWT] gin context keys of CLAIM_PHONE_NUMBER is nil")
		return ""
	}
	return c.Keys[CLAIM_PHONE_NUMBER].(string)
}

func GetClaimIsAdmin(c *gin.Context) bool {
	if _, ok := c.Keys[CLAIM_IS_ADMIN]; !ok {
		log.Errorf("[JWT] gin context keys of CLAIM_IS_ADMIN is nil")
		return false
	}
	return c.Keys[CLAIM_IS_ADMIN].(bool)
}
