package ios

type Ios interface {
	ShowIOSIP() string
}

type iosBuilder struct {
	ios string
}

func (ib *iosBuilder) ShowIOSIP() string {
	return "iOS IP is  " + ib.ios
}

func (ib *iosBuilder) PickClientsIOS(ios string) *iosBuilder {
	ib.ios = ios
	return ib
}
