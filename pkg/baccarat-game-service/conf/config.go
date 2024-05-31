package conf

/*参数说明
app.port // 应用端口
app.upload_file_path // 图片上传的临时文件夹目录，绝对路径！
app.cookie_key // 生成加密session
app.serve_type // 默认请使用GoServe
mysql.dsn // mysql 连接地址dsn
app.debug_mod // 开发模式建议设置为`true` 避免修改静态资源需要重启服务
*/

var AppJsonConfig = []byte(`
{
  "app": {
    "port": "8311",
    "debug_mod": "true"
  },
  "database": {
    "dsn": "user=postgres password=1234 dbname=baccarat host=127.0.0.1 port=5432 sslmode=disable"
  },
  "sccserver": {
    "url": "ws://127.0.0.1:8223/socketcluster/"
  }
}
`)
