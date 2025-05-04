package config

type AllConfig struct {
	Server     Server
	DataSource DataSource
}

type Server struct {
	Port  string
	Level string
}

type DataSource struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	Config   string
}

func (d *DataSource) Dsn() string {
	return d.Username + ":" + d.Password + "@tcp(" + d.Host + ":" + d.Port + ")/" + d.DBName + "?" + d.Config
}
