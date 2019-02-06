package util

var (
	mapid    map[string]int
	maplmtch map[string]chan int
)

// GoFn : wrapper of 'go func()' with concurrency limitation
func GoFn(fid string,
	n int,
	blockIfFull bool,
	f func(done <-chan int, id int, params ...interface{}),
	params ...interface{}) {

	if _, ok := mapid[fid]; !ok {
		mapid = make(map[string]int)
		// fPln(mapid[fid])
	}
	if _, ok := maplmtch[fid]; !ok {
		maplmtch = MakeMapOfStrChanInt(fid, n)
	}

	mapid[fid] = TerOp(mapid[fid] == n, 0, mapid[fid]).(int)

	if blockIfFull {
		maplmtch[fid] <- mapid[fid]
		go f(maplmtch[fid], mapid[fid], params...)
		mapid[fid]++
	} else {
		select {
		case maplmtch[fid] <- mapid[fid]:
			go f(maplmtch[fid], mapid[fid], params...)
			mapid[fid]++
		default:
			return
		}
	}
}

// GoFnReset : free the func-id's buffer values
func GoFnReset(fid string) {
	if _, ok := maplmtch[fid]; ok {
		delete(maplmtch, fid)
		delete(mapid, fid)
	}
}
