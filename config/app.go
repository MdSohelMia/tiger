package config

type AppConfig struct {
	//Application name
	//This Value name of your applications.
	Name string `mapstructure:"name"`
	// Application Environment
	Env string `mapstructure:"env"`
	//Application Debug Mode
	Debug *bool `mapstructure:"debug"`
	// Application Host
	Host string `mapstructure:"host"`
	//Application Port
	Port string `mapstructure:"port"`
	// Application Timezone
	Timezone string `mapstructure:"timezone"`
	//maintenance driver support only [redis & file]
	Maintenance string `mapstructure:"maintenance"`
}
