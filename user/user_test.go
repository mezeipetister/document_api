/*
 * Created on Wed May 30 2018
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

package user

import (
	"testing"

	"github.com/mezeipetister/document_api/dao"
)

const (
	serverAddress  = "localhost"
	dbName         = "DEMO"
	collectionName = "user"
)

func TestNewUser(t *testing.T) {
	db, _ := dao.New(serverAddress)
	defer db.CloseSession()
	if u1, err := New(db, dbName, collectionName); err == nil {
		u1.SetFName("Peter")
		u1.SetLName("Mezei")
		u1.SetEmail("mezeipietster@gmail.com")
		u1.Save()
		return
	} else {
		t.Errorf("Error occured during inserting new test user. Error message: %s", err)
	}
}
