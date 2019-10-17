package utils

// ToUnderScore string, XxYy to xx_yy , XxYY to xx_yy
func ToUnderScore(str string) string {
	data := make([]rune, 0, len(str)*2)
	for k, v := range str {
		if k > 0 && v >= 65 && v <= 90 && str[k-1] >= 97 && str[k-1] <= 122  {
			data = append(data, '_')
		}
		if v >= 65 && v <= 90 {
			data = append(data, v+32)
		} else  {
			data = append(data, v)
		}
	}
	return string(data[:])
}

// ToCamelCase string, xx_yy to XxYy
func ToCamelCase(str string) string {
	data := make([]rune, 0, len(str))
	for k,v := range str {
		if v == 95 {
			continue
		}
		if (k==0 || (k > 0 && str[k-1] == 95)) && (v >= 97 && v <= 122) {
			data = append(data, v-32)
		} else {
			data = append(data, v)
		}
	}
	return string(data[:])
}
