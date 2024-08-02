package tests

import (
	"benchmarks/base"
	"benchmarks/utils"
	"fmt"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

type Castle = base.Castle

type Kingdom = base.Kingdom

type Monster = base.Monster

type Bestiary = base.Bestiary

type MonsterWeakness = base.MonsterWeakness

type User = base.User

var UserModel = base.UserModel

var MonsterModel = base.MonsterModel

var KingdomModel = base.KingdomModel

var BestiaryModel = base.BestiaryModel

func SoTimeout(t *testing.T, f func() bool, timeout ...<-chan time.Time) {
	if len(timeout) == 0 {
		timeout = append(timeout, time.After(10*time.Second))
	}
	for {
		select {
		case <-timeout[0]:
			t.Fatalf("Timeout waiting for assertion to execute")
		default:
			ok := f()
			if ok {
				So(ok, ShouldBeTrue)
				return
			} else {
				time.Sleep(100 * time.Millisecond)
			}
		}
	}
}

func Benchmark(f1 func(), f2 func(), iterationCount ...int) {
	iterations := 10
	if len(iterationCount) > 0 {
		iterations = iterationCount[0]
	}
	t1Total := time.Duration(0)
	for i := 0; i < iterations; i++ {
		t1 := utils.TraceExecutionTime(f1)
		t1Total += t1
	}
	t2Total := time.Duration(0)
	for i := 0; i < iterations; i++ {
		t2 := utils.TraceExecutionTime(f2)
		t2Total += t2
	}
	fmt.Printf("\n      Driver : %fs\n      Elemental : %fs\n", t1Total.Seconds()/float64(iterations), t2Total.Seconds()/float64(iterations))
}
