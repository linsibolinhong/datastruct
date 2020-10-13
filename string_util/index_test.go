package string_util

import (
	"fmt"
	"math/rand"
	"reflect"
	"runtime"
	"strings"
	"testing"
	"time"
)

type IndexFunc func(s, p string) int

func generatePattern(slen, plen int) []string {
	dict := ""

	for i := 0; i < 2; i++ {
		dict += string('0'+i)
	}

	s := ""
	p := ""
	for i := 0; i < slen; i++ {
		s += string(dict[rand.Int() % len(dict)])
	}

	ps := rand.Int() % slen
	plen = rand.Int() % plen
	if ps + plen < slen {
		p = s[ps:ps+plen]
	} else {
		p = s[ps:]
	}

	if rand.Int() % 2 == 0 && false {
		count := 0
		for count < 100 {
			p = ""
			for i := 0; i < plen; i++ {
				p += string(dict[rand.Int()%len(dict)])
			}
			if strings.Index(s, p) >= 0 {
				break
			}
			count++
		}
	}

	return []string{s, p}
}

func Test_Index(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	funcs := []IndexFunc{
		SundayIndex,
		RabinKarpIndex,
		KMPIndex,
		BMIndex,
	}

	pts := [][]string{
		{
			"001100100110000001000000100",
			"00110000001000000100",
		},
	}

	for i := 0; i <100; i++ {
		pts = append(pts, generatePattern(1000*10,100))
	}

	for idx, pt := range pts {
		s := pt[0]
		p := pt[1]
		st := time.Now()
		ret := strings.Index(s, p)
		cost := time.Since(st)
		fmt.Printf("idx:%v, ret:%v, cost:%v， patten:%v, \n", idx, ret, cost, p)
		for _, fuc := range funcs {
			fucName := runtime.FuncForPC(reflect.ValueOf(fuc).Pointer()).Name()
			fucl := strings.Split(fucName, ".")
			fucName = fucl[len(fucl)-1]
			st := time.Now()
			k := fuc(s, p)
			if k != ret {
				fmt.Printf("falied, fucname:%v, %v:%v, s:%v\n", fucName, ret, k, s)
				panic(fucName + " errorr.....")
			}
			cost := time.Since(st)
			fmt.Printf("idx:%v, fuc:%v, ret:%v, cost:%v \n", idx, fucName, k, cost)
		}
	}

}
