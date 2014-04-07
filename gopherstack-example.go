package main

import (
	"fmt"
	"github.com/mindjiver/gopherstack"
	"os"
)

func main() {

	apiurl := os.Getenv("CLOUDSTACK_API_URL")
	if len(apiurl) == 0 {
		fmt.Println("Needed environment variable CLOUDSTACK_API_URL not found, exiting")
		os.Exit(1)
	}
	apikey := os.Getenv("CLOUDSTACK_API_KEY")
	if len(apikey) == 0 {
		fmt.Println("Needed environment variable CLOUDSTACK_API_KEY not found, exiting")
		os.Exit(1)
	}
	secret := os.Getenv("CLOUDSTACK_SECRET")
	if len(secret) == 0 {
		fmt.Println("Needed environment variable CLOUDSTACK_SECRET not found, exiting")
		os.Exit(1)
	}

	// Always validate any SSL certificates in the chain
	insecureskipverify := false
	cs := gopherstack.CloudstackClient{}.New(apiurl, apikey, secret, insecureskipverify)

	vmid := "19d2acfb-e281-4a13-8d62-e04ab501271d"
	response, err := cs.ListVirtualMachines(vmid)
	if err != nil {
		fmt.Errorf("Error listing virtual machine: %s", err)
		os.Exit(1)
	}

	if len(response.Listvirtualmachinesresponse.Virtualmachine) > 0 {
		ip := response.Listvirtualmachinesresponse.Virtualmachine[0].Nic[0].Ipaddress
		state := response.Listvirtualmachinesresponse.Virtualmachine[0].State
		fmt.Printf("%s has IP : %s and state : %s\n", vmid, ip, state)
	} else {
		fmt.Printf("No VM with UUID: %s found\n", vmid)
	}

}
