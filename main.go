package main

import (
	"github.com/Tackem-org/Global/system"
	"github.com/Tackem-org/Global/system/setupData"
	"github.com/Tackem-org/testing/static"
	"github.com/Tackem-org/testing/web"

	pbr "github.com/Tackem-org/Global/pb/registration"
)

var (
	Version string = "v0.0.0-devel"
	sd             = &setupData.SetupData{
		ServiceName: "test",
		ServiceType: "system",
		SingleRun:   false,
		StartActive: true,
		VerboseLog:  true,
		NavItems: []*pbr.NavItem{
			{LinkType: pbr.LinkType_Main, Title: "Test", Icon: "user", Path: "/"},
		},
		StaticFS: &static.FS,
		Paths: []*setupData.PathItem{
			{
				Path:       "/",
				Permission: "",
				Call:       web.RootPage,
			},
		},
	}
)

func main() {
	system.Run(Version, sd)
}
