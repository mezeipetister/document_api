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
	"testing"
	"unsafe"
)

func TestGetConfig(t *testing.T) {
	c := New()
	if unsafe.Sizeof(c) == 0 {
		t.Error("Config file is empty after initialization.")
	}
	if (len(c.getServerAddress()) == 0) ||
		(c.getServerPort() == 0) {
		t.Error("No server address and server port set in configuration file.")
	}
}

func TestGetSampleSettings(t *testing.T) {
	c := New()
	if serverAddress := c.getServerAddress(); serverAddress != "localhost" {
		t.Errorf("Server address is not the required in the test settings. Now: %s; expected: %s", serverAddress, "localhost")
	}
}
