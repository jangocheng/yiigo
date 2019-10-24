package yiigo

func init() {
	// init default logger
	logger = newLogger(&logConf{
		Path:       "app.log",
		MaxSize:    500,
		MaxBackups: 0,
		MaxAge:     0,
		Compress:   true,
	}, false)

	// load config file: env.toml
	loadConfigFile()

	debug := Env("app.debug").Bool(true)

	// init logger
	initLogger(debug)
	// init mailer
	initMailer()
	// init db
	initDB(debug)
	// init mongodb
	initMongoDB()
	// init redis
	initRedis()
}