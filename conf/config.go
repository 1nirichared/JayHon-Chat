package conf

var AppJsonConfig = []byte(`
{
  "app":{
      "port":"8366",
      "upload_file_path":"G:\\golang_project\\JayHonChat\\tempt_img",
      "cookie_key": "d84f8151e561d4aaddd0ef86de9ee6971d9179cdc9de425e8dafbf6a32bd2220",
      "serve_type":"GoServe",
      "debug_mod":"true"
  },
  "mysql":{
     "dsn": "root:123456@tcp(127.0.0.1:3306)/jayhon_chat?charset=utf8mb4&parseTime=True&loc=Local"
     }
    }
   `)
