package credit

import "math"

func Calculate(creditSum int64, time int64, percent int64) (payment, overpayment, amount int64) {
	creditSum_v := creditSum * 100 // переводим сумму кредита в копейки
	loanRate := round((float64(percent) / 12 / 100), 4)  // годовая процентная ставка
	const months = 12
	periods := time * months
	elevatedNumber := math.Pow((1.0 + loanRate), float64(periods))
	coeffAnnuity := round((loanRate * elevatedNumber / (elevatedNumber - 1)), 7)
	payment = int64(coeffAnnuity * float64(creditSum_v))
	amount = payment * periods
	overpayment = amount - creditSum_v
	return payment, overpayment, amount
}

func round(number float64, signsCount int) float64 {
	pow := math.Pow(10, float64(signsCount))
	new_number := math.Round(number * pow)
	return new_number / pow
}
