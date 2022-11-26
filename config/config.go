package config

import "github.com/spf13/viper"

type Auth struct {
	SecretKey      string `json:"secretKey"`
	ExpiredTime    int64  `json:"expiredTime"`
	RbacModelPath  string `json:"rbac-model-path"`
	RbacPolicyPath string `json:"rbac-policy-path"`
	CookiePath     string `json:"cookie-path"`
	CookieSecure   bool   `json:"cookie-secure"`
}

func AuthConfig() *Auth {
	return &Auth{
		SecretKey:      viper.GetString("app.auth.secretKey"),
		ExpiredTime:    viper.GetInt64("app.auth.expiredTime"),
		RbacModelPath:  viper.GetString("app.auth.rbac-model-path"),
		RbacPolicyPath: viper.GetString("app.auth.rbac-policy-path"),
		CookiePath:     viper.GetString("app.auth.cookie-path"),
		CookieSecure:   viper.GetBool("app.auth.cookie-secure"),
	}
}

type BaseConfig struct {
	BasePath  string
	BasicAuth string
}
