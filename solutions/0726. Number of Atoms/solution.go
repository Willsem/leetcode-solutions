import (
	"slices"
	"strconv"
)

func countOfAtoms(formula string) string {
	type state struct {
		elems map[string]int
	}
	var (
		stack = make([]state, 1, 1000)
		sIdx  = 0

		elem  string
		count = 0

		groupEnd bool
	)

	stack[sIdx].elems = make(map[string]int)

	flush := func() {
		if count == 0 {
			count = 1
		}

		if groupEnd {
			for elem, v := range stack[sIdx].elems {
				stack[sIdx-1].elems[elem] += v * count
			}
			stack[sIdx].elems = nil
			groupEnd = false
			sIdx-- // move up the stack
		} else if len(elem) > 0 {
			stack[sIdx].elems[elem] += count
			elem = ""
		}
		count = 0
	}

	for _, ru := range formula {
		switch {
		case ru >= 'A' && ru <= 'Z':
			flush()
			elem = string(ru)

		case ru >= 'a' && ru <= 'z':
			elem += string(ru)

		case ru >= '0' && ru <= '9':
			count = count*10 + int(ru-'0')

		case ru == '(':
			flush()
			sIdx++
			if sIdx == len(stack) {
				stack = append(stack, state{elems: make(map[string]int)})
			} else {
				stack[sIdx].elems = make(map[string]int)
			}

		case ru == ')':
			flush()
			groupEnd = true

		default:
			panic("invalid input")
		}
	}

	flush()

	names := make([]string, 0, len(stack[0].elems))
	for elem := range stack[0].elems {
		names = append(names, elem)
	}
	slices.Sort(names)

	out := make([]byte, 0, 2000)
	for _, elem := range names {
		out = append(out, []byte(elem)...)
		if stack[0].elems[elem] > 1 {
			out = append(out, []byte(strconv.Itoa(stack[0].elems[elem]))...)
		}
	}

	return string(out)
}
