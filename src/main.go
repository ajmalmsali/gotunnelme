package main

import (
  "gotunnelme"
  "fmt"
  "flag"
)

func main() {
  tunnel := gotunnelme.NewTunnel()
  fmt.Println("\nTunnel using http://ajm.al \n=========")

  subdomainPtr := flag.String("s", "www", "a subdomain")
  portPtr := flag.Int("p", 8080, "local port")
  hostPtr := flag.String("h", "localhost", "host to forward to")
  tlsPtr := flag.Bool("tls", false, "SSL or not")
  flag.Parse()

  url, getUrlErr := tunnel.GetUrl(*subdomainPtr)

  if getUrlErr != nil {
    fmt.Println(getUrlErr)
    fmt.Println("Subdomain Creation Issue!")
  } else {
    fmt.Println("Subdomain Created")
    fmt.Println("Get Subdomain:", url)
    fmt.Printf("Creating Tunnel on %s:%d \n", *hostPtr, *portPtr)

    tunnelErr := tunnel.CreateTunnel(*hostPtr, *portPtr, *tlsPtr)
    if tunnelErr != nil {
      fmt.Println(tunnelErr)
      fmt.Println("Tunnel Creation Failed!")
    } else {
      //fmt.Println("Get Tunnel:", tunnelErr)
      fmt.Println("Tunnel Created!")
    }
  }
}
