package op

// Deprecated: Use If instead of it.
// 這種用法回傳的型別就只能any，還要透過斷言才有辦法變成所想要的型別
func Ternary(statement bool, a, b any) any {
	if statement {
		return a
	}
	return b
}

// If 倚靠generic讓回傳的型別可與參數相同
func If[T any](cond bool, trueVal, falseVal T) T {
	if cond {
		return trueVal
	}
	return falseVal
}
