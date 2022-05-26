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
	spl := strings.Split(a.Space, ".")
	for i := 0; i < 4080; i++ {
		morph := []string{
			"16",
			"32",
			"48",
			"64",
			"80",
			"96",
			"112",
			"128",
			"144",
			"160",
			"176",
			"192",
			"208",
			"224",
			"240",
		}
		next, _ := strconv.Atoi(spl[2])
		spl[2] = strconv.Itoa(next + 1)
		joined := strings.Join(spl, ".")
		s = append(s, fmt.Sprintf("%s%s", joined, a.Cidr))
		for _, v := range morph {
			m := strings.Split(joined, ".")
			m[3] = strings.Replace(m[3], m[3], v, 1)
			rejoin := strings.Join(m, ".")
			s = append(s, fmt.Sprintf("%s%s", rejoin, a.Cidr))
		}
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
	for i := 2; i < 127; i++ {
		morph := []string{
			"64",
			"128",
			"192",
		}
		next, _ := strconv.Atoi(spl[2])
		spl[2] = strconv.Itoa(next + 1)
		joined := strings.Join(spl, ".")
		s = append(s, fmt.Sprintf("%s%s", joined, a.Cidr))
		for _, v := range morph {
			m := strings.Split(joined, ".")
			m[3] = v
			rejoin := strings.Join(m, ".")
			s = append(s, fmt.Sprintf("%s%s", rejoin, a.Cidr))
		}
	}

	for i := 129; i < 255; i++ {
		morph := []string{
			"64",
			"128",
			"192",
		}
		next, _ := strconv.Atoi(spl[2])
		spl[2] = strconv.Itoa(next + 1)
		joined := strings.Join(spl, ".")
		s = append(s, fmt.Sprintf("%s%s", joined, a.Cidr))
		for _, v := range morph {
			m := strings.Split(joined, ".")
			m[3] = v
			rejoin := strings.Join(m, ".")
			s = append(s, fmt.Sprintf("%s%s", rejoin, a.Cidr))
		}
	}
	// fmt.Println(s)
	if len(s) == 0 {
		return s, fmt.Errorf("not able to build any venti address spaces in BuildVentiNetworks")
	} else {
		return s, nil
	}
}
