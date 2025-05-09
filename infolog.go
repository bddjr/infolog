package infolog

import (
	"fmt"
	"log"
)

// ln
func Logln(prefix string, v ...any) {
	log.Println("[" + prefix + "] " + fmt.Sprint(v...))
}

func INFOln(v ...any) {
	Logln("INFO", v...)
}

func WARNln(v ...any) {
	Logln("WARN", v...)
}

func ERRORln(v ...any) {
	Logln("ERROR", v...)
}

// f
func Logf(prefix string, format string, v ...any) {
	log.Printf("["+prefix+"] "+format, v...)
}

func INFOf(format string, v ...any) {
	Logf("INFO", format, v...)
}

func WARNf(format string, v ...any) {
	Logf("WARN", format, v...)
}

func ERRORf(format string, v ...any) {
	Logf("ERROR", format, v...)
}

// .
func Log(prefix string, v ...any) {
	log.Print("[" + prefix + "] " + fmt.Sprint(v...))
}

func INFO(v ...any) {
	Log("INFO", v...)
}

func WARN(v ...any) {
	Log("WARN", v...)
}

func ERROR(v ...any) {
	Log("ERROR", v...)
}
