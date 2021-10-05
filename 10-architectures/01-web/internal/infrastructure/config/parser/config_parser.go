package parser

import (
	"reflect"

	"github.com/spf13/viper"

	"web-sample/internal/infrastructure/config/errors"
	"web-sample/internal/infrastructure/config/models"
	"web-sample/internal/infrastructure/config/parser/interfaces"
)

type configParser struct {
	filePath string
	fileName string
	fileType string
}

func NewConfigParser(filePath, fileName, fileType string) interfaces.ConfigParser {
	return &configParser{
		filePath: filePath,
		fileName: fileName,
		fileType: fileType,
	}
}

func (cp *configParser) Parse() (*models.Config, error) {
	v := viper.New()

	v.AddConfigPath(cp.filePath)
	v.SetConfigName(cp.fileName)
	v.SetConfigType(cp.fileType)

	bindEnvironmentVariables(v, reflect.TypeOf(&models.Config{}))

	if err := v.ReadInConfig(); err != nil {
		return nil, errors.NewConfigParseError(err)
	}

	config := models.NewConfig()

	if err := v.Unmarshal(&config); err != nil {
		return nil, errors.NewConfigParseError(err)
	}

	if err := config.Validate(); err != nil {
		return nil, errors.NewConfigParseError(err)
	}

	return config, nil
}

func bindEnvironmentVariables(v *viper.Viper, typ reflect.Type) error {
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		if field.Type.Kind() == reflect.Ptr || field.Type.Kind() == reflect.Struct {
			if err := bindEnvironmentVariables(v, field.Type); err != nil {
				return errors.NewConfigParseError(err)
			}
			continue
		}

		if value, ok := field.Tag.Lookup("env"); ok {
			if err := v.BindEnv(value); err != nil {
				return errors.NewConfigParseError(err)
			}
		}
	}

	return nil
}
