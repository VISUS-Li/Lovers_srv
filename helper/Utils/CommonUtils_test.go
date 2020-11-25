package Utils

import (
	"Lovers_srv/config"
	"testing"
	"time"
)


func TestVerifyUUIDFormat(t *testing.T) {
	config.GlobalConfig.RunMode = config.RUNMODE_DEV
	uuid := "a91559d8-0c99-11eb-9650-00163e2ed191"
	ErrorOutputf("[VerifyUUIDFormat] verify uuid:%s %v",uuid, VerifyUUIDFormat(uuid))
	uuid = "12345677901234567asdgfaf,.d"
	ErrorOutputf("[VerifyUUIDFormat] verify uuid:%s %v",uuid, VerifyUUIDFormat(uuid))
}

func TestGetNowDayStartEnd(t *testing.T){
	config.GlobalConfig.RunMode = config.RUNMODE_DEV
	start, end := GetNowDayStartEnd()
	InfoOutputf("start:%s",time.Unix(start, 0).Format("2006-01-02 15:04:05"))
	InfoOutputf("end:%s",time.Unix(end, 0).Format("2006-01-02 15:04:05"))
}