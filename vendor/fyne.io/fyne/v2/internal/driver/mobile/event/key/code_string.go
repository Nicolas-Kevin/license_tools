// Code generated by "stringer -type=Code"; DO NOT EDIT.

package key

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[CodeUnknown-0]
	_ = x[CodeA-4]
	_ = x[CodeB-5]
	_ = x[CodeC-6]
	_ = x[CodeD-7]
	_ = x[CodeE-8]
	_ = x[CodeF-9]
	_ = x[CodeG-10]
	_ = x[CodeH-11]
	_ = x[CodeI-12]
	_ = x[CodeJ-13]
	_ = x[CodeK-14]
	_ = x[CodeL-15]
	_ = x[CodeM-16]
	_ = x[CodeN-17]
	_ = x[CodeO-18]
	_ = x[CodeP-19]
	_ = x[CodeQ-20]
	_ = x[CodeR-21]
	_ = x[CodeS-22]
	_ = x[CodeT-23]
	_ = x[CodeU-24]
	_ = x[CodeV-25]
	_ = x[CodeW-26]
	_ = x[CodeX-27]
	_ = x[CodeY-28]
	_ = x[CodeZ-29]
	_ = x[Code1-30]
	_ = x[Code2-31]
	_ = x[Code3-32]
	_ = x[Code4-33]
	_ = x[Code5-34]
	_ = x[Code6-35]
	_ = x[Code7-36]
	_ = x[Code8-37]
	_ = x[Code9-38]
	_ = x[Code0-39]
	_ = x[CodeReturnEnter-40]
	_ = x[CodeEscape-41]
	_ = x[CodeDeleteBackspace-42]
	_ = x[CodeTab-43]
	_ = x[CodeSpacebar-44]
	_ = x[CodeHyphenMinus-45]
	_ = x[CodeEqualSign-46]
	_ = x[CodeLeftSquareBracket-47]
	_ = x[CodeRightSquareBracket-48]
	_ = x[CodeBackslash-49]
	_ = x[CodeSemicolon-51]
	_ = x[CodeApostrophe-52]
	_ = x[CodeGraveAccent-53]
	_ = x[CodeComma-54]
	_ = x[CodeFullStop-55]
	_ = x[CodeSlash-56]
	_ = x[CodeCapsLock-57]
	_ = x[CodeF1-58]
	_ = x[CodeF2-59]
	_ = x[CodeF3-60]
	_ = x[CodeF4-61]
	_ = x[CodeF5-62]
	_ = x[CodeF6-63]
	_ = x[CodeF7-64]
	_ = x[CodeF8-65]
	_ = x[CodeF9-66]
	_ = x[CodeF10-67]
	_ = x[CodeF11-68]
	_ = x[CodeF12-69]
	_ = x[CodePause-72]
	_ = x[CodeInsert-73]
	_ = x[CodeHome-74]
	_ = x[CodePageUp-75]
	_ = x[CodeDeleteForward-76]
	_ = x[CodeEnd-77]
	_ = x[CodePageDown-78]
	_ = x[CodeRightArrow-79]
	_ = x[CodeLeftArrow-80]
	_ = x[CodeDownArrow-81]
	_ = x[CodeUpArrow-82]
	_ = x[CodeKeypadNumLock-83]
	_ = x[CodeKeypadSlash-84]
	_ = x[CodeKeypadAsterisk-85]
	_ = x[CodeKeypadHyphenMinus-86]
	_ = x[CodeKeypadPlusSign-87]
	_ = x[CodeKeypadEnter-88]
	_ = x[CodeKeypad1-89]
	_ = x[CodeKeypad2-90]
	_ = x[CodeKeypad3-91]
	_ = x[CodeKeypad4-92]
	_ = x[CodeKeypad5-93]
	_ = x[CodeKeypad6-94]
	_ = x[CodeKeypad7-95]
	_ = x[CodeKeypad8-96]
	_ = x[CodeKeypad9-97]
	_ = x[CodeKeypad0-98]
	_ = x[CodeKeypadFullStop-99]
	_ = x[CodeKeypadEqualSign-103]
	_ = x[CodeF13-104]
	_ = x[CodeF14-105]
	_ = x[CodeF15-106]
	_ = x[CodeF16-107]
	_ = x[CodeF17-108]
	_ = x[CodeF18-109]
	_ = x[CodeF19-110]
	_ = x[CodeF20-111]
	_ = x[CodeF21-112]
	_ = x[CodeF22-113]
	_ = x[CodeF23-114]
	_ = x[CodeF24-115]
	_ = x[CodeHelp-117]
	_ = x[CodeMute-127]
	_ = x[CodeVolumeUp-128]
	_ = x[CodeVolumeDown-129]
	_ = x[CodeLeftControl-224]
	_ = x[CodeLeftShift-225]
	_ = x[CodeLeftAlt-226]
	_ = x[CodeLeftGUI-227]
	_ = x[CodeRightControl-228]
	_ = x[CodeRightShift-229]
	_ = x[CodeRightAlt-230]
	_ = x[CodeRightGUI-231]
	_ = x[CodeCompose-65536]
}

const (
	_Code_name_0 = "CodeUnknown"
	_Code_name_1 = "CodeACodeBCodeCCodeDCodeECodeFCodeGCodeHCodeICodeJCodeKCodeLCodeMCodeNCodeOCodePCodeQCodeRCodeSCodeTCodeUCodeVCodeWCodeXCodeYCodeZCode1Code2Code3Code4Code5Code6Code7Code8Code9Code0CodeReturnEnterCodeEscapeCodeDeleteBackspaceCodeTabCodeSpacebarCodeHyphenMinusCodeEqualSignCodeLeftSquareBracketCodeRightSquareBracketCodeBackslash"
	_Code_name_2 = "CodeSemicolonCodeApostropheCodeGraveAccentCodeCommaCodeFullStopCodeSlashCodeCapsLockCodeF1CodeF2CodeF3CodeF4CodeF5CodeF6CodeF7CodeF8CodeF9CodeF10CodeF11CodeF12"
	_Code_name_3 = "CodePauseCodeInsertCodeHomeCodePageUpCodeDeleteForwardCodeEndCodePageDownCodeRightArrowCodeLeftArrowCodeDownArrowCodeUpArrowCodeKeypadNumLockCodeKeypadSlashCodeKeypadAsteriskCodeKeypadHyphenMinusCodeKeypadPlusSignCodeKeypadEnterCodeKeypad1CodeKeypad2CodeKeypad3CodeKeypad4CodeKeypad5CodeKeypad6CodeKeypad7CodeKeypad8CodeKeypad9CodeKeypad0CodeKeypadFullStop"
	_Code_name_4 = "CodeKeypadEqualSignCodeF13CodeF14CodeF15CodeF16CodeF17CodeF18CodeF19CodeF20CodeF21CodeF22CodeF23CodeF24"
	_Code_name_5 = "CodeHelp"
	_Code_name_6 = "CodeMuteCodeVolumeUpCodeVolumeDown"
	_Code_name_7 = "CodeLeftControlCodeLeftShiftCodeLeftAltCodeLeftGUICodeRightControlCodeRightShiftCodeRightAltCodeRightGUI"
	_Code_name_8 = "CodeCompose"
)

var (
	_Code_index_1 = [...]uint16{0, 5, 10, 15, 20, 25, 30, 35, 40, 45, 50, 55, 60, 65, 70, 75, 80, 85, 90, 95, 100, 105, 110, 115, 120, 125, 130, 135, 140, 145, 150, 155, 160, 165, 170, 175, 180, 195, 205, 224, 231, 243, 258, 271, 292, 314, 327}
	_Code_index_2 = [...]uint8{0, 13, 27, 42, 51, 63, 72, 84, 90, 96, 102, 108, 114, 120, 126, 132, 138, 145, 152, 159}
	_Code_index_3 = [...]uint16{0, 9, 19, 27, 37, 54, 61, 73, 87, 100, 113, 124, 141, 156, 174, 195, 213, 228, 239, 250, 261, 272, 283, 294, 305, 316, 327, 338, 356}
	_Code_index_4 = [...]uint8{0, 19, 26, 33, 40, 47, 54, 61, 68, 75, 82, 89, 96, 103}
	_Code_index_6 = [...]uint8{0, 8, 20, 34}
	_Code_index_7 = [...]uint8{0, 15, 28, 39, 50, 66, 80, 92, 104}
)

func (i Code) String() string {
	switch {
	case i == 0:
		return _Code_name_0
	case 4 <= i && i <= 49:
		i -= 4
		return _Code_name_1[_Code_index_1[i]:_Code_index_1[i+1]]
	case 51 <= i && i <= 69:
		i -= 51
		return _Code_name_2[_Code_index_2[i]:_Code_index_2[i+1]]
	case 72 <= i && i <= 99:
		i -= 72
		return _Code_name_3[_Code_index_3[i]:_Code_index_3[i+1]]
	case 103 <= i && i <= 115:
		i -= 103
		return _Code_name_4[_Code_index_4[i]:_Code_index_4[i+1]]
	case i == 117:
		return _Code_name_5
	case 127 <= i && i <= 129:
		i -= 127
		return _Code_name_6[_Code_index_6[i]:_Code_index_6[i+1]]
	case 224 <= i && i <= 231:
		i -= 224
		return _Code_name_7[_Code_index_7[i]:_Code_index_7[i+1]]
	case i == 65536:
		return _Code_name_8
	default:
		return "Code(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
