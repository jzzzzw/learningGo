package model
import(
	"fmt"
)

type Account struct{
	No string
	balance float64
	password string
}
func (a *Account)Setno(str string){
	if (len(str)<6||len(str)>10){
		fmt.Printf("invalid length of Account no.\n")
	}else{
		a.No=str
	}

}
func (a *Account)Setbal(str float64){
	if (str<20){
		fmt.Printf("invalid number of Account balance.\n")
	}else{
		a.balance=str
	}

}
func (a *Account)Setpwd(str string){
	if (len(str)<6||len(str)>6){
		fmt.Printf("invalid length of Account password.\n")
	}else{
		a.password=str
	}
}
func NewAccount() *Account{
	return &Account {

	}
}