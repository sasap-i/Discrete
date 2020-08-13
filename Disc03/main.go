
// ----------------------------------------------------------------------
// Package : main
// Program : main.go
// Item : bufio, os, scanner, strconv, math, error, goto, import
// Date : 2020.5.18
// ----------------------------------------------------------------------

package main

import (
    "bufio"
    "fmt"
    "math"
    "os"
    "strconv"
)

const csName1 = "お名前を入力して下さい : "
const csName2 = "さん、こんにちは"
const csData1 = "最大値を入力して下さい : "
const csData2 = "しきい値を入力して下さい : "
const csData3 = "数値を入力して下さい : "
const ciChapterId = 33

var pkgScanner = bufio.NewScanner(os.Stdin)
var pkgErrMsg = "completed"

func bStdInputName (sName string) bool {
    var name string
    fmt.Printf("%s", sName)
    pkgScanner.Scan()
    name = pkgScanner.Text()
    fmt.Printf("%s%s\n", name, csName2)
    return true
}

func bStdInputData (sMax string, sBound string) bool {
    var err error
    var dStr string
    var iMax, iBound int
    fmt.Printf("%s", sMax)
    
Input1: 
    pkgScanner.Scan()
    dStr = pkgScanner.Text()
    if len(dStr) == 0 {
    	pkgErrMsg = csData3
    	fmt.Printf("%s", csData3)
    	goto Input1
    }
    iMax, err = strconv.Atoi(dStr)
    if err != nil {
    	pkgErrMsg = csData3
    	fmt.Printf("%s", csData3)
    	goto Input1
    }
    fmt.Printf("%s", sBound)
 
Input2:
    pkgScanner.Scan()
    dStr = pkgScanner.Text()
    if len(dStr) == 0 {
    	pkgErrMsg = csData3
    	fmt.Printf("%s", csData3)
    	goto Input2
    }
    iBound, err = strconv.Atoi(dStr)
    if err != nil {
    	pkgErrMsg = csData3
    	fmt.Printf("%s", csData3)
    	goto Input2
    }
    bArgCalculate(iMax, iBound)
    return true
}

func bArgCalculate (iMax int, iBound int) bool {
    var xi int
    var x, y, z float64

    xi = iBound
    for i := 0; i < iMax; i++ {
        if xi <= iBound {
            xi += 1
        } else {
            xi = xi / 2
            if xi <= 1 {
                xi = 1
                break
            }
        }
    }
    x = float64(xi)
    y = x * math.Log(x)
    z = x * x * math.Pi
    fmt.Printf("Result x = %f, y = %f , z = %.6f\n", x, y, z)
    return true
}

func main() {
    var riS, riX, riY int
    // fmt.Printf("Chapter %d Start\n", ciChapterId)
    {
    	riS, riX, riY = iArgRequest()
    	if riS == 0 {
    	    bStdInputName(csName1)
    	} else if riS == 1 {
    	    bStdInputName(csName1)
	    bStdInputData(csData1, csData2)
    	} else if riS == 2 {
    	    bArgCalculate(riX, riY)
    	} else {
    	    fmt.Printf("%s\n", pkgErrMsg)
    	}
    }
    // fmt.Printf("Chapter %d Exit\n", ciChapterId)
}
// ----------------------------------------------------------------------
