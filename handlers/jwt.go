package handlers

import (
	"crypto/rsa"
	"encoding/json"
	"net/http"
	"time"

	validator "gopkg.in/go-playground/validator.v9"

	"github.com/DataDog/datadog-go/statsd"
	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jws"
	bogcrypto "github.com/building-microservices-with-go/crypto"

	log "github.com/sirupsen/logrus"
)

const oneDay time.Duration = (1440 * time.Minute)

var validate = validator.New()
var defaultFields = log.Fields{
	"service": "jwt",
	"handler": "auth",
}

type LoginRequest struct {
	Username string `json:"username" validate:"email"`
	Password string `json:"password" validate:"max=36,min=8"`
}

type JWT struct {
	rsaPrivate *rsa.PrivateKey
	statsd     *statsd.Client
	logger     *log.Logger
}

// generateJWT creates a new JWT and signs it with the private key
func (j *JWT) generateJWT(request LoginRequest) []byte {
	claims := jws.Claims{}
	claims.SetExpiration(time.Now().Add(oneDay))
	claims.Set("userID", request.Username)
	claims.Set("accessLevel", "user")

	jwt := jws.NewJWT(claims, crypto.SigningMethodRS256)

	b, _ := jwt.Serialize(j.rsaPrivate)

	return b
}

func (j *JWT) Handle(rw http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		j.statsd.Incr("jwt.badmethod", nil, 1)

		j.logger.WithFields(defaultFields).Infof("Method: %s, not allowed", r.Method)
		rw.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	request := LoginRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		j.statsd.Incr("jwt.badrequest", nil, 1)

		j.logger.WithFields(defaultFields).Errorf("Error decoding request %v", err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	err = validate.Struct(request)
	if err != nil {
		j.statsd.Incr("jwt.badrequest", nil, 1)

		j.logger.WithFields(defaultFields).Errorf("Error validating request %s", err.Error())
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	j.logger.WithFields(defaultFields).Infof("Login request from %s", request.Username)
	jwt := j.generateJWT(request)

	j.statsd.Incr("jwt.success", nil, 1)

	rw.Write(jwt)
}

func NewJWT(logger *log.Logger, statsd *statsd.Client) *JWT {
	var err error
	rsaPrivate, err := bogcrypto.UnmarshalRSAPrivateKeyFromFile("./sample_key.priv")
	if err != nil {
		logger.WithFields(defaultFields).Fatalf("Unable to parse private key: %v", err)
	}

	return &JWT{
		rsaPrivate: rsaPrivate,
		logger:     logger,
		statsd:     statsd,
	}
}
