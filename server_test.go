package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPI(t *testing.T) {
	server := httptest.NewServer(New())
	defer server.Close()

	t.Run("GET /calculations", func(t *testing.T) {
		resp, err := http.Get(server.URL + "/calculations")
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("POST /calculations", func(t *testing.T) {
		type Expected struct {
			code int
			body string
		}

		cases := map[string]struct {
			input    string
			expected Expected
		}{
			"1+2=3": {
				input:    `{"a": 1, "b": 2, "op": "+"}`,
				expected: Expected{code: http.StatusCreated, body: `{"result": 3}`},
			},
			"1-2=-1": {
				input:    `{"a": 1, "b": 2, "op": "-"}`,
				expected: Expected{code: http.StatusCreated, body: `{"result": -1}`},
			},
			"1*2=2": {
				input:    `{"a": 1, "b": 2, "op": "*"}`,
				expected: Expected{code: http.StatusCreated, body: `{"result": 2}`},
			},
			"2/2=1": {
				input:    `{"a": 2, "b": 2, "op": "/"}`,
				expected: Expected{code: http.StatusCreated, body: `{"result": 1}`},
			},
		}

		for name, tt := range cases {
			tt := tt

			t.Run(name, func(t *testing.T) {
				resp, err := http.Post(server.URL+"/calculations", "application/json", strings.NewReader(tt.input))
				assert.NoError(t, err)
				defer resp.Body.Close()

				assert.Equal(t, tt.expected.code, resp.StatusCode)

				result, err := ioutil.ReadAll(resp.Body)
				assert.NoError(t, err)

				assert.JSONEq(t, tt.expected.body, string(result))
			})
		}
	})
}
