package runner

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/KenMGJ/advent-of-code-2018/internal/device"
	"github.com/KenMGJ/advent-of-code-2018/internal/sets"
)

const DAY_16_REGISTER_COUNT = 4

func (r *Runner) Day16Part1(lines []string) {
	p := parseDay16(lines)

	opMap := make(map[int]*sets.StringSet)

	cnt := 0

	for i := 0; i < len(p.Before); i++ {

		ops := sets.NewStringSet()

		for _, op := range device.Ops {

			r := []int{}
			r = append(r, p.Before[i]...)

			d := device.Device{
				R: r,
			}
			d.ExecOpName(op, p.Test[i].A, p.Test[i].B, p.Test[i].C)

			if d.R[0] == p.After[i][0] && d.R[1] == p.After[i][1] && d.R[2] == p.After[i][2] && d.R[3] == p.After[i][3] {
				ops.Add(op)
			}
		}

		if ops.Size() >= 3 {
			cnt++
		}

		opMap[p.Test[i].Op] = ops
	}

	fmt.Println(cnt)

	opcodeMap := make(map[string]int)
	for len(opMap) > 0 {
		for k, v := range opMap {
			if v.Size() == 1 {
				val := v.Vals()[0]
				opcodeMap[val] = k
				delete(opMap, k)

				for _, v := range opMap {
					v.Remove(val)
				}
			}
		}
	}
	fmt.Println(opcodeMap)
}

func (r *Runner) Day16Part2(lines []string) {
	p := parseDay16(lines)
	d := device.NewDevice(p.Program, DAY_16_REGISTER_COUNT)
	d.Run()
	fmt.Println(d.R[0])
}

type day16Lines struct {
	After   [][]int
	Before  [][]int
	Program []device.OpAndArgs
	Test    []device.OpAndArgs
}

func parseDay16(lines []string) *day16Lines {

	before, after, test, prgm := [][]int{}, [][]int{}, []device.OpAndArgs{}, []device.OpAndArgs{}

	beforeMatcher := regexp.MustCompile(`Before: \[(.*)\]`)
	afterMatcher := regexp.MustCompile(`^After:  \[(.*)\]$`)

	seenBefore := false
	for _, l := range lines {

		if l == "" {
			continue
		}

		matches := beforeMatcher.FindStringSubmatch(l)
		if len(matches) > 0 {

			b := mustConvertStringToIntSlice(matches[1], ", ")
			before = append(before, b)

			seenBefore = true
			continue
		}

		matches = afterMatcher.FindStringSubmatch(l)
		if len(matches) > 0 {

			a := mustConvertStringToIntSlice(matches[1], ", ")
			after = append(after, a)

			seenBefore = false
			continue
		}

		if seenBefore {
			t := mustConvertStringToIntSlice(l, " ")
			test = append(test, device.OpAndArgs{Op: t[0], A: t[1], B: t[2], C: t[3]})
		} else {
			p := mustConvertStringToIntSlice(l, " ")
			prgm = append(prgm, device.OpAndArgs{Op: p[0], A: p[1], B: p[2], C: p[3]})
		}
	}

	return &day16Lines{
		After:   after,
		Before:  before,
		Test:    test,
		Program: prgm,
	}

}

func mustConvertStringToIntSlice(str, sep string) []int {
	ss := strings.Split(str, sep)
	is := []int{}
	for _, s := range ss {
		i, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		is = append(is, i)
	}
	return is
}
