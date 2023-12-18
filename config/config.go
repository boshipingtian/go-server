package config

type Config struct {
	Datasource Datasource `yaml:"datasource"`
	Logger     Logger     `yaml:"logger"`
	System     System     `yaml:"system"`
}
