package testing

// MustPanic 有錯誤才是正常，沒有錯誤會引發panic("should panic")
func MustPanic(testFunc func()) {
	defer func() {
		if err := recover(); err == nil {
			panic("should panic")
		}
	}()
	testFunc()
}

func TestPanic(testFunc func()) (reason interface{}, isPanic bool) {
	defer func() {
		if err := recover(); err != nil {
			reason = err
			isPanic = true
		}
	}()
	testFunc()
	return nil, false
}
