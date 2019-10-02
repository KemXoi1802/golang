package iso8583

//BitType
type BitType int

//BitLength
type BitLength int

const (
	AN BitType = iota
	ANS
	BCD
	BINARY
	NOT_SPECIFIC
)

const (
	FIXED BitLength = iota
	LLVAR
	LLLVAR
)

type Field struct {
	FieldType string
	FieldLen  int
	FieldFix  bool
}

type FieldAttr struct {
	FieldType        BitType
	FieldLen         BitLength
	FieldMaxLength   int
	FieldDescription string
}

//Spec
var Spec = map[int]FieldAttr{
	1:  FieldAttr{FieldType: BINARY, FieldLen: FIXED, FieldMaxLength: 8, FieldDescription: "BIT MAP"},
	2:  FieldAttr{FieldType: BCD, FieldLen: LLVAR, FieldMaxLength: 19, FieldDescription: "PAN NUMBER"},
	3:  FieldAttr{FieldType: BCD, FieldLen: FIXED, FieldMaxLength: 6, FieldDescription: "PROCESSING CODE"},
	4:  FieldAttr{FieldType: AN, FieldLen: FIXED, FieldMaxLength: 12, FieldDescription: "TRANSACTION PRIMARY AMOUNT"},
	5:  FieldAttr{FieldType: BCD, FieldLen: FIXED, FieldMaxLength: 12, FieldDescription: "UNKNOWN"},
	6:  FieldAttr{FieldType: BCD, FieldLen: FIXED, FieldMaxLength: 12, FieldDescription: "UNKNOWN"},
	7:  FieldAttr{FieldType: BCD, FieldLen: FIXED, FieldMaxLength: 10, FieldDescription: "TRANSACTION DATE/TIME"},
	8:  FieldAttr{FieldType: BCD, FieldLen: FIXED, FieldMaxLength: 8, FieldDescription: "UNKNOWN"},
	9:  FieldAttr{FieldType: BCD, FieldLen: FIXED, FieldMaxLength: 8, FieldDescription: "UNKNOWN"},
	10: FieldAttr{FieldType: BCD, FieldLen: FIXED, FieldMaxLength: 8, FieldDescription: "UNKNOWN"},
	11: FieldAttr{FieldType: BCD, FieldLen: FIXED, FieldMaxLength: 6, FieldDescription: "TRACE NUMBER"},
	12: FieldAttr{FieldType: BCD, FieldLen: FIXED, FieldMaxLength: 6, FieldDescription: "TRANSACTION TIME(HHMMSS)"},
	13: FieldAttr{FieldType: BCD, FieldLen: FIXED, FieldMaxLength: 4, FieldDescription: "TRANSACTION DATE(MMDD)"},
	14: FieldAttr{FieldType: BCD, FieldLen: FIXED, FieldMaxLength: 4, FieldDescription: "EXPIRATION DATE"},
	15: FieldAttr{FieldType: BCD, FieldLen: FIXED, FieldMaxLength: 4, FieldDescription: "SETTLEMENT DATE"},
	16: FieldAttr{FieldType: BCD, FieldLen: FIXED, FieldMaxLength: 4, FieldDescription: "CURRENCY CONVERSION DATE"},
	17: FieldAttr{FieldType: BCD, FieldLen: FIXED, FieldMaxLength: 4, FieldDescription: "CAPTYRE DATE"},
	18: FieldAttr{FieldType: BCD, FieldLen: FIXED, FieldMaxLength: 4, FieldDescription: "MERCHANT TYPE OR MERCHANT CATEGORY CODE"},
	19: FieldAttr{FieldType: BCD, FieldLen: FIXED, FieldMaxLength: 3, FieldDescription: "ACQURING INSTITUTION"},
	20: FieldAttr{FieldType: BCD, FieldLen: FIXED, FieldMaxLength: 3, FieldDescription: "PAN EXTENDED"},
	21: FieldAttr{FieldType: BCD, FieldLen: FIXED, FieldMaxLength: 3, FieldDescription: "FORWARDING INSTITUTION"},
	22: FieldAttr{FieldType: BCD, FieldLen: FIXED, FieldMaxLength: 3, FieldDescription: "POINT OF SERVICE ENTRY MODE"},
	23: FieldAttr{FieldType: BCD, FieldLen: FIXED, FieldMaxLength: 3, FieldDescription: "PAN SEQUENCE NUMBER"},
	24: FieldAttr{FieldType: BCD, FieldLen: FIXED, FieldMaxLength: 4, FieldDescription: "NETWORK INTERNATIONAL IDENTIFIER (NII)"},
	25: FieldAttr{FieldType: BCD, FieldLen: FIXED, FieldMaxLength: 2, FieldDescription: "POINT OF SERVICE CONDITION CODE"},
	26: FieldAttr{FieldType: BCD, FieldLen: FIXED, FieldMaxLength: 2, FieldDescription: "POINT OF SERVICE CAPTURE CODE"},
	27: FieldAttr{FieldType: BCD, FieldLen: FIXED, FieldMaxLength: 1, FieldDescription: "AUTHORIZING IDENTIFICATION RESPONSE LENGTH"},
	28: FieldAttr{FieldType: BCD, FieldLen: FIXED, FieldMaxLength: 8, FieldDescription: "AMOUNT, TRANSACTION FEE"},
	29: FieldAttr{FieldType: BCD, FieldLen: FIXED, FieldMaxLength: 8, FieldDescription: "AMOUNT, SETTLEMENT FEE"},
	30: FieldAttr{FieldType: BCD, FieldLen: FIXED, FieldMaxLength: 8, FieldDescription: "AMOUNT, TRANSACTION PROCESSING FEE"},
	31: FieldAttr{FieldType: BCD, FieldLen: FIXED, FieldMaxLength: 8, FieldDescription: "AMOUNT, SETTLEMENT PROCESSING FEE"},
	32: FieldAttr{FieldType: BCD, FieldLen: LLVAR, FieldMaxLength: 11, FieldDescription: "ACQUIRING INSTITUTION IDENTIFUCATION CODE"},
	33: FieldAttr{FieldType: BCD, FieldLen: LLVAR, FieldMaxLength: 11, FieldDescription: "FORWARDING INSTITUTION IDENTIFICATION CODE"},
	34: FieldAttr{FieldType: BCD, FieldLen: LLVAR, FieldMaxLength: 28, FieldDescription: "PRIMARY ACCOUNT NUMBER, EXTENDED"},
	35: FieldAttr{FieldType: BCD, FieldLen: LLVAR, FieldMaxLength: 37, FieldDescription: "TRACK 2 DATA"},
	36: FieldAttr{FieldType: BCD, FieldLen: LLLVAR, FieldMaxLength: 104, FieldDescription: "TRACK 3 DATA"},
	37: FieldAttr{FieldType: AN, FieldLen: FIXED, FieldMaxLength: 12, FieldDescription: "RETRIEVAL REFERENCE NUMBER"},
	38: FieldAttr{FieldType: AN, FieldLen: FIXED, FieldMaxLength: 6, FieldDescription: "AUTHORIZATION IDENTIFICATION RESPONSE"},
	39: FieldAttr{FieldType: AN, FieldLen: FIXED, FieldMaxLength: 2, FieldDescription: "RESPONSE CODE"},
	40: FieldAttr{FieldType: AN, FieldLen: FIXED, FieldMaxLength: 3, FieldDescription: "SERVICE RESTRICTION CODE"},
	41: FieldAttr{FieldType: ANS, FieldLen: FIXED, FieldMaxLength: 8, FieldDescription: "CARD ACCEPTOR TERMINAL IDENTIFICATION"},
	42: FieldAttr{FieldType: ANS, FieldLen: FIXED, FieldMaxLength: 15, FieldDescription: "CARD ACCEPTOR IDENTIFICATION CODE"},
	43: FieldAttr{FieldType: ANS, FieldLen: FIXED, FieldMaxLength: 40, FieldDescription: "CARD ACCEPTOR NAME/LOCATION (1-23 STREET ADDRESS, 24-36 CITY, 37-38 STATE, 39-40 COUNTRY)"},
	44: FieldAttr{FieldType: AN, FieldLen: LLVAR, FieldMaxLength: 25, FieldDescription: "ADDITIONAL RESPONSE DATA"},
	45: FieldAttr{FieldType: AN, FieldLen: LLVAR, FieldMaxLength: 76, FieldDescription: "TRACK 1 DATA"},
	46: FieldAttr{FieldType: AN, FieldLen: LLLVAR, FieldMaxLength: 999, FieldDescription: "ADDITIONAL DATA (ISO)"},
	47: FieldAttr{FieldType: AN, FieldLen: LLLVAR, FieldMaxLength: 999, FieldDescription: "ADDITIONAL DATA (NATIONAL)"},
	48: FieldAttr{FieldType: AN, FieldLen: LLLVAR, FieldMaxLength: 999, FieldDescription: "ADDITIONAL DATA (PRIVATE)"},
	49: FieldAttr{FieldType: AN, FieldLen: FIXED, FieldMaxLength: 3, FieldDescription: "CURRENCY CODE, TRANSACTION"},
	50: FieldAttr{FieldType: AN, FieldLen: FIXED, FieldMaxLength: 3, FieldDescription: "CURRENCY CODE, SETTLEMENT"},
	51: FieldAttr{FieldType: AN, FieldLen: FIXED, FieldMaxLength: 3, FieldDescription: "CURRENCY CODE, CARDHOLDER BILLING"},
	52: FieldAttr{FieldType: BINARY, FieldLen: FIXED, FieldMaxLength: 16, FieldDescription: "PERSONAL IDENTIFICATION NUMBER DATA"},
	53: FieldAttr{FieldType: BCD, FieldLen: FIXED, FieldMaxLength: 18, FieldDescription: "SECURITY RELATED CONTROL INFORMATION"},
	54: FieldAttr{FieldType: AN, FieldLen: LLLVAR, FieldMaxLength: 120, FieldDescription: "ADDITIONAL AMOUNTS"},
	55: FieldAttr{FieldType: ANS, FieldLen: LLLVAR, FieldMaxLength: 999, FieldDescription: "ICC DATA – EMV HAVING MULTIPLE TAGS"},
	56: FieldAttr{FieldType: ANS, FieldLen: LLLVAR, FieldMaxLength: 999, FieldDescription: "RESERVED (ISO)"},
	57: FieldAttr{FieldType: ANS, FieldLen: LLLVAR, FieldMaxLength: 999, FieldDescription: "RESERVED (NATIONAL)"},
	58: FieldAttr{FieldType: ANS, FieldLen: LLLVAR, FieldMaxLength: 999, FieldDescription: "RESERVED (NATIONAL)"},
	59: FieldAttr{FieldType: ANS, FieldLen: LLLVAR, FieldMaxLength: 999, FieldDescription: "RESERVED (NATIONAL)"},
	60: FieldAttr{FieldType: AN, FieldLen: LLLVAR, FieldMaxLength: 7, FieldDescription: "RESERVED (NATIONAL) (E.G. SETTLEMENT REQUEST: BATCH NUMBER, ADVICE TRANSACTIONS: ORIGINAL TRANSACTION AMOUNT, BATCH UPLOAD: ORIGINAL MTI PLUS ORIGINAL RRN PLUS ORIGINAL STAN, ETC.)"},
	61: FieldAttr{FieldType: ANS, FieldLen: LLLVAR, FieldMaxLength: 999, FieldDescription: "RESERVED (PRIVATE) (E.G. CVV2/SERVICE CODE   TRANSACTIONS)"},
	62: FieldAttr{FieldType: ANS, FieldLen: LLLVAR, FieldMaxLength: 999, FieldDescription: "RESERVED (PRIVATE) (E.G. TRANSACTIONS: INVOICE NUMBER, KEY EXCHANGE TRANSACTIONS: TPK KEY, ETC.)"},
	63: FieldAttr{FieldType: ANS, FieldLen: LLLVAR, FieldMaxLength: 999, FieldDescription: "RESERVED (PRIVATE)"},
	64: FieldAttr{FieldType: BINARY, FieldLen: FIXED, FieldMaxLength: 16, FieldDescription: "MESSAGE AUTHENTICATION CODE (MAC)"},
	// 65:  FieldAttr{BINARY, FIXED, 16, "EXTENDED BITMAP INDICATOR"},
	// 66:  FieldAttr{BCD, FIXED, 1, "SETTLEMENT CODE"},
	// 67:  FieldAttr{BCD, FIXED, 2, "EXTENDED PAYMENT CODE"},
	// 68:  FieldAttr{BCD, FIXED, 3, "RECEIVING INSTITUTION COUNTRY CODE"},
	// 69:  FieldAttr{BCD, FIXED, 3, "SETTLEMENT INSTITUTION COUNTRY CODE"},
	// 70:  FieldAttr{BCD, FIXED, 3, "NETWORK MANAGEMENT INFORMATION CODE"},
	// 71:  FieldAttr{BCD, FIXED, 4, "MESSAGE NUMBER"},
	// 72:  FieldAttr{ANS, LLLVAR, 999, "LAST MESSAGE'S NUMBER"},
	// 73:  FieldAttr{BCD, FIXED, 6, "ACTION DATE (YYMMDD)"},
	// 74:  FieldAttr{BCD, FIXED, 10, "NUMBER OF CREDITS"},
	// 75:  FieldAttr{BCD, FIXED, 10, "CREDITS, REVERSAL NUMBER"},
	// 76:  FieldAttr{BCD, FIXED, 10, "NUMBER OF DEBITS"},
	// 77:  FieldAttr{BCD, FIXED, 10, "DEBITS, REVERSAL NUMBER"},
	// 78:  FieldAttr{BCD, FIXED, 10, "TRANSFER NUMBER"},
	// 79:  FieldAttr{BCD, FIXED, 10, "TRANSFER, REVERSAL NUMBER"},
	// 80:  FieldAttr{BCD, FIXED, 10, "NUMBER OF INQUIRIES"},
	// 81:  FieldAttr{BCD, FIXED, 10, "NUMBER OF AUTHORIZATIONS"},
	// 82:  FieldAttr{BCD, FIXED, 12, "CREDITS, PROCESSING FEE AMOUNT"},
	// 83:  FieldAttr{BCD, FIXED, 12, "CREDITS, TRANSACTION FEE AMOUNT"},
	// 84:  FieldAttr{BCD, FIXED, 12, "DEBITS, PROCESSING FEE AMOUNT"},
	// 85:  FieldAttr{BCD, FIXED, 12, "DEBITS, TRANSACTION FEE AMOUNT"},
	// 86:  FieldAttr{BCD, FIXED, 15, "TOTAL AMOUNT OF CREDITS"},
	// 87:  FieldAttr{BCD, FIXED, 15, "CREDITS, REVERSAL AMOUNT"},
	// 88:  FieldAttr{BCD, FIXED, 15, "TOTAL AMOUNT OF DEBITS"},
	// 89:  FieldAttr{BCD, FIXED, 15, "DEBITS, REVERSAL AMOUNT"},
	// 90:  FieldAttr{BCD, FIXED, 42, "ORIGINAL DATA ELEMENTS"},
	// 91:  FieldAttr{AN, FIXED, 1, "FILE UPDATE CODE"},
	// 92:  FieldAttr{BCD, FIXED, 2, "FILE SECURITY CODE"},
	// 93:  FieldAttr{BCD, FIXED, 5, "RESPONSE INDICATOR"},
	// 94:  FieldAttr{AN, FIXED, 7, "SERVICE INDICATOR"},
	// 95:  FieldAttr{AN, FIXED, 42, "REPLACEMENT AMOUNTS"},
	// 96:  FieldAttr{AN, FIXED, 8, "MESSAGE SECURITY CODE"},
	// 97:  FieldAttr{BCD, FIXED, 16, "NET SETTLEMENT AMOUNT"},
	// 98:  FieldAttr{ANS, FIXED, 25, "PAYEE"},
	// 99:  FieldAttr{BCD, LLVAR, 11, "SETTLEMENT INSTITUTION IDENTIFICATION CODE"},
	// 100: FieldAttr{BCD, LLVAR, 11, "RECEIVING INSTITUTION IDENTIFICATION CODE"},
	// 101: FieldAttr{ANS, FIXED, 17, "FILE NAME"},
	// 102: FieldAttr{ANS, LLVAR, 28, "ACCOUNT IDENTIFICATION 1"},
	// 103: FieldAttr{ANS, LLVAR, 28, "ACCOUNT IDENTIFICATION 2"},
	// 104: FieldAttr{ANS, LLLVAR, 100, "TRANSACTION DESCRIPTION"},
	// 105: FieldAttr{ANS, LLLVAR, 999, "RESERVED FOR ISO USE"},
	// 106: FieldAttr{ANS, LLLVAR, 999, "RESERVED FOR ISO USE"},
	// 107: FieldAttr{ANS, LLLVAR, 999, "RESERVED FOR ISO USE"},
	// 108: FieldAttr{ANS, LLLVAR, 999, "RESERVED FOR ISO USE"},
	// 109: FieldAttr{ANS, LLLVAR, 999, "RESERVED FOR ISO USE"},
	// 110: FieldAttr{ANS, LLLVAR, 999, "RESERVED FOR ISO USE"},
	// 111: FieldAttr{ANS, LLLVAR, 999, "RESERVED FOR ISO USE"},
	// 112: FieldAttr{ANS, LLLVAR, 999, "RESERVED FOR NATIONAL USE"},
	// 113: FieldAttr{BCD, LLVAR, 11, "RESERVED FOR NATIONAL USE"},
	// 114: FieldAttr{ANS, LLLVAR, 999, "RESERVED FOR NATIONAL USE"},
	// 115: FieldAttr{ANS, LLLVAR, 999, "RESERVED FOR NATIONAL USE"},
	// 116: FieldAttr{ANS, LLLVAR, 999, "RESERVED FOR NATIONAL USE"},
	// 117: FieldAttr{ANS, LLLVAR, 999, "RESERVED FOR NATIONAL USE"},
	// 118: FieldAttr{ANS, LLLVAR, 999, "RESERVED FOR NATIONAL USE"},
	// 119: FieldAttr{ANS, LLLVAR, 999, "RESERVED FOR NATIONAL USE"},
	// 120: FieldAttr{ANS, LLLVAR, 999, "RESERVED FOR PRIVATE USE"},
	// 121: FieldAttr{ANS, LLLVAR, 999, "RESERVED FOR PRIVATE USE"},
	// 122: FieldAttr{ANS, LLLVAR, 999, "RESERVED FOR PRIVATE USE"},
	// 123: FieldAttr{ANS, LLLVAR, 999, "RESERVED FOR PRIVATE USE"},
	// 124: FieldAttr{ANS, LLLVAR, 255, "RESERVED FOR PRIVATE USE"},
	// 125: FieldAttr{ANS, LLVAR, 50, "RESERVED FOR PRIVATE USE"},
	// 126: FieldAttr{ANS, LLVAR, 6, "RESERVED FOR PRIVATE USE"},
	// 127: FieldAttr{ANS, LLLVAR, 999, "RESERVED FOR PRIVATE USE"},
	// 128: FieldAttr{BINARY, FIXED, 16, "MESSAGE AUTHENTICATION CODE"},
}
