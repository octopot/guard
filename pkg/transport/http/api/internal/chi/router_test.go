//go:generate echo $PWD/$GOPACKAGE/$GOFILE
//go:generate mockgen -package chi_test -destination $PWD/pkg/transport/http/api/internal/chi/mock_api_test.go github.com/kamilsk/guard/pkg/transport/http/api/internal API
package chi_test
