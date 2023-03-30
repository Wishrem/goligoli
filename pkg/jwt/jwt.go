package jwt

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/wishrem/goligoli/pkg/conf"
	"github.com/wishrem/goligoli/pkg/e"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Claims struct {
	jwt.StandardClaims
	UserID int64
	Roles  []string
}

func (c *Claims) Valid() error {
	if !c.VerifyExpiresAt(time.Now().Unix(), true) {
		return status.Error(codes.PermissionDenied, "Token Has Expired")
	}
	if !c.VerifyIssuer(conf.App.JWT.Issuer, true) {
		return status.Error(codes.PermissionDenied, "Invalid Issuer")
	}
	return nil
}

func Parse(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(conf.App.JWT.Secret), nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, status.Error(codes.PermissionDenied, e.UNAUTHENTICATED)
	}
}

func Generate(userID int64, roles []string) (string, error) {
	expAt := time.Now().Add(conf.App.JWT.Exp)
	claims := &Claims{
		jwt.StandardClaims{
			ExpiresAt: expAt.Unix(),
			Issuer:    conf.App.JWT.Issuer,
		},
		userID,
		roles,
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod(conf.App.JWT.SigningMethod), claims)
	ss, err := token.SignedString([]byte(conf.App.JWT.Secret))
	if err != nil {
		return "", status.Error(codes.Internal, e.INTERNAL)
	}
	return ss, nil
}
