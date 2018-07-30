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
	"context"

	"github.com/mezeipetister/document_api/model"
	"github.com/mezeipetister/document_api/pkg/db"
)

func main() {
	client := db.NewClient()
	defer client.Disconnect(context.Background())

	d := model.NewDocument(client)
	d.Name = "HelloBello!"

	c := context.Background()
	d.SetLog(c, "New document created")

	d.Save()

	d.SetLog(c, "hellobello")
	d.SetLog(c, "hellobello2")
	d.SetLog(c, "hellobello3")
	d.SetLog(c, "hellobello4")

	d.Remove(c)

	// time.Sleep(time.Second * 5)

	// d.Remove()
}
