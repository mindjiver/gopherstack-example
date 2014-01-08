package main

import (
//	"flag"
	"os"
	"fmt"
	"github.com/mitchellh/packer/builder/cloudstack"
)

func main() {
//	request := flag.String("command", "listVirtualMachines", "List Virtual Machines")
//	flag.Parse()

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

	serviceofferingid := "a7f96693-f86e-4e35-92e7-44870f4146dc"
	templateid :=        "9a0ddd35-5e4a-4675-b668-9c7b89124636"
	zoneid :=            "489e5147-85ba-4f28-a78d-226bf03db47c"
	networkids :=        []string{"9ab9719e-1f03-40d1-bfbe-b5dbf598e27f"}

	cs := cloudstack.CloudStackClient{}.New(apiurl, apikey, secret)
	key, _ := cs.CreateSSHKeyPair("packer-key-pair")
	fmt.Printf("key : %s", key)

	vmid, _ := cs.DeployVirtualMachine(serviceofferingid, templateid, zoneid, networkids, "packer-key-pair", "packer", "")
	fmt.Printf("vmid : %s", vmid)

//	cs.Templates
	// rootdeviceid, err  := cloudstack.StopVirtualMachine(vmid)
	// templateid, err = cloudstack.CreateTemplate("Packer Generated Template", "packer-generated-template", osid, rootdeviceid)
	// _, err = cloudstack.VirtualMachineState(vmid)

//	cs.DestroyVirtualMachine(vmid)
	cs.DeleteSSHKeyPair("packer-key-pair")
}
