package utils

import (
	"fmt"

	"github.com/mwohlert/gonetworkmanager"
)

// DeleteWirelessConnections deletes all Connections in NetworkManager with the
// type 802-11-wireless
func DeleteWirelessConnections() error {
	settings, err := gonetworkmanager.NewSettings()
	if err != nil {
		return err
	}

	connections, err := settings.ListConnections()
	if err != nil {
		return err
	}
	for _, connection := range connections {
		connectionType := fmt.Sprintf("%v", connection.GetSettings()["connection"]["type"])

		if connectionType == "802-11-wireless" {
			connection.Delete()
		}
	}
	return nil
}
