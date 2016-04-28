package fp

import (
	"time"
)

func are_all_trueN(params []interface{}, predicates ...FuncBoolN) bool {
	trueCnt := 0
	for _, pred := range predicates {
		if pred(params...) {
			trueCnt++
		} else {
			break
		}
	}

	return len(predicates) == trueCnt
}

func are_all_true2(paramA, paramB interface{}, predicates ...FuncBool2) bool {
	trueCnt := 0
	for _, pred := range predicates {
		if pred(paramA, paramB) {
			trueCnt++
		} else {
			break
		}
	}

	return len(predicates) == trueCnt
}

func are_all_true1(paramA interface{}, predicates ...FuncBool1) bool {
	trueCnt := 0
	for _, pred := range predicates {
		if pred(paramA) {
			trueCnt++
		} else {
			break
		}
	}

	return len(predicates) == trueCnt
}

func are_all_tick_true2(paramA, paramB time.Time, predicates ...FuncTickBool) bool {
	trueCnt := 0
	for _, pred := range predicates {
		if pred(paramA, paramB) {
			trueCnt++
		} else {
			break
		}
	}

	return len(predicates) == trueCnt
}
