package MsgPush

import (
	"Lovers_srv/config"
	"testing"
)

func TestStartWSServer(t *testing.T) {
	config.Init(config.MSG_PUSH_SRV_NAME)
	go StartWSServer(config.WSConfig.WSListenAddr)
	select {

	}
}