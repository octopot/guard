package cmd

import (
	"database/sql"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/kamilsk/go-kit/pkg/fn"
	"github.com/kamilsk/guard/pkg/storage"
	"github.com/pkg/errors"
	"github.com/rakyll/statik/fs"
	"github.com/rubenv/sql-migrate"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	_ "github.com/kamilsk/guard/pkg/storage/migrations"
)

// Migrate applies database migrations.
var Migrate = &cobra.Command{
	Use:   "migrate",
	Short: "Apply database migration",
	Args:  cobra.RangeArgs(0, 2),
	RunE: func(cmd *cobra.Command, args []string) error {
		config := new(migrationConfig)
		config.Read(args)

		migrate.SetTable(cnf.Union.MigrationConfig.Table)
		migrate.SetSchema(cnf.Union.MigrationConfig.Schema)

		layer := storage.Must(storage.Database(cnf.Union.DatabaseConfig))
		src := &migrate.HttpFileSystemMigrationSource{
			FileSystem: Statik{}.Must("/" + layer.Dialect()),
		}

		var runner = run
		if cnf.Union.MigrationConfig.DryRun {
			runner = dryRun
		}
		return runner(container{cmd, src, layer}, config)
	},
}

func init() {
	v := viper.New()
	v.SetEnvPrefix("migration")
	fn.Must(
		func() error { return v.BindEnv("table") },
		func() error { return v.BindEnv("schema") },
		func() error {
			v.SetDefault("table", defaults["table"])
			v.SetDefault("schema", defaults["schema"])
			return nil
		},
		func() error {
			flags := Migrate.Flags()
			flags.StringVarP(&cnf.Union.MigrationConfig.Table,
				"table", "t", v.GetString("table"), "migration table name")
			flags.StringVarP(&cnf.Union.MigrationConfig.Schema,
				"schema", "s", v.GetString("schema"), "migration schema")
			flags.UintVarP(&cnf.Union.MigrationConfig.Limit,
				"limit", "l", 0, "limit the number of migrations (0 = unlimited)")
			flags.BoolVarP(&cnf.Union.MigrationConfig.DryRun,
				"dry-run", "", false, "do not apply migration, just print them")
			return nil
		},
	)
	db(Migrate)
}

func dryRun(c container, cnf *migrationConfig) error {
	plan, _, err := migrate.PlanMigration(c.storage.Database(), c.storage.Dialect(), c.resource, cnf.direction, cnf.limit)
	if err != nil {
		return err
	}
	for _, m := range plan {
		var queries []string
		if cnf.direction == migrate.Up {
			c.printer.Printf("==> Would apply migration %s (up)\n", m.Id)
			queries = m.Up
		} else {
			c.printer.Printf("==> Would apply migration %s (down)\n", m.Id)
			queries = m.Down
		}
		for _, query := range queries {
			c.printer.Println(query)
		}
	}
	return nil
}

func run(c container, cnf *migrationConfig) error {
	count, err := migrate.ExecMax(c.storage.Database(), c.storage.Dialect(), c.resource, cnf.direction, cnf.limit)
	if err != nil {
		return err
	}
	c.printer.Printf("Applied %d migration(s)!\n", count)
	return nil
}

type container struct {
	printer interface {
		Printf(format string, v ...interface{})
		Println(v ...interface{})
	}
	resource migrate.MigrationSource
	storage  interface {
		Database() *sql.DB
		Dialect() string
	}
}

// TODO issue#refactoring join with pkg/config.MigrationConfig
type migrationConfig struct {
	direction migrate.MigrationDirection
	limit     int
}

// Read tries to read migration configuration from the passed arguments:
// - 0 index must contain the direction (up by default)
// - 1 index must contain the number of migrations (0 by default)
func (c *migrationConfig) Read(args []string) error {
	c.direction, c.limit = migrate.Up, 0
	if len(args) > 0 {
		switch {
		case strings.EqualFold(args[0], "up"):
			c.direction = migrate.Up
		case strings.EqualFold(args[0], "down"):
			c.direction = migrate.Down
		default:
			return errors.Errorf("invalid direction %q", args[0])
		}
		if len(args) == 2 {
			limit, err := strconv.Atoi(args[1])
			if err != nil {
				return errors.Wrap(err, "limit arg must be a valid integer")
			}
			c.limit = limit
		}
	}
	return nil
}

// Statik provides possibility to choose root directory of github.com/rakyll/statik/fs.
type Statik struct {
	Dir    string
	origin http.FileSystem
}

// Must build new Statik instance.
func (statik Statik) Must(root string) *Statik {
	if root != "" && !strings.HasPrefix(root, "/") {
		panic(errors.Errorf("root directory must be prefixed by slash: %s", root))
	}
	origin, err := fs.New()
	if err != nil {
		panic(errors.Wrap(err, "trying to instantiate a Statik"))
	}
	statik.Dir, statik.origin = root, origin
	return &statik
}

// Open implements http.FileSystem interface.
func (statik *Statik) Open(name string) (http.File, error) {
	return statik.origin.Open(filepath.Join(statik.Dir, name))
}
