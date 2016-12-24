package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/ttacon/libphonenumber"
)

var PR_DUMP = spew.Dump
var PR_INFO = fmt.Println

var validation_result = map[libphonenumber.ValidationResult]string{
	libphonenumber.IS_POSSIBLE:          "IS_POSSIBLE",
	libphonenumber.INVALID_COUNTRY_CODE: "INVALID_COUNTRY_CODE",
	libphonenumber.TOO_SHORT:            "TOO_SHORT",
	libphonenumber.TOO_LONG:             "TOO_LONG",
}

var phone_type = map[libphonenumber.PhoneNumberType]string{
	libphonenumber.FIXED_LINE:           "FIXED_LINE",
	libphonenumber.MOBILE:               "MOBILE",
	libphonenumber.FIXED_LINE_OR_MOBILE: "FIXED_LINE_OR_MOBILE",
	libphonenumber.TOLL_FREE:            "TOLL_FREE",
	libphonenumber.PREMIUM_RATE:         "PREMIUM_RATE",
	libphonenumber.SHARED_COST:          "SHARED_COST",
	libphonenumber.VOIP:                 "VOIP",
	libphonenumber.PERSONAL_NUMBER:      "PERSONAL_NUMBER",
	libphonenumber.PAGER:                "PAGER",
	libphonenumber.UAN:                  "UAN",
	libphonenumber.VOICEMAIL:            "VOICEMAIL",
	libphonenumber.UNKNOWN:              "UNKNOWN",
}

func main() {
	num, err := libphonenumber.Parse("8401679936867", "VN")
	if err != nil {
		panic(err)
	}

	PR_INFO("****Validation Results:****")
	PR_INFO("Result from IsPossibleNumberWithReason():", validation_result[libphonenumber.IsPossibleNumberWithReason(num)])
	PR_INFO("Result from GetNumberType(): ", phone_type[libphonenumber.GetNumberType(num)])

	PR_INFO("****Formatting Results:****")
	PR_INFO(num.String())
	PR_DUMP(num)
}
