//go:generate echo $PWD/$GOPACKAGE/$GOFILE
//go:generate mockgen -package api_test -destination mock_service_test.go go.octolab.org/ecosystem/guard/internal/transport/http/api Service
package api_test
