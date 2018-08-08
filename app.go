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
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/mongodb/mongo-go-driver/bson/objectid"

	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/mezeipetister/document_api/document"
	"github.com/mezeipetister/document_api/pkg/common"
	"github.com/mezeipetister/document_api/pkg/db"
)

// Response template
type Response struct {
	Message string
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&Response{Message: "Ok!"})
}

func createDocument(w http.ResponseWriter, r *http.Request) {
	client := db.NewClient()
	defer client.Disconnect(context.Background())

	d := document.NewDocument(client)

	// Parse the request body to a document struct
	rbody, _ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(rbody, &d); err != nil {
		panic(err)
	}

	d.Save()

	w.WriteHeader(http.StatusOK)

	type response struct {
		DocumentID string `json:"document_id"`
	}

	json.NewEncoder(w).Encode(&response{DocumentID: d.ID})
}

func getDocument(w http.ResponseWriter, r *http.Request) {
	client := db.NewClient()
	defer client.Disconnect(context.Background())

	vars := mux.Vars(r)
	result := document.GetDocumentByID(context.Background(), client, vars["id"])

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(&result)
}

func findDocument(w http.ResponseWriter, r *http.Request) {
	client := db.NewClient()
	defer client.Disconnect(context.Background())

	type query struct {
		Query string
	}

	q := query{}

	// Parse the request body to a document struct
	rbody, _ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(rbody, &q); err != nil {
		panic(err)
	}
	result := document.FindDocuments(context.Background(), client, q.Query)

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(&result)
}

func uploadDocument(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("uploadfile")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	filePath := objectid.New().Hex() + filepath.Ext(handler.Filename)
	f, err := os.OpenFile("./public/files/"+filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)

	type result struct {
		FileName string `json:"file_name"`
		FilePath string `json:"file_path"`
		FileSize int64  `json:"file_size"`
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&result{
		FileName: handler.Filename,
		FilePath: filePath,
		FileSize: handler.Size,
	})
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", index).Methods("GET")
	r.HandleFunc("/document/create", createDocument).Methods("POST")
	r.HandleFunc("/document/get/{id}", getDocument).Methods("GET")
	r.HandleFunc("/document/find", findDocument).Methods("POST")
	r.HandleFunc("/document/upload", uploadDocument).Methods("POST")

	port := ":" + strconv.Itoa(common.Config.Server.Port)
	log.Fatal(http.ListenAndServe(port, r))

	// u := user.NewUser(client)
	// u.Email = "mezeipetister@gmail.com"
	// u.SetPassword(context.Background(), "HelloBello!")
	// u.Save()

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
