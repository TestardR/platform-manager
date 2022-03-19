package http

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/TestardR/platformmanager/pkg/mock"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type handlerCaseService struct {
	h      handler
	status int
}

func TestHandler_GetService(t *testing.T) {
	t.Parallel()
	mc := gomock.NewController(t)
	t.Cleanup(func() { mc.Finish() })

	tests := map[string]handlerCaseService{
		"fail-get-deployements": handleGetServiceCaseFailGetDeployments(mc),
		"success":               handleGetServiceCaseOk(mc),
	}

	for tn, tc := range tests {
		tn, tc := tn, tc
		t.Run(tn, func(t *testing.T) {
			t.Parallel()

			router := gin.New()
			router.GET("", tc.h.GetServices)
			ts := httptest.NewServer(router)
			t.Cleanup(func() { ts.Close() })

			req, err := http.NewRequest(http.MethodGet, ts.URL, nil)
			if err != nil {
				t.Error(err)
			}
			resp, err := ts.Client().Do(req)
			require.NoError(t, err)

			resp.Body.Close()
			assert.Equal(t, tc.status, resp.StatusCode)
		})
	}
}

func handleGetServiceCaseOk(mc *gomock.Controller) handlerCaseService {
	mpm := mock.NewMockManagerer(mc)

	mpm.EXPECT().GetDeployments(gomock.Any(), gomock.Any()).Return(buildDeployments(), nil)

	return handlerCaseService{
		h: handler{
			pm: mpm,
		},
		status: http.StatusOK,
	}
}

func handleGetServiceCaseFailGetDeployments(mc *gomock.Controller) handlerCaseService {
	ml := mock.NewMockLogger(mc)
	mpm := mock.NewMockManagerer(mc)

	mpm.EXPECT().GetDeployments(gomock.Any(), gomock.Any()).Return(nil, errors.New("mock"))
	ml.EXPECT().Error(gomock.Any())

	return handlerCaseService{
		h: handler{
			log: ml,
			pm:  mpm,
		},
		status: http.StatusInternalServerError,
	}
}

type handlerCaseGetServicePerApplicationGroup struct {
	h                handler
	applicationGroup string
	status           int
}

func TestHandler_ZombieStatus(t *testing.T) {
	t.Parallel()
	mc := gomock.NewController(t)
	t.Cleanup(func() { mc.Finish() })

	tests := map[string]handlerCaseGetServicePerApplicationGroup{
		"fail-get-deployements": handleGetServicePerApplicationGroupCaseFailGetDpls(mc),
		"success":               handleGetServicePerApplicationGroupCaseSuccess(mc),
	}

	for tn, tc := range tests {
		tn, tc := tn, tc
		t.Run(tn, func(t *testing.T) {
			t.Parallel()

			w := httptest.NewRecorder()
			ctx, router := gin.CreateTestContext(w)

			router.GET("/:applicationGroup", tc.h.GetServicesPerApplicationGroup)
			uri := fmt.Sprintf("/%s", tc.applicationGroup)

			var err error
			ctx.Request, err = http.NewRequest("GET", uri, nil)
			require.NoError(t, err)

			router.ServeHTTP(w, ctx.Request)

			if w.Code != tc.status {
				t.Errorf("Expected status %d, got %d", tc.status, w.Code)
			}
		})
	}
}

func handleGetServicePerApplicationGroupCaseFailGetDpls(mc *gomock.Controller) handlerCaseGetServicePerApplicationGroup {
	ml := mock.NewMockLogger(mc)
	mpm := mock.NewMockManagerer(mc)

	ml.EXPECT().Info(gomock.Any())
	mpm.EXPECT().GetDeploymentsPerLabel(gomock.Any(), gomock.Any(), "applicationGroup", "mock").Return(nil, errors.New("mock"))
	ml.EXPECT().Error(gomock.Any())

	return handlerCaseGetServicePerApplicationGroup{
		h: handler{
			log: ml,
			pm:  mpm,
		},
		applicationGroup: "mock",
		status:           http.StatusInternalServerError,
	}
}

func handleGetServicePerApplicationGroupCaseSuccess(mc *gomock.Controller) handlerCaseGetServicePerApplicationGroup {
	ml := mock.NewMockLogger(mc)
	mpm := mock.NewMockManagerer(mc)

	ml.EXPECT().Info(gomock.Any())
	mpm.EXPECT().GetDeploymentsPerLabel(gomock.Any(), gomock.Any(), "applicationGroup", "mock").Return(buildDeployments(), nil)

	return handlerCaseGetServicePerApplicationGroup{
		h: handler{
			log: ml,
			pm:  mpm,
		},
		applicationGroup: "mock",
		status:           http.StatusOK,
	}
}

func buildDeployments() []v1.Deployment {
	dpl := v1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: "demo-deployment",
		},
		Spec: v1.DeploymentSpec{
			Replicas: int32Ptr(2),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"applicationGroup": "test",
				},
			},
		},
	}

	return []v1.Deployment{dpl}
}

func int32Ptr(i int32) *int32 { return &i }
