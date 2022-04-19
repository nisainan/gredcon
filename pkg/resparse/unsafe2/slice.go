/*
 * @Author: gitsrc
 * @Date: 2020-07-10 14:30:51
 * @LastEditors: gitsrc
 * @LastEditTime: 2020-07-10 14:50:22
 * @FilePath: /redcon/pkg/resparse/unsafe2/slice.go
 */
// Copyright 2016 CodisLabs. All Rights Reserved.
// Licensed under the MIT (MIT-LICENSE.txt) license.

package unsafe2

import "redcon/pkg/resparse/sync2/atomic2"

type Slice interface {
	Type() string

	Buffer() []byte
	reclaim()

	Slice2(beg, end int) Slice
	Slice3(beg, end, cap int) Slice
	Parent() Slice
}

var maxOffheapBytes atomic2.Int64

func MaxOffheapBytes() int64 {
	return maxOffheapBytes.Int64()
}

func SetMaxOffheapBytes(n int64) {
	maxOffheapBytes.Set(n)
}

const MinOffheapSlice = 1024 * 16

func MakeSlice(n int) Slice {
	if n >= MinOffheapSlice {
		if s := newCGoSlice(n, false); s != nil {
			return s
		}
	}
	return newGoSlice(n)
}

func MakeOffheapSlice(n int) Slice {
	if n >= 0 {
		return newCGoSlice(n, true)
	}
	panic("make slice with negative size")
}

func FreeSlice(s Slice) {
	if s != nil {
		s.reclaim()
	}
}
