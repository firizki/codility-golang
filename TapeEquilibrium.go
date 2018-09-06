package solution

func Solution(A []int) int {
    var fix_sum, cur_sum, res_temp, val_return int
    for i:=0;i<len(A);i++{
        fix_sum += A[i]
        if A[i] < 0 {
            val_return += ^A[i]+1
        } else {
            val_return += A[i]
        }
    }
    for i:=0;i<len(A)-1;i++{
        cur_sum += A[i]
        res_temp = cur_sum - (fix_sum - cur_sum)
        if res_temp < 0 {
            res_temp = ^res_temp+1
        }
        if res_temp < val_return {
            val_return = res_temp
        }
    }

    return val_return
}
