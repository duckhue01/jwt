package jwt

import (
	"fmt"
	"time"

	"github.com/fatih/color"
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

func isExpire(exp time.Time) *IsExpireResp {
	res := &IsExpireResp{
		now:   time.Now(),
		expAt: exp,
	}

	if res.expAt.Before(res.now) {
		res.isExpire = true
	} else {
		res.isExpire = false
	}

	return res
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

		res := isExpire(exp.Time)
		if res.isExpire {
			color.Red(fmt.Sprintf("Token expire time: %s, now: %s\n", res.expAt.Format(time.ANSIC), res.now.Format(time.ANSIC)))
		} else {
			color.Green(fmt.Sprintf("Token expire time: %s, now: %s\n", res.expAt.Format(time.ANSIC), res.now.Format(time.ANSIC)))
		}
	}

	if len(requiredClaims) > 0 {
		claims := token.Claims.(jwt.MapClaims)
		for _, v := range requiredClaims {
			res, ok := claims[v]
			if ok {
				color.Green(fmt.Sprintf("Expected claim: %v: %v", v, res))
			} else {
				color.Red(fmt.Sprintf("Missing claim: %v", v))
			}
		}
	}
	return nil
}
