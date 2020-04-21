//go:generate echo $PWD/$GOPACKAGE/$GOFILE
//go:generate mockgen -package chi_test -destination mock_api_test.go go.octolab.org/ecosystem/guard/internal/transport/http/api/internal API
package chi_test
