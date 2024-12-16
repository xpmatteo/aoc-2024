package day9

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
