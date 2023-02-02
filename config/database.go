package config

type DatabaseConfig struct {
	Default    string      `mapstructure:"default"`
	Migrations string      `mapstructure:"migrations"`
	Tidb       TidbConfig  `mapstructure:"tidb"`
	Mysql      MysqlConfig `mapstructure:"mysql"`
	Pgsql      MysqlConfig `mapstructure:"pgsql"`
}

type TidbConfig struct {
	Host      string `mapstructure:"host"`
	Username  string `mapstructure:"username"`
	Password  string `mapstructure:"password"`
	Database  string `mapstructure:"database"`
	Charset   string `mapstructure:"charset"`
	Collation string `mapstructure:"collation"`
	Engine    string `mapstructure:"engine"`
}
type MysqlConfig struct {
	Host      string `mapstructure:"host"`
	Username  string `mapstructure:"username"`
	Password  string `mapstructure:"password"`
	Database  string `mapstructure:"database"`
	Charset   string `mapstructure:"charset"`
	Collation string `mapstructure:"collation"`
	Engine    string `mapstructure:"engine"`
}
type PgsqlConfig struct {
	Host      string `mapstructure:"host"`
	Username  string `mapstructure:"username"`
	Password  string `mapstructure:"password"`
	Database  string `mapstructure:"database"`
	Charset   string `mapstructure:"charset"`
	Collation string `mapstructure:"collation"`
	Engine    string `mapstructure:"engine"`
}

// Redis Config
type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
}
