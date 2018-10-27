//go:generate echo $PWD/$GOPACKAGE/$GOFILE
//go:generate mockgen -package guard_test -destination $PWD/pkg/service/guard/mock_storage_test.go github.com/kamilsk/guard/pkg/service/guard Storage
package guard_test
