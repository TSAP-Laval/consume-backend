package metricsmodule_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tsap-laval/consume-backend/app/modules/metricsmodule"
)

func TestControllerSpec(t *testing.T) {

	t.Run("MetricsController returns 400 when TeamID is bad", func(t *testing.T) {
		c := metricsmodule.MetricsController{}

		mux := http.NewServeMux()
		server := httptest.NewServer(mux)
		defer server.Close()

		mux.HandleFunc("/", c.CreateMetric)

		var jsonStr = []byte(`{}`)
		req, _ := http.NewRequest("POST", server.URL+"/asdf", bytes.NewBuffer(jsonStr))

		client := &http.Client{}
		resp, err := client.Do(req)

		if err != nil {
			t.Fatal(err)
		}

		if resp.StatusCode != http.StatusBadRequest {
			t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, resp.StatusCode)
		}
	})
}
