package infolog_test

import (
	"errors"
	"testing"

	"github.com/bddjr/infolog"
)

func Test(t *testing.T) {
	infolog.INFO("info", "1")
	infolog.INFOf("infof%v", 2)
	infolog.INFOln("infoln", "3")

	infolog.ERROR("error", "1")
	infolog.ERRORf("errorf%v", errors.New("2"))
	infolog.ERRORln("errorln", "3")

	infolog.WARN("warn", "1")
	infolog.WARNf("warnf%v", 2)
	infolog.WARNln("warnln", "3")
}
