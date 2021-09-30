package main

import "fmt"

func main()  {
	a :=[5]string{"I","am","stupid","and","weak"}
	for k,_ := range  a{
		if k ==2{
			a[2]="smart"
		}
		if k ==4{
			a[4]="strong"
		}
	}
	fmt.Println(a)
}
