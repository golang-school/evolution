package logger

type Config struct {
	AppName       string `envconfig:"APP_NAME"    required:"true"`
	AppVersion    string `envconfig:"APP_VERSION" required:"true"`
	Level         string `default:"error"         envconfig:"LOGGER_LEVEL"`
	PrettyConsole bool   `default:"false"         envconfig:"LOGGER_PRETTY_CONSOLE"`
}

func Init(c Config) {
	// Настраиваем level и форматирование
}
