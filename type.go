package main

type EnumBox struct {
	Version   string     `yaml:"version"`
	Imports   []string   `yaml:"imports"`
	Variables []Variable `yaml:"variables"`
}

type Variable struct {
	Name  string `yaml:"name"`
	Type  string `yaml:"type"`
	Value string `yaml:"value"`
}
