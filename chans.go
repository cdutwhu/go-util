package util

/* NEED golang generic ... */

// MakeArrOfChanInt : make a batch of chan int
func MakeArrOfChanInt(n, nBuf int) []chan int {
	chans := make([]chan int, n)
	for i := range chans {
		chans[i] = make(chan int, nBuf)
	}
	return chans
}

// MakeMapOfStrChanInt : make a map of chan int
func MakeMapOfStrChanInt(key string, nBuf int) map[string]chan int {
	chmap := make(map[string]chan int)
	chmap[key] = make(chan int, nBuf)
	return chmap
}

// MakeArrOfChanStr : make a batch of chan string
func MakeArrOfChanStr(n, nBuf int) []chan string {
	chans := make([]chan string, n)
	for i := range chans {
		chans[i] = make(chan string, nBuf)
	}
	return chans
}
