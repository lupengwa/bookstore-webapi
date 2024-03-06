package db

type PostgreSQLDbConfig struct {
	Url      string `validate:"required"`
	Port     string `validate:"required"`
	DbName   string `validate:"required"`
	User     string `validate:"required"`
	Password string
	SslMode  string
}
