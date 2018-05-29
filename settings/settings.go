/*
 * Created on Sat May 26 2018
 * Copyright (c) 2018 Peter Mezei
 *
 * License AGPL v3.0
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published
 * by the Free Software Foundation, either version 3 of the License.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>
 *
 * For more information please contact me
 * via github.com
 */

package settings

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

const (
	// Config file name to use (import/export)
	configFileName string = "config.json"

	// errorMessagePanic to display once panic occures.
	errorMsgPanic string = "Oo. An error occured while config file parsed."
)

// Config
var configuration *Config

// Config struct.
type Config struct {
	ServerAddress   string `json:"server_address"`
	ServerPort      int    `json:"server_port"`
	DBServerAddress string `json:"db_server_address"`
	DBServerPort    int    `json:"db_server_port"`
}

// Interface for config.
type Interface interface {
	getServerAddress() string
	getServerPort() int
	getDBAddress() string
	getDBPort() int
}

// New config instance.
func New() *Config {
	getConfig()
	return configuration
}

// Get server address from the config.
func (c *Config) getServerAddress() string {
	return c.ServerAddress
}

// Get server port from the config.
func (c *Config) getServerPort() int {
	return c.ServerPort
}

// Get DB server address from the config.
func (c *Config) getDBAddress() string {
	return c.DBServerAddress
}

// Get DB server port from the config.
func (c *Config) getDBPort() int {
	return c.DBServerPort
}

// return the read configs
func getConfig() {

	// Check if the config file exists. If not, then create it
	// with default values.
	if _, err := os.Stat(configFileName); os.IsNotExist(err) {

		// Create default values.
		content, _ := json.MarshalIndent(
			&Config{
				ServerAddress:   "localhost",
				ServerPort:      8080,
				DBServerAddress: "localhost",
				DBServerPort:    0,
			},
			"", "    ") // no prefix, but 4 spaces indent

		// Write default vaules to the new config file.
		ioutil.WriteFile(configFileName, content, 0755)
	}

	// Read the config file. If no error, then return.
	if file, err := ioutil.ReadFile(configFileName); err == nil {
		err = json.Unmarshal(file, &configuration)
		return
	}

	// If something went wrong, and the config.json is not readable
	// then panic.
	panic(errorMsgPanic)
}
