/*
 * Created on Fri Jul 20 2018
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

package model

import (
	"fmt"

	"github.com/mezeipetister/document_api/pkg/setting"
)

// Document model
type Document struct {
	ID          string
	Name        string
	Description string
	File        string
	Folder      string
	Partners    []Partner
	DueDate     string
	Tasks       []Task
	Comments    []Comment
	Changelog   []Log
	Status      bool
}

// Remove document
func (d *Document) Remove() error {
	fmt.Println(setting.AppVersion)
	return nil
}

// SetName document
func (d *Document) SetName(name string) error {
	return nil
}

// SetDescription set the description of the document
func (d *Document) SetDescription(description string) error {
	return nil
}

// SetDueDate to the document
func (d *Document) SetDueDate(date string) error {
	return nil
}

// NewDocument return a new, empty document
func NewDocument() *Document {
	return &Document{}
}
