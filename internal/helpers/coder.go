// функция реализует кодирование
// на входе - натуральное число
// на выходе - циферно-буквенная последовательность

package helpers

import (	
	"strings"
)

const base62Chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func Base62Encode(n uint64) string {

    if n == 0 {
        return "0"
    }

    var sb strings.Builder

    for n > 0 {
        sb.WriteByte(base62Chars[n%62])
        n /= 62
    }
	
    // переворачиваем, потому что строили с конца

	b := sb.String()
	bytes := []byte(b)

    for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
        bytes[i], bytes[j] = bytes[j], bytes[i]
    }

    return string(bytes)
}