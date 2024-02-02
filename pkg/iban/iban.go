package iban

import (
	"strings"
	"task-payment-system/pkg/utils"
)

const (
	COUNTRY_CODE = "BY"   //2-знаковый международный код страны
	BANK_CODE    = "TEST" //4-знаковый код банка BIC
	BALANCE_CODE = "1111" //4-знаковый номер балансового счета, по классификации счетов согласно утвержденным нормам бухгалтерского учета банков
)

// TODO: генерация формата IBAN с контрольным числом
func GenerateIban() string {
	builder := strings.Builder{}
	builder.WriteString(COUNTRY_CODE)
	builder.WriteString("00") //2-знаковое контрольное число
	builder.WriteString(BANK_CODE)
	builder.WriteString(BALANCE_CODE)

	individualNumber := utils.GenerateRandomNumberSequence(16)
	builder.WriteString(individualNumber)

	return builder.String()
}
