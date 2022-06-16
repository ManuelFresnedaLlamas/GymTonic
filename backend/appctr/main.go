package appctr

func Start() {
	prepareCfg() // first
	prepareLog() // second
	prepareUpload()
	prepareDB()
}
