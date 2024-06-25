//Danny Radosevich
//UWYO CEDAR
//A Seperate file to hold the map values

package main

func getHexFromMap(kp uint16) string {
	//our mapping of the keycodes read in from /dev/input to the hex to write to /dev/hidg0
	keyCodeMap := map[uint16]string{
		1:   "29", // Escape
		2:   "1E", // one
		3:   "1F", // two
		4:   "20", // three
		5:   "21", // four
		6:   "22", // five
		7:   "23", // six
		8:   "24", // seven
		9:   "25", // eight
		10:  "26", // nine
		11:  "27", // zero
		12:  "2D", // minus
		13:  "2E", // equal
		14:  "2A", // Delete
		15:  "2B", // Tab
		16:  "14", // +q
		17:  "1A", // +w
		18:  "08", // +e
		19:  "15", // +r
		20:  "17", // +t
		21:  "1C", // +y
		22:  "18", // +u
		23:  "0C", // +i
		24:  "12", // +o
		25:  "13", // +p
		26:  "2F", // bracketleft
		27:  "30", // bracketright
		28:  "28", // Return
		29:  "E0", // Control
		30:  "04", // +a
		31:  "16", // +s
		32:  "07", // +d
		33:  "09", // +f
		34:  "0A", // +g
		35:  "0B", // +h
		36:  "0D", // +j
		37:  "0E", // +k
		38:  "0F", // +l
		39:  "33", // semicolon
		40:  "34", // apostrophe
		41:  "35", // grave
		42:  "E1", // Shift
		43:  "31", // backslash
		44:  "1D", // +z
		45:  "1B", // +x
		46:  "06", // +c
		47:  "19", // +v
		48:  "05", // +b
		49:  "11", // +n
		50:  "10", // +m
		51:  "36", // comma
		52:  "37", // period
		53:  "38", // slash
		54:  "E1", // Shift
		55:  "55", // KP_Multiply
		56:  "38", // Alt
		57:  "2C", // space
		58:  "39", // CtrlL_Lockv Caps lock
		59:  "3A", // F1
		60:  "3B", // F2
		61:  "3C", // F3
		62:  "3D", // F4
		63:  "3E", // F5
		64:  "3F", // F6
		65:  "40", // F7
		66:  "41", // F8
		67:  "42", // F9
		68:  "43", // F10
		69:  "53", // Num_Lock //prev 45
		70:  "47", // Scroll_Lock //prev 46
		71:  "5F", // KP_7 //prev 47
		72:  "60", // KP_8 //prev 48
		73:  "61", // KP_9 //prev 49
		74:  "56", // KP_Subtract //prev 4A
		75:  "5C", // KP_4 //prev 4B
		76:  "5D", // KP_5 //prev 4C
		77:  "5E", // KP_6 //prev 4D
		78:  "57", // KP_Add //prev 4E
		79:  "59", // KP_1 //prev 4F
		80:  "5A", // KP_2 //prev 50
		81:  "5B", // KP_3 //prev 51
		82:  "62", // KP_0 // prev 52
		83:  "63", // KP_Period //prev 53
		84:  "0F", // Last_Console
		86:  "56", // less
		87:  "57", // F11
		88:  "58", // F12
		96:  "58", // KP_Enter
		97:  "E0", // Control
		98:  "54", // KP_Divide
		99:  "00", // nul
		100: "38", // Alt
		101: "B7", // Break
		102: "4A", // Find home //prev B5
		103: "52", // Up //52 previously
		104: "4B", // Prior page up //prev 49
		105: "50", // Left //4B previously
		106: "4F", // Right //4D prev
		107: "4D", // Select end //prev 50
		108: "51", // Down
		109: "4E", // Next page down //prev 51
		110: "49", // Insert //prev 4C
		111: "4C", // Remove delete //prev 4A
		112: "54", // Macro
		113: "68", // F13
		114: "69", // F14
		115: "75", // Help
		116: "00", // Do
		117: "6F", // F17
		118: "56", // KP_MinPlus
		119: "48", // Pause
		121: "63", // KP_Period
		125: "38", // Windows
		126: "38", // Windows
		128: "00", // nul
		129: "00", // nul
		130: "00", // nul
		131: "00", // nul
		132: "00", // nul
		133: "00", // nul
		134: "00", // nul
		135: "00", // nul
		136: "00", // nul
		137: "00", // nul
		138: "00", // nul
		139: "00", // nul
		140: "00", // nul
		141: "00", // nul
		142: "00", // nul
		143: "00", // nul
		144: "00", // nul
		145: "00", // nul
		146: "00", // nul
		147: "00", // nul
		148: "00", // nul
		149: "00", // nul
		150: "00", // nul
		151: "00", // nul
		152: "00", // nul
		153: "00", // nul
		154: "00", // nul
		155: "00", // nul
		156: "00", // nul
		157: "00", // nul
		158: "00", // nul
		159: "00", // nul
		160: "00", // nul
		161: "00", // nul
		162: "00", // nul
		163: "00", // nul
		164: "00", // nul
		165: "00", // nul
		166: "00", // nul
		167: "00", // nul
		168: "00", // nul
		169: "00", // nul
		170: "00", // nul
		171: "00", // nul
		172: "00", // nul
		173: "00", // nul
		174: "00", // nul
		175: "00", // nul
		176: "00", // nul
		177: "00", // nul
		178: "00", // nul
		179: "00", // nul
		180: "00", // nul
		181: "00", // nul
		182: "00", // nul
		183: "00", // nul
		184: "00", // nul
		185: "00", // nul
		186: "00", // nul
		187: "00", // nul
		188: "00", // nul
		189: "00", // nul
		190: "00", // nul
		191: "00", // nul
		192: "00", // nul
		193: "00", // nul
		194: "00", // nul
		195: "00", // nul
		196: "00", // nul
		197: "00", // nul
		198: "00", // nul
		199: "00", // nul
		200: "00", // nul
		201: "00", // nul
		202: "00", // nul
		203: "00", // nul
		204: "00", // nul
		205: "00", // nul
		206: "00", // nul
		207: "00", // nul
		208: "00", // nul
		209: "00", // nul
		210: "00", // nul
		211: "00", // nul
		212: "00", // nul
		213: "00", // nul
		214: "00", // nul
		215: "00", // nul
		216: "00", // nul
		217: "00", // nul
		218: "00", // nul
		219: "00", // nul
		220: "00", // nul
		221: "00", // nul
		222: "00", // nul
		223: "00", // nul
		224: "00", // nul
		225: "00", // nul
		226: "00", // nul
		227: "00", // nul
		228: "00", // nul
		229: "00", // nul
		230: "00", // nul
		231: "00", // nul
		232: "00", // nul
		233: "00", // nul
		234: "00", // nul
		235: "00", // nul
		236: "00", // nul
		237: "00", // nul
		238: "00", // nul
		239: "00", // nul
		240: "00", // nul
		241: "00", // nul
		242: "00", // nul
		243: "00", // nul
		244: "00", // nul
		245: "00", // nul
		246: "00", // nul
		247: "00", // nul
		248: "00", // nul
		249: "00", // nul
		250: "00", // nul
		251: "00", // nul
		252: "00", // nul
		253: "00", // nul
		254: "00", // nul
		255: "00", // nul
	}
	return keyCodeMap[kp]
}
