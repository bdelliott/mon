package main

import (
	"flag"
	"fmt"
	"github.com/bdelliott/mon/pkg/openstack/auth"
	"github.com/bdelliott/mon/pkg/rackspace/monitoring"
	"github.com/rackspace/gophercloud/openstack"
	"os"
)

func main() {

	flag.Usage = func() {
		fmt.Printf("Usage: %s cloudName\n", os.Args[0])
	}
	flag.Parse()

	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	args := flag.Args()
	cloudName := args[0]

	authOptions := auth.FromCloudsYaml(cloudName)
	openstackClient, err := openstack.AuthenticatedClient(authOptions)
	if err != nil {
		panic(err)
	}

	zones, err := monitoring.GetZones(authOptions.TenantID, openstackClient.TokenID)
	if err != nil {
		panic(err)
	}

	for _, zone := range zones {
		fmt.Println("Zone id: ", zone.Id)
	}

}
