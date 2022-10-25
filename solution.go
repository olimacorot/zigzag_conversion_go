package zigzag_conversion

import (
    "strings"
)

func convert(s string, numRows int) string {

    if numRows == 1{
        return s
    } 

    total := 0
    position := 0
    increment := false
    zigzag := make([]string, numRows)

    for total < len(s) {
        zigzag[position] = zigzag[position] + string(s[total])

        if position == 0 || position >= (numRows-1) {
            increment = !increment
        }

        if increment {
            position += 1
        } else {
            position -= 1
        }

        total += 1
    }

    return strings.Join(zigzag, "")
}
