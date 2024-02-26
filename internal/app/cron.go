package app

import (
	"fmt"
	"time"
)

func startCron() {
	now := time.Now()
	firstTime := time.Date(now.Year(), now.Month(), now.Day(), 16, 05, 0, 0, now.Location())

	if now.After(firstTime) {
		firstTime = firstTime.Add(24 * time.Hour)
	}

	delay := firstTime.Sub(now)
	timer := time.NewTimer(delay)

	task := func() {
		fmt.Println("task is processing")
	}

	go func() {
		for {
			<-timer.C
			task()
			timer.Reset(24 * time.Hour)
		}
	}()

	select {}
}

//type currencyRateCronWorker struct {
//	ratesGetter   controller.CurrencyRateService
//	currencyCodes []string
//	l             kitlog.LoggerInterface
//	timer         *time.Timer
//}
//
//func newCurrencyRateCron(
//	ratesGetter controller.CurrencyRateService,
//	currencyCodes []string,
//	l kitlog.LoggerInterface,
//) *currencyRateCronWorker {
//	return &currencyRateCronWorker{
//		ratesGetter,
//		currencyCodes,
//		l,
//		nil,
//	}
//}
//
//// updateCurrencyRate is used to write currency rate in redis (rate_kzt_to_kgs) to get it after. For example in integration inservice to convert money from KGS to KZT and vice versa
//func (w *currencyRateCronWorker) start(hour, min, sec int) {
//	now := time.Now()
//	loc := time.FixedZone("UTC+6", 6*60*60)
//	firstTime := time.Date(now.Year(), now.Month(), now.Day(), hour, min, sec, 0, loc)
//
//	if now.After(firstTime) {
//		firstTime = firstTime.Add(24 * time.Hour)
//	}
//
//	delay := firstTime.Sub(now)
//
//	w.timer = time.NewTimer(delay)
//	wg := sync.WaitGroup{}
//
//	for {
//		select {
//		case <-w.timer.C:
//			w.processRates(&wg)
//			w.timer.Reset(24 * time.Hour)
//		}
//	}
//}
//
//func (w *currencyRateCronWorker) stop() {
//	w.timer.Stop()
//}
//
//func (w *currencyRateCronWorker) processRates(wg *sync.WaitGroup) {
//	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
//	defer cancel()
//
//	rates, err := w.ratesGetter.FetchRates(ctx)
//	if err != nil {
//		w.l.Error("can not fetch currency rates:", err)
//	}
//
//	for _, code := range w.currencyCodes {
//		code := code
//		wg.Add(1)
//		go func() {
//			defer wg.Done()
//			err := w.updateRate(ctx, code, rates)
//			if err != nil {
//				w.l.Errorf("error on save rate of %s: %v", code, err)
//			}
//		}()
//	}
//	wg.Wait()
//}
//
//func (w *currencyRateCronWorker) updateRate(ctx context.Context, code string, rates national_bank.Items) error {
//	rate, err := rates.GetRate(code)
//	if err != nil {
//		return fmt.Errorf("rates.GetRate: %w", err)
//	}
//	rate += rate * 3 / 100
//	_, errCode := w.ratesGetter.Create(ctx, dto.CurrencyRateHistoryCreateRequest{
//		CurrencyCode: code,
//		Rate:         rate,
//	})
//	if errCode != nil {
//		return fmt.Errorf("save rate of %s, %v", code, errCode)
//	}
//	return nil
//}
