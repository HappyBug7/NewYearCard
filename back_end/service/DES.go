package service

import (
	"math"
	"strings"
)

type DES struct {
}

var Base64 = [64]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "+", "/"}
var Base64Map = map[string]int{"A": 0, "B": 1, "C": 2, "D": 3, "E": 4, "F": 5, "G": 6, "H": 7, "I": 8, "J": 9, "K": 10, "L": 11, "M": 12, "N": 13, "O": 14, "P": 15, "Q": 16, "R": 17, "S": 18, "T": 19, "U": 20, "V": 21, "W": 22, "X": 23, "Y": 24, "Z": 25, "a": 26, "b": 27, "c": 28, "d": 29, "e": 30, "f": 31, "g": 32, "h": 33, "i": 34, "j": 35, "k": 36, "l": 37, "m": 38, "n": 39, "o": 40, "p": 41, "q": 42, "r": 43, "s": 44, "t": 45, "u": 46, "v": 47, "w": 48, "x": 49, "y": 50, "z": 51, "0": 52, "1": 53, "2": 54, "3": 55, "4": 56, "5": 57, "6": 58, "7": 59, "8": 60, "9": 61, "+": 62, "/": 63}

var sBox = [8][4][16]int{
	{{14, 4, 13, 1, 2, 15, 11, 8, 3, 10, 6, 12, 5, 9, 0, 7},
		{0, 15, 7, 4, 14, 2, 13, 1, 10, 6, 12, 11, 9, 5, 3, 8},
		{4, 1, 14, 8, 13, 6, 2, 11, 15, 12, 9, 7, 3, 10, 5, 0},
		{15, 12, 8, 2, 4, 9, 1, 7, 5, 11, 3, 14, 10, 0, 6, 13}},
	{{15, 1, 8, 14, 6, 11, 3, 4, 9, 7, 2, 13, 12, 0, 5, 10},
		{3, 13, 4, 7, 15, 2, 8, 14, 12, 0, 1, 10, 6, 9, 11, 5},
		{0, 14, 7, 11, 10, 4, 13, 1, 5, 8, 12, 6, 9, 3, 2, 15},
		{13, 8, 10, 1, 3, 15, 4, 2, 11, 6, 7, 12, 0, 5, 14, 9}},
	{{10, 0, 9, 14, 6, 3, 15, 5, 1, 13, 12, 7, 11, 4, 2, 8},
		{13, 7, 0, 9, 3, 4, 6, 10, 2, 8, 5, 14, 12, 11, 15, 1},
		{13, 6, 4, 9, 8, 15, 3, 0, 11, 1, 2, 12, 5, 10, 14, 7},
		{1, 10, 13, 0, 6, 9, 8, 7, 4, 15, 14, 3, 11, 5, 2, 12}},
	{{7, 13, 14, 3, 0, 6, 9, 10, 1, 2, 8, 5, 11, 12, 4, 15},
		{13, 8, 11, 5, 6, 15, 0, 3, 4, 7, 2, 12, 1, 10, 14, 9},
		{10, 6, 9, 0, 12, 11, 7, 13, 15, 1, 3, 14, 5, 2, 8, 4},
		{3, 15, 0, 6, 10, 1, 13, 8, 9, 4, 5, 11, 12, 7, 2, 14}},
	{{2, 12, 4, 1, 7, 10, 11, 6, 8, 5, 3, 15, 13, 0, 14, 9},
		{14, 11, 2, 12, 4, 7, 13, 1, 5, 0, 15, 10, 3, 9, 8, 6},
		{4, 2, 1, 11, 10, 13, 7, 8, 15, 9, 12, 5, 6, 3, 0, 14},
		{11, 8, 12, 7, 1, 14, 2, 13, 6, 15, 0, 9, 10, 4, 5, 3}},
	{{12, 1, 10, 15, 9, 2, 6, 8, 0, 13, 3, 4, 14, 7, 5, 11},
		{10, 15, 4, 2, 7, 12, 9, 5, 6, 1, 13, 14, 0, 11, 3, 8},
		{9, 14, 15, 5, 2, 8, 12, 3, 7, 0, 4, 10, 1, 13, 11, 6},
		{4, 3, 2, 12, 9, 5, 15, 10, 11, 14, 1, 7, 6, 0, 8, 13}},
	{{4, 11, 2, 14, 15, 0, 8, 13, 3, 12, 9, 7, 5, 10, 6, 1},
		{13, 0, 11, 7, 4, 9, 1, 10, 14, 3, 5, 12, 2, 15, 8, 6},
		{1, 4, 11, 13, 12, 3, 7, 14, 10, 15, 6, 8, 0, 5, 9, 2},
		{6, 11, 13, 8, 1, 4, 10, 7, 9, 5, 0, 15, 14, 2, 3, 12}},
	{{13, 2, 8, 4, 6, 15, 11, 1, 10, 9, 3, 14, 5, 0, 12, 7},
		{1, 15, 13, 8, 10, 3, 7, 4, 12, 5, 6, 11, 0, 14, 9, 2},
		{7, 11, 4, 1, 9, 12, 14, 2, 0, 6, 10, 13, 15, 3, 5, 8},
		{2, 1, 14, 7, 4, 10, 8, 13, 15, 12, 9, 0, 3, 5, 6, 11}}}

var PC1 = [8][7]int{{57, 49, 41, 33, 25, 17, 9}, {1, 58, 50, 42, 34, 26, 18}, {10, 2, 59, 51, 43, 35, 27}, {19, 11, 3, 60, 52, 44, 36}, {63, 55, 47, 39, 31, 23, 15}, {7, 62, 54, 46, 38, 30, 22}, {14, 6, 61, 53, 45, 37, 29}, {21, 13, 5, 28, 20, 12, 4}}
var PC2 = [8][6]int{{14, 17, 11, 24, 1, 5}, {3, 28, 15, 6, 21, 10}, {23, 19, 12, 4, 26, 8}, {16, 7, 27, 20, 13, 2}, {41, 52, 31, 37, 47, 55}, {30, 40, 51, 45, 33, 48}, {44, 49, 39, 56, 34, 53}, {46, 42, 50, 36, 29, 32}}

var IP = [8][8]int{{58, 50, 42, 34, 26, 18, 10, 2}, {60, 52, 44, 36, 28, 20, 12, 4}, {62, 54, 46, 38, 30, 22, 14, 6}, {64, 56, 48, 40, 32, 24, 16, 8}, {57, 49, 41, 33, 25, 17, 9, 1}, {59, 51, 43, 35, 27, 19, 11, 3}, {61, 53, 45, 37, 29, 21, 13, 5}, {63, 55, 47, 39, 31, 23, 15, 7}}
var IP_ = [8][8]int{{40, 8, 48, 16, 56, 24, 64, 32}, {39, 7, 47, 15, 55, 23, 63, 31}, {38, 6, 46, 14, 54, 22, 62, 30}, {37, 5, 45, 13, 53, 21, 61, 29}, {36, 4, 44, 12, 52, 20, 60, 28}, {35, 3, 43, 11, 51, 19, 59, 27}, {34, 2, 42, 10, 50, 18, 58, 26}, {33, 1, 41, 9, 49, 17, 57, 25}}

var E = [8][6]int{{32, 1, 2, 3, 4, 5}, {4, 5, 6, 7, 8, 9}, {8, 9, 10, 11, 12, 13}, {12, 13, 14, 15, 16, 17}, {16, 17, 18, 19, 20, 21}, {20, 21, 22, 23, 24, 25}, {24, 25, 26, 27, 28, 29}, {28, 29, 30, 31, 32, 31}}

var LS = [16]int{1, 1, 2, 2, 2, 2, 2, 2, 1, 2, 2, 2, 2, 2, 2, 1}

var P = [32]int{16, 7, 20, 21, 29, 12, 28, 17, 1, 15, 23, 26, 5, 18, 31, 10, 2, 8, 24, 14, 32, 27, 3, 9, 19, 13, 30, 6, 22, 11, 4, 25}

var ori_key = [64]int{0, 1, 1, 1, 0, 0, 1, 1, 0, 1, 1, 0, 0, 0, 0, 1, 0, 1, 1, 0, 0, 1, 0, 0, 0, 1, 1, 0, 0, 0, 1, 0, 0, 1, 1, 1, 0, 1, 0, 1, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 1, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 1}

func (d *DES) leftCycle(ori_list [28]int, num int) (t [28]int) {
	temp := [28]int{}
	for i := 0; i < 28; i++ {
		temp[i] = ori_list[(i+num)%28]
	}
	return temp
}

func (d *DES) subKeysGenerate(left [28]int, right [28]int) (s [8][6]int) {
	new_turn_key := [56]int{}
	for i := 0; i < 56; i++ {
		if i < 28 {
			new_turn_key[i] = left[i]
		} else {
			new_turn_key[i] = right[i-28]
		}
	}

	subKey := [8][6]int{}

	for i := 0; i < 8; i++ {
		for j := 0; j < 6; j++ {
			subKey[i][j] = new_turn_key[PC2[i][j]-1]
		}
	}
	return subKey
}

func (d *DES) SubKeysGenerate() (s [16][8][6]int) {
	real_key := [56]int{}
	for i := 0; i < 8; i++ {
		for j := 0; j < 7; j++ {
			real_key[i*7+j] = ori_key[PC1[i][j]-1]
		}
	}

	left := [28]int{}
	right := [28]int{}

	for i := 0; i < 28; i++ {
		left[i] = real_key[i]
		right[i] = real_key[28+i]
	}

	subKey := [16][8][6]int{}

	left_list := [][28]int{}
	right_list := [][28]int{}

	for i := 0; i < 16; i++ {
		left = d.leftCycle(left, LS[i])
		right = d.leftCycle(right, LS[i])
		left_list = append(left_list, left)
		right_list = append(right_list, right)
		subKey[i] = d.subKeysGenerate(left, right)
	}

	return subKey
}

func (d *DES) E_process(Einfo [32]int) (e [8][6]int) {
	E_processed := [8][6]int{}
	for i := 0; i < 8; i++ {
		for j := 0; j < 6; j++ {
			E_processed[i][j] = Einfo[E[i][j]-1]
		}
	}
	return E_processed
}

func (d *DES) XOR_process(processed_info [8][6]int, key [8][6]int) (x [8][6]int) {
	XOR_processed := [8][6]int{}
	for i := 0; i < 8; i++ {
		for j := 0; j < 6; j++ {
			XOR_processed[i][j] = (processed_info[i][j] + key[i][j]) % 2
		}
	}
	return XOR_processed
}

func (d *DES) SBOX_process(processed_info [8][6]int) (s [32]int) {
	SBox_processed := [32]int{}
	s_slice := [4]int{}
	for i := 0; i < 8; i++ {
		row := 2*processed_info[i][0] + 1*processed_info[i][5]
		column := 8*processed_info[i][1] + 4*processed_info[i][2] + 2*processed_info[i][3] + 1*processed_info[i][4]
		num := sBox[i][row][column]
		s_slice = d.Ten2two(num)
		for j := 0; j < 4; j++ {
			SBox_processed[4*i+j] = s_slice[3-j]
		}
	}
	return SBox_processed
}

func (d *DES) Ten2two(num int) (s [4]int) {
	s_slice := [4]int{}
	num_list := []int{}
	for {
		if num > 0 {
			rem := num % 2
			num_list = append(num_list, rem)
			num = int(math.Floor(float64(num) / 2))
		} else {
			break
		}
	}
	total := len(num_list)
	for j := 0; j < 4; j = j + 1 {
		if j >= total {
			s_slice[3-j] = 0
		} else {
			s_slice[3-j] = num_list[j]
		}
	}
	return s_slice
}

func (d *DES) P_process(processed_info [32]int) (p [32]int) {
	P_processed := [32]int{}
	for i := 0; i < 32; i++ {
		P_processed[i] = processed_info[P[i]-1]
	}
	return P_processed
}

func (d *DES) Final_XOR_process(processed_info [32]int, ori_info [32]int) (r [32]int) {
	R_ := [32]int{}
	for i := 0; i < 32; i++ {
		R_[i] = (processed_info[i] + ori_info[i]) % 2
	}
	return R_
}

func (d *DES) F(L [32]int, R [32]int, turn int, subKey [16][8][6]int) (l, r [32]int) {
	L_ := R

	R_E_processed := d.E_process(R)

	turn_Key := subKey[turn]

	XOR_processed := d.XOR_process(R_E_processed, turn_Key)

	SBox_processed := d.SBOX_process(XOR_processed)

	P_processed := d.P_process(SBox_processed)

	R_ := d.Final_XOR_process(P_processed, L)

	if turn != 15 {
		return L_, R_
	} else {
		return R_, L_
	}
}

func (d *DES) re_F(L [32]int, R [32]int, turn int, subKey [16][8][6]int) (l, r [32]int) {
	L_ := R

	R_E_processed := d.E_process(R)

	turn_Key := subKey[15-turn]

	XOR_processed := d.XOR_process(R_E_processed, turn_Key)

	SBox_processed := d.SBOX_process(XOR_processed)

	P_processed := d.P_process(SBox_processed)

	R_ := d.Final_XOR_process(P_processed, L)

	if turn != 15 {
		return L_, R_
	} else {
		return R_, L_
	}
}

func (d *DES) IP_replasement(ori_Info [64]int) (l, r [32]int) {
	IP_processed := [64]int{}
	for i := 0; i < 64; i++ {
		IP_processed[i] = ori_Info[IP[int(math.Floor(float64(i)/8))][i%8]-1]
	}
	left := [32]int{}
	right := [32]int{}
	for i := 0; i < 64; i++ {
		if i < 32 {
			left[i] = IP_processed[i]
		} else {
			right[i-32] = IP_processed[i]
		}
	}
	return left, right
}

func (d *DES) IP__replasement(left, right [32]int) (i [64]int) {
	IP__processed := [64]int{}
	for i := 0; i < 64; i++ {
		num := IP_[int(math.Floor(float64(i)/8))][i%8] - 1
		if num < 32 {
			IP__processed[i] = left[num]
		} else {
			IP__processed[i] = right[num-32]
		}
	}
	return IP__processed
}

func (d *DES) Answer_translate(processed_info [64]int) (a string) {
	var builder strings.Builder
	for i := 0; i < int(math.Floor(float64(len(processed_info))/8)); i = i + 1 {
		Ascii_num := 0
		for j := 0; j < 8; j = j + 1 {
			Ascii_num = Ascii_num + int(int(math.Pow(2, float64(7-j)))*processed_info[i*8+j])
		}
		if Ascii_num != 0 {
			builder.WriteString(string(Ascii_num))
		}
	}
	answer := builder.String()
	return answer
}

func (d *DES) X642two(num int) (s [8]int) {
	s_slice := [8]int{}
	num_list := []int{}
	for {
		if num > 0 {
			rem := num % 2
			num_list = append(num_list, rem)
			num = int(math.Floor(float64(num) / 2))
		} else {
			break
		}
	}
	total := len(num_list)
	for j := 0; j < 8; j = j + 1 {
		if j >= total {
			s_slice[7-j] = 0
		} else {
			s_slice[7-j] = num_list[j]
		}
	}
	return s_slice
}

func (d *DES) Simple_Base64_Decode(info string) (i [64]int) {
	answer := []int{}
	for i := 0; i < 10; i++ {
		temp := d.X642two(Base64Map[string(info[i])])
		answer = append(answer, temp[2:]...)
	}
	temp := d.X642two(Base64Map[string(info[10])])
	answer = append(answer, temp[2:6]...)
	return [64]int(answer)
}

func (d *DES) Simple_Base64_Encode(info [64]int) (s string) {
	var builder strings.Builder
	for i := 0; i < 10; i++ {
		Ascii_num := 0
		for j := 0; j < 6; j++ {
			Ascii_num = Ascii_num + int(int(math.Pow(2, float64(5-j)))*info[i*6+j])
		}
		builder.WriteString(Base64[Ascii_num])
	}
	Ascii_num := 0
	for j := 0; j < 4; j++ {
		Ascii_num = Ascii_num + int(int(math.Pow(2, float64(5-j)))*info[60+j])
	}
	builder.WriteString(Base64[Ascii_num])
	builder.WriteString("=")
	answer := builder.String()
	return answer
}

func (d *DES) Base64_Encode_process(info [][64]int) (s string) {
	var builder strings.Builder
	for i := 0; i < len(info); i++ {
		turn_info := d.Simple_Base64_Encode(info[i])
		builder.WriteString(turn_info)
	}
	answer_string := builder.String()
	return answer_string
}

func (d *DES) Base64_Decode_process(info string) (i [][64]int) {
	answer := [][64]int{}
	for i := 0; i < int(math.Floor(float64(len(info))/12)); i++ {
		answer = append(answer, d.Simple_Base64_Decode(info[i*12:i*12+12]))
	}
	return answer
}

func (d *DES) Ori_info_process(ori_info string) (i [][64]int) {
	result := [][64]int{}
	for i := 0; i < int(math.Floor(float64(len(ori_info))/8)); i++ {
		temp := []int{}
		for j := 0; j < 8; j++ {
			temp_arr := d.X642two(int(ori_info[i*8+j]))
			temp = append(temp, temp_arr[:]...)
		}
		result = append(result, [64]int(temp))
	}
	if len(ori_info)%8 != 0 {
		temp := []int{}
		for i := 0; i < len(ori_info)%8; i++ {
			temp_arr := d.X642two(int(ori_info[int(math.Floor(float64(len(ori_info))/8))*8+i]))
			temp = append(temp, temp_arr[:]...)
		}
		for i := 0; i < 8-len(ori_info)%8; i++ {
			arr := [8]int{0, 0, 0, 0, 0, 0, 0, 0}
			temp = append(temp, arr[:]...)
		}
		result = append(result, [64]int(temp))
	}
	return result
}

func (d *DES) Simple_DES_process(mode int, ori_info [64]int) (p [64]int) {

	subKey := d.SubKeysGenerate()

	left_info, right_info := d.IP_replasement(ori_info)
	if mode == 1 {
		for i := 0; i < 16; i++ {
			left_info, right_info = d.F(left_info, right_info, i, subKey)
		}
	} else {
		for i := 0; i < 16; i++ {
			left_info, right_info = d.re_F(left_info, right_info, i, subKey)
		}
	}

	Processed_info := d.IP__replasement(left_info, right_info)

	return Processed_info

	// return d.Answer_translate(Processed_info)

	// return d.Answer_translate(d.Simple_Base64_Decode(d.Simple_Base64_Encode(Processed_info)))
}

func (d *DES) Final_DES_Encode_process(info string) (p string) {
	pre_processed_info := d.Ori_info_process(info)

	processed_info := [][64]int{}
	for i := 0; i < len(pre_processed_info); i++ {
		processed_info = append(processed_info, d.Simple_DES_process(1, pre_processed_info[i]))
	}

	Base64_processed := d.Base64_Encode_process(processed_info)

	return Base64_processed
}

func (d *DES) Final_DES_Decode_process(info string) (p string) {
	pre_processed_info := d.Base64_Decode_process(info)

	processed_info := [][64]int{}
	for i := 0; i < len(pre_processed_info); i++ {
		processed_info = append(processed_info, d.Simple_DES_process(0, pre_processed_info[i]))
	}
	var builder strings.Builder
	for i := 0; i < len(processed_info); i++ {
		builder.WriteString(d.Answer_translate(processed_info[i]))
	}
	Final_info := builder.String()

	return Final_info
}
