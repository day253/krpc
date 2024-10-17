package registry_zookeeper

const (
	DefaultRegistrySeparater = ","
)

type Conf struct {
	Enabled   bool
	Metabase  string
	TimeoutMs int
	User      string
	Password  string
}
