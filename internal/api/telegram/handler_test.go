package telegram

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/mock/gomock"
)

func TestTelegramHandler(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Api Telegram Handler Suite")
}

var _ = Describe("telegram handler test", func() {
	var (
		ctrl        *gomock.Controller
		target      *Handler
		mockService *MockService
		router      *echo.Echo
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockService = NewMockService(ctrl)
		target = NewHandler(mockService)
		router = echo.New()

		router.POST("/apis/v1/telegrams/alerts", target.AlertJob)
	},
	)

	AfterEach(func() {
		ctrl.Finish()
	},
	)

	Describe("AlertJob test", func() {
		It("successfully", func() {
			w := httptest.NewRecorder()

			req, _ := http.NewRequest(http.MethodPost, "/apis/v1/telegrams/alerts", nil)
			req.Header.Set("Content-Type", "application/json")

			mockService.EXPECT().alertJob(gomock.Any()).Return(nil)

			router.ServeHTTP(w, req)

			Expect(w.Code).To(Equal(http.StatusOK))
		})
	},
	)
})
