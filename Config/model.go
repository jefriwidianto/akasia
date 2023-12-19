package Config

const (
	ENVIRONMENT_PATH = "/Environment/"
	Localhost        = "Local"
	PathMigration    = "../Repository/Migration/"
)

type DbSqlConfigName string

const (
	// Database Connection Constant name
	DATABASE_MAIN DbSqlConfigName = "DBmain"
)

// ftroct for collect data object Config environment ".yml"
type Environment struct {
	Databases database `yaml:"databases"`
}

type database struct {
	Username           string `yaml:"username"`
	Password           string `yaml:"password"`
	Port               string `yaml:"port"`
	Engine             string `yaml:"engine"`
	Host               string `yaml:"host"`
	Maximum_connection int    `yaml:"maximum_connection"`
}
