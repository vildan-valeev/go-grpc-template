package bff

import (
	stdErr "errors"
	"go-grpc-template/pkg/errors"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GRPCError(err error) error {
	var (
		pgErr *pgconn.PgError
		e     *errors.Error
	)

	switch {
	case err == nil:
		return nil
	case stdErr.As(err, &e):
		log.Error().Err(err).Msg(e.Message)

		return status.Error(GetCode(e.Code), e.Code)
	case stdErr.Is(err, pgx.ErrNoRows):
		return status.Error(codes.NotFound, errors.NotFound)
	case stdErr.As(err, &pgErr):
		log.Error().Err(err).Msg(pgErr.Message)

		return status.Error(codes.Internal, errors.Internal)

	default:
		log.Error().Err(err).Msg("internal")

		return status.Error(codes.Internal, errors.Internal)
	}
}

func GetCode(msg string) codes.Code {
	v, ok := code[msg]
	if !ok {
		return codes.Unknown
	}

	return v
}

//nolint:gochecknoglobals
var code = map[string]codes.Code{
	errors.Canceled:           codes.Canceled,           // 1.
	errors.Unknown:            codes.Unknown,            // 2.
	errors.DeadlineExceeded:   codes.DeadlineExceeded,   // 4.
	errors.NotFound:           codes.NotFound,           // 5.
	errors.AlreadyExists:      codes.AlreadyExists,      // 6.
	errors.PermissionDenied:   codes.PermissionDenied,   // 7.
	errors.ResourceExhausted:  codes.ResourceExhausted,  // 8.
	errors.FailedPrecondition: codes.FailedPrecondition, // 9.
	errors.Aborted:            codes.Aborted,            // 10.
	errors.OutOfRange:         codes.OutOfRange,         // 11.
	errors.Unimplemented:      codes.Unimplemented,      // 12.
	errors.Internal:           codes.Internal,           // 13.
	errors.Unavailable:        codes.Unavailable,        // 14.
	errors.DataLoss:           codes.DataLoss,           // 15.
	errors.Unauthenticated:    codes.Unauthenticated,    // 16.

	errors.InvalidArgument: codes.InvalidArgument,
	//errors.InvalidArgumentPhoneNumbers:                  codes.InvalidArgument,
	//errors.InvalidArgumentOrderID:                       codes.InvalidArgument,
	//errors.InvalidArgumentOrder:                         codes.InvalidArgument,
	//errors.InvalidArgumentPayer:                         codes.InvalidArgument,
	//errors.InvalidArgumentPayDate:                       codes.InvalidArgument,
	//errors.InvalidArgumentPayDateBeforeNow:              codes.InvalidArgument,
	//errors.InvalidArgumentAccount:                       codes.InvalidArgument,
	//errors.InvalidArgumentSumma:                         codes.InvalidArgument,
	//errors.InvalidArgumentSummaLength:                   codes.InvalidArgument,
	//errors.InvalidArgumentDocument:                      codes.InvalidArgument,
	//errors.InvalidArgumentCSC:                           codes.InvalidArgument,
	//errors.InvalidArgumentUnit:                          codes.InvalidArgument,
	//errors.InvalidArgumentFinancialResponsibilityCenter: codes.InvalidArgument,
	//errors.InvalidArgumentBudgetFunds:                   codes.InvalidArgument,
	//errors.InvalidArgumentMatching:                      codes.InvalidArgument,
	//errors.InvalidArgumentRecipient:                     codes.InvalidArgument,
	//errors.InvalidArgumentOperationType:                 codes.InvalidArgument,
	//errors.InvalidArgumentINN:                           codes.InvalidArgument,
	//errors.InvalidArgumentCurrency:                      codes.InvalidArgument,
	//errors.InvalidArgumentFileID:                        codes.InvalidArgument,
	//errors.InvalidArgumentOrganization:                  codes.InvalidArgument,
	//errors.InvalidArgumentDepartment:                    codes.InvalidArgument,
	//errors.InvalidArgumentCostBudget:                    codes.InvalidArgument,
	//errors.InvalidArgumentInvestmentArticle:             codes.InvalidArgument,
	//errors.InvalidArgumentComment:                       codes.InvalidArgument,
	//errors.InvalidArgumentQueryLen:                      codes.InvalidArgument,
	//errors.InvalidArgumentEmail:                         codes.InvalidArgument,
	//errors.NotFoundOrder:                                codes.NotFound,
	//errors.PermissioneEmployeeOnly:                      codes.PermissionDenied,
	//errors.BadAuthorizationString:                       codes.Unauthenticated,
	//errors.ProductionDivisionNotFound:                   codes.NotFound,
	//errors.InvalidArgumentProductionDivisionID:          codes.InvalidArgument,
	//errors.InvalidArgumentFRCID:                         codes.InvalidArgument,
	//errors.InvalidCashFlowBudgetID:                      codes.InvalidArgument,
	//errors.InvalidDateFormat:                            codes.InvalidArgument,
	//errors.InvalidSumFormat:                             codes.InvalidArgument,
	//errors.AccountDetailsRequiredField:                  codes.NotFound,
	//errors.AccountDetailsOnlyOne:                        codes.AlreadyExists,
	//errors.InvalidCollborator:                           codes.NotFound,
	//errors.EmptyCollaboratorsList:                       codes.NotFound,
	//errors.NotFoundAccountFile:                          codes.NotFound,
	//errors.MoreThanOneAccountFile:                       codes.OutOfRange,
	//errors.NotFoundUser:                                 codes.NotFound,
}
