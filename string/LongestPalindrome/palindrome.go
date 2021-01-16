/*


 */

package LongestPalindrome

func LongestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}

	max := 1
	maxStr := s[0:1]
	rear := len(s) - 1

	for start := 0; start < rear; start++ {
		for end := rear; end-start+1 > max; end-- {
			i := start
			j := end
			for i < j {
				if s[i] == s[j] {
					i++
					j--
				} else {
					break
				}
			}

			if i >= j {
				max = end - start + 1
				maxStr = s[start : end+1]
			}
		}
	}
	return maxStr
}
