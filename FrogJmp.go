package solution

func Solution(X int, Y int, D int) int {
    result := (Y-X)/D
    if (Y-X)%D > 0 {
        result++
    }
    return result
}
