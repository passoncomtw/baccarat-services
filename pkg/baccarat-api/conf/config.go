package conf

var AppJsonConfig = []byte(`
{
  "swagger": {
    "basePath": "/",
    "host": "localhost:8332",
    "version": "0.9.0",
    "title": "百家樂 API",
    "description": "百家樂 API"
  },
  "app": {
    "port": "8332",
    "debug_mod": "true"
  },
  "database": {
    "dsn": "user=postgres password=1234 dbname=postgres host=127.0.0.1 port=5432 sslmode=disable"
  }
}
`)
