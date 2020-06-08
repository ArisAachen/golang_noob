package main

import (
	"reflect"
	"sync"
	"unsafe"
)

const (
	memoryLackOut = "memory lack out"
	fileType      = "fileType"
)

//memory queue
type MemoryQueue struct {
	//data operate position
	queryPos  int
	termPos   int
	insertPos int

	//data size
	dataSize int

	//data count
	dataCount int

	//memory
	memory    []byte
	queueLock sync.RWMutex
}

//data tag
type tagDataHead struct {
	dataType   string
	dataLength int
}

func newMemoryQueue() *MemoryQueue {
	newQueue := new(MemoryQueue)
	return newQueue
}

//reset memory queue
func (memoryQueue *MemoryQueue) ResetMemoryQueue() {
	memoryQueue.queueLock.Lock()
	//data size and count reset
	memoryQueue.dataSize = 0
	memoryQueue.dataCount = 0

	//data pos reset
	memoryQueue.insertPos = 0
	memoryQueue.queryPos = 0
	memoryQueue.termPos = 0

	//memory reset
	memoryQueue.memory = make([]byte, 200, 600)

	memoryQueue.queueLock.Unlock()
}

//init queue pos according to config file
func (memoryQueue *MemoryQueue) InitMemoryQueue(queryPos, termPos, insertPos, count int) {
	memoryQueue.queueLock.Lock()


	memoryQueue.queueLock.Unlock()
}

//get queue count
func (memoryQueue *MemoryQueue) GetCount() int {
	memoryQueue.queueLock.RLock()
	count := memoryQueue.dataCount
	memoryQueue.queueLock.RUnlock()
	return count
}

//check if is empty
func (memoryQueue *MemoryQueue) IsEmpty() bool {
	isEmpty := memoryQueue.GetCount() == 0
	return isEmpty
}

//push data
func (memoryQueue *MemoryQueue) Push(data string) {
	if len(data) == 0 {
		return
	}
	memoryQueue.queueLock.Lock()
	//struct data
	dataInsert := tagDataHead{
		dataType:   fileType,
		dataLength: len(data),
	}
	//check capacity,if use out,apply larger memory
	dataSize := int(unsafe.Sizeof(dataInsert))
	memoryQueue.checkCap(dataSize + len(data))

	//transform struct to slice header
	var dataSlice reflect.SliceHeader
	dataSlice.Cap = dataSize
	dataSlice.Len = dataSize
	dataSlice.Data = uintptr(unsafe.Pointer(&dataInsert))
	//transform slice head to byte
	dateByte := *(*[]byte)(unsafe.Pointer(&dataSlice))

	//push data into memory
	memberCopy(memoryQueue.memory, memoryQueue.insertPos, dateByte, 0, len(dateByte))
	//copy data
	if len(data) != 0 {
		memberCopy(memoryQueue.memory, memoryQueue.insertPos+len(dateByte), []byte(data), 0, len(data))
	}

	//refresh message
	memoryQueue.dataCount++
	memoryQueue.insertPos += len(dateByte) + len(data)
	memoryQueue.termPos = max(memoryQueue.insertPos, memoryQueue.termPos)
	memoryQueue.dataSize += len(data) + dataSize

	memoryQueue.queueLock.Unlock()
	return
}

//pop
func (memoryQueue *MemoryQueue) Pop() string {
	//check if is empty
	if memoryQueue.GetCount() == 0 {
		return ""
	}

	//check if query equals terminal
	memoryQueue.queueLock.Lock()
	if memoryQueue.queryPos == memoryQueue.termPos {
		memoryQueue.queryPos = 0
		memoryQueue.termPos = memoryQueue.insertPos
	}
	//get query data
	var dataQuery tagDataHead
	sliceSize := int(unsafe.Sizeof(dataQuery))
	dataByte := memoryQueue.memory[memoryQueue.queryPos : memoryQueue.queryPos+sliceSize]
	//transform query data
	dataSlice := (*reflect.SliceHeader)(unsafe.Pointer(&dataByte))
	dataQuery = *(*tagDataHead)(unsafe.Pointer(dataSlice.Data))

	//get back data
	memoryQueue.dataCount--
	memoryQueue.queryPos += sliceSize
	data := memoryQueue.memory[memoryQueue.queryPos : memoryQueue.queryPos+dataQuery.dataLength]
	memoryQueue.queryPos += dataQuery.dataLength
	memoryQueue.dataSize -= sliceSize + dataQuery.dataLength

	memoryQueue.queueLock.Unlock()
	return string(data)
}

//check cap
func (memoryQueue *MemoryQueue) checkCap(length int) {
	//check if memory can accommodate our data
	if memoryQueue.dataSize+length > cap(memoryQueue.memory) {
		panic(memoryLackOut)
	} else if memoryQueue.insertPos == memoryQueue.termPos && memoryQueue.insertPos+length > cap(memoryQueue.memory) {
		if memoryQueue.queryPos > length {
			memoryQueue.insertPos = 0
		} else {
			panic(memoryLackOut)
		}
	} else if memoryQueue.insertPos < memoryQueue.termPos && memoryQueue.insertPos+length > memoryQueue.queryPos {
		panic(memoryLackOut)
	}

	//if memory cannot accommodate the data,apply large memory
	defer func() {
		//check if be sent error is memory lack out
		if err := recover(); err == memoryLackOut {
			//apply new byte memory
			increase := cap(memoryQueue.memory) + min(cap(memoryQueue.memory)/2, length*10)
			newMemory := make([]byte, cap(memoryQueue.memory)+length, increase)

			//data migration
			interval := memoryQueue.termPos - memoryQueue.queryPos
			if interval > 0 {
				memberCopy(newMemory, 0, memoryQueue.memory, memoryQueue.queryPos, interval)
			}
			if memoryQueue.insertPos > memoryQueue.termPos {
				memberCopy(newMemory, interval, memoryQueue.memory, memoryQueue.termPos, memoryQueue.insertPos-memoryQueue.termPos)
			} else {
				memberCopy(newMemory, interval, memoryQueue.memory, 0, memoryQueue.insertPos)
			}

			//reset queue info
			memoryQueue.memory = newMemory
			memoryQueue.queryPos = 0
			memoryQueue.termPos = len(memoryQueue.memory)
			memoryQueue.insertPos = len(memoryQueue.memory)
			memoryQueue.dataSize = len(memoryQueue.memory)
		}
	}()

	return
}
