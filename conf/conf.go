package conf

import (
	"os"
	"sync"

	"gopkg.in/yaml.v3"
)

type Conf struct {
	Version  int    `yaml:"version"`
	Debug    bool   `yaml:"debug"`
	URL      string `yaml:"url"`
	Domain   string `yaml:"domain"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

var (
	conf     Conf
	confLock sync.RWMutex
)

func Load(filePath string) error {
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	var newConf Conf
	err = yaml.Unmarshal(bytes, &newConf)
	if err != nil {
		return err
	}

	confLock.Lock()
	conf = newConf
	confLock.Unlock()

	return nil
}

func Get() Conf {
	confLock.RLock()
	defer confLock.RUnlock()

	return conf
}
