package env

import "os"

var (
  DbPassword = os.Getenv("DB_PASSWORD")
  DbUser     = os.Getenv("DB_USER")
  DbName     = os.Getenv("DB_NAME")
  DbHost     = os.Getenv("DB_HOST")
  DbPort     = os.Getenv("DB_PORT")
)