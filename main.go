package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

func main() {

	args := os.Args[1]

	forBroad := strings.Split(args, "/")
	class, _ := strconv.Atoi(forBroad[1])

	forBroad = strings.Split(forBroad[0], ".")
	check := checkOctectClass(forBroad[0], class)
	if check {
		network, ranges, broadcast, subnetmask := calcIP(forBroad[0], forBroad[1], forBroad[2], forBroad[3], class)
		printOut(network, ranges, broadcast, subnetmask)
	} else {
		errors("IP Class out of Index")
	}
}

func calcNet(net int) int {
	if net >= 24 {
		pars := 32 - net
		res := math.Pow(2, float64(pars))
		return int(res)
	} else if net >= 16 {
		pars := 24 - net
		res := math.Pow(2, float64(pars))
		return int(res)
	} else if net >= 8 {
		pars := 16 - net
		res := math.Pow(2, float64(pars))
		return int(res)
	} else {
		return 0
	}
}

func divIP(octet string, calc_net int) int {
	n, _ := strconv.Atoi(octet)
	res := n / calc_net
	rounDown := math.Round(float64(res))
	result := rounDown * float64(calc_net)
	return int(result)

}

func forBroadcast(octet string, calc_net int) int {
	if calc_net == 256 {
		return 255
	} else {
		n, _ := strconv.Atoi(octet)
		div_res := n / calc_net
		roundUp := math.Ceil(float64(div_res))
		result := (roundUp + 1) * float64(calc_net)
		return int(result) - 1
	}
}

func calcIP(a string, b string, c string, d string, class int) (network []string, ranges string, broadcast []string, subnetmask []string) {

	calc_net := calcNet(class)
	lastNetmask := 256 - calc_net
	if lastNetmask == 0 {
		lastNetmask = 0
	}

	if class >= 24 {
		lastBroadcast := forBroadcast(d, calc_net)
		forNetwork := divIP(d, calc_net)
		subnetmask = append(subnetmask, "255", "255", "255", strconv.Itoa(lastNetmask))
		network = append(network, a, b, c, strconv.Itoa(forNetwork))
		broadcast = append(broadcast, a, b, c, strconv.Itoa(lastBroadcast))
		ranges = fmt.Sprintf("%s.%s.%s.%d - %s.%s.%s.%d", a, b, c, (forNetwork + 1), a, b, c, (lastBroadcast - 1))

	} else if class >= 16 {
		lastBroadcast := forBroadcast(c, calc_net)
		forNetwork := divIP(c, calc_net)
		subnetmask = append(subnetmask, "255", "255", strconv.Itoa(lastNetmask), "0")
		network = append(network, a, b, strconv.Itoa(forNetwork), "0")
		broadcast = append(broadcast, a, b, strconv.Itoa(lastBroadcast), "255")
		ranges = fmt.Sprintf("%s.%s.%d.%d - %s.%s.%d.%d", a, b, forNetwork, 1, a, b, lastBroadcast, 254)

	} else if class >= 8 {
		lastBroadcast := forBroadcast(b, calc_net)
		forNetwork := divIP(b, calc_net)
		subnetmask = append(subnetmask, "255", strconv.Itoa(lastNetmask), "0", "0")
		network = append(network, a, strconv.Itoa(forNetwork), "0", "0")
		broadcast = append(broadcast, a, strconv.Itoa(lastBroadcast), "255", "255")
		ranges = fmt.Sprintf("%s.%d.%d.%d - %s.%d.%d.%d", a, forNetwork, 0, 1, a, lastBroadcast, 255, 254)

	} else {
		log.Println("Sorry invalid netmask IP Address")
	}

	return network, ranges, broadcast, subnetmask
}

func printOut(network []string, ranges string, broadcast []string, subnet []string) {
	var Cyan = "\033[36m"

	t := table.NewWriter()
	t.AppendRow(table.Row{"Network		", strings.Join(network, ".")})
	t.AppendRow(table.Row{"Range		", ranges})
	t.AppendRow(table.Row{"Broadcast	", strings.Join(broadcast, ".")})
	t.AppendRow(table.Row{"Subnetmask	", strings.Join(subnet, ".")})
	t.SetStyle(table.StyleLight)
	t.Style().Title.Align = text.AlignCenter
	t.SetTitle(Cyan + "arche")

	fmt.Println("\n" + t.Render() + "\n")
}

func errors(txt string) {
	var Cyan = "\033[36m"

	t := table.NewWriter()
	t.AppendRow(table.Row{"error  ", txt})
	t.SetStyle(table.StyleLight)
	t.Style().Title.Align = text.AlignCenter
	t.SetTitle(Cyan + "arche")

	fmt.Println("\n" + t.Render() + "\n")
}

func checkOctectClass(a string, net int) bool {
	var forA []string
	var forB []string
	var forC []string

	for classA := 1; classA <= 127; classA++ {
		forA = append(forA, strconv.Itoa(classA))
	}

	for classB := 128; classB <= 191; classB++ {
		forB = append(forB, strconv.Itoa(classB))
	}

	for classC := 192; classC <= 255; classC++ {
		forC = append(forC, strconv.Itoa(classC))
	}

	if net >= 8 && checkTrue(forA, a) {
		return true
	} else if net >= 16 && checkTrue(forB, a) {
		return true
	} else if net >= 24 && checkTrue(forC, a){
		return true
	} else {
		return false
	}

}

func checkTrue(elems []string, containt string) bool {
	for _, data := range elems {
		check := strings.Contains(data, containt)
		if check {
			return true
		}
	}

	return false

}
