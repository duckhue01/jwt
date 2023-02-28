package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/pflag"
)

const (
	IsExpireFlag = "is-expire"
	IsValidFlag  = "is-valid"
	ContainFlag  = "contain"
)

type IsExpireResp struct {
	expAt    time.Time
	now      time.Time
	isExpire bool
}

func IsExpire(exp time.Time) *IsExpireResp {
	res := &IsExpireResp{
		now:   time.Now(),
		expAt: exp,
	}

	if res.expAt.Before(exp) {
		res.isExpire = false

	} else {
		res.isExpire = true

	}
	return res
}

func IsContain(claims jwt.Claims) map[string]bool {
	return make(map[string]bool)
}

// todo: implement later
func IsValid(token string) bool {
	return true
}

func VerifyHandler(eTok string, flags *pflag.FlagSet) error {
	token, err := jwt.Parse(eTok, nil)
	if token == nil {
		return fmt.Errorf("malformed token: %w", err)
	}

	checkExpire, _ := flags.GetBool(IsExpireFlag)
	requiredClaims, err := flags.GetStringSlice(ContainFlag)

	if err != nil {
		return fmt.Errorf("can't get %s flag: %w", IsExpireFlag, err)
	}

	if checkExpire {
		exp, err := token.Claims.GetExpirationTime()
		if err != nil {
			return fmt.Errorf("can't get claim: %w", err)
		}

		res := IsExpire(exp.Time)
		fmt.Println(res.expAt.Format(time.ANSIC))
		fmt.Println(res.now.Format(time.ANSIC))
		fmt.Println(res.isExpire)
	}

	if len(requiredClaims) > 0 {
		claims := token.Claims.(jwt.MapClaims)

		for _, v := range requiredClaims {
			_, ok := claims[v]
			if !ok {
				fmt.Println("don't have", v)
			}
		}
	}
	return nil
}
