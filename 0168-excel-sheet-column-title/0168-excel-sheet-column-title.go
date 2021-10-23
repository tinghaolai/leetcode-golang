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

func extract_2(n int, title string) string {
    for n > 26 {
        divNum := (n - 1) / 26
        n = (n - 1) % 26 + 1
        title = extract(divNum, title)
    }

    if n < 1 {
        return title
    }

    return title + convertAsc_2(n - 1)
}


func convertAsc_2(value int) string {
    return string(rune('A' + value))
}
