package comm

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Common struct {
	HTTPAddr  string `yaml:"http_address"`
	HTTPPort  int    `yaml:"http_port"`
	KeyFactor string `yaml:"key_factor"`
	Debug     bool   `yaml:"debug"`
}

type Database struct {
	Driver   string `yaml:"driver"`
	URL      string `yaml:"url"`
	UserName string `yaml:"user_name"`
	Password string `yaml:"password"`
}

// ParseCfg config from file
func ParseCfg(cfg interface{}) error {
	t, err := ioutil.ReadFile(gAppcfg)
	//fmt.Println(string(t))
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(t, cfg)
	if err != nil {
		return err
	}
	return err
}
