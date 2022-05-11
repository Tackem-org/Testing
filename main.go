package main

import (
	"github.com/Tackem-org/Global/system"
	"github.com/Tackem-org/Global/system/setupData"
	"github.com/Tackem-org/testing/static"
	"github.com/Tackem-org/testing/web"

	pbc "github.com/Tackem-org/Global/pb/config"
	pbr "github.com/Tackem-org/Global/pb/registration"
)

const (
	masterConfigFile = "test.json"
	logFile          = "test.log"
)

var (
	Version    string = "v0.0.0-devel"
	Commit     string
	CommitDate string
	sd         = &setupData.SetupData{

		ServiceName: "test",
		ServiceType: "system",
		SingleRun:   false,
		StartActive: true,
		VerboseLog:  true,
		ConfigItems: []*pbr.ConfigItem{
			{
				Key:          "user.password.minimum",
				DefaultValue: "8",
				Type:         pbc.ValueType_Uint,
				Label:        "Password Minimum Length",
				HelpText:     "what is the minimum password length",
				InputType:    pbr.InputType_INumber,
				InputAttributes: &pbr.InputAttributes{
					Other: map[string]string{
						"min": "1",
						"max": "16",
					},
				},
			},
		},
		NavItems: []*pbr.NavItem{
			{LinkType: pbr.LinkType_User, Title: "Test", Icon: "user", Path: "/test"},
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
	system.Run(sd, masterConfigFile, logFile, Version, Commit, CommitDate)
}
