/*
 * Created on Sat Jul 21 2018
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

package common

import (
	"io/ioutil"
	"sync"

	"github.com/BurntSushi/toml"
)

var once sync.Once

func init() {
	// Singleton!?
	once.Do(func() {
		// TODO: error handling!
		data, _ := ioutil.ReadFile("config.toml")
		if _, err := toml.Decode(string(data), &Config); err != nil {
			// handle error
		}
	})
}

// Config ...
var Config *setting

// Setting type
type setting struct {
	Title, Version string
	DB             db
	Server         server
}

type db struct {
	DBName             string `toml:"db_name"`
	CollectionDocument string `toml:"collection_document"`
	CollectionUser     string `toml:"collection_user"`
}

type server struct {
	Port int `toml:"port"`
}
