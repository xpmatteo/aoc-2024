package day9

import (
	"math"
	"slices"
	"strconv"
)

type block int

const emptyBlock = block(-1)
const e = emptyBlock

type disk []block

func parseDisk(input string) disk {
	var result disk
	isFile := true
	nextId := block(0)
	for _, valAscii := range input {
		value := valAscii - '0'
		for i := 0; i < int(value); i++ {
			if isFile {
				result = append(result, nextId)
			} else {
				result = append(result, emptyBlock)
			}
		}
		if isFile {
			nextId++
		}
		isFile = !isFile
	}
	return result
}

func checksum(solution disk) int {
	var result int
	for i, val := range solution {
		if val != emptyBlock {
			result += i * int(val)
		}
	}
	return result
}

func compact(d disk, steps int) {
	if steps == -1 {
		steps = math.MaxInt
	}
	left, right := 0, len(d)-1
	for i := 0; i < steps && left < right; {
		if d[left] != e {
			left++
		} else if d[right] == e {
			right--
		} else {
			d[left], d[right] = d[right], d[left]
			i++
		}
	}
}

func parseSolution(s string) disk {
	var result disk
	for _, val := range s {
		value := block(val - '0')
		if val == '.' {
			value = emptyBlock
		}
		result = append(result, value)
	}
	return result
}

type fileId int

const emptySpace = fileId(-1)

type blockSpan struct {
	fileId fileId // -1 for empty
	len    int
}

func (s blockSpan) IsEmpty() bool {
	return s.fileId == emptySpace
}

func (s blockSpan) IsFile() bool {
	return s.fileId != emptySpace
}

func (s blockSpan) fits(len int) bool {
	return s.IsEmpty() && s.len >= len
}

type disk2 []blockSpan

func parseDisk2(s string) disk2 {
	var result disk2
	id := fileId(0)
	isFile := true
	for _, val := range s {
		value := val - '0'
		if isFile {
			result = append(result, file(id, int(value)))
			id++
		} else {
			result = append(result, empty(int(value)))
		}
		isFile = !isFile
	}
	return result
}

func compact2(d disk2) disk2 {
	right := len(d) - 1
	for ; right >= 0; right-- {
		f := d[right]
		if f.IsFile() {
			fileLen := f.len
			left := 0
			for left < right && !d[left].fits(fileLen) {
				left++
			}
			if left == right {
				// no fit
				continue
			}
			emptyLen := d[left].len
			if emptyLen == fileLen {
				d.swap(left, right)
			}
			if emptyLen > fileLen {
				d = d.split(left, fileLen, emptyLen-fileLen)
				right++
				d.swap(left, right)
				d.compactEmpties()
			}
		}
	}
	return d
}

func empty(len int) blockSpan {
	return blockSpan{
		len:    len,
		fileId: emptySpace,
	}
}

func file(id fileId, len int) blockSpan {
	return blockSpan{
		len:    len,
		fileId: id,
	}
}

func (d disk2) String() string {
	var result string
	for _, span := range d {
		for range span.len {
			if span.IsEmpty() {
				result += "."
			} else {
				result += strconv.Itoa(int(span.fileId))
			}
		}
	}
	return result
}

func (d disk2) toDisk1() disk {
	var result disk
	for _, span := range d {
		for range span.len {
			if span.IsEmpty() {
				result = append(result, emptyBlock)
			} else {
				result = append(result, block(span.fileId))
			}
		}
	}
	return result
}

func (d2 disk2) checksum() int {
	d1 := d2.toDisk1()
	return checksum(d1)
}

func (d disk2) swap(i int, j int) {
	d[i], d[j] = d[j], d[i]
}

func (d disk2) split(index int, lenLeft int, lenRight int) disk2 {
	d[index].len = lenRight
	d = slices.Insert(d, index, empty(lenLeft))
	return d
}

func (d disk2) compactEmpties() {
	for i := 0; i+1 < len(d); i++ {
		if d[i].IsEmpty() && d[i+1].IsEmpty() {
			d[i].len += d[i+1].len
			d[i+1].len = 0
		}
	}
}
