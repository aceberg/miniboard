package models

// Conf - web gui config
type Conf struct {
	Host     string
	Port     string
	Theme    string
	Color    string
	Icon     string
	ConfPath string
	YamlPath string
	ColorOn  string
	ColorOff string
	BtnWidth string
}

// Host - panel element
type Host struct {
	Name  string `yaml:"name"`
	Addr  string `yaml:"addr"`
	Port  string `yaml:"port"`
	URL   string `yaml:"url"`
	Icon  string `yaml:"icon"`
	State bool
}

// Panel - board element
type Panel struct {
	Name  string `yaml:"name"`
	Scan  bool   `yaml:"scan"`
	Hosts []Host `yaml:"hosts"`
}

// Links - all links
type Links struct {
	Panels []Panel `yaml:"panels"`
}

// GuiData - web gui data
type GuiData struct {
	Config Conf
	Themes []string
	Links  Links
}
