package config

import (
	"flag"
	"os"
)

type Config struct {
	personalAccessToken string
	domainName          string
}

func Get() *Config {
	conf := &Config{}

	flag.StringVar(&conf.personalAccessToken, "AccessToken", os.Getenv("ACCESS_TOKEN"), "third-party service Access token")
	flag.StringVar(&conf.domainName, "thirdpartyDomainName", os.Getenv("DOMAIN_NAME"), "third-party service Domain Name")

	flag.Parse()

	return conf
}

func (c *Config) GetPersonalAccessToken() string {
	return c.personalAccessToken
}

func (c *Config) GetDomainName() string {
	return c.domainName
}
