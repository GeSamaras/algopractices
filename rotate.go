func rotatedDigits(n int) int {
    dp:=make([]int,n+1)
    valid:=0
    for i:=0;i<=n;i++{
        if i/10!=0{
            if dp[i%10]==0 || dp[i/10]==0 {
                continue   
            }else if dp[i%10]==-1 && dp[i/10]==-1{
                dp[i]=-1
                continue
            }else{
                dp[i]=1
                valid++  
            }
        }else{
            res:=validDigit(i)
            if res==1{
                valid++
            }
            dp[i]=res
        }
    }
    return valid
}

func validDigit(digit int) int{
    switch digit{
        case 2,5,6,9:
            return 1
        case 0,1,8:
            return -1
        default:
            return 0
    }
}
