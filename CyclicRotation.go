package solution

func Solution(A []int, K int) []int {
    length := len(A)
    if length > 0 {
        if K > length{
            K = K%length
        }
        A = append(A[length-K:length],A[0:length-K]...)
    }
    return A
}
