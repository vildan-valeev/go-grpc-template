package verify_test

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"gitlab.tn.ru/superapp/order/orders/pkg/verify"
)

func TestAccount(t *testing.T) {
	type args struct {
		acc string
	}
	tests := []struct {
		name      string
		args      args
		assertion require.ErrorAssertionFunc
	}{
		{
			name: "Empty",
			args: args{
				acc: "",
			},
			assertion: require.Error,
		},
		{
			name: "less len",
			args: args{
				acc: "1234567891234567891",
			},
			assertion: require.Error,
		},
		{
			name: "more len",
			args: args{
				acc: "123456712345671234567",
			},
			assertion: require.Error,
		},
		{
			name: "bad char",
			args: args{
				acc: "1acc2",
			},
			assertion: require.Error,
		},
		{
			name: "ok",
			args: args{
				acc: "01234567891234567891",
			},
			assertion: require.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.assertion(t, verify.Account(tt.args.acc))
		})
	}
}

func TestComment(t *testing.T) {
	type args struct {
		comment string
	}
	tests := []struct {
		name      string
		args      args
		assertion require.ErrorAssertionFunc
	}{
		{
			name: "empty",
			args: args{
				comment: "",
			},
			assertion: require.NoError,
		},
		{
			name: "ok",
			args: args{
				comment: "comment",
			},
			assertion: require.NoError,
		},
		{
			name: "more len",
			args: args{
				comment: strings.Repeat("comment", 200),
			},
			assertion: require.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.assertion(t, verify.Comment(tt.args.comment))
		})
	}
}

func TestCurrency(t *testing.T) {
	type args struct {
		cur string
	}
	tests := []struct {
		name      string
		args      args
		assertion require.ErrorAssertionFunc
	}{
		{
			name: "empty",
			args: args{
				cur: "",
			},
			assertion: require.Error,
		},
		{
			name: "bad currency",
			args: args{
				cur: "BADCUR",
			},
			assertion: require.Error,
		},
		{
			name: "ok",
			args: args{
				cur: "RUB",
			},
			assertion: require.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.assertion(t, verify.Currency(tt.args.cur))
		})
	}
}

func TestEmails(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name      string
		args      args
		assertion require.ErrorAssertionFunc
	}{
		{
			name: "nil",
			args: args{
				email: "",
			},
			assertion: require.Error,
		},
		{
			name: "ok",
			args: args{
				email: "my@mail.com",
			},
			assertion: require.NoError,
		},
		{
			name: "bad email",
			args: args{
				email: "email",
			},
			assertion: require.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.assertion(t, verify.Email(tt.args.email))
		})
	}
}

func TestPayDate(t *testing.T) {
	type args struct {
		paydate int64
	}
	tests := []struct {
		name      string
		args      args
		assertion require.ErrorAssertionFunc
	}{
		{
			name: "bad payday",
			args: args{
				paydate: 1616202978,
			},
			assertion: require.Error,
		},
		{
			name: "zero payday",
			args: args{
				paydate: 0,
			},
			assertion: require.Error,
		},
		{
			name: "ok payday",
			args: args{
				paydate: time.Now().UTC().Unix() + 1000,
			},
			assertion: require.NoError,
		},
		{
			name: "ok now payday",
			args: args{
				paydate: time.Now().UTC().Unix(),
			},
			assertion: require.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.assertion(t, verify.PayDate(tt.args.paydate))
		})
	}
}

func TestSumma(t *testing.T) {
	type args struct {
		s uint64
	}
	tests := []struct {
		name      string
		args      args
		assertion require.ErrorAssertionFunc
	}{
		{
			name: "zero",
			args: args{
				s: 0,
			},
			assertion: require.Error,
		},
		{
			name: "1",
			args: args{
				s: 1,
			},
			assertion: require.NoError,
		},
		{
			name: "ok",
			args: args{
				s: 21436543,
			},
			assertion: require.NoError,
		},
		{
			name: "max",
			args: args{
				s: 999_999_999_999_999,
			},
			assertion: require.NoError,
		},
		{
			name: "more size",
			args: args{
				s: 999_999_999_999_999_1,
			},
			assertion: require.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.assertion(t, verify.Summa(tt.args.s))
		})
	}
}

func TestUUID(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name      string
		args      args
		assertion require.ErrorAssertionFunc
	}{
		{
			name: "empty",
			args: args{
				id: "",
			},
			assertion: require.Error,
		},
		{
			name: "bad",
			args: args{
				id: "ewiuhtnt9tcnuc",
			},
			assertion: require.Error,
		},
		{
			name: "ok",
			args: args{
				id: "dab24110-1d74-462e-8b67-57b566fb2f65",
			},
			assertion: require.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.assertion(t, verify.UUID(tt.args.id))
		})
	}
}
