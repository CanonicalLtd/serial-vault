// -*- Mode: Go; indent-tabs-mode: t -*-

/*
 * Copyright (C) 2017-2018 Canonical Ltd
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 3 as
 * published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package usso

import (
	"errors"
	"log"
	"strings"
	"time"

	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/juju/usso/openid"
)

// NewJWTToken creates a new JWT from the verified OpenID response
func NewJWTToken(resp *openid.Response) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	token.Claims[ClaimsUsername] = resp.SReg["nickname"]
	token.Claims[ClaimsName] = resp.SReg["fullname"]
	token.Claims[ClaimsEmail] = resp.SReg["email"]
	token.Claims[ClaimsIdentity] = resp.ID
	token.Claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		log.Printf("Error signing the JWT: %v", err.Error())
	}
	return tokenString, err
}

func keyFunc(token *jwt.Token) (interface{}, error) {
	return []byte(jwtSecret), nil
}

// VerifyJWT checks that we have a valid token
func VerifyJWT(jwtToken string) (*jwt.Token, error) {

	token, err := jwt.Parse(jwtToken, keyFunc)

	return token, err
}

// AddJWTCookie sets the JWT as a cookie
func AddJWTCookie(jwtToken string, w http.ResponseWriter) {

	// Set the JWT as a bearer token
	// (In practice, the cookie will be used more as clicking on a page link will not send the auth header)
	w.Header().Set("Authorization", "Bearer "+jwtToken)

	expireCookie := time.Now().Add(time.Hour * 1)
	cookie := http.Cookie{Name: JWTCookie, Value: jwtToken, Expires: expireCookie, HttpOnly: true}
	http.SetCookie(w, &cookie)
}

// JWTExtractor extracts the JWT from a request and returns the token string.
// The token is not verified.
func JWTExtractor(r *http.Request) (string, error) {

	// Get the JWT from the header
	jwtToken := r.Header.Get("Authorization")
	splitToken := strings.Split(jwtToken, " ")
	if len(splitToken) != 2 {
		jwtToken = ""
	} else {
		jwtToken = splitToken[1]
	}

	// Check the cookie if we don't have a bearer token in the header
	if jwtToken == "" {
		cookie, err := r.Cookie(JWTCookie)
		if err != nil {
			log.Println("Cannot find the JWT")
			return "", errors.New("Cannot from the JWT")
		}
		jwtToken = cookie.Value
	}

	return jwtToken, nil
}