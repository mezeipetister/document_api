/*
 * Created on Sat Aug 18 2018
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

mod document_service;
pub mod document;

extern crate protobuf;
extern crate grpc;
extern crate futures;
extern crate futures_cpupool;

fn main() {
    let mut d1 = document::DocumentScheme::new();
    d1.set_title("Hello Bello".to_string());
    println!("{}", d1.take_title());
}