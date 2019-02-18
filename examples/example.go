/*
 * The MIT License
 *
 * Copyright 2018 kinsey40.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 * THE SOFTWARE.
 *
 * File:   example.go
 * Author: kinsey40
 *
 * Created on 13 January 2019, 11:05
 *
 * This file contains lots of examples for how you use the Pbar package in
 * a variety of scenarios. The user is encouraged to copy these examples!
 * Note that this file does not contain functionality present within the Pbar
 * module and is included for user convience.
 *
 */

package main

import (
	"time"

	"github.com/kinsey40/pbar"
)

func iterateUsingArray() {
	x := [5]int{1, 2, 3, 4, 5}
	p, err := pbar.Pbar(x)
	if err != nil {
		panic(err)
	}

	p.SetDescription("Array")
	p.Initialize()
	for range x {
		time.Sleep(time.Second * 1)
		p.Update()
	}
}

func iterateUsingString() {
	x := "abcde"
	p, err := pbar.Pbar(x)
	if err != nil {
		panic(err)
	}

	p.SetDescription("String")
	p.Initialize()
	for range x {
		time.Sleep(time.Second * 1)
		p.Update()
	}
}

func iterateUsingSlice() {
	x := []int{1, 2, 3, 4, 5}
	p, err := pbar.Pbar(x)
	if err != nil {
		panic(err)
	}

	p.SetDescription("Slice")
	p.Initialize()
	for range x[:] {
		time.Sleep(time.Second * 1)
		p.Update()
	}
}

func iterateUsingChannel() {
	size := 5
	x := make(chan int, size)
	for index := 0; index < size; index++ {
		x <- index
	}

	close(x)
	p, err := pbar.Pbar(x)
	if err != nil {
		panic(err)
	}

	p.SetDescription("Channel")
	p.Initialize()
	for range x {
		time.Sleep(time.Second * 1)
		p.Update()
	}
}

func iterateUsingMap() {
	x := map[string]string{"1": "a", "2": "b", "3": "c", "4": "d", "5": "e"}
	p, err := pbar.Pbar(x)
	if err != nil {
		panic(err)
	}

	p.SetDescription("Map")
	p.Initialize()
	for range x {
		time.Sleep(time.Second * 1)
		p.Update()
	}
}

func iterateUsingValues() {
	p, err := pbar.Pbar(10)
	if err != nil {
		panic(err)
	}

	p.SetDescription("Values")
	p.Initialize()
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second * 1)
		p.Update()
	}
}

func main() {
	iterateUsingArray()
	iterateUsingString()
	iterateUsingMap()
	iterateUsingChannel()
	iterateUsingSlice()
	iterateUsingValues()
}
