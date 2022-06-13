package errors

import (
	"errors"
	"fmt"
)

const (
	Canceled                                     = "canceled"
	Unknown                                      = "unknown_error"
	DeadlineExceeded                             = "deadline_exceeded"
	NotFound                                     = "not_found"
	AlreadyExists                                = "already_exists"
	PermissionDenied                             = "permission_denied"
	ResourceExhausted                            = "resource_exhausted"
	FailedPrecondition                           = "failed_precondition"
	Aborted                                      = "aborted"
	OutOfRange                                   = "out_of_range"
	Unimplemented                                = "unimplemented"
	Internal                                     = "internal"
	Unavailable                                  = "unavailable"
	DataLoss                                     = "data_loss"
	Unauthenticated                              = "unauthenticated"
	InvalidArgument                              = "invalid_argument"
	InvalidStatus                                = "invalid_status"
	InvalidArgumentPhoneNumbers                  = "invalid_phone"
	InvalidArgumentOrderID                       = "invalid_order_id"
	InvalidArgumentOrder                         = "invalid_order"
	InvalidArgumentPayer                         = "invalid_payer"
	InvalidArgumentPayDate                       = "invalid_pay_date"
	InvalidArgumentPayDateBeforeNow              = "invalid_pay_date_before_now"
	InvalidArgumentAccount                       = "invalid_account"
	InvalidArgumentSumma                         = "invalid_summa"
	InvalidArgumentSummaLength                   = "invalid_summa_length"
	InvalidArgumentDocument                      = "invalid_document"
	InvalidArgumentCSC                           = "invalid_csc"
	InvalidArgumentUnit                          = "invalid_unit"
	InvalidArgumentFinancialResponsibilityCenter = "invalid_financial_responsibility_center"
	InvalidArgumentBudgetFunds                   = "invalid_budget_funds"
	InvalidArgumentMatching                      = "invalid_matching"
	InvalidArgumentRecipient                     = "invalid_recipient"
	InvalidArgumentOperationType                 = "invalid_operation_type"
	InvalidArgumentINN                           = "invalid_inn"
	InvalidArgumentCurrency                      = "invalid_currency"
	InvalidArgumentFileID                        = "invalid_file_id"
	InvalidArgumentOrganization                  = "invalid_organization"
	InvalidArgumentDepartment                    = "invalid_department"
	InvalidArgumentCostBudget                    = "invalid_cost_budget"
	InvalidArgumentInvestmentArticle             = "invalid_investment_article"
	InvalidArgumentComment                       = "invalid_comment"
	InvalidArgumentQueryLen                      = "invalid_query_len"
	InvalidArgumentEmail                         = "invalid_email"
	InvalidArgumentDirectoryType                 = "invalid_directory_type"
	InvalidArgumentDirectoryID                   = "invalid_directory_id"
	NotFoundOrder                                = "not_found_order"
	PermissioneEmployeeOnly                      = "permission_denied_employee_only"
	BadAuthorizationString                       = "bad_authorization_string"
	ProductionDivisionNotFound                   = "production_division_not_found"
	InvalidArgumentProductionDivisionID          = "invalid_product_division_id"
	InvalidArgumentFRCID                         = "invalid_frc_id"
	InvalidCashFlowBudgetID                      = "invalid_cash_flow_budget_id"
	InvalidDateFormat                            = "invalid_date_format"
	InvalidSumFormat                             = "invalid_sum_format"
	AccountDetailsRequiredField                  = "account_details_required_field"
	AccountDetailsOnlyOne                        = "account_details_must_be_in_a_single_copy"
	InvalidCollborator                           = "invalid_collaborator"
	EmptyCollaboratorsList                       = "empty_collaborators_list"
	NotFoundAccountFile                          = "not_found_account_file"
	MoreThanOneAccountFile                       = "more_than_one_account_file"
	NotFoundUser                                 = "not_found_user"
)

// Error represents an application-specific error. Application errors can be
// unwrapped by the caller to extract out the code & message.
type Error struct {
	// Machine-readable error code.
	Code string

	// Human-readable error message.
	Message string
}

// Error implements the error interface.
func (e *Error) Error() string {
	return fmt.Sprintf("error: code=%s message=%s", e.Code, e.Message)
}

// ErrorCode unwraps an application error and returns its code.
// Non-application errors always return Internal.
func ErrorCode(err error) string {
	var e *Error

	if err == nil {
		return ""
	} else if errors.As(err, &e) {
		return e.Code
	}

	return Internal
}

// ErrorMessage unwraps an application error and returns its message.
// Non-application errors always return "internal".
func ErrorMessage(err error) string {
	var e *Error

	if err == nil {
		return ""
	} else if errors.As(err, &e) {
		return e.Message
	}

	return Internal
}

// Errorf is a helper function to return an Error with a given code and formatted message.
func Errorf(code string, format string, args ...interface{}) *Error {
	return &Error{
		Code:    code,
		Message: fmt.Sprintf(format, args...),
	}
}
