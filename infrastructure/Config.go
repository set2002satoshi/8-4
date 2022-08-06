package infrastructure

type Config struct {
	DB struct {
		Production struct {
			Host     string
			Username string
			Password string
			DBName   string
		}
		Test struct {
			Host     string
			Username string
			Password string
			DBName   string
		}
	}
	Routing struct {
		Port string
	}
}

func NewConfig() *Config {
	c := new(Config)

	c.DB.Production.Host = "localhost"
	c.DB.Production.Username = "docker"
	c.DB.Production.Password = "pass"
	c.DB.Production.DBName = "main"

	c.DB.Test.Host = "localhost"
	c.DB.Test.Username = "docker"
	c.DB.Test.Password = "pass"
	c.DB.Test.DBName = "main"

	c.Routing.Port = ":8080"

	return c
}
