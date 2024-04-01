package config

type ConfigInterface struct {
	BaseDataPath string `json:"baseDataPath"`
}

var Config *ConfigInterface = &ConfigInterface{
	BaseDataPath: "/usr/local/notepad",
}
