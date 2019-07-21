package week3

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// Input
	//fmt.Print("Enter series of integer numbers: ")
	//r := bufio.NewReader(os.Stdin)
	//line, err := r.ReadString('\n')
	//if err != nil {
	//	log.Fatal(err)
	//}
	//numbers, err := parseInts(line)
	numbers := []int{576, 949, 829, 3, 417, 607, 240, 350, 250, 900, 368, 333, 642, 136, 433, 840, 373, 899, 275, 913, 717, 587, 524, 681, 818, 307, 424, 297, 702, 345, 516, 334, 674, 479, 664, 878, 960, 735, 945, 202, 389, 282, 867, 490, 606, 683, 746, 146, 288, 18, 981, 237, 187, 961, 441, 341, 200, 572, 74, 652, 863, 989, 532, 938, 379, 406, 110, 701, 898, 387, 73, 590, 536, 703, 991, 473, 975, 603, 875, 13, 650, 332, 716, 308, 694, 207, 123, 94, 364, 567, 670, 244, 124, 751, 722, 725, 359, 528, 499, 604, 69, 993, 59, 495, 846, 854, 634, 24, 378, 727, 755, 552, 348, 963, 90, 665, 710, 218, 305, 940, 367, 580, 338, 853, 737, 190, 861, 85, 964, 699, 720, 169, 589, 398, 402, 298, 847, 14, 258, 779, 833, 347, 882, 979, 659, 775, 914, 155, 189, 677, 886, 252, 491, 687, 507, 238, 4, 813, 542, 831, 644, 227, 760, 176, 481, 232, 353, 537, 488, 804, 632, 159, 286, 764, 51, 675, 178, 444, 753, 932, 806, 705, 31, 65, 15, 35, 91, 130, 885, 120, 944, 284, 817, 214, 540, 917, 645, 342, 142, 556, 502, 467, 635, 802, 789, 647, 680, 713, 170, 62, 612, 763, 568, 521, 750, 290, 897, 98, 500, 682, 254, 543, 76, 692, 757, 462, 698, 783, 380, 439, 761, 721, 476, 177, 445, 937, 245, 771, 357, 165, 892, 483, 711, 1, 890, 195, 153, 821, 145, 431, 93, 592, 707, 855, 230, 301, 432, 226, 619, 557, 978, 747, 291, 143, 796, 25, 420, 549, 180, 734, 578, 57, 982, 183, 38, 822, 118, 504, 443, 869, 236, 182, 844, 484, 954, 405, 538, 196, 112, 261, 381, 947, 464, 21, 693, 995, 194, 893, 436, 974, 151, 653, 871, 164, 668, 630, 651, 594, 663, 905, 53, 828, 72, 732, 598, 518, 998, 983, 437, 208, 6, 824, 107, 229, 172, 622, 456, 283, 346, 106, 584, 501, 87, 616, 382, 219, 421, 265, 304, 27, 733, 404, 127, 75, 588, 626, 648, 125, 7, 111, 786, 534, 494, 866, 383, 263, 427, 643, 5, 767, 724, 908, 309, 314, 412, 188, 391, 241, 748, 704, 319, 709, 759, 697, 140, 919, 874, 531, 950, 715, 527, 137, 372, 999, 686, 166, 480, 46, 843, 26, 181, 133, 729, 374, 167, 600, 595, 129, 888, 134, 400, 396, 211, 411, 158, 996, 361, 144, 33, 957, 515, 928, 561, 312, 790, 655, 517, 458, 770, 797, 68, 766, 135, 174, 492, 386, 9, 881, 80, 40, 470, 510, 809, 198, 88, 615, 103, 474, 930, 777, 280, 669, 442, 459, 916, 17, 132, 102, 45, 104, 726, 339, 426, 184, 551, 649, 738, 859, 148, 887, 920, 204, 691, 119, 526, 407, 448, 422, 912, 901, 318, 657, 306, 248, 52, 931, 925, 618, 700, 54, 503, 47, 86, 673, 139, 520, 429, 826, 29, 108, 918, 213, 689, 78, 67, 413, 116, 81, 83, 952, 676, 355, 929, 625, 902, 397, 776, 313, 344, 468, 640, 331, 523, 638, 49, 303, 834, 451, 641, 879, 575, 255, 922, 482, 939, 302, 156, 370, 191, 762, 115, 839, 970, 736, 316, 43, 393, 780, 247, 388, 97, 326, 157, 122, 419, 564, 596, 246, 105, 371, 477, 621, 740, 55, 60, 714, 10, 325, 508, 505, 239, 581, 8, 934, 392, 216, 825, 41, 138, 781, 32, 794, 706, 921, 425, 210, 718, 672, 744, 941, 559, 360, 623, 586, 832, 936, 447, 984, 948, 530, 435, 117, 816, 434, 438, 79, 636, 896, 161, 20, 942, 573, 328, 233, 154, 550, 64, 571, 773, 324, 39, 212, 466, 830, 453, 12, 798, 463, 953, 82, 173, 163, 679, 842, 509, 168, 719, 849, 268, 812, 512, 289, 795, 608, 601, 343, 994, 472, 617, 891, 909, 28, 541, 749, 327, 730, 548, 42, 756, 880, 973, 856, 574, 966, 394, 201, 384, 583, 440, 745, 915, 579, 369, 375, 160, 923, 678, 570, 810, 889, 56, 34, 100, 819, 857, 478, 958, 955, 609, 723, 903, 403, 274, 287, 637, 731, 741, 506, 452, 639, 992, 662, 96, 1000, 454, 186, 836, 77, 876, 712, 311, 203, 986, 11, 837, 260, 791, 300, 563, 66, 959, 852, 965, 956, 365, 562, 267, 251, 800, 967, 613, 366, 904, 61, 633, 827, 401, 907, 801, 71, 48, 988, 685, 690, 317, 555, 23, 602, 883, 294, 457, 614, 951, 242, 807, 987, 792, 243, 788, 206, 293, 266, 315, 884, 742, 985, 205, 235, 895, 50, 628, 752, 565, 498, 624, 646, 299, 253, 971, 497, 769, 320, 894, 330, 525, 872, 408, 113, 845, 362, 708, 193, 569, 728, 851, 296, 399, 620, 271, 946, 149, 696, 269, 585, 910, 877, 911, 415, 610, 390, 969, 793, 860, 820, 37, 519, 553, 943, 356, 418, 449, 599, 593, 684, 228, 785, 84, 395, 179, 270, 695, 423, 805, 656, 838, 976, 257, 547, 487, 924, 935, 841, 627, 95, 109, 906, 223, 591, 450, 968, 171, 469, 175, 217, 546, 862, 926, 234, 848, 414, 262, 2, 220, 754, 352, 511, 803, 141, 496, 349, 340, 977, 128, 121, 765, 249, 385, 743, 358, 671, 545, 44, 455, 460, 336, 582, 858, 337, 22, 351, 197, 814, 224, 70, 471, 808, 799, 209, 772, 231, 114, 811, 285, 605, 782, 416, 126, 661, 980, 279, 864, 321, 152, 36, 554, 363, 461, 870, 544, 410, 778, 430, 221, 868, 514, 513, 774, 222, 631, 147, 629, 758, 99, 63, 199, 865, 329, 354, 667, 409, 475, 560, 688, 823, 558, 185, 377, 539, 259, 972, 768, 256, 835, 739, 225, 30, 850, 927, 493, 446, 192, 658, 522, 292, 376, 486, 997, 295, 654, 323, 962, 89, 815, 787, 990, 322, 529, 666, 277, 577, 485, 278, 784, 428, 272, 273, 335, 16, 933, 535, 660, 131, 276, 162, 19, 215, 566, 264, 92, 101, 873, 465, 150, 533, 58, 489, 597, 281, 611, 310}
	//if err != nil {
	//	log.Fatal(err)
	//}
	// Create a channel
	c := make(chan []int)
	m := make(chan []int)
	resultChan := make(chan []int)

	// Divide
	numbersLen := len(numbers)
	divideSlice := numbersLen / 4
	index := 0
	for i := 0; i < 3; i++ {
		go sortSubArray(numbers[index:index+divideSlice], c)
		index = index + divideSlice
	}
	go sortSubArray(numbers[index:], c)

	sub1 := <-c
	sub2 := <-c
	go merge(sub1, sub2, m)

	sub3 := <-c
	sub4 := <-c
	go merge(sub3, sub4, m)

	go merge(<-m, <-m, resultChan)
	fmt.Println("Sort result: ", <-resultChan)
}

func parseInts(s string) ([]int, error) {
	var (
		fields  = strings.Fields(s)
		numbers = make([]int, len(fields))
		err     error
	)
	for i, f := range fields {
		numbers[i], err = strconv.Atoi(f)
		if err != nil {
			return nil, err
		}
	}
	return numbers, nil
}

func sortSubArray(slice []int, c chan []int) {
	fmt.Println("Will sorting following subarray: ", slice)
	sort.Ints(slice)
	c <- slice
}

func merge(left, right []int, m chan []int) {
	//make an []int of size of length of left + length of right
	result := make([]int, len(left)+len(right))

	for i := 0; len(left) > 0 || len(right) > 0; i++ {

		if len(left) > 0 && len(right) > 0 {
			if left[0] < right[0] {
				result[i] = left[0]
				left = left[1:]
			} else {
				result[i] = right[0]
				right = right[1:]
			}
		} else if len(left) > 0 {
			result[i] = left[0]
			left = left[1:]
		} else if len(right) > 0 {
			result[i] = right[0]
			right = right[1:]
		}
	}
	m <- result
}