package utils

import (
	"time"
)

func TraceExecutionTime(function func()) time.Duration {
	timer := time.Now()
	function()
	elapsed := time.Since(timer)
	return elapsed
}

func TraceExecutionTimeWithArgs(function func(...interface{}), args ...interface{}) time.Duration {
	timer := time.Now()
	function(args...)
	elapsed := time.Since(timer)
	return elapsed
}