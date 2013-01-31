/*
Solution to problem 1 in Project Eular: http://projecteuler.net/problem=1

Copyright (C) 2013  David Murphy <dave@schwuk.com>

This program is free software: you can redistribute it and/or modify it under
the terms of the GNU General Public License as published by the Free Software
Foundation, either version 3 of the License, or (at your option) any later
version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY
WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A
PARTICULAR PURPOSE.  See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with
this program.  If not, see <http://www.gnu.org/licenses/>.

*/
package main

import "fmt"

func Solution() int {
	total := 0
	for i := 1; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			total += i
		}
	}
	return total
}

func main() {
	fmt.Println(Solution())
}
