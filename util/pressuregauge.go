package util

type PressureGauge struct {
	ch chan struct{}
}

func NewPressureGauge(limit int) *PressureGauge {
	ch := make(chan struct{}, limit)
	for i := 0; i < limit; i++ {
		ch <- struct{}{}
	}
	return &PressureGauge{
		ch: ch,
	}
}
func (pg *PressureGauge) Process(f func()) {
	select {
	case <-pg.ch:
		go f()
		pg.ch <- struct{}{}
	}
}
