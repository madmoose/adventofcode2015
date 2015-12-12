package main

import "fmt"

type element struct {
	length int
	decay  []int
	value  string
}

var elements = []element{
	element{0, []int{}, ""}, // dummy
	element{4, []int{63}, "1112"},
	element{7, []int{64, 62}, "1112133"},
	element{12, []int{65}, "111213322112"},
	element{12, []int{66}, "111213322113"},
	element{4, []int{68}, "1113"},
	element{5, []int{69}, "11131"},
	element{12, []int{84, 55}, "111311222112"},
	element{6, []int{70}, "111312"},
	element{8, []int{71}, "11131221"},
	element{10, []int{76}, "1113122112"},
	element{10, []int{77}, "1113122113"},
	element{14, []int{82}, "11131221131112"},
	element{12, []int{78}, "111312211312"},
	element{14, []int{79}, "11131221131211"},
	element{18, []int{80}, "111312211312113211"},
	element{42, []int{81, 29, 91}, "111312211312113221133211322112211213322112"},
	element{42, []int{81, 29, 90}, "111312211312113221133211322112211213322113"},
	element{26, []int{81, 30}, "11131221131211322113322112"},
	element{14, []int{75, 29, 92}, "11131221133112"},
	element{28, []int{75, 32}, "1113122113322113111221131221"},
	element{14, []int{72}, "11131221222112"},
	element{24, []int{73}, "111312212221121123222112"},
	element{24, []int{74}, "111312212221121123222113"},
	element{5, []int{83}, "11132"},
	element{7, []int{86}, "1113222"},
	element{10, []int{87}, "1113222112"},
	element{10, []int{88}, "1113222113"},
	element{8, []int{89, 92}, "11133112"},
	element{2, []int{1}, "12"},
	element{9, []int{3}, "123222112"},
	element{9, []int{4}, "123222113"},
	element{23, []int{2, 61, 29, 85}, "12322211331222113112211"},
	element{2, []int{5}, "13"},
	element{6, []int{28}, "131112"},
	element{32, []int{24, 33, 61, 29, 91}, "13112221133211322112211213322112"},
	element{32, []int{24, 33, 61, 29, 90}, "13112221133211322112211213322113"},
	element{8, []int{7}, "13122112"},
	element{3, []int{8}, "132"},
	element{5, []int{9}, "13211"},
	element{6, []int{10}, "132112"},
	element{10, []int{21}, "1321122112"},
	element{18, []int{22}, "132112211213322112"},
	element{18, []int{23}, "132112211213322113"},
	element{6, []int{11}, "132113"},
	element{10, []int{19}, "1321131112"},
	element{8, []int{12}, "13211312"},
	element{7, []int{13}, "1321132"},
	element{8, []int{14}, "13211321"},
	element{12, []int{15}, "132113212221"},
	element{20, []int{18}, "13211321222113222112"},
	element{34, []int{16}, "1321132122211322212221121123222112"},
	element{34, []int{17}, "1321132122211322212221121123222113"},
	element{20, []int{20}, "13211322211312113211"},
	element{10, []int{6, 61, 29, 92}, "1321133112"},
	element{7, []int{26}, "1322112"},
	element{7, []int{27}, "1322113"},
	element{11, []int{25, 29, 92}, "13221133112"},
	element{13, []int{25, 29, 67}, "1322113312211"},
	element{21, []int{25, 29, 85}, "132211331222113112211"},
	element{17, []int{25, 29, 68, 61, 29, 89}, "13221133122211332"},
	element{2, []int{61}, "22"},
	element{1, []int{33}, "3"},
	element{4, []int{40}, "3112"},
	element{7, []int{41}, "3112112"},
	element{14, []int{42}, "31121123222112"},
	element{14, []int{43}, "31121123222113"},
	element{7, []int{38, 39}, "3112221"},
	element{4, []int{44}, "3113"},
	element{6, []int{48}, "311311"},
	element{8, []int{54}, "31131112"},
	element{10, []int{49}, "3113112211"},
	element{16, []int{50}, "3113112211322112"},
	element{28, []int{51}, "3113112211322112211213322112"},
	element{28, []int{52}, "3113112211322112211213322113"},
	element{9, []int{47, 38}, "311311222"},
	element{12, []int{47, 55}, "311311222112"},
	element{12, []int{47, 56}, "311311222113"},
	element{16, []int{47, 57}, "3113112221131112"},
	element{18, []int{47, 58}, "311311222113111221"},
	element{24, []int{47, 59}, "311311222113111221131221"},
	element{23, []int{47, 60}, "31131122211311122113222"},
	element{16, []int{47, 33, 61, 29, 92}, "3113112221133112"},
	element{6, []int{45}, "311312"},
	element{5, []int{46}, "31132"},
	element{15, []int{53}, "311322113212221"},
	element{6, []int{38, 29, 89}, "311332"},
	element{10, []int{38, 30}, "3113322112"},
	element{10, []int{38, 31}, "3113322113"},
	element{3, []int{34}, "312"},
	element{27, []int{36}, "312211322212221121123222113"},
	element{27, []int{35}, "312211322212221121123222112"},
	element{5, []int{37}, "32112"},
}

func next(in []int) []int {
	var out = make([]int, len(elements))

	for i := 1; i != len(elements); i++ {
		for _, e := range elements[i].decay {
			out[i] += in[e]
		}
	}

	return out
}

func main() {
	var s string
	if n, err := fmt.Scanf("%s\n", &s); n != 1 {
		panic(err)
	}

	var e int
	for e = 1; e != len(elements); e++ {
		if s == elements[e].value {
			break
		}
	}
	if e == len(elements) {
		panic(s)
	}

	var lengths = make([]int, len(elements))
	for i := 1; i != len(lengths); i++ {
		lengths[i] = elements[i].length
	}

	for i := 0; i != 50; i++ {
		lengths = next(lengths)
	}
	fmt.Println(lengths[e])
}
