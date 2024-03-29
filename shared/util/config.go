package util

import "github.com/spf13/viper"

type Config struct {
	PortApp          string `mapstructure:"PORT_APP"`
	DBDsn            string `mapstructure:"DB_DSN"`
	AppName          string `mapstructure:"APP_NAME"`
	GOEnv            string `mapstructure:"GO_ENV"`
	TokenSymetricKey string `mapstructure:"TOKEN_SYMETRIC_KEY"`
	AWSRegion        string `mapstructure:"AWS_REGION"`
	AWSS3Bucket      string `mapstructure:"AWS_S3_BUCKET"`
	AWSAccessKey     string `mapstructure:"AWS_ACCESS_KEY"`
	AWSSecretKey     string `mapstructure:"AWS_SECRET_KEY"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
