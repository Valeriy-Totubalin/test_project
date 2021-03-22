package config

type DBMySql struct {
	host     string
	user     string
	password string
	port     string
	name     string
}

func NewDBMysql() *DBMySql {
	return &DBMySql{
		host:     getEnv("DB_HOST", "127.0.0.1"),
		port:     getEnv("DB_PORT", "3306"),
		user:     getEnv("DB_USER_NAME", "root"),
		password: getEnv("DB_PASSWORD", ""),
		name:     getEnv("DB_NAME", "db"),
	}
}

func (db *DBMySql) GetHost() string {
	return db.host
}

func (db *DBMySql) GetUser() string {
	return db.user
}

func (db *DBMySql) GetPassword() string {
	return db.password
}

func (db *DBMySql) GetPort() string {
	return db.port
}

func (db *DBMySql) GetName() string {
	return db.name
}
