package config

type DBMySql struct {
	host     string
	user     string
	password string
	port     string
	name     string
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
