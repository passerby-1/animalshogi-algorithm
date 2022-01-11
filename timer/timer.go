package timer

import "time"

/*
// How to use in main()

timeChan := time.NewTimer(time.Second * 59)
tickChan := time.NewTimer(time.Second * 1)
resetChan := make(chan bool)
resetCompreteChan := make(chan bool)

// を定義して Timer 関数に渡す

// How to reset timer

// リセット動作のトリガーを引く関数に resetChan と resetCompreteChan を渡して以下の通りにする
func somefunction(resetChan chan bool, resetCompreteChan chan bool) {

	go func() { // リセットが完了する度に、resetChan を false へ戻すため
		for {
			select {
			case <-resetCompreteChan:
				resetChan <- false
			default:
			}
		}
	}()

	for {

		time.Sleep(time.Second * 60)
		resetChan <- true // にするとタイマーリセット

	}

}

*/

func Timer(timeChan *time.Timer, tickChan *time.Timer, resetChan chan bool, resetCompreteChan chan bool) {

	count := 0

	for {
		select {
		case <-timeChan.C:
			// 1分 (時間制限オーバーしないために59秒) に 1 回起こしたい動作を書く
			// timeChan を渡してやればここでなくても良い
			resetTimer(timeChan, time.Second*59, resetCompreteChan)

		case <-tickChan.C:
			// 毎秒ごとに起こしたい動作を書く, count を秒数として利用可能
			resetTimer(tickChan, time.Second, resetCompreteChan)
			count++

		case <-resetChan:
			resetTimer(timeChan, time.Second*59, resetCompreteChan)
			resetTimer(tickChan, time.Second, resetCompreteChan)
			count = 0
			resetCompreteChan <- true

		}
	}

}

func resetTimer(timer *time.Timer, d time.Duration, resetCompreteChan chan bool) {

	if !timer.Stop() {
		select {
		case <-timer.C:
		default:
			resetCompreteChan <- false
		}
	}

	timer.Reset(d)
}
