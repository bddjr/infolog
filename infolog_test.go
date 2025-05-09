package infolog_test

import (
	"errors"
	"testing"

	"github.com/bddjr/infolog"
)

func Test(t *testing.T) {
	infolog.INFO("info")
	infolog.INFOf("infof %v", 123)
	infolog.INFOln("infoln")

	infolog.ERROR("error")
	infolog.ERRORf("errorf %v", errors.New("123"))
	infolog.ERRORln("errorln")

	infolog.WARN("warn")
	infolog.WARNf("warnf %v", 123)
	infolog.WARNln("warnln")
}
