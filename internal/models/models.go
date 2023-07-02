package models

// Conf - web gui config
type Conf struct {
	Host       string
	Port       string
	Theme      string
	Color      string
	Icon       string
	ConfPath   string
	YamlPath   string
	NodePath   string
	ColorOn    string
	ColorOff   string
	BtnWidth   string
	WebRefresh string
}

// Host - panel element
type Host struct {
	Name  string `yaml:"name"`
	Addr  string `yaml:"addr"`
	Port  string `yaml:"port"`
	URL   string `yaml:"url"`
	Icon  string `yaml:"icon"`
	State bool   `yaml:"state"`
}

// Panel - tab element
type Panel struct {
	Name    string       `yaml:"name"`
	Scan    bool         `yaml:"scan"`
	Timeout string       `yaml:"timeout"`
	Hosts   map[int]Host `yaml:"hosts"`
}

// Tab - board element
type Tab struct {
	Name   string         `yaml:"name"`
	Panels map[int]string `yaml:"panels"`
}

// MonPanel - uptime element
type MonPanel struct {
	Retries int      `yaml:"retries"`
	Notify  []string `yaml:"notify"`
}

// MonData - for monitoring results
type MonData struct {
	Panel string
	Host  string
	Addr  string
	Port  string
	Date  string
	State bool
}

// Uptime - board element
type Uptime struct {
	Enabled bool                `yaml:"enabled"`
	Notify  map[string]string   `yaml:"notify"`
	Panels  map[string]MonPanel `yaml:"panels"`
}

// Links - all links
type Links struct {
	Tabs   map[int]Tab      `yaml:"tabs"`
	Panels map[string]Panel `yaml:"panels"`
	Uptime Uptime           `yaml:"uptime"`
}

// GuiData - web gui data
type GuiData struct {
	Config     Conf
	Themes     []string
	Links      Links
	Panels     map[int]Panel
	CurrentTab string
	TabEdit    int
	UptimeMon  []MonData
}
