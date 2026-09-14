package pkg

type Options struct {
	Dir       string
	Ext       string
	Name      string
	Username  string
	Password  string
	Port      uint16
	Extension bool
	Network   bool
}

func (o Options) authenticated() bool {
	return o.Username != "" && o.Password != ""
}

func (o Options) title() string {
	if o.Name != "" {
		return o.Name
	}
	return "Server"
}
