package main

import tm "github.com/buger/goterm"

func Clear() {
	tm.Clear()
	tm.MoveCursor(1, 1)
	tm.Flush()
}
