package pkgAdEnv

const constPathEnv = "../.env"

type AdEnv struct {
	ConstEnv   string `json:"ConstEnv"`
	DBHost     string `json:"DB_HOST"`
	DBUser     string `json:"DB_USER"`
	DBPass     string `json:"DB_PASSWORD"`
	DBName     string `json:"DB_NAME"`
	DBPort     string `json:"DB_PORT"`
	DBProtocal string `json:"DB_PROTOCOL"`
	DBDriver   string `json:"DB_DRIVER"`
	ValKey     string `json:"EN_KEY"`
	BindHost   string `json:"LDAP_SERVER_IP"`
	BindPass   string `json:"LDAP_BIND_PASSWORD"`
	BindPort   string `json:"LDAP_SERVER_PORT"`
	BindDN     string `json:"LDAP_BIND_DN"`
	BindBaseDN string `json:"LDAP_BASE_DN"`
}
