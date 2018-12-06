package util

var (
	mapid    map[string]int
	maplmtch map[string]chan int
)

// GoFn : wrapper of 'go func()' with concurrency limitation
func GoFn(fid string,
	n int,
	f func(done <-chan int, id int, params ...interface{}),
	params ...interface{}) {

	if _, ok := mapid[fid]; !ok {
		mapid = make(map[string]int)
	}
	if _, ok := maplmtch[fid]; !ok {
		maplmtch = MakeMapOfStrChanInt(fid, n)
	}

	/* unblock */
	// select {
	// case maplmtch[fid] <- mapid[fid]:
	// 	go f(maplmtch[fid], mapid[fid], params...)
	// 	mapid[fid]++
	// default:
	// 	mapid[fid] = 0
	// }

	/* block if full */
	if mapid[fid] == n {
		mapid[fid] = 0
	}
	maplmtch[fid] <- mapid[fid]
	go f(maplmtch[fid], mapid[fid], params...)
	mapid[fid]++
}

// GoFnReset : free the func-id's buffer values
func GoFnReset(fid string) {
	if _, ok := maplmtch[fid]; ok {
		delete(maplmtch, fid)
		delete(mapid, fid)
	}
}
