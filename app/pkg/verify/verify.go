package verify

import (
	"errors"
	"regexp"
	"time"

	"github.com/google/uuid"
	"golang.org/x/text/currency"
)

var (
	ErrInvalidAccount       = errors.New("invalid account")
	ErrInvalidEmail         = errors.New("invalid email")
	ErrInvalidSumma         = errors.New("invalid summa")
	ErrInvalidUUID          = errors.New("invalid uuid")
	ErrInvalidCurrency      = errors.New("invalid currency")
	ErrInvalidLengthComment = errors.New("invalid length comment")
	ErrPayDateBeforeNow     = errors.New("pay date before now")
)

const (
	EmailValidationRegexp = "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
	AccountLen            = 20
	MinSum                = 1
	MaxSum                = 999_999_999_999_999
)

// Account проверяет счет на длину и символы.
func Account(acc string) error {
	if len(acc) != AccountLen {
		return ErrInvalidAccount
	}

	for _, r := range acc {
		if r < '0' || r > '9' {
			return ErrInvalidAccount
		}
	}

	return nil
}

// Comment проверяет на максимальную длину комментария.
func Comment(comment string) error {
	if len(comment) > 210 {
		return ErrInvalidLengthComment
	}

	return nil
}

// Currency проверяет трехбуквенный код валюты ISO 4217.
// Возвращает ошибку, если cur не правильно сформирован или не является распознанным кодом валюты.
func Currency(cur string) error {
	if _, err := currency.ParseISO(cur); err != nil {
		return ErrInvalidCurrency
	}

	return nil
}

// Email проверяет email на соответствие регулярному выражению.
func Email(email string) error {
	emailValidator := regexp.MustCompile(EmailValidationRegexp)

	if !emailValidator.MatchString(email) {
		return ErrInvalidEmail
	}

	return nil
}

// PayDate проверяет что дата платежа заполнена и не раньше текущей (текущую можно).
func PayDate(paydate int64) error {
	payDate := time.Unix(paydate, 0).UTC()
	now := time.Now().UTC()
	testDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)

	if payDate.Before(testDate) {
		return ErrPayDateBeforeNow
	}

	return nil
}

// UUID проверяет id на uuid.
func UUID(id string) error {
	if _, err := uuid.Parse(id); err != nil || id == "" {
		return ErrInvalidUUID
	}

	return nil
}

func TextCheck(id string) error {
	//if _, err := uuid.Parse(id); err != nil || id == "" {
	//	return ErrInvalidUUID
	//}

	return nil
}
