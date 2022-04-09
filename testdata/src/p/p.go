package p

func goFuncWithRecover() {
	go func() { // want "goroutine should have recover in defer func"
	}()
	go func() {
		defer func() {
			recover()
		}()
	}()

	go func() {
		defer func() {
			_ = recover()
		}()
	}()
	go func() {
		defer func() {
			if r := recover(); r != nil {
			}
		}()
	}()
}
