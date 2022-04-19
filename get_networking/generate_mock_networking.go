package get_networking

import (
	"fmt"
	"strconv"
	"strings"
)

type SwitchData struct {
	Size   string
	Space  string
	Cidr   string
	Region string
}

func MakeTheSwitch(s SwitchData) ([]string, error) {
	switch s.Size {
	case "grande":
		if len(s.Cidr) != 0 {
			networks, err := BuildGrandeNetworks(s)
			if err != nil {
				return networks, err
			}
			return networks, nil
		} else {
			s.Cidr = "/24"
			networks, err := BuildGrandeNetworks(s)
			if err != nil {
				return networks, err
			}
			return networks, nil
		}

	case "venti":
		if len(s.Cidr) != 0 {
			networks, err := BuildGrandeNetworks(s)
			if err != nil {
				return networks, err
			}
			return networks, nil
		} else {
			s.Cidr = "/23"
			networks, err := BuildGrandeNetworks(s)
			if err != nil {
				return networks, err
			}
			return networks, nil
		}
	default:
		if len(s.Cidr) != 0 {
			networks, err := BuildGrandeNetworks(s)
			if err != nil {
				return networks, err
			}
			return networks, nil
		} else {
			s.Cidr = "/24"
			networks, err := BuildGrandeNetworks(s)
			if err != nil {
				return networks, err
			}
			return networks, nil
		}
	}
}

//BuildGrandeNetworks builds a 16 oz virtual network
func BuildGrandeNetworks(a SwitchData) ([]string, error) {
	var s []string
	for i := 0; i < 128; i++ {
		next := strings.Replace(strings.Split(a.Space, ".")[2], strings.Split(a.Space, ".")[2], strconv.Itoa(i+2), 2)
		s = append(s, next)
	}
	if len(s) == 0 {
		return s, fmt.Errorf("not able to build any grande address spaces in BuildGrandeNetworks")
	} else {
		return s, nil
	}
}

//BuildVentiNetworks builds a 20oz virtual network
func BuildVentiNetworks(a SwitchData) ([]string, error) {
	var s []string
	for i := 0; i < 128; i++ {
		next := strings.Replace(strings.Split(a.Space, ".")[2], strings.Split(a.Space, ".")[2], strconv.Itoa(i+1), 2)
		s = append(s, next)
	}
	if len(s) == 0 {
		return s, fmt.Errorf("not able to build any venti address spaces in BuildVentiNetworks")
	} else {
		return s, nil
	}
}
