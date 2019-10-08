package iso8583

//BitType define type of message
type BitType int

//BitLength define length type of message
type BitLength int

const (
	//An is a type of a field
	An BitType = iota
	//Ans is a type of a field
	Ans
	//Bcd is a type of a field
	Bcd
	//Binary is a type of a field
	Binary
)

const (
	//Fixed is a length type of field
	Fixed BitLength = iota
	//Llvar is a length type of field
	Llvar
	//Lllvar is a length type of field
	Lllvar
)

//FieldAttr defined attr of a field
type FieldAttr struct {
	FieldType        BitType
	FieldLen         BitLength
	FieldMaxLength   int
	FieldDescription string
}

//Spec define spec for each bAnk is different
var Spec = map[int]FieldAttr{
	1:  FieldAttr{FieldType: Binary, FieldLen: Fixed, FieldMaxLength: 8, FieldDescription: "BIT MAP"},
	2:  FieldAttr{FieldType: Bcd, FieldLen: Llvar, FieldMaxLength: 19, FieldDescription: "PAn NUMBER"},
	3:  FieldAttr{FieldType: Bcd, FieldLen: Fixed, FieldMaxLength: 6, FieldDescription: "PROCESSING CODE"},
	4:  FieldAttr{FieldType: An, FieldLen: Fixed, FieldMaxLength: 12, FieldDescription: "TRAnsACTION PRIMARY AMOUNT"},
	5:  FieldAttr{FieldType: Bcd, FieldLen: Fixed, FieldMaxLength: 12, FieldDescription: "UNKNOWN"},
	6:  FieldAttr{FieldType: Bcd, FieldLen: Fixed, FieldMaxLength: 12, FieldDescription: "UNKNOWN"},
	7:  FieldAttr{FieldType: Bcd, FieldLen: Fixed, FieldMaxLength: 10, FieldDescription: "TRAnsACTION DATE/TIME"},
	8:  FieldAttr{FieldType: Bcd, FieldLen: Fixed, FieldMaxLength: 8, FieldDescription: "UNKNOWN"},
	9:  FieldAttr{FieldType: Bcd, FieldLen: Fixed, FieldMaxLength: 8, FieldDescription: "UNKNOWN"},
	10: FieldAttr{FieldType: Bcd, FieldLen: Fixed, FieldMaxLength: 8, FieldDescription: "UNKNOWN"},
	11: FieldAttr{FieldType: Bcd, FieldLen: Fixed, FieldMaxLength: 6, FieldDescription: "TRACE NUMBER"},
	12: FieldAttr{FieldType: Bcd, FieldLen: Fixed, FieldMaxLength: 6, FieldDescription: "TRAnsACTION TIME(HHMMSS)"},
	13: FieldAttr{FieldType: Bcd, FieldLen: Fixed, FieldMaxLength: 4, FieldDescription: "TRAnsACTION DATE(MMDD)"},
	14: FieldAttr{FieldType: Bcd, FieldLen: Fixed, FieldMaxLength: 4, FieldDescription: "EXPIRATION DATE"},
	15: FieldAttr{FieldType: Bcd, FieldLen: Fixed, FieldMaxLength: 4, FieldDescription: "SETTLEMENT DATE"},
	16: FieldAttr{FieldType: Bcd, FieldLen: Fixed, FieldMaxLength: 4, FieldDescription: "CURRENCY CONVERSION DATE"},
	17: FieldAttr{FieldType: Bcd, FieldLen: Fixed, FieldMaxLength: 4, FieldDescription: "CAPTYRE DATE"},
	18: FieldAttr{FieldType: Bcd, FieldLen: Fixed, FieldMaxLength: 4, FieldDescription: "MERCHAnT TYPE OR MERCHAnT CATEGORY CODE"},
	19: FieldAttr{FieldType: Bcd, FieldLen: Fixed, FieldMaxLength: 3, FieldDescription: "ACQURING INSTITUTION"},
	20: FieldAttr{FieldType: Bcd, FieldLen: Fixed, FieldMaxLength: 3, FieldDescription: "PAn EXTENDED"},
	21: FieldAttr{FieldType: Bcd, FieldLen: Fixed, FieldMaxLength: 3, FieldDescription: "FORWARDING INSTITUTION"},
	22: FieldAttr{FieldType: Bcd, FieldLen: Fixed, FieldMaxLength: 3, FieldDescription: "POINT OF SERVICE ENTRY MODE"},
	23: FieldAttr{FieldType: Bcd, FieldLen: Fixed, FieldMaxLength: 3, FieldDescription: "PAn SEQUENCE NUMBER"},
	24: FieldAttr{FieldType: Bcd, FieldLen: Fixed, FieldMaxLength: 4, FieldDescription: "NETWORK INTERNATIONAL IDENTIFIER (NII)"},
	25: FieldAttr{FieldType: Bcd, FieldLen: Fixed, FieldMaxLength: 2, FieldDescription: "POINT OF SERVICE CONDITION CODE"},
	26: FieldAttr{FieldType: Bcd, FieldLen: Fixed, FieldMaxLength: 2, FieldDescription: "POINT OF SERVICE CAPTURE CODE"},
	27: FieldAttr{FieldType: Bcd, FieldLen: Fixed, FieldMaxLength: 1, FieldDescription: "AUTHORIZING IDENTIFICATION RESPONSE LENGTH"},
	28: FieldAttr{FieldType: Bcd, FieldLen: Fixed, FieldMaxLength: 8, FieldDescription: "AMOUNT, TRAnsACTION FEE"},
	29: FieldAttr{FieldType: Bcd, FieldLen: Fixed, FieldMaxLength: 8, FieldDescription: "AMOUNT, SETTLEMENT FEE"},
	30: FieldAttr{FieldType: Bcd, FieldLen: Fixed, FieldMaxLength: 8, FieldDescription: "AMOUNT, TRAnsACTION PROCESSING FEE"},
	31: FieldAttr{FieldType: Bcd, FieldLen: Fixed, FieldMaxLength: 8, FieldDescription: "AMOUNT, SETTLEMENT PROCESSING FEE"},
	32: FieldAttr{FieldType: Bcd, FieldLen: Llvar, FieldMaxLength: 11, FieldDescription: "ACQUIRING INSTITUTION IDENTIFUCATION CODE"},
	33: FieldAttr{FieldType: Bcd, FieldLen: Llvar, FieldMaxLength: 11, FieldDescription: "FORWARDING INSTITUTION IDENTIFICATION CODE"},
	34: FieldAttr{FieldType: Bcd, FieldLen: Llvar, FieldMaxLength: 28, FieldDescription: "PRIMARY ACCOUNT NUMBER, EXTENDED"},
	35: FieldAttr{FieldType: Bcd, FieldLen: Llvar, FieldMaxLength: 37, FieldDescription: "TRACK 2 DATA"},
	36: FieldAttr{FieldType: Bcd, FieldLen: Lllvar, FieldMaxLength: 104, FieldDescription: "TRACK 3 DATA"},
	37: FieldAttr{FieldType: An, FieldLen: Fixed, FieldMaxLength: 12, FieldDescription: "RETRIEVAL REFERENCE NUMBER"},
	38: FieldAttr{FieldType: An, FieldLen: Fixed, FieldMaxLength: 6, FieldDescription: "AUTHORIZATION IDENTIFICATION RESPONSE"},
	39: FieldAttr{FieldType: An, FieldLen: Fixed, FieldMaxLength: 2, FieldDescription: "RESPONSE CODE"},
	40: FieldAttr{FieldType: An, FieldLen: Fixed, FieldMaxLength: 3, FieldDescription: "SERVICE RESTRICTION CODE"},
	41: FieldAttr{FieldType: Ans, FieldLen: Fixed, FieldMaxLength: 8, FieldDescription: "CARD ACCEPTOR TERMINAL IDENTIFICATION"},
	42: FieldAttr{FieldType: Ans, FieldLen: Fixed, FieldMaxLength: 15, FieldDescription: "CARD ACCEPTOR IDENTIFICATION CODE"},
	43: FieldAttr{FieldType: Ans, FieldLen: Fixed, FieldMaxLength: 40, FieldDescription: "CARD ACCEPTOR NAME/LOCATION (1-23 STREET ADDRESS, 24-36 CITY, 37-38 STATE, 39-40 COUNTRY)"},
	44: FieldAttr{FieldType: An, FieldLen: Llvar, FieldMaxLength: 25, FieldDescription: "ADDITIONAL RESPONSE DATA"},
	45: FieldAttr{FieldType: An, FieldLen: Llvar, FieldMaxLength: 76, FieldDescription: "TRACK 1 DATA"},
	46: FieldAttr{FieldType: An, FieldLen: Lllvar, FieldMaxLength: 999, FieldDescription: "ADDITIONAL DATA (ISO)"},
	47: FieldAttr{FieldType: An, FieldLen: Lllvar, FieldMaxLength: 999, FieldDescription: "ADDITIONAL DATA (NATIONAL)"},
	48: FieldAttr{FieldType: An, FieldLen: Lllvar, FieldMaxLength: 999, FieldDescription: "ADDITIONAL DATA (PRIVATE)"},
	49: FieldAttr{FieldType: An, FieldLen: Fixed, FieldMaxLength: 3, FieldDescription: "CURRENCY CODE, TRAnsACTION"},
	50: FieldAttr{FieldType: An, FieldLen: Fixed, FieldMaxLength: 3, FieldDescription: "CURRENCY CODE, SETTLEMENT"},
	51: FieldAttr{FieldType: An, FieldLen: Fixed, FieldMaxLength: 3, FieldDescription: "CURRENCY CODE, CARDHOLDER BILLING"},
	52: FieldAttr{FieldType: Binary, FieldLen: Fixed, FieldMaxLength: 16, FieldDescription: "PERSONAL IDENTIFICATION NUMBER DATA"},
	53: FieldAttr{FieldType: Bcd, FieldLen: Fixed, FieldMaxLength: 18, FieldDescription: "SECURITY RELATED CONTROL INFORMATION"},
	54: FieldAttr{FieldType: An, FieldLen: Lllvar, FieldMaxLength: 120, FieldDescription: "ADDITIONAL AMOUNTS"},
	55: FieldAttr{FieldType: Ans, FieldLen: Lllvar, FieldMaxLength: 999, FieldDescription: "ICC DATA – EMV HAVING MULTIPLE TAGS"},
	56: FieldAttr{FieldType: Ans, FieldLen: Lllvar, FieldMaxLength: 999, FieldDescription: "RESERVED (ISO)"},
	57: FieldAttr{FieldType: Ans, FieldLen: Lllvar, FieldMaxLength: 999, FieldDescription: "RESERVED (NATIONAL)"},
	58: FieldAttr{FieldType: Ans, FieldLen: Lllvar, FieldMaxLength: 999, FieldDescription: "RESERVED (NATIONAL)"},
	59: FieldAttr{FieldType: Ans, FieldLen: Lllvar, FieldMaxLength: 999, FieldDescription: "RESERVED (NATIONAL)"},
	60: FieldAttr{FieldType: An, FieldLen: Lllvar, FieldMaxLength: 7, FieldDescription: "RESERVED (NATIONAL) (E.G. SETTLEMENT REQUEST: BATCH NUMBER, ADVICE TRAnsACTIONS: ORIGINAL TRAnsACTION AMOUNT, BATCH UPLOAD: ORIGINAL MTI PLUS ORIGINAL RRN PLUS ORIGINAL STAn, ETC.)"},
	61: FieldAttr{FieldType: Ans, FieldLen: Lllvar, FieldMaxLength: 999, FieldDescription: "RESERVED (PRIVATE) (E.G. CVV2/SERVICE CODE   TRAnsACTIONS)"},
	62: FieldAttr{FieldType: Ans, FieldLen: Lllvar, FieldMaxLength: 999, FieldDescription: "RESERVED (PRIVATE) (E.G. TRAnsACTIONS: INVOICE NUMBER, KEY EXCHAnGE TRAnsACTIONS: TPK KEY, ETC.)"},
	63: FieldAttr{FieldType: Ans, FieldLen: Lllvar, FieldMaxLength: 999, FieldDescription: "RESERVED (PRIVATE)"},
	64: FieldAttr{FieldType: Binary, FieldLen: Fixed, FieldMaxLength: 16, FieldDescription: "MESSAGE AUTHENTICATION CODE (MAC)"},
	// 65:  FieldAttr{Binary, Fixed, 16, "EXTENDED BITMAP INDICATOR"},
	// 66:  FieldAttr{Bcd, Fixed, 1, "SETTLEMENT CODE"},
	// 67:  FieldAttr{Bcd, Fixed, 2, "EXTENDED PAYMENT CODE"},
	// 68:  FieldAttr{Bcd, Fixed, 3, "RECEIVING INSTITUTION COUNTRY CODE"},
	// 69:  FieldAttr{Bcd, Fixed, 3, "SETTLEMENT INSTITUTION COUNTRY CODE"},
	// 70:  FieldAttr{Bcd, Fixed, 3, "NETWORK MAnAGEMENT INFORMATION CODE"},
	// 71:  FieldAttr{Bcd, Fixed, 4, "MESSAGE NUMBER"},
	// 72:  FieldAttr{Ans, Lllvar, 999, "LAST MESSAGE'S NUMBER"},
	// 73:  FieldAttr{Bcd, Fixed, 6, "ACTION DATE (YYMMDD)"},
	// 74:  FieldAttr{Bcd, Fixed, 10, "NUMBER OF CREDITS"},
	// 75:  FieldAttr{Bcd, Fixed, 10, "CREDITS, REVERSAL NUMBER"},
	// 76:  FieldAttr{Bcd, Fixed, 10, "NUMBER OF DEBITS"},
	// 77:  FieldAttr{Bcd, Fixed, 10, "DEBITS, REVERSAL NUMBER"},
	// 78:  FieldAttr{Bcd, Fixed, 10, "TRAnsFER NUMBER"},
	// 79:  FieldAttr{Bcd, Fixed, 10, "TRAnsFER, REVERSAL NUMBER"},
	// 80:  FieldAttr{Bcd, Fixed, 10, "NUMBER OF INQUIRIES"},
	// 81:  FieldAttr{Bcd, Fixed, 10, "NUMBER OF AUTHORIZATIONS"},
	// 82:  FieldAttr{Bcd, Fixed, 12, "CREDITS, PROCESSING FEE AMOUNT"},
	// 83:  FieldAttr{Bcd, Fixed, 12, "CREDITS, TRAnsACTION FEE AMOUNT"},
	// 84:  FieldAttr{Bcd, Fixed, 12, "DEBITS, PROCESSING FEE AMOUNT"},
	// 85:  FieldAttr{Bcd, Fixed, 12, "DEBITS, TRAnsACTION FEE AMOUNT"},
	// 86:  FieldAttr{Bcd, Fixed, 15, "TOTAL AMOUNT OF CREDITS"},
	// 87:  FieldAttr{Bcd, Fixed, 15, "CREDITS, REVERSAL AMOUNT"},
	// 88:  FieldAttr{Bcd, Fixed, 15, "TOTAL AMOUNT OF DEBITS"},
	// 89:  FieldAttr{Bcd, Fixed, 15, "DEBITS, REVERSAL AMOUNT"},
	// 90:  FieldAttr{Bcd, Fixed, 42, "ORIGINAL DATA ELEMENTS"},
	// 91:  FieldAttr{An, Fixed, 1, "FILE UPDATE CODE"},
	// 92:  FieldAttr{Bcd, Fixed, 2, "FILE SECURITY CODE"},
	// 93:  FieldAttr{Bcd, Fixed, 5, "RESPONSE INDICATOR"},
	// 94:  FieldAttr{An, Fixed, 7, "SERVICE INDICATOR"},
	// 95:  FieldAttr{An, Fixed, 42, "REPLACEMENT AMOUNTS"},
	// 96:  FieldAttr{An, Fixed, 8, "MESSAGE SECURITY CODE"},
	// 97:  FieldAttr{Bcd, Fixed, 16, "NET SETTLEMENT AMOUNT"},
	// 98:  FieldAttr{Ans, Fixed, 25, "PAYEE"},
	// 99:  FieldAttr{Bcd, Llvar, 11, "SETTLEMENT INSTITUTION IDENTIFICATION CODE"},
	// 100: FieldAttr{Bcd, Llvar, 11, "RECEIVING INSTITUTION IDENTIFICATION CODE"},
	// 101: FieldAttr{Ans, Fixed, 17, "FILE NAME"},
	// 102: FieldAttr{Ans, Llvar, 28, "ACCOUNT IDENTIFICATION 1"},
	// 103: FieldAttr{Ans, Llvar, 28, "ACCOUNT IDENTIFICATION 2"},
	// 104: FieldAttr{Ans, Lllvar, 100, "TRAnsACTION DESCRIPTION"},
	// 105: FieldAttr{Ans, Lllvar, 999, "RESERVED FOR ISO USE"},
	// 106: FieldAttr{Ans, Lllvar, 999, "RESERVED FOR ISO USE"},
	// 107: FieldAttr{Ans, Lllvar, 999, "RESERVED FOR ISO USE"},
	// 108: FieldAttr{Ans, Lllvar, 999, "RESERVED FOR ISO USE"},
	// 109: FieldAttr{Ans, Lllvar, 999, "RESERVED FOR ISO USE"},
	// 110: FieldAttr{Ans, Lllvar, 999, "RESERVED FOR ISO USE"},
	// 111: FieldAttr{Ans, Lllvar, 999, "RESERVED FOR ISO USE"},
	// 112: FieldAttr{Ans, Lllvar, 999, "RESERVED FOR NATIONAL USE"},
	// 113: FieldAttr{Bcd, Llvar, 11, "RESERVED FOR NATIONAL USE"},
	// 114: FieldAttr{Ans, Lllvar, 999, "RESERVED FOR NATIONAL USE"},
	// 115: FieldAttr{Ans, Lllvar, 999, "RESERVED FOR NATIONAL USE"},
	// 116: FieldAttr{Ans, Lllvar, 999, "RESERVED FOR NATIONAL USE"},
	// 117: FieldAttr{Ans, Lllvar, 999, "RESERVED FOR NATIONAL USE"},
	// 118: FieldAttr{Ans, Lllvar, 999, "RESERVED FOR NATIONAL USE"},
	// 119: FieldAttr{Ans, Lllvar, 999, "RESERVED FOR NATIONAL USE"},
	// 120: FieldAttr{Ans, Lllvar, 999, "RESERVED FOR PRIVATE USE"},
	// 121: FieldAttr{Ans, Lllvar, 999, "RESERVED FOR PRIVATE USE"},
	// 122: FieldAttr{Ans, Lllvar, 999, "RESERVED FOR PRIVATE USE"},
	// 123: FieldAttr{Ans, Lllvar, 999, "RESERVED FOR PRIVATE USE"},
	// 124: FieldAttr{Ans, Lllvar, 255, "RESERVED FOR PRIVATE USE"},
	// 125: FieldAttr{Ans, Llvar, 50, "RESERVED FOR PRIVATE USE"},
	// 126: FieldAttr{Ans, Llvar, 6, "RESERVED FOR PRIVATE USE"},
	// 127: FieldAttr{Ans, Lllvar, 999, "RESERVED FOR PRIVATE USE"},
	// 128: FieldAttr{Binary, Fixed, 16, "MESSAGE AUTHENTICATION CODE"},
}
