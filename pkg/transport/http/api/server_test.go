//go:generate echo $PWD/$GOPACKAGE/$GOFILE
//go:generate mockgen -package api_test -destination $PWD/pkg/transport/http/api/mock_service_test.go github.com/kamilsk/guard/pkg/transport/http/api Service
package api_test
