package libs

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// Hash type
type Hash map[interface{}]interface{}

// JSON type
type JSON map[string]interface{}

var (
	// Config map
	Config Hash
)

// Configure is configuration method
func Configure() {
	data, err := ioutil.ReadFile("config/default.yml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(data, &Config)
	if err != nil {
		panic(err)
	}

	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
	}
	log.Println(env)

	data, err = ioutil.ReadFile("config/" + env + ".yml")
	if err != nil {
		panic(err)
	}

	var override Hash
	err = yaml.Unmarshal(data, &override)
	if err != nil {
		panic(err)
	}

	recursiveOverride(&Config, override)
}

func recursiveOverride(conf *Hash, over Hash) {
	for k, v := range *conf {
		// 上書き側設定の項目の存在確認
		o, ok := over[k]
		if ok == false {
			continue
		}

		// Hash であるか確認
		if vHash, ok := v.(Hash); ok {
			recursiveOverride(&vHash, o.(Hash))
		} else {
			(*conf)[k] = o
		}
	}
}
