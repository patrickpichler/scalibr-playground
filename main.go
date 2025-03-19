package main

import (
	"context"
	"fmt"

	scalibr "github.com/google/osv-scalibr"
	"github.com/google/osv-scalibr/artifact/image/layerscanning/image"
	dl "github.com/google/osv-scalibr/detector/list"
	el "github.com/google/osv-scalibr/extractor/filesystem/list"
	"github.com/google/osv-scalibr/plugin"
)

func main() {
	s := scalibr.New()
	fmt.Println("get image start")
	img, err := image.FromRemoteName("docker.io/library/alpine:latest", image.DefaultConfig())
	if err != nil {
		panic(err)
	}
	fmt.Println("get image end")

	capas := &plugin.Capabilities{
		OS:            plugin.OSLinux,
		Network:       plugin.NetworkAny,
		DirectFS:      false,
		RunningSystem: false,
	}

	fmt.Println("scan container start")
	res, err := s.ScanContainer(context.Background(), img, &scalibr.ScanConfig{
		FilesystemExtractors: el.FromCapabilities(capas),
		Detectors:            dl.FromCapabilities(capas),
		Capabilities:         capas,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("scan container end")

	fmt.Printf("res.Inventories: %v\n", len(res.Inventories))
	for _, inv := range res.Inventories {
		fmt.Printf("inv.Name: %v\n", inv.Name)
		for _, loc := range inv.Locations {
			fmt.Printf("loc: %v\n", loc)
		}
	}
}
