package numbers

import "math"

func IsPrime(num int) bool{
	for i:=2 ; i<= int(math.Floor(math.Sqrt(float64(num)))) ;i++ {
		if num % i ==0{
		   return false
		}

	}
	
   return num>1
}
