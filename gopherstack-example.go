package main

import (
	"fmt"
	"github.com/mindjiver/gopherstack"
	"os"
	"time"
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

	serviceofferingid := "a7f96693-f86e-4e35-92e7-44870f4146dc"
	templateid := "9a0ddd35-5e4a-4675-b668-9c7b89124636"
	zoneid := "489e5147-85ba-4f28-a78d-226bf03db47c"
	networkids := []string{"9ab9719e-1f03-40d1-bfbe-b5dbf598e27f"}

	key_pair_name := "packer-key-pair"
	displayname := "packer-testing"

	cs := gopherstack.CloudStackClient{}.New(apiurl, apikey, secret)
	cs.CreateSSHKeyPair(key_pair_name)

	cs.ListProjects("")

	vmid, jobid, _ := cs.DeployVirtualMachine(serviceofferingid, templateid, zoneid, networkids, key_pair_name, displayname, "", "")
	cs.WaitForAsyncJob(jobid, 2*time.Minute)

	ip, state, _ := cs.ListVirtualMachines(vmid)
	fmt.Printf("%s has IP : %s and state : %s", vmid, ip, state)

	//jobid, _ = cs.StopVirtualMachine(vmid)
	// cs.WaitForAsyncJob(jobid, 5*time.Minute)

	_, state, _ = cs.ListVirtualMachines(vmid)
	fmt.Printf("%s has IP : %s and state : %s", vmid, ip, state)

	volumeId, _ := cs.ListVolumes(vmid)
	fmt.Printf("VM has volume id : %s", volumeId)

	// jobid, _ = cs.DestroyVirtualMachine(vmid)
	// cs.WaitForAsyncJob(jobid, 5*time.Minute)

	// _, state, _ = cs.VirtualMachineState(vmid)
	// fmt.Printf("%s has IP : %s and state : %s", vmid, ip, state)

	cs.DeleteSSHKeyPair(key_pair_name)
}
