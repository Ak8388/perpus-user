package jwt

import (
	"latihangolanguser/entity"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func JWT(id string, nim string) (*entity.TokenAkses, error) {
	errEnv := godotenv.Load()

	if errEnv != nil {
		return nil, errEnv
	}

	form := os.Getenv("FORM")
	secret := os.Getenv("SECRET")

	ta := &entity.TokenAkses{}
	exp := time.Now().Add(time.Hour * 1)
	uid, err := uuid.NewRandom()

	if err != nil {
		return nil, err
	}

	refuid, err2 := uuid.NewRandom()

	if err2 != nil {
		return nil, err2
	}

	ta.ExpAt = exp.Unix()
	ta.UidTok = uid.String()
	ta.UidRefTok = refuid.String()

	Claims := entity.Claims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    form,
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: ta.ExpAt,
			Id:        ta.UidTok,
		},
		Uid: id,
		Nim: nim,
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		Claims,
	)

	sigAccs, err3 := token.SignedString([]byte(secret))

	if err3 != nil {
		return nil, err3
	}

	RefClaims := entity.Claims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    form,
			IssuedAt:  time.Now().Unix(),
			NotBefore: time.Now().Unix(),
			Id:        ta.UidRefTok,
		},
		Uid: id,
		Nim: nim,
	}

	tokenRef := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		RefClaims,
	)

	refSig, errref := tokenRef.SignedString([]byte(secret))

	if errref != nil {
		return nil, errref
	}

	ta.RefAksesToken = refSig
	ta.TokenAkses = sigAccs
	return ta, nil
}
