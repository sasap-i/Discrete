
// ----------------------------------------------------------------------
// Package     : main
// Program     : main.go
// Version     : 1.0.7
// Execution   : main < input.txt > output.txt
// Annotation  : recursive call
//             : parallel execution
//             : medium buffer I/O
//             : continuity oriented
// Date        : 2020.7.17
// @Author     : Hidekazu Sasaki
// ----------------------------------------------------------------------

package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
    "strings"
    "sync"
)

type UfRow struct {
    colMax int
    colSum float64
}

const csData1 = "データを読み込めません 。"
const csData2 = "データ型の例外が発生しました。"
const csData3 = "データ値の例外が発生しました。"
const csData4 = "データを書き込めません 。"
const csDiscVersion = "1.0.7"
const ciReadMaxSize = 10000
const ciWriteMaxSize = 10000

var pkgReader = bufio.NewReaderSize(os.Stdin, ciReadMaxSize)
var pkgWriter = bufio.NewWriterSize(os.Stdout, ciWriteMaxSize)
var pkgRowWg sync.WaitGroup
var pkgColWg sync.WaitGroup
var pkgRowMu sync.Mutex
var pkgColMu sync.Mutex
var pkgErrMsg string

var ufRowMax int
var ufRow []UfRow
var ufBrl []byte
var ufBuf []byte
var ufCol []string
var ufStr string
var ufWrl int
var ufErr = false

func iStdInputRowInit () bool {
    var err error
    var bLine []byte
    var n int
    
    bLine, _, err = pkgReader.ReadLine()
    if err != nil {
	    pkgErrMsg = csData1
        log.Print(err)
        return true
    }
    if len(bLine) == 0 {
        pkgErrMsg = csData1
        return true
    }
    n, err = strconv.Atoi(string(bLine))
    if err != nil {
        pkgErrMsg = csData2
        log.Print(err)
    	return true
    }
    if n < 1 {
        pkgErrMsg = csData3
        return true
    }
    
    ufRow = make([]UfRow, n)
    ufRowMax = n
    return false
}

func bStdInputColInit (iRow int) bool {
    var err error
    var bLine []byte
    var n int
    ufRow[iRow].colSum = 0
    ufRow[iRow].colMax = 0
    
    bLine, _, err = pkgReader.ReadLine()
    if err != nil {
        pkgErrMsg = csData1
        log.Print(err)
        return true
    }
    if len(bLine) == 0 {
        pkgErrMsg = csData1
        return true
    }
    n, err = strconv.Atoi(string(bLine))
    if err != nil {
        pkgErrMsg = csData2
        log.Print(err)
        return true
    }
    if n < 1 {
        pkgErrMsg = csData3
        return true
    }
    
    ufRow[iRow].colMax = n
    ufCol = make([]string, n)
    return false
}

func vStdInputSum (iRow int, iCol int, splStr string) {
    var err error
    var iNext int
    var fx float64
    defer pkgColMu.Unlock()
    
    pkgColMu.Lock()
    fx, err = strconv.ParseFloat(splStr, 64)
    if err != nil {
        pkgErrMsg = csData2
        log.Print(err)
    } else if fx > 0 {
        ufRow[iRow].colSum += fx * fx
    }
    
    iNext = iCol + 1
    if iNext < ufRow[iRow].colMax {
        go vStdInputSum(iRow, iNext, ufCol[iNext])
    } else {
        pkgColWg.Done()
    }
}

func bStdInputData (iRow int) bool {
    var err error
    var isContinue bool
    ufBuf = make([]byte, 0)
    
ReadLineLabel:
    ufBrl = make([]byte, ciReadMaxSize)
    ufBrl, isContinue, err = pkgReader.ReadLine()
    if err != nil {
        pkgErrMsg = csData1
        log.Print(err)
        return true
    }
    
    ufBuf = append(ufBuf, ufBrl...)
    if isContinue {
        goto ReadLineLabel
    }
    if len(ufBuf) == 0 {
        pkgErrMsg = csData1
        return true
    }
    if ufRow[iRow].colMax == 0 {
        return true
    }
    
    ufCol = strings.Split(string(ufBuf), " ")
    pkgColWg.Add(1)
    go vStdInputSum(iRow, 0, ufCol[0])
    pkgColWg.Wait()
    return false
}

func vInputRecur (iRow int) {
    var iNext int
    defer pkgRowMu.Unlock()
    
    pkgRowMu.Lock()
    ufErr = bStdInputColInit(iRow)
    ufErr = bStdInputData(iRow)
    
    iNext = iRow + 1
    if iNext < ufRowMax {
        go vInputRecur(iNext)
    } else {
        pkgRowWg.Done()
    }
}

func vStdOutputCol (iRow int) {
    var iNext int
    defer pkgColMu.Unlock()
    
    pkgColMu.Lock()
    ufCol[iRow] = strconv.FormatFloat(ufRow[iRow].colSum, 'f', 0, 64)
    
    iNext = iRow + 1
    if iNext < ufRowMax {
        go vStdOutputCol(iNext)
    } else {
        pkgColWg.Done()
    }
}

func vStdOutputRowInit () {
    ufCol = make([]string, ufRowMax)
    pkgColWg.Add(1)
    go vStdOutputCol(0)
    pkgColWg.Wait()
    
    ufStr = strings.Join(ufCol, "\n")
    ufWrl = len(ufStr)
}

func bStdOutputData (iBegin int, iEnd int) bool {
    var err error
    defer pkgWriter.Flush()
    
    _, err = pkgWriter.WriteString(ufStr[iBegin:iEnd])
    if err != nil {
        pkgErrMsg = csData4
        log.Print(err)
        return true
    }
    return false
}

func vOutputRecur (iBegin int) {
    var iNext int
    defer pkgRowMu.Unlock()
    
    pkgRowMu.Lock()
    iNext = ufWrl
    if iBegin + ciWriteMaxSize < ufWrl {
        iNext = iBegin + ciWriteMaxSize
    }
    ufErr = bStdOutputData(iBegin, iNext)
    
    if iNext < ufWrl {
        go vOutputRecur(iNext)
    } else {
        pkgWriter.Flush()
        pkgRowWg.Done()
    }
}

func main () { 
    fmt.Printf("----- Discrete02  Start Version : %s\n", csDiscVersion)
    ufErr = iStdInputRowInit()
    if ufErr { 
        fmt.Fprintf(os.Stderr, "Discrete02  Error : %s\n", pkgErrMsg)
        os.Exit(1)
    }
    pkgRowWg.Add(1)
    go vInputRecur(0)
    pkgRowWg.Wait()
    
    vStdOutputRowInit()
    pkgRowWg.Add(1)
    go vOutputRecur(0)
    pkgRowWg.Wait()
    
    if ufErr {
        fmt.Fprintf(os.Stderr, "Discrete02 Error : %s\n", pkgErrMsg)
        os.Exit(1)
    }
    fmt.Printf("\n----- Discrete02 Completed Length : %d\n", ufWrl)
}

// ----------------------------------------------------------------------
//     Copyright 2020, Hidekazu Sasaki
// ----------------------------------------------------------------------
