// SPDX-License-Identifier: Apache-2.0
// Copyright 2022 Intel Corporation
// Copyright 2019 free5GC.org

package main

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"

	"github.com/omec-project/udm/logger"
	"github.com/omec-project/udm/service"
)

var UDM = &service.UDM{}

var appLog *logrus.Entry

func init() {
	appLog = logger.AppLog
}

func main() {
	app := cli.NewApp()
	app.Name = "HEXA_udm"
	fmt.Print(app.Name, "\n")
	app.Usage = "-free5gccfg common configuration file -udmcfg udm configuration file"
	app.Action = action
	app.Flags = UDM.GetCliCmd()
	if err := app.Run(os.Args); err != nil {
		appLog.Errorf("UDM Run error: %v", err)
	}
}

func action(c *cli.Context) error {
	if err := UDM.Initialize(c); err != nil {
		logger.CfgLog.Errorf("%+v", err)
		return fmt.Errorf("Failed to initialize !!")
	}

	UDM.Start()

	return nil
}
