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

//	serviceofferingid := "62fc8ae5-06ac-4021-bed6-90dfdca6b6b5"
	serviceofferingid := "a7f96693-f86e-4e35-92e7-44870f4146dc"
	templateid :=        "9a0ddd35-5e4a-4675-b668-9c7b89124636"
	zoneid :=            "489e5147-85ba-4f28-a78d-226bf03db47c"
	networkid :=         "9ab9719e-1f03-40d1-bfbe-b5dbf598e27f"

	// ipxe boot
	//templateid :=        "26de0a07-eee6-4b00-9c4f-fdb7b29f6ba2"
	// only needed for booting from ISO to kickstart entire OS.
	//diskofferingid :=    "ef781d7f-f8e8-4f73-985c-e0b0a8ef8d48"
	// CentOS 6.0
	// osid := "144"

	cs := cloudstack.CloudStackClient{}.New(apiurl, apikey, secret)
	cs.CreateSSHKeyPair("packer-key-pair")
	vmid, _ := cs.DeployVirtualMachine(serviceofferingid, templateid, zoneid, networkid, "packer-key-pair", "packer-automation-test", "")
	fmt.Printf("vmid : %s", vmid)

	// rootdeviceid, err  := cloudstack.StopVirtualMachine(vmid)
	// templateid, err = cloudstack.CreateTemplate("Packer Generated Template", "packer-generated-template", osid, rootdeviceid)
	// _, err = cloudstack.VirtualMachineState(vmid)

//	cs.DestroyVirtualMachine(vmid)
	cs.DeleteSSHKeyPair("packer-key-pair")

	// _, err = cloudstack.DeleteSSHKeyPair("packer-key-pair")
}
