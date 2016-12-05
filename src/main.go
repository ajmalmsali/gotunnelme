package main

import (
  "gotunnelme"
  "fmt"
  "flag"
)

func main() {
  tunnel := gotunnelme.NewTunnel()
  fmt.Println("\nTunnel using http://ajm.al \n=========")
  subdomainPtr := flag.String("s", "project1", "a subdomain")
  portPtr := flag.Int("port", 8080, "local port")
  flag.Parse()

  url, getUrlErr := tunnel.GetUrl(*subdomainPtr)
  if getUrlErr != nil {
    fmt.Println(getUrlErr)
    fmt.Println("Url Creation Issue!")
  } else {
    fmt.Println("Url Created")
    fmt.Println("Get Url:", url)
    fmt.Printf("Creating Tunnel... on Port %d \n", *portPtr)
    
    tunnelErr := tunnel.CreateTunnel(*portPtr)
    if tunnelErr != nil {
      fmt.Println(tunnelErr)
      fmt.Println("Tunnel Creation Failed!")
    } else {
      fmt.Println("Get Url:", url)
      fmt.Println("Tunnel Created!")
    }
  }
}
