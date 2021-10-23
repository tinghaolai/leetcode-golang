package leetcode

func convertToTitle(n int) string {
    return extract(n, "")
}

func extract(n int, title string) string {
    for n > 26 {
        divNum := n / 26
        n %= 26
        if n == 0 {
            n = 26
            divNum --
        }

        title = extract(divNum, title)
    }

    if n < 1 {
        return title
    }

    return title + convertAsc(n)
}


func convertAsc(value int) string {
    return string(rune('A' - 1 + value))
}
