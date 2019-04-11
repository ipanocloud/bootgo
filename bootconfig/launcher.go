package bootconfig

//init bootgo env
func InitEnv() {

	//env init
	EnvInit()

	//db init
	DbInit()

	//log init
	LogInit()
}
