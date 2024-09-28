package conf

var AppJsonConfig = []byte(`
{
  "app":{
      "port":"8366",
      "upload_file_path":"G:\\golang_project\\JayHonChat\\tempt_img",
      "serve_type":"GOServe",
      "debug_mod":"true"
  },
  "mysql":{
     "dsn": "root:123456@tcp(127.0.0.1:3306)/jayhon_chat?charset=utf8mb4&parseTime=True&loc=Local"
     }
    }
   `)
