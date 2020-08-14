package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	/*
							x := 4
							fmt.Printf("%d",adjacentElementsProduct(x))
							diamond(x)

							var x = []int{1, 3, 2, 1}
											fmt.Printf("diff %d",makeArrayConsecutive2(x))
											fmt.Printf("%v is %t\n",x,almostIncreasingSequence(x))
											x = []int{1, 3, 2}
											fmt.Printf("diff %d",makeArrayConsecutive2(x))
											fmt.Printf("%v is %t\n",x,almostIncreasingSequence(x))

											x = []int{10, 1, 2, 3, 4, 5}
											fmt.Printf("%v is %t\n",x,almostIncreasingSequence(x))
											fmt.Printf("%v is %t\n",x,almostIncreasingSequence(x))
											x = []int{1, 2, 3, 4, 3, 6}
											fmt.Printf("%v is %t\n",x,almostIncreasingSequence(x))
											x := []int{1, 2, 5, 3, 5}
											x := []int{1, 2, 1, 2}
											fmt.Printf("Before %v\n",x)
											fmt.Printf("%v is %t\n",x,almostIncreasingSequence(x))
											x := []string{"aba",  "aa",  "ad",  "vcd",  "aba"}
											fmt.Printf("%q => %q\n",x,allLongestStrings(x))
											s1 := "aabcc"
											s2 := "adcaa"
											fmt.Println("s1 ",s1," s2 ",s2," has ",commonCharacterCount(s1,s2))
											fmt.Println(1634,isLucky(1634))
											y := []int{-1, 150, 190, 170, -1, -1, 160, 180}
											fmt.Println(y," => ",sortByHeight(y))
										x := "foo(bar(baz))blim"
										fmt.Println(x," => ",reverseInParentheses(x))
										fmt.Println(x," => ",Reverse(x))

										var picture = []string{"abc", "ded"}
										var result = addBorder(picture)
										for _, v := range result {
											fmt.Println(v)
										}
									var a = []int{1, 2, 4, 3, 3}
									var b = []int{1, 2, 3, 3, 3}
									fmt.Printf("areSimilar(a,b) %t", areSimilar(a, b))

								var a = []int{-1000, 0, -2, 0}
								fmt.Printf("arrayChange(a) %d", arrayChange(a))

						var x string = "aabb"
						fmt.Printf("palindromeRearranging(\"%s\") = %t", x, palindromeRearranging(x))
					var inputArray = []int{2, 4, 1, 0}
					fmt.Printf("arrayMaximalAdjacentDifference(\"%v\") = %d", inputArray, arrayMaximalAdjacentDifference(inputArray))

				var ip_string = "172.16.254.1"
				fmt.Printf("isIPv4Address(%s) = %t", ip_string, isIPv4Address(ip_string))

			var x = []int{1, 4, 10, 6, 2}
			fmt.Printf("%v expected 7 got %d\n", x, avoidObstacles(x))
			x = []int{1000, 999}
			fmt.Printf("%v expected 6 got %d\n", x, avoidObstacles(x))
			x = []int{5, 8, 9, 13, 14}
			fmt.Printf("%v expected 6 got %d\n", x, avoidObstacles(x))
			x = []int{19, 32, 11, 23}
			fmt.Printf("%v expected 3 got %d\n", x, avoidObstacles(x))

		var x = [][]int{{7, 4, 0, 1},
			{5, 6, 2, 2},
			{6, 10, 7, 8},
			{1, 4, 2, 0}}
		y := boxBlur(x)
		for i := 0; i < len(y); i++ {
			for j := 0; j < len(y[0]); j++ {
				fmt.Printf("%d ", y[i][j])
			}
			fmt.Printf("\n")
		}
	*/
	input1 := [][]bool{{true, false, false},
		{false, true, false},
		{false, false, false}}
	matrix := minesweeper(input1)

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Printf("\n")
	}
}

//#24
/*
	In the popular Minesweeper game you have a board with some mines and those cells that don't contain
	a mine have a number in it that indicates the total number of mines in the neighboring cells. Starting
	off with some arrangement of mines we want to create a Minesweeper game setup.

	Example

	For

	matrix = {{true, false, false},
			{false, true, false},
			{false, false, false}}

	the output should be

	minesweeper(matrix) = [[1, 2, 1],
						[2, 1, 1],
						[1, 1, 1]]

	Check out the image below for better understanding:

	Input/Output

		[execution time limit] 4 seconds (go)

		[input] array.array.boolean matrix

		A non-empty rectangular matrix consisting of boolean values - true if the corresponding cell
		contains a mine, false otherwise.

		Guaranteed constraints:
		2 ≤ matrix.length ≤ 100,
		2 ≤ matrix[0].length ≤ 100.

		[output] array.array.integer
			Rectangular matrix of the same size as matrix each cell of which contains an integer
			equal to the number of mines in the neighboring cells. Two cells are called neighboring
			if they share at least one corner.



*/
func minesweeper(matrix [][]bool) [][]int {
	var len1 int = len(matrix[0])
	var len2 int = len(matrix)
	var result = make([][]int, len2)
	for i := range result {
		result[i] = make([]int, len1)
	}
	for i := 0; i < len2; i++ {
		for j := 0; j < len1; j++ {
			var sum int = 0
			// fmt.Printf("For %dx%d checking\n", i, j)

			for k := i - 1; k < i+2; k++ {
				for l := j - 1; l < j+2; l++ {
					if k < 0 || k > len2-1 || l < 0 || l > len1-1 {
						continue
					}
					// fmt.Printf(" %dx%d ", k, l)
					if matrix[k][l] {
						sum++
					}

				}
			}
			if matrix[i][j] {
				result[i][j] = sum - 1
			} else {
				result[i][j] = sum
			}

		}
	}
	return result
}

//#23
/*

	Last night you partied a little too hard. Now there's a black and white photo of you that's
	about to go viral! You can't let this ruin your reputation, so you want to apply the box
	blur algorithm to the photo to hide its content.

	The pixels in the input image are represented as integers. The algorithm distorts the input
	image in the following way: Every pixel x in the output image has a value equal to the
	average value of the pixel values from the 3 × 3 square that has its center at x,
	including x itself. All the pixels on the border of x are then removed.

	Return the blurred image as an integer, with the fractions rounded down.

	Example

	For

	image = [[1, 1, 1],
			[1, 7, 1],
			[1, 1, 1]]

	the output should be boxBlur(image) = [[1]].

	To get the value of the middle pixel in the input 3 × 3 square: (1 + 1 + 1 + 1 + 7 + 1 + 1 + 1 + 1) = 15 / 9 = 1.66666 = 1. The border pixels are cropped from the final result.

	For

	image = [[7, 4, 0, 1],
			[5, 6, 2, 2],
			[6, 10, 7, 8],
			[1, 4, 2, 0]]

	the output should be

	boxBlur(image) = [[5, 4],
					[4, 4]]

	There are four 3 × 3 squares in the input image, so there should be four integers in the blurred output.
	To get the first value: (7 + 4 + 0 + 5 + 6 + 2 + 6 + 10 + 7) = 47 / 9 = 5.2222 = 5. The other three
	integers are obtained the same way, then the surrounding integers are cropped from the final result.

	Input/Output

		[execution time limit] 4 seconds (go)

		[input] array.array.integer image

		An image, stored as a rectangular matrix of non-negative integers.

		Guaranteed constraints:
		3 ≤ image.length ≤ 100,
		3 ≤ image[0].length ≤ 100,
		0 ≤ image[i][j] ≤ 255.

		[output] array.array.integer
			A blurred image represented as integers, obtained through the process in the description.


	Input:

	image:
	[[1,1,1],
	[1,7,1],
	[1,1,1]]

	Expected Output:

	[[1]]

	Input:

	image:
	[[0,18,9],
	[27,9,0],
	[81,63,45]]

	Expected Output:

	[[28]]
	Input:

	image:
	[[36,0,18,9],
	[27,54,9,0],
	[81,63,72,45]]

	Expected Output:

	[[40,30]]
	Input:

	image:
	[[7,4,0,1],
	[5,6,2,2],
	[6,10,7,8],
	[1,4,2,0]]

	Expected Output:

	[[5,4],
	[4,4]]
	Input:

	image:
	[[36,0,18,9,9,45,27],
	[27,0,54,9,0,63,90],
	[81,63,72,45,18,27,0],
	[0,0,9,81,27,18,45],
	[45,45,27,27,90,81,72],
	[45,18,9,0,9,18,45],
	[27,81,36,63,63,72,81]]

	Expected Output:

	[[39,30,26,25,31],
	[34,37,35,32,32],
	[38,41,44,46,42],
	[22,24,31,39,45],
	[37,34,36,47,59]]
	Input:

	image:
	[[36,0,18,9,9,45,27],
	[27,0,254,9,0,63,90],
	[81,255,72,45,18,27,0],
	[0,0,9,81,27,18,45],
	[45,45,227,227,90,81,72],
	[45,18,9,255,9,18,45],
	[27,81,36,127,255,72,81]]

	Expected Output:

	[[82,73,48,25,31],
	[77,80,57,32,32],
	[81,106,88,68,42],
	[44,96,103,89,45],
	[59,113,137,126,80]]
*/
func boxBlur(image [][]int) [][]int {
	var len1 int = len(image[0]) - 2
	var len2 int = len(image) - 2
	var blurred = make([][]int, len2)
	for i := range blurred {
		blurred[i] = make([]int, len1)
	}
	// fmt.Printf("blurred created of %dx%d\n", len(blurred), len(blurred[0]))
	for i := 1; i < len(image)-1; i++ {
		for j := 1; j < len(image[0])-1; j++ {
			// fmt.Printf("Here for %dx%d\n", i, j)
			var sum int = 0
			// fmt.Printf("k = %d to %d  l = %d to %d\n", i-1, i+1, j-1, j+1)
			for k := i - 1; k < i+2; k++ {
				for l := j - 1; l < j+2; l++ {
					// fmt.Printf("%d %d ", k, l)
					sum += image[k][l]
				}
				// fmt.Printf("\n")
			}
			// fmt.Printf("Populating blurred[%d][%d] \n", i-1, j-1)
			blurred[i-1][j-1] = int(math.Floor(float64(sum / 9)))
			// fmt.Printf("\n")
		}
	}

	return blurred
}

//#22
/*
	You are given an array of integers representing coordinates of obstacles situated on a straight line.

	Assume that you are jumping from the point with coordinate 0 to the right. You are allowed only to make
	jumps of the same length represented by some integer.

	Find the minimal length of the jump enough to avoid all the obstacles.

	Example

	For inputArray = [5, 3, 6, 7, 9], the output should be
	avoidObstacles(inputArray) = 4.

	Check out the image below for better understanding:

	Input/Output

		[execution time limit] 4 seconds (go)

		[input] array.integer inputArray

		Non-empty array of positive integers.

		Guaranteed constraints:
		2 ≤ inputArray.length ≤ 1000,
		1 ≤ inputArray[i] ≤ 1000.

		[output] integer
			The desired length.

			Input:

inputArray: [2, 3]

Output:

	2

	Expected Output:

	4

	Input:

	inputArray: [1000, 999]

	Output:

	2

	Expected Output:

	6

	Input:

	inputArray: [5, 8, 9, 13, 14]

	Output:

	10

	Expected Output:

	6
	Input:

	inputArray: [1, 4, 10, 6, 2]

	Output:

	5

	Expected Output:

	7
	Input:

	inputArray: [1000, 999]

	Output:

	4

	Expected Output:

	6
	Input:

	inputArray: [5, 8, 9, 13, 14]

	Output:

	10

	Expected Output:

	6

*/
func avoidObstacles(inputArray []int) int {
	var distance int
	var data_dic_org = make(map[int]bool)
	var max_val int = 0
	var min_val int = inputArray[0]
	for _, v := range inputArray {
		data_dic_org[v] = false
		lowM := primeMultiples(v)
		for key, _ := range lowM {
			//start multiple loop
			// fmt.Printf("Lowest of %d is %d\n", v, key)
			if key < min_val {
				min_val = key
			}
			if key > max_val {
				max_val = key
			}
			//end of multiple loop
		}
	}

	loopForward := false
	// fmt.Printf("max_val %d\n", max_val)
	// for distance = 1; distance <= max_val+1; distance++ {
	for distance = 1; distance <= max_val+1; distance++ {
		loopForward = false
		// fmt.Printf("Trying with d %d\n", distance)

		for j := 0; j <= max_val+1; j += distance {
			// fmt.Printf("Inc %d\n", j)

			_, is_still_present := data_dic_org[j]

			if is_still_present {
				// fmt.Printf("%d is present \n", j)
				loopForward = true
				// distance = j
				break
			}
		}
		if !loopForward {
			break
		}
	}
	if !loopForward {
		return distance
	} else {
		return len(inputArray)
	}
}

func primeMultiples(p_number int) map[int]bool {
	v_multiple := make(map[int]bool)
	if p_number <= 2 {
		v_multiple[2] = false
		return v_multiple
	}
	for i := 2; i <= int(p_number/2); i++ {
		// _, is_present := v_multiple[i]
		if primeCheck(i) {
			// fmt.Printf("Returning %d\n", i)
			// fmt.Printf("primeCheck(%d) = %t\n", i, primeCheck(i))
			v_multiple[i] = false
		}
	}
	v_multiple[p_number] = false
	return v_multiple
}

func primeCheck(p_number int) bool {
	var flag bool = true

	for i := 2; i <= int(p_number/2); i++ {
		if p_number%i == 0 {
			return false
		}
	}
	return flag
}

//#21
/*
	An IP address is a numerical label assigned to each device (e.g., computer, printer) participating in a computer network that uses the Internet Protocol for communication. There are two versions of the Internet protocol, and thus two versions of addresses. One of them is the IPv4 address.

	Given a string, find out if it satisfies the IPv4 address naming rules.

	Example

		For inputString = "172.16.254.1", the output should be
		isIPv4Address(inputString) = true;

		For inputString = "172.316.254.1", the output should be
		isIPv4Address(inputString) = false.

		316 is not in range [0, 255].

		For inputString = ".254.255.0", the output should be
		isIPv4Address(inputString) = false.

		There is no first number.

	Input/Output

		[execution time limit] 4 seconds (go)

		[input] string inputString

		A string consisting of digits, full stops and lowercase English letters.

		Guaranteed constraints:
		1 ≤ inputString.length ≤ 30.

		[output] boolean
			true if inputString satisfies the IPv4 address naming rules, false otherwise.

*/
func isIPv4Address(inputString string) bool {
	var flag bool = true
	var s = strings.Split(inputString, ".")
	if len(s) != 4 {
		return false
	}
	for _, v := range s {
		x, err := strconv.Atoi(v)
		if err != nil {
			return false
		}
		if len(v) != len(strconv.Itoa(x)) {
			return false
		}
		if !(x >= 0 && x <= 255) {
			return false
		}
	}
	return flag
}

//#20
/*
	Given an array of integers, find the maximal absolute difference between any two of its adjacent elements.

	Example

	For inputArray = [2, 4, 1, 0], the output should be
	arrayMaximalAdjacentDifference(inputArray) = 3.

	Input/Output

		[execution time limit] 4 seconds (go)

		[input] array.integer inputArray

		Guaranteed constraints:
		3 ≤ inputArray.length ≤ 10,
		-15 ≤ inputArray[i] ≤ 15.

		[output] integer
			The maximal absolute difference.
*/
func arrayMaximalAdjacentDifference(inputArray []int) int {
	var max_diff int = 0

	for i := 1; i < len(inputArray); i++ {
		var s, l int
		if inputArray[i-1] >= inputArray[i] {
			s = inputArray[i]
			l = inputArray[i-1]
		} else {
			s = inputArray[i-1]
			l = inputArray[i]
		}

		if max_diff < (l - s) {
			max_diff = (l - s)
		}
	}

	return max_diff
}

//#19
/*
	Call two arms equally strong if the heaviest weights they each are able to lift are equal.

	Call two people equally strong if their strongest arms are equally strong (the strongest
		arm can be both the right and the left), and so are their weakest arms.

	Given your and your friend's arms' lifting capabilities find out if you two are equally strong.

	Example

		For yourLeft = 10, yourRight = 15, friendsLeft = 15, and friendsRight = 10, the output should be
		areEquallyStrong(yourLeft, yourRight, friendsLeft, friendsRight) = true;
		For yourLeft = 15, yourRight = 10, friendsLeft = 15, and friendsRight = 10, the output should be
		areEquallyStrong(yourLeft, yourRight, friendsLeft, friendsRight) = true;
		For yourLeft = 15, yourRight = 10, friendsLeft = 15, and friendsRight = 9, the output should be
		areEquallyStrong(yourLeft, yourRight, friendsLeft, friendsRight) = false.

	Input/Output

		[execution time limit] 4 seconds (go)

		[input] integer yourLeft

		A non-negative integer representing the heaviest weight you can lift with your left arm.

		Guaranteed constraints:
		0 ≤ yourLeft ≤ 20.

		[input] integer yourRight

		A non-negative integer representing the heaviest weight you can lift with your right arm.

		Guaranteed constraints:
		0 ≤ yourRight ≤ 20.

		[input] integer friendsLeft

		A non-negative integer representing the heaviest weight your friend can lift with his or her left arm.

		Guaranteed constraints:
		0 ≤ friendsLeft ≤ 20.

		[input] integer friendsRight

		A non-negative integer representing the heaviest weight your friend can lift with his or her right arm.

		Guaranteed constraints:
		0 ≤ friendsRight ≤ 20.

		[output] boolean
			true if you and your friend are equally strong, false otherwise.

*/
func areEquallyStrong(yourLeft int, yourRight int, friendsLeft int, friendsRight int) bool {
	if (yourLeft == friendsLeft && yourRight == friendsRight) || (yourLeft == friendsRight && yourRight == friendsLeft) {
		return true
	} else {
		return false
	}
}

//#18
/*
	Given a string, find out if its characters can be rearranged to form a palindrome.

	Example

	For inputString = "aabb", the output should be
	palindromeRearranging(inputString) = true.

	We can rearrange "aabb" to make "abba", which is a palindrome.

	Input/Output

		[execution time limit] 4 seconds (go)

		[input] string inputString

		A string consisting of lowercase English letters.

		Guaranteed constraints:
		1 ≤ inputString.length ≤ 50.

		[output] boolean
			true if the characters of the inputString can be rearranged to form a palindrome, false otherwise.

*/
func palindromeRearranging(inputString string) bool {
	var flag bool = true
	m1 := make(map[rune]int)
	for _, i := range inputString {
		_, is_present := m1[i]

		if is_present {
			m1[i]++
		} else {
			m1[i] = 1
		}
	}

	var miss_count int = 1
	for _, val := range m1 {
		if val%2 != 0 && miss_count == 1 {
			miss_count--
		} else if val%2 != 0 && miss_count == 0 {
			return false
		}
	}

	return flag
}

//#17
/*
	You are given an array of integers. On each move you are allowed to increase exactly
	one of its element by one. Find the minimal number of moves required to obtain a strictly
	increasing sequence from the input.

	Example

		For inputArray = [1, 1, 1], the output should be
		arrayChange(inputArray) = 3.

	Input/Output

		[execution time limit] 4 seconds (go)

		[input] array.integer inputArray

		Guaranteed constraints:
		3 ≤ inputArray.length ≤ 105,
		-105 ≤ inputArray[i] ≤ 105.

		[output] integer
			The minimal number of moves needed to obtain a strictly increasing sequence from inputArray.
			It's guaranteed that for the given test cases the answer always fits signed 32-bit integer type.

*/
func arrayChange(inputArray []int) int {
	var no_of_change int = 0
	for index, value := range inputArray {
		if index == 0 {
			continue
		}
		if inputArray[index-1] < value {
			continue
		} else {
			var new_value int
			new_value = inputArray[index-1] + 1
			no_of_change += int(math.Abs(float64(new_value - value)))
			inputArray[index] = new_value
		}
	}
	return no_of_change
}

//#16
/*
	Two arrays are called similar if one can be obtained from another by swapping at most one
	pair of elements in one of the arrays.

	Given two arrays a and b, check whether they are similar.

	Example

		For a = [1, 2, 3] and b = [1, 2, 3], the output should be
		areSimilar(a, b) = true.

		The arrays are equal, no need to swap any elements.

		For a = [1, 2, 3] and b = [2, 1, 3], the output should be
		areSimilar(a, b) = true.

		We can obtain b from a by swapping 2 and 1 in b.

		For a = [1, 2, 2] and b = [2, 1, 1], the output should be
		areSimilar(a, b) = false.

		Any swap of any two elements either in a or in b won't make a and b equal.

	Input/Output

	[execution time limit] 4 seconds (go)

	[input] array.integer a

	Array of integers.

	Guaranteed constraints:
	3 ≤ a.length ≤ 105,
	1 ≤ a[i] ≤ 1000.

	[input] array.integer b

	Array of integers of the same length as a.

	Guaranteed constraints:
	b.length = a.length,
	1 ≤ b[i] ≤ 1000.

	[output] boolean
		true if a and b are similar, false otherwise.

*/
func areSimilar(a []int, b []int) bool {
	var flag bool = true
	m1 := make(map[int]int)
	m2 := make(map[int]int)

	for _, i := range a {
		_, is_present := m1[i]

		if is_present {
			m1[i]++
		} else {
			m1[i] = 1
		}
	}
	for _, i := range b {
		_, is_present := m2[i]

		if is_present {
			m2[i]++
		} else {
			m2[i] = 1
		}
	}

	if len(m1) != len(m2) {
		return false
	}
	fmt.Print(m1)
	fmt.Print(m2)
	for key, value1 := range m1 {
		value2, is_present2 := m2[key]
		if !is_present2 {
			return false
		}
		//check character count
		if value1 != value2 {
			return false
		}
	}

	var mistake_cnt int = 2
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			// fmt.Printf("Mistake at %d %d != %d\n", i, a[i], b[i])
			mistake_cnt--
		}
		if mistake_cnt < 0 {
			break
		}
	}
	if mistake_cnt < 0 {
		return false
	}
	return flag
}

//#15
/*
	Given a rectangular matrix of characters, add a border of asterisks(*) to it.

	picture = ["abc",
			"ded"]

	addBorder(picture) = ["*****",
						"*abc*",
						"*ded*",
						"*****"]
*/
func addBorder(picture []string) []string {
	var max_length int = 0
	var new_picture []string
	for _, x := range picture {
		if max_length < len(x) {
			max_length = len(x)
		}
		new_picture = append(new_picture, "*"+x+"*")
	}
	var stars = strings.Repeat("*", max_length+2)
	var temp = []string{stars}
	temp = append(temp, new_picture...)
	temp = append(temp, stars)
	return temp
}

//#14
/*
	Several people are standing in a row and need to be divided into two teams.
	The first person goes into team 1, the second goes into team 2, the third goes
	into team 1 again, the fourth into team 2, and so on.

	You are given an array of positive integers - the weights of the people. Return
	an array of two integers, where the first element is the total weight of team 1,
	and the second element is the total weight of team 2 after the division is
	complete.
*/
func alternatingSums(a []int) []int {
	var sum = []int{0, 0}
	for i := 0; i < len(a); i++ {
		if i%2 == 0 {
			sum[0] += a[i]
		} else {
			sum[1] += a[i]
		}
	}

	return sum
}

//#13
/*
	Write a function that reverses characters in (possibly nested) parentheses in
	the input string.
	Input strings will always be well-formed with matching ()s.

	func reverseInParentheses(inputString string) string {
		clone_slice := []rune(inputString)
		final_slice := []rune()

		for i,_:=range clone_slice{

		}
	}
*/
func reverseInParentheses(inputString string) string {
	b_count := 0
	b_start := -1
	b_end := -1
	looking_form_end := false
	for i, x := range inputString {

		if x == '(' {
			b_count++
			b_start = i
			looking_form_end = true
		} else if x == ')' && looking_form_end {
			looking_form_end = false
			b_count++
			b_end = i
		}
	}
	fmt.Println(" inputString ", inputString, "b_start ", b_start, " b_end ", b_end, " b_count ", b_count)
	//if count 2 just 2 brackets then
	if b_count >= 2 && b_end > -1 && b_start > -1 {
		//recursion
		inputString_arr := []rune(inputString)
		corrected_slice := inputString_arr[0:b_start]
		rev_string := Reverse(string(inputString_arr[b_start+1 : b_end]))
		corrected_slice = append(corrected_slice, []rune(rev_string)...)
		corrected_slice = append(corrected_slice, inputString_arr[b_end+1:]...)

		return reverseInParentheses(string(corrected_slice))
	} else /* if b_count==0*/ {
		return inputString
	}
}

func Reverse(s string) string {
	runes := []rune(s)
	// fmt.Println(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

//#12
/*
	Some people are standing in a row in a park.
	There are trees between them which cannot be moved.
	Your task is to rearrange the people by their heights
	in a non-descending order without moving the trees.
	People can be very tall!
*/
func sortByHeight(a []int) []int {
	var temp_slice []int
	final_slice := a
	for _, x := range a {
		if x >= 0 {
			temp_slice = append(temp_slice, x)
		}
	}
	//sort temp
	for i, _ := range temp_slice {
		for j := i + 1; j < len(temp_slice); j++ {
			if temp_slice[i] > temp_slice[j] {
				temp_slice[i], temp_slice[j] = temp_slice[j], temp_slice[i]
			}
		}
	}
	j := 0
	for i, _ := range final_slice {
		if final_slice[i] >= 0 {
			final_slice[i] = temp_slice[j]
			j++
		}
	}
	// fmt.Println(temp_slice)
	return final_slice
}

//#11
/*
	Ticket numbers usually consist of an even number of digits. A ticket number is
	considered lucky if the sum of the first half of the digits is equal to the sum
	of the second half.
*/
func isLucky(n int) bool {
	copy_n := n
	n_length := 0
	for ; copy_n > 0; copy_n /= 10 {
		n_length++
	}
	half := float64(n_length / 2)
	first_half := n / int(math.Pow(10, half))
	second_half := n % int(math.Pow(10, half))
	// fmt.Println("n_length/2 ",n_length/2," (10^half) ",math.Pow(10,half)," n/(10^half) ",n/math.Pow(10,half))
	// fmt.Println("first_half ",first_half," second_half ",second_half)
	return digitSum(first_half) == digitSum(second_half)
}

func digitSum(x int) int {
	l_sum := 0
	for ; x > 0; x /= 10 {
		l_sum += x % 10
	}
	return l_sum
}

//#10
/*
	Given two strings, find the number of common characters between them.
*/
func commonCharacterCount(s1 string, s2 string) int {
	var l_sum int = 0
	var l_s1_len int = len(s1)
	var l_s2_len int = len(s2)
	s1_arr := []rune(s1)
	s2_arr := []rune(s2)

	for i := 0; i < l_s1_len; i++ {
		for j := i + 1; j < l_s1_len; j++ {
			if s1_arr[i] > s1_arr[j] {
				s1_arr[i], s1_arr[j] = s1_arr[j], s1_arr[i]
			}
		}
	}

	for i := 0; i < l_s2_len; i++ {
		for j := i + 1; j < l_s2_len; j++ {
			if s2_arr[i] > s2_arr[j] {
				s2_arr[i], s2_arr[j] = s2_arr[j], s2_arr[i]
			}
		}
	}

	for i := 0; i < l_s1_len; i++ {
		for j := 0; j < l_s2_len; j++ {
			if s1_arr[i] == s2_arr[j] {
				l_sum++
				s2_arr[j] = '\n'
				break
			}
		}
	}

	return l_sum
}

//#9
/*
	Given an array of strings, return another array containing all
	of its longest strings.
*/
func allLongestStrings(inputArray []string) []string {
	var max_size int = 0
	var first_index int = -1
	for i := 0; i < len(inputArray); i++ {
		if max_size < len(inputArray[i]) {
			max_size = len(inputArray[i])
			first_index = i
		}
	}

	slice := []string{inputArray[first_index]}
	for i := first_index + 1; i < len(inputArray); i++ {
		if len(inputArray[i]) == max_size {
			slice = append(slice, inputArray[i])
		}
	}
	return slice
}

//#8
/*
	After becoming famous, the CodeBots decided to move into a new building together.
	Each of the rooms has a different cost, and some of them are free, but there's a
	rumour that all the free rooms are haunted! Since the CodeBots are quite superstitious,
	they refuse to stay in any of the free rooms, or any of the rooms below any of
	the free rooms.

	Given matrix, a rectangular matrix of integers, where each value represents the
	cost of the room, your task is to return the total sum of all rooms that are
	suitable for the CodeBots (ie: add up all the values that don't appear below a 0).
*/
func matrixElementsSum(matrix [][]int) int {
	var l_sum int = 0
	for i := 0; i < len(matrix[0]); i++ {
		for j := 0; j < len(matrix); j++ {
			if matrix[j][i] != 0 {
				l_sum += matrix[j][i]
			} else {
				break
			}
		}
	}
	return l_sum
}

//#7
/*
	Given a sequence of integers as an array, determine whether it is possible to obtain a
	strictly increasing sequence by removing no more than one element from the array.

	Note: sequence a0, a1, ..., an is considered to be a strictly increasing if a0 < a1 < ... < an.
	Sequence containing only one element is also considered to be strictly increasing.
*/
func almostIncreasingSequence(sequence []int) bool {
	removedOne := false
	for i := 0; i < (len(sequence) - 1); i++ {
		if sequence[i] >= sequence[i+1] {
			if removedOne {
				return false
			} else {
				removedOne = true
				if i != 0 && i < len(sequence)-2 {
					if (sequence[i-1] < sequence[i+1]) || (sequence[i] < sequence[i+2]) {
						continue
					} else {
						return false
					}
				}
			}
		}
	}
	return true
}

func almostIncreasingSequence2(sequence []int) bool {
	// fmt.Println("Started",sequence)
	flag1, index1 := checkSequence(sequence)
	fmt.Printf("1st %v status %t at %d\n", sequence, flag1, index1)

	var arr_size int = len(sequence)
	if !flag1 && index1+1 <= arr_size {
		// slice = sequence_l[0:0]
		slice := make([]int, len(sequence[:index1]))
		copy(slice, sequence[:index1])
		slice = append(slice, sequence[(index1+1):]...)
		flag1, index1 = checkSequence(slice)
		fmt.Printf("2nd %v status after %t at %d\n", slice, flag1, index1)
	}
	if flag1 {
		return true
	}
	// fmt.Printf("\nhere %d of %v\n",index1,sequence)
	if !flag1 && index1-1 >= 0 {
		slice := make([]int, len(sequence[:index1-1]))
		copy(slice, sequence[:index1])
		slice = append(slice, sequence[index1:]...)
		flag1, index1 = checkSequence(slice)
		// fmt.Printf("%v status after %t at %d\n",slice,flag1,index1)
	}
	return flag1
}

func checkSequence(sequence []int) (bool, int) {
	var l_fail_index int = -1
	var arr_size int = len(sequence)

	for i := 1; i < arr_size; i++ {
		if sequence[i-1] >= sequence[i] {
			//need to investigatre whether first term is wrong or l_second
			//2 digit case
			// fmt.Printf("At %d\n",i)
			if arr_size == 2 {
				return true, i
			} else if i == arr_size-1 { /*else if arr_size==3{
					if sequence[0]>sequence[2]{return false,l_fail_index}
				}*/
				//last term can be avoided
				return false, i
			} else if sequence[i-1] <= sequence[i+1] {
				// fmt.Printf("equal case\n")
				l_fail_index = i
				break
			} else if sequence[i] <= sequence[i+1] {
				l_fail_index = i - 1
				break
			} else {
				//no future
				return false, i
			}
		}
	}
	if l_fail_index == -1 {
		return true, l_fail_index
	} else {
		// fmt.Printf("bad at %d\n",l_fail_index)
		return false, l_fail_index
	}
}

//#6
/*
	Ratiorg got statues of different sizes as a present from CodeMaster for his birthday,
	each statue having an non-negative integer size. Since he likes to make things perfect,
	he wants to arrange them from smallest to largest so that each statue will be bigger
	than the previous one exactly by 1. He may need some additional statues to be able to
	accomplish that. Help him figure out the minimum number of additional statues needed.
*/

func makeArrayConsecutive2(statues []int) int {
	var temp int = 0
	if len(statues) == 1 {
		return temp
	}
	var a_big int = statues[0]
	var a_small int = statues[0]
	for i := 1; i < len(statues); i++ {
		if a_big < statues[i] {
			a_big = statues[i]
		}
		if a_small > statues[i] {
			a_small = statues[i]
		}
	}
	fmt.Printf("max is %d small is %d\n", a_big, a_small)
	temp = a_big - a_small + 1 - len(statues)

	return temp
}

//#5 diamond (interesting polygon
func diamond(n int) {
	for h := 1; h <= n; h++ {
		for i := 1; i <= h; i++ {
			fmt.Print("*")
		}
		for i := 1; i <= h-1; i++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	for h := 1; h <= n-1; h++ {
		for i := n - h; i >= 1; i-- {
			fmt.Print("*")
		}
		for i := n - h - 1; i >= 1; i-- {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

//#3
func checkPalindrome(inputString string) bool {
	var lnth int = len(inputString)
	lnth--
	for i := 0; i <= lnth; i++ {
		if inputString[i] == inputString[lnth-i] || i == (lnth-i) {
			continue
		} else {
			return false
		}
	}
	return true
}

//#4
func adjacentElementsProduct(inputArray []int) int {
	max := 0
	temp_int := 0
	fentry := true
	for i := 0; i < len(inputArray)-1; i++ {
		temp_int = inputArray[i] * inputArray[i+1]
		if fentry {
			max = temp_int
			fentry = false
		} else if temp_int > max {
			max = temp_int
		}
	}
	return max
}
