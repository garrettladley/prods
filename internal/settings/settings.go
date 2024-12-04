package settings

import "github.com/caarlos0/env/v11"

type Settings struct {
	App `envPrefix:"APP_"`
	DB  `envPrefix:"DB_"`
}

func Load() (Settings, error) {
	return env.ParseAs[Settings]()
}
