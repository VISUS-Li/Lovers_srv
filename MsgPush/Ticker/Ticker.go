package Ticker

import "time"

const RunTickerInterval = time.Second * 1

var RunTicker *time.Ticker
func NewRunTicker(){
	RunTicker = time.NewTicker(RunTickerInterval)
}

func StopTicker(){
	RunTicker.Stop()
}