package numbers

import "fmt"

func Convert(num int) string {
	chars := []string{
		"〇",
		"一",
		"二",
		"三",
		"四",
		"五",
		"六",
		"七",
		"八",
		"九",
		"十",
		"百",
		"千",
		"万",
	}
	bits := []int{}
	if num == 0 {
		return chars[0]
	}
	for num != 0 {
		bits = append(bits, num%10)
		num /= 10
	}

	// fmt.Printf("bits = %v\n", bits)

	reslist := []int{}
	var m, n int
	l := len(bits)
	k := l - 1
	for i := 0; i < l; i++ {
		if bits[l-i-1] == 0 {
			m = 1
		} else {
			if m == 1 {
				reslist = append(reslist, 0)
				n++
			}
			reslist = append(reslist, bits[l-i-1])
			n++
			if l-i-1 != 0 {
				reslist = append(reslist, k+9)
				n++
			}
			m = 0
		}
		k--
	}

	var res string

	for i := 0; i < len(reslist); i++ {
		res += chars[reslist[i]]
	}

	// fmt.Printf("len = %d\n", n)
	return res
}

func Convert2(num int) string {
	chars := []string{
		"〇",
		"一",
		"二",
		"三",
		"四",
		"五",
		"六",
		"七",
		"八",
		"九",
		"十",
		"百",
		"千",
		"万",
	}
	bits := []int{}
	if num == 0 {
		return chars[0]
	}
	for num != 0 {
		bits = append(bits, num%10)
		num /= 10
	}

	fmt.Printf("bits = %v\n", bits)

	reslist := []int{}
	var m int
	l := len(bits)
	k := l - 1
	for i := 0; i < l; i++ {
		if bits[l-i-1] == 0 {
			m = 1
		} else {
			if m == 1 {
				reslist = append(reslist, 0)
			}
			reslist = append(reslist, bits[l-i-1])
			if l-i-1 != 0 {
				reslist = append(reslist, k+9)
			}
			m = 0
		}
		k--
	}

	var res string

	for i := 0; i < len(reslist); i++ {
		res += chars[reslist[i]]
	}
	return res
}
