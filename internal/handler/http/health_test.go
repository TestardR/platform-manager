package http

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type handlerCaseHealth struct {
	h      handler
	status int
}

func TestHandler_Health(t *testing.T) {
	t.Parallel()
	mc := gomock.NewController(t)
	t.Cleanup(func() { mc.Finish() })

	tests := map[string]handlerCaseHealth{
		"health-ok": handlerHealthCaseOk(mc),
	}

	for tn, tc := range tests {
		tn, tc := tn, tc
		t.Run(tn, func(t *testing.T) {
			t.Parallel()

			ts, uri := handlerHealthServer(tc.h)
			t.Cleanup(func() { ts.Close() })

			req, err := http.NewRequest(http.MethodGet, uri, nil)
			require.NoError(t, err)

			resp, err := ts.Client().Do(req)
			require.NoError(t, err)

			resp.Body.Close()
			assert.Equal(t, tc.status, resp.StatusCode)

		})
	}
}

func handlerHealthCaseOk(mc *gomock.Controller) handlerCaseHealth {
	return handlerCaseHealth{
		h:      handler{},
		status: http.StatusOK,
	}
}

func handlerHealthServer(h handler) (*httptest.Server, string) {
	router := gin.New()

	router.GET("", h.Health)
	ts := httptest.NewServer(router)

	return ts, ts.URL
}
