func isValid(s string) bool {
    type Stack struct {
        items []rune
    }

    stack := Stack{}

    push := func(ch rune) {
        stack.items = append(stack.items, ch)
    }

    pop := func() rune {
        n := len(stack.items)
        ch := stack.items[n-1]
        stack.items = stack.items[:n-1]
        return ch
    }

    isEmpty := func() bool {
        return len(stack.items) == 0
    }

    pairs := map[rune]rune{
        ')' : '(',
        ']' : '[',
        '}' : '{',
    }

for _, ch := range s {
    if open, ok := pairs[ch]; ok {
        if isEmpty() || pop() != open {
            return false
        }
    } else {
        push (ch)
    }
}
return isEmpty()

}
