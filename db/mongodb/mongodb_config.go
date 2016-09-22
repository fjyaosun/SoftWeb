package mongodb

var mongoConf struct {
	URL          string
	MGO_CONN_CAP int
	DB           string
}

func configInit() {
	mongoConf.URL = "127.0.0.1:27017"
	mongoConf.MGO_CONN_CAP = 256
	mongoConf.DB = "cloudConfig"
}
