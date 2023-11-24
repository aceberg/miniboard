package conf

import (
	"github.com/spf13/viper"

	"github.com/aceberg/miniboard/internal/auth"
	"github.com/aceberg/miniboard/internal/check"
	"github.com/aceberg/miniboard/internal/models"
)

// Get - read config from file or env
func Get(path string) (config models.Conf, authConf auth.Conf) {

	viper.SetDefault("HOST", "0.0.0.0")
	viper.SetDefault("PORT", "8849")
	viper.SetDefault("THEME", "flatly")
	viper.SetDefault("COLOR", "dark")
	viper.SetDefault("COLORON", "#89ff89")
	viper.SetDefault("COLOROFF", "#ff7171")
	viper.SetDefault("BTNWIDTH", "180px")
	viper.SetDefault("WEBREFRESH", "60")
	viper.SetDefault("DBTRIMDAYS", "30")

	viper.SetDefault("AUTH_USER", "")
	viper.SetDefault("AUTH_PASSWORD", "")
	viper.SetDefault("AUTH_EXPIRE", "7d")

	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	check.IfError(err)

	viper.AutomaticEnv() // Get ENVIRONMENT variables

	config.Host, _ = viper.Get("HOST").(string)
	config.Port, _ = viper.Get("PORT").(string)
	config.Theme, _ = viper.Get("THEME").(string)
	config.Color, _ = viper.Get("COLOR").(string)
	config.ColorOn, _ = viper.Get("COLORON").(string)
	config.ColorOff, _ = viper.Get("COLOROFF").(string)
	config.BtnWidth, _ = viper.Get("BTNWIDTH").(string)
	config.WebRefresh, _ = viper.Get("WEBREFRESH").(string)
	config.DBTrimDays, _ = viper.Get("DBTRIMDAYS").(string)

	authConf.Auth = viper.GetBool("AUTH")
	authConf.User, _ = viper.Get("AUTH_USER").(string)
	authConf.Password, _ = viper.Get("AUTH_PASSWORD").(string)
	authConf.ExpStr, _ = viper.Get("AUTH_EXPIRE").(string)

	authConf.Expire = auth.ToTime(authConf.ExpStr)
	config.Auth = authConf.Auth

	return config, authConf
}

// Write - write config to file
func Write(config models.Conf, authConf auth.Conf) {

	viper.SetConfigFile(config.ConfPath)
	viper.SetConfigType("yaml")

	viper.Set("host", config.Host)
	viper.Set("port", config.Port)
	viper.Set("theme", config.Theme)
	viper.Set("color", config.Color)
	viper.Set("coloron", config.ColorOn)
	viper.Set("coloroff", config.ColorOff)
	viper.Set("btnwidth", config.BtnWidth)
	viper.Set("webrefresh", config.WebRefresh)
	viper.Set("dbtrimdays", config.DBTrimDays)

	viper.Set("auth", authConf.Auth)
	viper.Set("auth_user", authConf.User)
	viper.Set("auth_password", authConf.Password)
	viper.Set("auth_expire", authConf.ExpStr)

	err := viper.WriteConfig()
	check.IfError(err)
}
