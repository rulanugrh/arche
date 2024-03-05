package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

  args := os.Args[1]
  
  forBroad := strings.Split(args, "/")
  class, _ := strconv.Atoi(forBroad[1])
  
  forBroad = strings.Split(forBroad[0], ".")
  network, ranges, broadcast, subnetmask := calcIP(forBroad[0], forBroad[1], forBroad[2], forBroad[3], class)
  fmt.Printf("Network : %s\n", strings.Join(network, "."))
  fmt.Printf("Ranges : %s\n", ranges)
  fmt.Printf("Broadcast : %s\n", strings.Join(broadcast, "."))
  fmt.Printf("Subnetmask : %s\n", strings.Join(subnetmask, "."))
}

func calcNetwork(net int) (class string) {
  if net >= 24 {
    class = "C"
    return class
  } else if net >= 16 {
    class = "B"
    return class
  } else if net >= 8 {
    class = "A"
    return class
  } else {
    class = "invalid class"
    return class
  }
}

func calcNet( net int ) int {
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

func divIP(octet string, calc_net int) int{
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
    result := (roundUp+1) * float64(calc_net)
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
    broadcast = append(broadcast, a, b, strconv.Itoa(lastBroadcast),"255")
    ranges = fmt.Sprintf("%s.%s.%d.%d - %s.%s.%d.%d", a, b, forNetwork, 1, a, b, lastBroadcast, 254)
  } else if class >= 8 {
    lastBroadcast := forBroadcast(b, calc_net)
    forNetwork := divIP(b, calc_net)
    subnetmask = append(subnetmask, "255", strconv.Itoa(lastNetmask), "0", "0")
    network = append(network, a, strconv.Itoa(forNetwork), "0", "0")
    broadcast = append(broadcast, a, strconv.Itoa(lastBroadcast),"255", "255")
    ranges = fmt.Sprintf("%s.%d.%d.%d - %s.%d.%d.%d", a, forNetwork, 0, 1, a, lastBroadcast, 255, 254)
    

  } else {
    log.Println("Sorry invalid netmask IP Address")
  }

  return network, ranges, broadcast, subnetmask
}
