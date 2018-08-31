package solution

func Solution(A []int) int {
    var elmt int
    for i:=0;i<len(A);i++{
        elmt ^= A[i]
    }
    return elmt
}
