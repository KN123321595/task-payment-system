package iban

import (
	"strings"
	"task-payment-system/pkg/utils"
)

const (
	CountryCode = "BY"
	BankCode    = "TEST"
	BalanceCode = "1111"
)

// TODO генерация формата IBAN с контрольным числом
func GenerateIban() string {
	builder := strings.Builder{}
	builder.WriteString(CountryCode)
	builder.WriteString("00") //контрольное число
	builder.WriteString(BankCode)
	builder.WriteString(BalanceCode)

	individualNumber := utils.GenerateRandomSequence(16)
	builder.WriteString(individualNumber)

	return builder.String()
}
