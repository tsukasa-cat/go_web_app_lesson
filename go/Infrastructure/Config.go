package Infrastructure

type Config struct {
    DataBase struct {
        Development struct {
            Host string
            Username string
            Password string
            DBName string
        }
    }
}


func NewConfig() *Config {
    c := new(Config)
    c.DataBase.Development.Host = "go_web_app_db_1"
    c.DataBase.Development.Username = "admin"
    c.DataBase.Development.Password = "password"
    c.DataBase.Development.DBName = "go_db"
    return c
}