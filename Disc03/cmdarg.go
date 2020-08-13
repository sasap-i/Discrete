
// ----------------------------------------------------------------------
// Package : main
// Program : cmdarg.go
// Item : package, args
// Date : 2020.5.20
// ----------------------------------------------------------------------

package main

import (
    "os"
    "strconv"
)

const csSwitch1 = "引数を３個指定して下さい : "
const csSwitch2 = "整数を指定して下さい : "
const csSwitch3 = "正しいリクエストを指定して下さい : "
const csSwitch4 = "リクエストに応じた引数を指定して下さい : "

func iArgRequest () (int, int, int) {
	var err error
	var riS, riX, riY int
    if len(os.Args) < 2 {
    	pkgErrMsg = csSwitch1
    	return -1, 0, 0
    }
    riS, err = strconv.Atoi(os.Args[1])
    if err != nil {
    	pkgErrMsg = csSwitch2
    	return -1, 0, 0
    }
    
    if riS < 0 || riS > 2 {
    	pkgErrMsg = csSwitch3
    	return -1, 0, 0
    }	
	if riS == 2 {
		if len(os.Args) == 4 {
			riX, err = strconv.Atoi(os.Args[2])
			if err != nil {
				pkgErrMsg = csSwitch2
				return -1, 0, 0
			}
			riY, err = strconv.Atoi(os.Args[3])
			if err != nil {
				pkgErrMsg = csSwitch2
				return -1, 0, 0
			}
		} else {
			pkgErrMsg = csSwitch4
			return -1, 0, 0
		}
	}
	
    return riS, riX, riY
}

// ----------------------------------------------------------------------
