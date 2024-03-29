package configs

type LogConfig struct {
	Level     string `mapstructure:"level"`
	TimeZone  string `mapstructure:"time_zone"`
	LogName   string `mapstructure:"log_name"`
	LogSuffix string `mapstructure:"log_suffix"`
	MaxBackup int    `mapstructure:"max_backup"`
}
