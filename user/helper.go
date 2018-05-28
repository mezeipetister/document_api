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

package user

import "golang.org/x/crypto/bcrypt"

// Generate hash from given string; using bcrypt hash generator.
// It's a bit slow - approx. 1500ms, but it's a built-in slowlyness.
func hashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), 14)
}

// Compare a given string - convert to hash - to a given hash.
// It's a bit slow - approx. 1500ms, but it's a built-in slowlyness.
func checkPasswordHash(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
