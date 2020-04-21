//go:generate echo $PWD/$GOPACKAGE/$GOFILE
//go:generate mockgen -package guard_test -destination mock_storage_test.go go.octolab.org/ecosystem/guard/internal/service/guard Storage
package guard_test
