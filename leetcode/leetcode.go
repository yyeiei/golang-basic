package main

import (
	"fmt"
	"strings"
)

func longestValidParentheses1() {
	fmt.Println(longestValidParentheses("(()"))            // () 2
	fmt.Println(longestValidParentheses(")()())"))         // ()() 4
	fmt.Println(longestValidParentheses(")(())))(())())")) // (())() 6
	fmt.Println(longestValidParentheses(")()())"))         // ()() 4
	fmt.Println(longestValidParentheses("()(())"))         // ()(()) 6
}

// 32--
func longestValidParentheses(s string) int {
	maxans := 0
	s = strings.TrimRight(strings.TrimLeft(s, ")"), "(")
	dp := make([]int, len(s))
	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				if i >= 2 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			} else if i-dp[i-1] > 0 && s[i-dp[i-1]-1] == '(' {
				dp[i] = dp[i-1] + 2
				if (i - dp[i-1]) >= 2 {
					dp[i] = dp[i] + dp[i-dp[i-1]-2]
				}
			}
			if dp[i] > maxans {
				maxans = dp[i]
			}
		}
	}
	return maxans

	// 非连续
	// s = strings.TrimLeft(s, ")")
	// s = strings.TrimRight(s, "(")
	// left, right := 0, 0
	// for i := 0; i < len(s); i++ {
	// 	if s[i] == '(' {
	// 		left++
	// 	} else {
	// 		right++
	// 	}
	// }
	// if left < right {
	// 	return left * 2
	// }

	// return right * 2
}

func findClosestElements1() {
	fmt.Println(findClosestElements([]int{0, 1, 2, 2, 2, 3, 6, 8, 8, 9}, 5, 9))
	fmt.Println(findClosestElements([]int{0, 2, 2, 3, 4, 6, 7, 8, 9, 9}, 4, 5))
	fmt.Println(findClosestElements([]int{0, 0, 1, 2, 3, 3, 4, 7, 7, 8}, 3, 5))
	fmt.Println(findClosestElements([]int{1, 2, 3, 4, 5}, 4, 3))
	fmt.Println(findClosestElements([]int{1, 2, 3, 4, 5}, 4, -1))
	fmt.Println(findClosestElements([]int{0, 1, 1, 1, 2, 3, 6, 7, 8, 9}, 9, 4))
	fmt.Println(findClosestElements(
		[]int{-9978, -9965, -9894, -9887, -9877, -9861, -9861, -9860, -9859, -9798, -9795, -9786, -9773, -9766, -9732, -9726, -9686, -9628, -9622, -9600, -9596, -9575, -9554, -9546, -9508, -9505, -9502, -9489, -9455, -9402, -9385, -9327, -9312, -9281, -9280, -9253, -9208, -9207, -9167, -9163, -9161, -9161, -9160, -9153, -9115, -9103, -9054, -9024, -9019, -9001, -8999, -8966, -8947, -8932, -8899, -8892, -8878, -8861, -8846, -8796, -8783, -8745, -8688, -8664, -8638, -8636, -8616, -8598, -8586, -8574, -8559, -8558, -8555, -8510, -8505, -8501, -8476, -8426, -8412, -8371, -8339, -8319, -8274, -8262, -8253, -8231, -8151, -8150, -8147, -8081, -8078, -8063, -8027, -8000, -7994, -7986, -7963, -7873, -7871, -7864, -7856, -7811, -7798, -7744, -7731, -7718, -7668, -7642, -7635, -7629, -7594, -7585, -7584, -7494, -7448, -7425, -7408, -7386, -7365, -7331, -7329, -7323, -7307, -7303, -7270, -7251, -7247, -7231, -7226, -7197, -7168, -7161, -7158, -7124, -7108, -7088, -7064, -7043, -7018, -7000, -6999, -6985, -6983, -6967, -6963, -6962, -6957, -6913, -6903, -6894, -6882, -6872, -6869, -6843, -6838, -6837, -6832, -6824, -6780, -6755, -6753, -6749, -6705, -6692, -6670, -6657, -6654, -6609, -6596, -6565, -6541, -6539, -6523, -6516, -6435, -6424, -6418, -6410, -6351, -6324, -6302, -6299, -6282, -6217, -6202, -6162, -6138, -6121, -6104, -6048, -6029, -5992, -5991, -5984, -5963, -5941, -5936, -5920, -5916, -5891, -5882, -5878, -5842, -5825, -5811, -5796, -5760, -5754, -5754, -5718, -5706, -5687, -5685, -5667, -5654, -5611, -5592, -5585, -5547, -5534, -5529, -5522, -5520, -5445, -5421, -5405, -5400, -5391, -5385, -5365, -5353, -5348, -5336, -5315, -5315, -5246, -5186, -5178, -5130, -5115, -5069, -5042, -5030, -4978, -4962, -4953, -4951, -4947, -4945, -4940, -4917, -4909, -4898, -4895, -4875, -4863, -4820, -4818, -4814, -4783, -4758, -4751, -4732, -4717, -4703, -4692, -4683, -4654, -4648, -4645, -4638, -4614, -4590, -4555, -4553, -4527, -4518, -4515, -4509, -4502, -4472, -4470, -4467, -4424, -4391, -4385, -4383, -4375, -4374, -4367, -4332, -4305, -4291, -4262, -4246, -4243, -4209, -4187, -4165, -4152, -4140, -4139, -4128, -4082, -4073, -4064, -4054, -4031, -4003, -3987, -3986, -3984, -3945, -3939, -3906, -3880, -3852, -3839, -3835, -3820, -3812, -3811, -3802, -3745, -3730, -3727, -3710, -3709, -3671, -3663, -3657, -3629, -3602, -3600, -3592, -3533, -3459, -3436, -3410, -3398, -3368, -3367, -3347, -3342, -3326, -3320, -3316, -3293, -3278, -3271, -3259, -3256, -3246, -3177, -3175, -3106, -3100, -3094, -3091, -3088, -3082, -3080, -3050, -3033, -2990, -2977, -2934, -2895, -2879, -2860, -2854, -2847, -2844, -2842, -2819, -2802, -2736, -2686, -2679, -2671, -2644, -2635, -2604, -2597, -2571, -2562, -2556, -2528, -2493, -2491, -2405, -2384, -2357, -2349, -2334, -2313, -2279, -2219, -2208, -2173, -2155, -2134, -2124, -2111, -2101, -2082, -2040, -1989, -1981, -1956, -1921, -1890, -1885, -1822, -1787, -1786, -1774, -1740, -1698, -1673, -1666, -1632, -1616, -1615, -1611, -1594, -1551, -1548, -1477, -1473, -1448, -1321, -1296, -1276, -1255, -1255, -1247, -1242, -1211, -1195, -1172, -1131, -1107, -1100, -1081, -1076, -1064, -1044, -1032, -1022, -993, -990, -945, -940, -924, -909, -895, -883, -867, -860, -857, -849, -845, -822, -807, -803, -793, -756, -753, -714, -652, -629, -603, -598, -593, -578, -526, -483, -477, -458, -455, -430, -407, -401, -378, -316, -310, -309, -296, -270, -246, -223, -212, -195, -96, -45, -35, -1, 28, 34, 39, 66, 86, 109, 119, 145, 171, 174, 197, 212, 229, 286, 291, 301, 354, 355, 358, 399, 404, 410, 418, 427, 455, 460, 535, 575, 593, 595, 604, 638, 643, 646, 652, 692, 712, 719, 768, 796, 803, 825, 874, 886, 903, 945, 948, 1025, 1042, 1046, 1056, 1071, 1085, 1095, 1136, 1143, 1150, 1153, 1163, 1183, 1189, 1199, 1208, 1226, 1243, 1311, 1319, 1366, 1387, 1440, 1457, 1466, 1485, 1498, 1531, 1535, 1541, 1549, 1584, 1686, 1719, 1737, 1762, 1793, 1811, 1820, 1839, 1878, 1884, 1900, 1913, 1948, 1964, 1966, 1988, 2007, 2009, 2034, 2048, 2048, 2097, 2115, 2133, 2155, 2159, 2198, 2254, 2271, 2299, 2299, 2317, 2321, 2338, 2363, 2394, 2397, 2407, 2444, 2566, 2578, 2591, 2604, 2659, 2669, 2691, 2698, 2709, 2732, 2750, 2762, 2776, 2798, 2826, 2833, 2839, 2841, 2869, 2872, 2874, 2943, 3014, 3027, 3049, 3053, 3065, 3106, 3106, 3117, 3183, 3187, 3191, 3229, 3237, 3239, 3265, 3288, 3326, 3347, 3370, 3380, 3423, 3438, 3439, 3463, 3494, 3501, 3507, 3508, 3545, 3574, 3576, 3586, 3593, 3597, 3599, 3654, 3666, 3674, 3725, 3726, 3741, 3750, 3760, 3767, 3772, 3785, 3838, 3854, 3879, 3889, 3894, 3919, 3931, 3937, 3943, 3951, 4007, 4023, 4033, 4088, 4123, 4123, 4209, 4223, 4256, 4269, 4300, 4317, 4320, 4331, 4336, 4337, 4340, 4342, 4342, 4351, 4355, 4366, 4371, 4382, 4396, 4411, 4434, 4442, 4459, 4517, 4537, 4544, 4545, 4557, 4574, 4584, 4649, 4654, 4713, 4807, 4843, 4848, 4860, 4864, 4893, 4984, 5018, 5025, 5044, 5069, 5073, 5078, 5079, 5088, 5095, 5155, 5170, 5187, 5203, 5215, 5230, 5242, 5254, 5262, 5311, 5332, 5334, 5342, 5356, 5389, 5417, 5418, 5422, 5422, 5431, 5444, 5444, 5466, 5500, 5509, 5549, 5572, 5601, 5608, 5609, 5632, 5634, 5635, 5636, 5653, 5654, 5665, 5675, 5703, 5723, 5748, 5771, 5796, 5805, 5816, 5861, 5862, 5871, 5879, 5881, 5909, 5914, 5919, 5987, 5988, 6012, 6065, 6096, 6098, 6109, 6113, 6119, 6125, 6198, 6239, 6242, 6246, 6299, 6300, 6325, 6328, 6331, 6343, 6344, 6357, 6361, 6385, 6399, 6423, 6463, 6484, 6487, 6488, 6573, 6613, 6615, 6621, 6623, 6625, 6641, 6656, 6696, 6708, 6755, 6825, 6825, 6863, 6897, 6925, 6939, 6974, 6996, 7014, 7036, 7039, 7121, 7128, 7143, 7180, 7198, 7204, 7226, 7245, 7253, 7276, 7284, 7324, 7342, 7351, 7391, 7392, 7432, 7459, 7472, 7474, 7478, 7480, 7502, 7551, 7572, 7594, 7617, 7622, 7633, 7646, 7664, 7681, 7691, 7711, 7717, 7729, 7742, 7742, 7745, 7770, 7776, 7782, 7804, 7828, 7840, 7841, 7856, 7865, 7897, 7916, 7938, 7950, 7978, 7994, 7994, 7995, 8004, 8005, 8011, 8117, 8191, 8197, 8232, 8244, 8289, 8290, 8293, 8321, 8323, 8329, 8381, 8383, 8423, 8478, 8486, 8494, 8524, 8524, 8550, 8552, 8570, 8573, 8597, 8597, 8604, 8642, 8670, 8693, 8710, 8730, 8753, 8815, 8829, 8877, 8961, 8968, 8994, 9018, 9025, 9035, 9038, 9040, 9046, 9050, 9158, 9232, 9251, 9254, 9268, 9273, 9286, 9309, 9327, 9397, 9430, 9449, 9457, 9487, 9508, 9523, 9539, 9557, 9571, 9589, 9652, 9722, 9799, 9803, 9816, 9820, 9824, 9826, 9835, 9839, 9860, 9869, 9877, 9897, 9911, 9922, 9922, 9927, 9948, 9948, 9967, 9987},
		936, 7022))
}

// 658--
func findClosestElements(arr []int, k int, x int) []int {
	begin, end := 0, len(arr)-1
	mid := -1
	for begin < end && end-begin > k-1 {
		mid = begin + (end-begin)/2
		if arr[mid] == x || (mid > 0 && mid < len(arr)-1 && x > arr[mid-1] && x < arr[mid+1]) {
			break
		}
		if x > arr[mid] {
			begin = mid
		} else {
			end = mid
		}
	}

	begin = mid - k + 1
	end = mid + k - 1
	if begin < 0 {
		begin = 0
	}
	if end > len(arr)-1 {
		end = len(arr) - 1
	}
	for end-begin > k-1 {
		if x-arr[begin] <= arr[end]-x {
			end--
		} else {
			begin++
		}
	}

	return arr[begin : begin+k]

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}
