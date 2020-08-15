package traffic_stubs

type Info struct {
	Name string
	Version string
}
type Infos struct {
	Count int
	Details []Info
}

type Quote struct {
	By string
	Message string
}
type Quotes struct {
	Count int
	Details []Quote
}
