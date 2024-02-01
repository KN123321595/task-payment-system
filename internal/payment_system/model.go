package payment_system

import (
	"encoding/json"
	"fmt"
)

type Account struct {
	Iban     string `json:"iban"`
	Name     string `json:"name"`
	Balance  Money  `json:"balance"`
	IsActive bool   `json:"is_active"`
	Status   string `json:"status"`
}

type TransferRequest struct {
	SenderIban   string `json:"sender_iban"`
	ReceiverIban string `json:"receiver_iban"`
	Money        Money  `json:"money"`
}

type Money struct {
	Kopecks int
}

func (m *Money) MarshalJSON() ([]byte, error) {
	rubles := m.Kopecks / 100
	kopecks := m.Kopecks % 100

	displayValue := fmt.Sprintf("%d.%02d", rubles, kopecks)
	return json.Marshal(displayValue)
}

func (m *Money) UnmarshalJSON(data []byte) error {
	var displayValue string
	if err := json.Unmarshal(data, &displayValue); err != nil {
		return err
	}

	var rubles, kopecks int
	n, err := fmt.Sscanf(displayValue, "%d.%02d", &rubles, &kopecks)
	if err != nil || n != 2 {
		return fmt.Errorf("incorrect string format")
	}

	m.Kopecks = rubles*100 + kopecks
	return nil
}
