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

package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Config file name to use (import/export)
const configFileName string = "config.json"

// Config
var configuration config

// config members are exported
// because this way the JSON parser can
// manage the parsing.
type config struct {
	Server string `json:"server"`
	Port   int    `json:"port"`
}

// return the read configs
func getConfig() {

	// Check if the config file exists. If not, then create it
	// with default values.
	if _, err := os.Stat(configFileName); os.IsNotExist(err) {

		// Create default values.
		content, _ := json.MarshalIndent(&config{
			Server: "localhost",
			Port:   8080,
		}, "", "    ") // no prefix, but 4 spaces indent

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
	panic("Oo. An error occured while config file parsed.")
}
