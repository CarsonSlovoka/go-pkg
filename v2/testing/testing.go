package testing

// MustPanic 有錯誤才是正常，沒有錯誤會引發panic("should panic")
// If you know that you will get a panic, then you can call this function that will do recover.
func MustPanic(testFunc func()) {
	defer func() {
		if err := recover(); err == nil {
			panic("should panic")
		}
	}()
	testFunc()
}

// TestPanic 在測試的時候，我們會預期有panic，但如果想要確定panic的原因是否真的如同我們猜測，就可以使用這個函數幫忙
func TestPanic(testFunc func()) (reason any, isPanic bool) {
	defer func() {
		if err := recover(); err != nil {
			reason = err
			isPanic = true
		}
	}()
	testFunc()
	return nil, false
}
