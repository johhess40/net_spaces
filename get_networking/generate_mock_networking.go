package get_networking

import "fmt"

type SwitchData struct {
	Size  string
	Space string
	Cidr  string
}

func MakeTheSwitch(s SwitchData) {
	switch s.Size {
	case "grande":
		if len(s.Cidr) != 0 {
			BuildGrandeNetworks(s)
		} else {
			s.Cidr = "/24"
			BuildGrandeNetworks(s)
		}

	case "venti":
		if len(s.Cidr) != 0 {
			BuildGrandeNetworks(s)
		} else {
			s.Cidr = "/23"
			BuildGrandeNetworks(s)
		}
	default:
		if len(s.Cidr) != 0 {
			BuildGrandeNetworks(s)
		} else {
			s.Cidr = "/24"
			BuildGrandeNetworks(s)
		}
	}
}

//BuildGrandeNetworks builds a 16 oz virtual network
func BuildGrandeNetworks(a SwitchData) []string {
	var s []string
	next := fmt.Sprintf("%s%s", a.Space, a.Cidr)
	return s
}

//BuildVentiNetworks builds a 20oz virtual network
func BuildVentiNetworks(a SwitchData) []string {
	var s []string
	return s
}
