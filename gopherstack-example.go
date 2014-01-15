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

	serviceofferingid := "e43adbcf-4004-40fb-a452-766aaf3a55f3"
	templateid := "b34f2d7b-2bec-497e-a18e-06d0de94526e"
	diskofferingid := "df0a9e58-5c1b-4943-8fdf-2469ce945b5b"
	zoneid := "489e5147-85ba-4f28-a78d-226bf03db47c"
	networkids := []string{"9ab9719e-1f03-40d1-bfbe-b5dbf598e27f"}
	userdata := "#!ipxe\nchain http://10.4.128.74/cloudstack.ipxe\n"

	key_pair_name := "packer-key-pair"
	displayname := "peter-packer-testing"

	cs := gopherstack.CloudStackClient{}.New(apiurl, apikey, secret)

	cs.CreateSSHKeyPair(key_pair_name)
	vmid, jobid, _ := cs.DeployVirtualMachine(serviceofferingid, templateid,
		zoneid, "", diskofferingid, displayname, networkids,
		key_pair_name,  "", userdata)
	cs.WaitForAsyncJob(jobid, 2*time.Minute)

	ip, state, _ := cs.ListVirtualMachines(vmid)
	fmt.Printf("%s has IP : %s and state : %s", vmid, ip, state)

	jobid, _ = cs.DetachIso(vmid)
	cs.WaitForAsyncJob(jobid, 2*time.Minute)

	ip, state, _ = cs.ListVirtualMachines(vmid)
	fmt.Printf("%s has IP : %s and state : %s", vmid, ip, state)

	cs.DeleteSSHKeyPair(key_pair_name)
}
