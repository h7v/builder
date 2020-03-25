package builder

type server struct {
	android string
	ios     string
}

type Server interface {
	ShowAndroidIP() string
	ShowIOSIP() string
}

type ServerBuilder interface {
	PickClientsAndroid(string) ServerBuilder
	PickClientsIOS(string) ServerBuilder
	Build() Server
}

type serverBuilder struct {
	android string
	ios     string
}

func (sb *serverBuilder) Build() *server {
	return &server{
		android: sb.android,
		ios:     sb.ios,
	}
}

func (sb *serverBuilder) PickClientsAndroid(str string) *serverBuilder {
	sb.android = str
	return sb
}

func (sb *serverBuilder) PickClientsIOS(str string) *serverBuilder {
	sb.ios = str
	return sb
}
func New() *serverBuilder {
	return &serverBuilder{}
}
