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
			networks, err := BuildVentiNetworks(s)
			if err != nil {
				return networks, err
			}
			return networks, nil
		} else {
			s.Cidr = "/23"
			networks, err := BuildVentiNetworks(s)
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
	s = append(s, a.Space)
	for i := 0; i < 354; i++ {
		spl := strings.Split(a.Space, ".")
		next := strings.Replace(spl[2], spl[2], strconv.Itoa(i+1), 2)
		spl[2] = next
		joined := strings.Join(spl, ".")
		s = append(s, fmt.Sprintf("%s%s", joined, a.Cidr))
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
	spl := strings.Split(a.Space, ".")
	for i := 0; i < 128; i++ {
		if i%2 == 0 && spl[2] != "0" {
			next, _ := strconv.Atoi(spl[2])
			spl[2] = strconv.Itoa(next + i)
			joined := strings.Join(spl, ".")
			s = append(s, fmt.Sprintf("%s%s", joined, a.Cidr))
		} else if i%2 == 0 && spl[2] == "0" {
			next := strings.Replace(spl[2], spl[2], strconv.Itoa(i), 2)
			spl[2] = next
			joined := strings.Join(spl, ".")
			s = append(s, fmt.Sprintf("%s%s", joined, a.Cidr))
		}
	}
	if len(s) == 0 {
		return s, fmt.Errorf("not able to build any venti address spaces in BuildVentiNetworks")
	} else {
		return s, nil
	}
}
