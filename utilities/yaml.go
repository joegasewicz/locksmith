package utilities

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type users struct {
	Name     string
	Role     string
	Email    string
	Password string
}

type yamlScheme struct {
	Roles []string
	Users []users
}

type Yaml struct {
	file []byte
	Yaml yamlScheme
}

func NewYaml() *Yaml {
	return &Yaml{}
}

func (y *Yaml) Get(fileName string) {
	f, err := os.ReadFile(fileName)
	if err != nil {
		log.Println("locksmith.Yaml file not found so using default roles")
		return
	}
	y.file = f
}

func (y *Yaml) Do() error {
	var ys yamlScheme
	if err := yaml.Unmarshal(y.file, &ys); err != nil {
		return err
	}
	y.Yaml = ys
	return nil
}
