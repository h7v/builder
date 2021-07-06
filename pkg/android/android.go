package android

type Android interface {
	ShowAndroidIP()
}

type androidBuilder struct {
	android string
}

func (ab *androidBuilder) ShowAndroidIP() string {
	return "Android IP is" + ab.android
}

func (ab *androidBuilder) PickClientsAndroid(android string) *androidBuilder {
	ab.android = android
	return ab
}
