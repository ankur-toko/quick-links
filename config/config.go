package config

type ConfigProvider interface {
	Get(string) (string, bool)
	Set(string, string)
	Reload() error // reload from the file again
}

var configObject = GetConfigProvider()

func Get(k string) (string, bool) {
	return configObject.Get(k)
}

func Set(k string, v string) {
	configObject.Set(k, v)
}

func Reload() {
	configObject.Reload()
}
