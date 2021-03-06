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

package keypair_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/CanonicalLtd/serial-vault/datastore"
	"github.com/CanonicalLtd/serial-vault/service"
	"github.com/CanonicalLtd/serial-vault/service/keypair"
	check "gopkg.in/check.v1"
)

func (s *KeypairSuite) TestAPIKeypairsHandler(c *check.C) {

	tests := []KeypairTest{
		{"GET", "/api/keypairs", nil, 400, "application/json; charset=UTF-8", 0, false, false, 0},
		{"GET", "/api/keypairs", nil, 200, "application/json; charset=UTF-8", datastore.Admin, false, true, 2},
		{"GET", "/api/keypairs", nil, 400, "application/json; charset=UTF-8", datastore.Standard, false, false, 0},
		{"GET", "/api/keypairs", nil, 400, "application/json; charset=UTF-8", 0, true, false, 0},
	}

	for _, t := range tests {
		if t.EnableAuth {
			datastore.Environ.Config.EnableUserAuth = true
		}

		w := sendAdminAPIRequest(t.Method, t.URL, bytes.NewReader(t.Data), t.Permissions, c)
		c.Assert(w.Code, check.Equals, t.Code)
		c.Assert(w.Header().Get("Content-Type"), check.Equals, t.Type)

		result, err := parseListResponse(w)
		c.Assert(err, check.IsNil)
		c.Assert(result.Success, check.Equals, t.Success)
		c.Assert(len(result.Keypairs), check.Equals, t.List)

		datastore.Environ.Config.EnableUserAuth = false
	}
}

func (s *KeypairSuite) TestAPISyncKeypairsHandler(c *check.C) {
	datastore.ReEncryptKeypair = mockReEncryptKeypair

	k := keypair.SyncRequest{Secret: "NewKeystoreSecretInTheFactory"}
	data, _ := json.Marshal(k)

	tests := []KeypairTest{
		{"POST", "/api/keypairs/sync", data, 400, "application/json; charset=UTF-8", 0, false, false, 0},
		{"POST", "/api/keypairs/sync", data, 200, "application/json; charset=UTF-8", datastore.SyncUser, true, true, 2},
		{"POST", "/api/keypairs/sync", data, 200, "application/json; charset=UTF-8", datastore.Admin, true, true, 2},
		{"POST", "/api/keypairs/sync", data, 400, "application/json; charset=UTF-8", datastore.Standard, true, false, 0},
	}

	for _, t := range tests {
		if t.EnableAuth {
			datastore.Environ.Config.EnableUserAuth = true
		}

		w := sendAdminAPIRequest(t.Method, t.URL, bytes.NewReader(t.Data), t.Permissions, c)
		c.Assert(w.Code, check.Equals, t.Code)
		c.Assert(w.Header().Get("Content-Type"), check.Equals, t.Type)

		result, err := parseSyncResponse(w)
		c.Assert(err, check.IsNil)
		c.Assert(result.Success, check.Equals, t.Success)
		c.Assert(len(result.Keypairs), check.Equals, t.List)

		datastore.Environ.Config.EnableUserAuth = false
	}
}

func sendAdminAPIRequest(method, url string, data io.Reader, permissions int, c *check.C) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest(method, url, data)

	switch permissions {
	case datastore.Admin:
		r.Header.Set("user", "sv")
		r.Header.Set("api-key", "ValidAPIKey")
	case datastore.SyncUser:
		r.Header.Set("user", "sync")
		r.Header.Set("api-key", "ValidAPIKey")
	case datastore.Standard:
		r.Header.Set("user", "user1")
		r.Header.Set("api-key", "ValidAPIKey")
	default:
		break
	}

	service.AdminRouter().ServeHTTP(w, r)

	return w
}

func parseSyncResponse(w *httptest.ResponseRecorder) (keypair.SyncResponse, error) {
	// Check the JSON response
	result := keypair.SyncResponse{}
	err := json.NewDecoder(w.Body).Decode(&result)
	return result, err
}

func mockReEncryptKeypair(keypair datastore.Keypair, newSecret string) (string, string, error) {
	return "Base64SealedKey", "Base64SAuthKey", nil
}
