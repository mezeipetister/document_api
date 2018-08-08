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

	"github.com/mezeipetister/document_api/pkg/db"
	"github.com/mezeipetister/document_api/user"
)

// func Demo() string {
// 	return "Demo"
// }

func main() {

	client := db.NewClient()
	defer client.Disconnect(context.Background())

	u := user.NewUser(client)
	u.Email = "mezeipetister@gmail.com"
	u.SetPassword(context.Background(), "HelloBello!")
	u.Save()

	// d := document.NewDocument(client)
	// d.Name = "DemoDemo"
	// d.Description = "Peti"

	// c := context.Background()
	// d.SetLog(c, "New document created")

	// d.Save()

	// d.SetLog(c, "hellobello")
	// d.SetLog(c, "hellobello2")
	// d.SetLog(c, "hellobello3")
	// d.SetLog(c, "hellobello4")

	// d.SetComment(c, "First comment")
	// d.SetComment(c, "Second comment")
	// d.SetComment(c, "Third comment")

	// d.SetPartner(c, "A")
	// d.SetPartner(c, "B")
	// d.SetPartner(c, "C")

	// t := document.Task{
	// 	Title: "Demo Task Title",
	// }
	// d.SetTask(c, t)

	// result := model.FindDocument(c, client, "Kriszti")
	// fmt.Println(result.Logs[0].Message)
}
