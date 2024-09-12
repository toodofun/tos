package config

type Cfg func(c *Config)

func WithPort(port int) Cfg {
	return func(c *Config) {
		c.Server.Port = port
	}
}

func WithDebug(debug bool) Cfg {
	return func(c *Config) {
		c.Server.Debug = debug
	}
}

func WithStorageRoot(root string) Cfg {
	return func(c *Config) {
		c.Storage.Root = root
	}
}
