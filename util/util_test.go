/*
 * This file is part of Chihaya.
 *
 * Chihaya is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * Chihaya is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with Chihaya.  If not, see <http://www.gnu.org/licenses/>.
 */

package util

import (
	"math/rand"
	"testing"
)

func TestMin(t *testing.T) {
	for i := 0; i < 100; i++ {
		a := rand.Int()
		b := rand.Int()
		gotMin := Min(a, b)

		var actualMin int
		if b > a {
			actualMin = a
		} else {
			actualMin = b
		}

		if actualMin != gotMin {
			t.Fatalf("Min value (%d) is wrong for a=%d and b=%d!", gotMin, a, b)
		}
	}
}

func TestMax(t *testing.T) {
	for i := 0; i < 100; i++ {
		a := rand.Int()
		b := rand.Int()
		gotMax := Max(a, b)

		var actualMax int
		if b < a {
			actualMax = a
		} else {
			actualMax = b
		}

		if actualMax != gotMax {
			t.Fatalf("Max value (%d) is wrong for a=%d and b=%d!", gotMax, a, b)
		}
	}
}

func TestBtoa(t *testing.T) {
	for i := 0; i < 100; i++ {
		var b bool

		var actualResult string

		if rand.Intn(2) == 1 {
			b = true
			actualResult = "1"
		} else {
			b = false
			actualResult = "0"
		}

		gotResult := Btoa(b)
		if actualResult != gotResult {
			t.Fatalf("Got wrong string (%s) for bool %t!", gotResult, b)
		}
	}
}

func TestIntn(t *testing.T) {
	for i := 1; i < 2000; i++ {
		genInt := Intn(i)

		if genInt < 0 || genInt >= i {
			t.Fatalf("Generated random integer (%d) does not fall in the range [0, %d)!", genInt, i)
		}
	}
}

func TestRandStringBytes(t *testing.T) {
	for i := 0; i < 10; i++ {
		n := rand.Intn(100000)
		randomString := RandStringBytes(n)

		if len(randomString) != n {
			t.Fatalf("String (length %d) not of required length (%d)!", len(randomString), n)
		}
	}
}

func TestRand(t *testing.T) {
	for i := 0; i < 10; i++ {
		min := rand.Intn(1000)
		max := rand.Intn(1000) + min
		randomInt := Rand(min, max)

		if randomInt < min || randomInt > max {
			t.Fatalf("Integer %d is outside specified range (%d - %d)", randomInt, min, max)
		}
	}
}
