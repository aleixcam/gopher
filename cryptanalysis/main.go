package main

import (
	. "fmt"
)

const key string = "FRIENDSOFGO"

var alpha = [26]string{
	"A",
	"B",
	"C",
	"D",
	"E",
	"F",
	"G",
	"H",
	"I",
	"J",
	"K",
	"L",
	"M",
	"N",
	"O",
	"P",
	"Q",
	"R",
	"S",
	"T",
	"U",
	"V",
	"W",
	"X",
	"Y",
	"Z",
}

func main() {
	cases := []string{
		"FKZECDDC",
		"AZDE RO ECSUZNKW",
		"RLMVGH S JFMFFEB",
		"UYX WHFCG",
		"XVOYVG S TWOSSUASSJG",
		"F CWRT WAAJ GUT, ZV E TDDOCE TFI, NEE DOOD... OH NJ I TRUACI UT HZDMY ZSF. WKPJC ATNFWGMODX, JBVVNABL LFTD I LVGVSS HOXV, PEIH OCS ZVJZZ JVUKH AOQYFZC NJSWSYH YYM IILD UFROHKQG RPHWWK. RZIQRT WZS GGHYCM, VREWZ XVWJJ UEADYSI ZC XKMEY VWQWKH UCIRF WG HMK SRGQVR'V MZYOAFKM ARDHCS, ZVJ UMEGK KHFX, OS RZQBUWR XVOHV AXNWACS CWYY MRBXYV UUKJI BS QHKHWUM FE MRGLJS UROSVB",
	}

	for index, message := range cases {
		decrypted := decrypt(message)
		encrypted := encrypt(decrypted)

		Println(index)
		Println(decrypted)
		Println(encrypted)
	}
}

func decrypt(encrypted string) (decrypted string) {
	var index int
	for _, msgByte := range []byte(encrypted) {
		if isAlpha(msgByte) {
			decrypted += alpha[(msgByte-keyByte(index)+26)%26]
			index++
		} else {
			decrypted += string([]byte{msgByte})
		}
	}

	return
}

func encrypt(decrypted string) (encrypted string) {
	var index int
	for _, msgByte := range []byte(decrypted) {
		if isAlpha(msgByte) {
			encrypted += alpha[(msgByte+keyByte(index))%26]
			index++
		} else {
			encrypted += string([]byte{msgByte})
		}
	}

	return
}

func keyByte(index int) byte {
	return key[index%len(key)]
}

func isAlpha(byte byte) bool {
	return byte >= 'A' && byte <= 'Z'
}
