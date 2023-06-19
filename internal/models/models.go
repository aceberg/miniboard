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
	DefTab   string
}

// Host - panel element
type Host struct {
	// ID    string `yaml:"id"`
	Name  string `yaml:"name"`
	Addr  string `yaml:"addr"`
	Port  string `yaml:"port"`
	URL   string `yaml:"url"`
	Icon  string `yaml:"icon"`
	State bool   `yaml:"state"`
}

// Panel - tab element
type Panel struct {
	Name string `yaml:"name"`
	Scan bool   `yaml:"scan"`
	// Hosts []Host `yaml:"hosts"`
	Hosts map[int]Host `yaml:"hosts"`
}

// Tab - board element
type Tab struct {
	Name   string   `yaml:"name"`
	Panels []string `yaml:"panels"`
}

// Links - all links
type Links struct {
	Tabs   map[string]Tab   `yaml:"tabs"`
	Panels map[string]Panel `yaml:"panels"`
}

// GuiData - web gui data
type GuiData struct {
	Config     Conf
	Themes     []string
	Links      Links
	CurrentTab string
}
