package models

// Conf - web gui config
type Conf struct {
	Host     string
	Port     string
	Theme    string
	Color    string
	Icon     string
	ConfPath string
}

// Host - panel element
type Host struct {
	Addr string
	Port string
	Icon string
}

// Panel - board element
type Panel struct {
	ID    int
	Name  string
	Hosts []Host
}

// GuiData - web gui data
type GuiData struct {
	Config Conf
	Themes []string
	Panels []Panel
}
