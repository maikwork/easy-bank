package model

type DBSetting struct {
	Type     string `json:"type"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	DBName   string `json:"dbname"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Config struct {
	DBSetting DBSetting `json:"dbsetting"`
}
