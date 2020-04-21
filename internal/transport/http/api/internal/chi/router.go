package chi

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"go.octolab.org/ecosystem/guard/internal/transport/http/api/internal"
)

// Configure TODO issue#docs
func Configure(api internal.API) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Post("/api/v1/license/check", api.CheckLicenseV1)
	return r
}
