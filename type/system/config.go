package system

// Config stores general setting parameters that are loaded from
// enviroment variables, a dotenv file, and yaml files
type Config struct {
	Env               string `envconfig:"env" valid:"required"`
	Host              string `envconfig:"host" valid:"required"`
	BasicAuthUsername string `envconfig:"basic_auth_username" valid:"required"`
	BasicAuthPassword string `envconfig:"basic_auth_password" valid:"required"`
	HoneybadgerAPIKey string `envconfig:"honeybadger_api_key"`
	SampleDatabaseURL string `envconfig:"sample_database_url" valid:"required"`
}

// IsDevelopment returns whether the application is running as a development mode
func (c *Config) IsDevelopment() bool {
	return c.Env == "development"
}

// IsProduction returns whether the application is running as a production mode
func (c *Config) IsProduction() bool {
	return c.Env == "production"
}
