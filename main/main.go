package main
import (
	"fmt"
	"math/rand"
	"math"
	"strings"
	"learningGo/model"
	"time"
	_"sort"
	"strconv"
	"reflect"
	"os"
	"bufio"
	"io"
	_"io/ioutil"
	_"os"
	_"flag"
	_"encoding/json"
	_"runtime"
	"sync"
)
func main(){
	// var str  string
	// fmt.Printf("please input: ..\n")
	// fmt.Scanf("%s\n",&str)
	// switch strings.ToUpper(str){
	// 	case "A","B","C","D","E","F","G":
	// 		fmt.Printf("the result is %s\n",strings.ToUpper(str))
	// 	default:
	// 		fmt.Printf("nothing output..\n")
	// }
	// var i,j int
	// for i =1 ; i<10 ;i++{
	// 	for j =1;j<=i;j++{
	// 		fmt.Printf("%d * %d = %d\t",j,i,j*i)
	// 	}
	// 	fmt.Printf("\n")
	// }
	// var i,a,b,c int
	// fmt.Printf("input a number...\n")
	// fmt.Scanf("%d\n",&i)
	// a = i/100
	// b = i%100/10
	// c = i%10
	// fmt.Printf("a=%d,b=%d,c=%d\n",a,b,c)
	// if (i==a*a*a+b*b*b+c*c*c){
	// 	fmt.Printf("ok\n")
	// }else {
	// 	fmt.Printf("well..\n")
	// }
	/* var a,b string
		fmt.Printf("input your ...\n")
		fmt.Scanf("%s\n",&a)
		fmt.Printf("input your password...\n")
		fmt.Scanf("%s\n",&b)
		if(a=="张三" && b=="1234"){
				fmt.Printf("ok\n")
			}else {
				fmt.Printf("well..\n")
			} */
	// var a int
	// fmt.Printf("input your month...\n")
	// fmt.Scanf("%d\n",&a)
	// switch a{
	// case 1,3,5,7,8,10,12:
	// 	fmt.Printf("31days\n")
	// case 2:
	// 	fmt.Printf("28 or 29 days\n")
	// default:
	// 	fmt.Printf("30days\n")
	// }
	// var hgt,wgt int
	// fmt.Printf("input your weight...\n")
	// fmt.Scanf("%d\n",&wgt)
	// fmt.Printf("input your height...\n")
	// fmt.Scanf("%d\n",&hgt)
	// if(math.Abs(float64((hgt-108)*2-wgt))<=10.0){
	// 	fmt.Printf("%f ,OK\n",float64((hgt-108)*2-wgt))
	// }else {
	// 	fmt.Printf("%f, not OK\n",float64((hgt-108)*2-wgt))
	// }
//    var i ,j int 
		
// 		for{
// 		if(i!=95){
// 			i = rand.Intn(100)+1 //Intn前闭后开
// 			j++
// 			fmt.Printf("%d %d\n",i,j)
// 		}else
// 		{
// 			break
// 		}
		
// 	}
// 		fmt.Printf("%d %d\n",i,j)
	// fmt.Println("ok1")
	// goto label1
	// fmt.Println("ok2")
	// label1:
	// fmt.Println("ok3")
	// var a,b int
	// var oper string
	// var res int 
	// fmt.Printf("imput a number..\n")
	// fmt.Scanf("%d\n",&a)
	// fmt.Printf("imput a number..\n")
	// fmt.Scanf("%d\n",&b)
	// fmt.Printf("imput a operator..\n")
	// fmt.Scanf("%s\n",&oper)
	// res = calucate(oper,a,b)
	// fmt.Printf("%d %s %d ,the result is %d..\n",a,oper,b,res)
	// test(4)
	// var a,b int
	// a=20
	// b=10
	// // swap(&a,&b)
	// a,b=swap(a,b)
	// fmt.Printf("%d %d\n",a,b)
	// res :=func(n1 int,n2 int)int{
	// 	return  n1+n2
	// }(10,20)
	
	// fmt.Printf("res is %d\n",res)
	// a := func (n1 int,n2 int)int{
	// 	return n1+n2
	// }
	// b:= a(10,20)
	// fmt.Printf("res is %T  %d\n",a,b)
	// var str ,suf string
	// fmt.Printf("input a string pls..\n")
	// fmt.Scanf("%s\n",&str)
	// fmt.Printf("input a suffix pls..\n")
	// fmt.Scanf("%s\n",&suf)
	// f := makeSuffix(suf);
	// fmt.Printf("%v\n",f(str))
	//sum(10,20)
	// str := strings.Split("1|2|3|4","|")
	// fmt.Printf("%v %T\n",str,str)
	// for i:=0;i<len(str);i++{
	// 	fmt.Printf("str[%d] is %v\n",i,str[i])
	// }
	// var a ,b ,c int 
	// for{
	// 	fmt.Printf("pls input year ..\n")
	// 	fmt.Scanf("%d\n",&a)
	// 	fmt.Printf("pls input month ..\n")
	// 	fmt.Scanf("%d\n",&b)
	// 	if(b <=0|| b>=13){
	// 		fmt.Printf("month input error..\n")
	// 		continue
	// 	}
	// 	fmt.Printf("pls input day ..\n")
	// 	fmt.Scanf("%d\n",&c)
	// 	if (checkDay(a,b,c)==true){
	// 		fmt.Printf("the date is %d-%d-%d\n",a,b,c)
	// 		break
	// 	}else{
	// 		fmt.Printf("day input error..\n")
	// 		continue
	// 	}

	// }
	rand.Seed(time.Now().UnixNano())
	// var guess ,i int	
	// a := rand.Intn(10)
	// i =1
	// fmt.Printf("let's start guess 1 to 10..\n")
	// for {
	// 	fmt.Scanf("%d\n",&guess)
	// 	if(guess == a){
	// 		fmt.Printf("Got it...\n")
	// 		break;
	// 	}else{
	// 		fmt.Printf("Wrong. Again pls(current count %d)...\n",i+1)
	// 	}
	// 	if(i==10){
	// 		break
	// 	}else{
	// 		i++
	// 		continue
	// 	}
	// }
	// switch i{
	// case 1:
	// 	fmt.Printf("Geninus...\n")
	// case 2,3:
	// 	fmt.Printf("Clever...\n")
	// case 4,5,6,7,8,9:
	// 	fmt.Printf("Just so so...\n")
	// case 10:
	// 	fmt.Printf("Ok...\n")
	// default:
	// 	fmt.Printf("emmmm...\n")
	// }
	// fmt.Printf("the num is %d,your count is %d\n",a,i)
	// sum :=0
	// count :=1
	// for i:=1;i<=100;i++{
	// 	if(checkPrime(i) ==true){
	// 		fmt.Printf("%d \t",i)
	// 		count++
	// 		sum+=i
	// 	}
	// 	if(count%6==0){
	// 		count = 1
	// 		fmt.Printf("\n")
	// 	}
	// }
	// fmt.Printf("\nsum is %d\n",sum)
	// var str string
	// var daydiff int64
	// label1:
	// fmt.Printf("input a date afer 1990-01-01..\n")
	// fmt.Scanf("%v\n",&str)
	// dayinput,_ := time.ParseInLocation("2006-01-02 15:04:05",str+" 00:00:00",time.Local)
	// stdday,_ :=time.ParseInLocation("2006-01-02 15:04:05","1990-01-01 00:00:00",time.Local)
	// daydiff = (dayinput.Unix() - stdday.Unix())/60/60/24
	// //fmt.Printf("input day is %v, standard day is %v, day diff is %d\n",dayinput,stdday,daydiff)
	// fmt.Printf("input day is %v, standard day is %v, day diff is %d\n",dayinput.Format("2006-01-02"),stdday.Format("2006-01-02"),daydiff)
	// if(daydiff<0){
	// 	fmt.Printf("date input error. Again pls..\n")
	// 	goto label1
	// }else{
	// 	switch daydiff%5{
	// 	case 0,1,2:
	// 		fmt.Printf("fish day\n")
	// 	case 3,4:
	// 		fmt.Printf("no fish day\n")
	// 	}
	// }
	// fmt.Printf("------小小计算器------\n")
	// fmt.Printf("1.plus\n")
	// fmt.Printf("2.minus\n")
	// fmt.Printf("3.multi\n")
	// fmt.Printf("4.divide\n")
	// fmt.Printf("0.quit\n")
	// fmt.Printf("pls choose...\n")
	// var a int
	// fmt.Scanf("%d\n",&a)
	// switch a{
	// case 1:
	// 	fmt.Printf("10+5=%d\n",10+5)
	// case 2:
	// 	fmt.Printf("10-5=%d\n",10-5)
	// case 3:
	// 	fmt.Printf("10*5=%d\n",10*5)
	// case 4:
	// 	fmt.Printf("10/5=%d\n",10/5)
	// case 0:
	// 	break;
	// }
	// var a int
	// a='a'
	// for i:=1;i<27;i++{
	// 	fmt.Printf("%c ",a+i-1)
	// }
	// fmt.Printf("\n")
	// a='Z'
	// for j:=1;j<27;j++{
	// 	fmt.Printf("%c ",a-j+1)
	// }
	// fmt.Printf("\n")
	// var score[5] float64
	//  for i:=0;i<len(score);i++{
	// 	 fmt.Printf("input the score\n")
	// 	 fmt.Scanf("%v\n",&score[i])
	//  }
	//  fmt.Printf("the whole score are %v\n",score)
	// var array1[3] int =[3]int {1,2,3}
	// var array2 =[3]int {1,2,3}
	// var array3 =[...] int{6,7,8}
	// var array4 =[4] string {1:"a",2:"2a",0:"0a"}
	// fmt.Printf("%v\n%v\n%v\n%v\n",array1,array2,array3,array4)
	// for _,b := range array2{
	// 	fmt.Printf("%v\n",b)
	// }
	// testArray(&array2) 
	// testArray(&array3)
	// fmt.Printf("%v\n%v\n%v\n%v\n",array1,array2,array3,array4)
	// var arr1 [26]byte
	// for a,_ :=range arr1 {
	// 	arr1[a] = byte('A' +  a)
	// }
	// fmt.Printf("%c\n",arr1)
	// arr :=[10]int{10,20,30,40,50,60,80,70}
	// var sum  int
	// for _,b := range arr{
	// 	sum+=b
	// }
	// fmt.Printf("max is %v,index is %v\n",sum,sum/len(arr))
	// var arr [13]int
	// //var tmp int
	// for a,_:= range arr {
	// 	arr[a]=rand.Intn(100)
	// }
	// fmt.Printf("arr is %v\n",arr)
	// for i:=0;i<len(arr)/2;i++{
	// 	// tmp=arr[i]
	// 	// arr[i]=arr[len(arr)-i-1]
	// 	// arr[len(arr)-i-1]=tmp
	// 	arr[i],arr[len(arr)-i-1]=arr[len(arr)-i-1],arr[i]
	// }
	// fmt.Printf("arr is %v\n",arr)
	// var intarr [10]int = [10]int{1,2,3,4,5,11,22,33,44,55}
	// slice := intarr[1:5]
	// //var slice1 []int = make([]int ,4,10)
	// // slice[1] =10
	// for i:=0;i<len(slice);i++{
	// 	fmt.Printf("%v\n",slice[i])
	// }
	// for a,b := range slice{
	// 	fmt.Printf("slice[%v] is %v\n",a,b)
	// }
	//fmt.Printf("%v\n%v\n%v\n%v\n%v\n%v\n%v\n",slice,len(slice),cap(slice),intarr,&slice[1],&intarr[2],&slice[0])
	// var arr [4][6]int
	// arr[1][2]=1
	// arr[2][1]=2
	// arr[2][3]=3
	// for i:=0;i<len(arr);i++{
	// 	for j:=0;j<len(arr[i]);j++{
	// 		fmt.Printf("%v ",arr[i][j])
	// 	}
	// 	fmt.Printf("\n")
	// }
	// for i,v := range arr{
	// 	for j,v2 :=range v{
	// 		fmt.Printf("arr[%d][%d] is %v ",i,j,v2)
	// 	}
	// 	fmt.Printf("\n")
	// }
	rand.Seed(time.Now().UnixNano())
	// var arr = make([]int ,10)
	// var avg float64
	// var sum int 
	// for i:=0;i<len(arr);i++{
	// 	arr[i] = rand.Intn(100)+1
	// }
	// fmt.Printf("original array is %v\n",arr)
	// idx,max:=findMax(arr)
	// fmt.Printf("max value is %v,index is %v\n",max,idx)
	// sortArrayDesc(&arr)
	// fmt.Printf("sorted array is %v\n",arr)
	// for i:=0;i<len(arr);i++{
	// 	sum+=arr[i]
	// }
	// avg = float64(sum)/float64(len(arr))
	// fmt.Printf("sum value is %v,avg value is %.2f\n",sum,avg)
	// findInArray(55,arr)
	// sortArrayAsc(&arr)
	// fmt.Printf("\n")
	// arr =append (arr,rand.Intn(100)+1)
	// fmt.Printf("original array is %v\n",arr)
	// sortArrayAsc(&arr)
	// fmt.Printf("sorted array is %v\n",arr)
	// fmt.Printf("\n")
	// var arr2 [3][4]int
	// for a1:=0;a1<len(arr2);a1++{
	// 	for a2:=0;a2<len(arr2[a1]);a2++{
	// 		// fmt.Printf("input value for arr2[%v][%v]:\n",a1+1,a2+1)
	// 		// fmt.Scanf("%d\n",&arr2[a1][a2])
	// 		arr2[a1][a2]= rand.Intn(100)+1
	// 	}
	// }
	// for a1:=0;a1<len(arr2);a1++{
	// 	fmt.Printf("%v\n",arr2[a1])
	// }
	// for a,_:=range arr2{
	// 	for b,_:= range arr2[a]{
	// 		if a==0||a==len(arr2)-1{
	// 			arr2[a][b]=0
	// 		}else if b==0||b==len(arr2[a])-1{
	// 			arr2[a][b]=0
	// 		}
	// 	}

	// }
	// for a1:=0;a1<len(arr2);a1++{
	// 	fmt.Printf("%v\n",arr2[a1])
	// }
	// fmt.Printf("\n")
	// var arr3 [4][5]int
	// for a1:=0;a1<len(arr3);a1++{
	// 	for a2:=0;a2<len(arr3[a1]);a2++{
	// 		arr3[a1][a2]= rand.Intn(100)+1
	// 	}
	// }
	// for a1:=0;a1<len(arr3);a1++{
	// 	fmt.Printf("%v\n",arr3[a1])
	// }
	// fmt.Printf("\n")
	// for i:=0;i<2;i++{
	// 	for j:=0;j<len(arr3[i]);j++{
	// 		arr3[i][j],arr3[3-i][j]=arr3[3-i][j],arr3[i][j]
	// 	}
	// }
	// for a1:=0;a1<len(arr3);a1++{
	// 	fmt.Printf("%v\n",arr3[a1])
	// }
	// fmt.Printf("\n")
	// var arr4 = []int{1,3,5,7,9}
	// sortArrayDesc(&arr)
	// fmt.Printf("%v\n",arr4)
	// fmt.Printf("\n")
	// var arr5 = []string{"CC","BB","AA","DD","BB","AA","DD","BB","AA","DD","BB","AA","DD"}
	// findString("AA",arr5)
	// fmt.Printf("\n")

	// var arr6 []int 
	// for i:=0;i<50;i++{
	// 	arr6 =append(arr6,rand.Intn(100)+1)
	// }
	// fmt.Printf("%v\n",arr6)
	// sortArrayAsc(&arr6)
	// fmt.Printf("%v\n",arr6)
	// binaryFind(90,arr6,0,len(arr6))
	// fmt.Printf("\n")

	// var arr7 []int
	// for i:=0;i<10;i++{
	// 	arr7 =append(arr7,rand.Intn(100)+1)
	// }
	// fmt.Printf("%v\n",arr7)
	// idx1,max :=findMax(arr7)
	// idx2,min :=findMin(arr7)
	// fmt.Printf("max is %v,index is %v\n",max,idx1)
	// fmt.Printf("min is %v,index is %v\n",min,idx2)
	// fmt.Printf("\n")

	// var sum_arr7 int
	// for i:=0;i<len(arr7);i++{
	// 	sum_arr7+=arr7[i]
	// }
	// fmt.Printf("avg value is %v\n",sum_arr7/len(arr7))
	// var arr_up ,arr_down []int
	// for i:=0;i<len(arr7);i++{
	// 	if arr7[i]>sum_arr7/len(arr7){
	// 		arr_up = append(arr_up,arr7[i])
	// 	}else if arr7[i]<sum_arr7/len(arr7){
	// 		arr_down = append(arr_down,arr7[i])
	// 	}
	// }
	// fmt.Printf("large than average are %v,the count is %v\n",arr_up,len(arr_up))
	// fmt.Printf("less than average are %v,the count is %v\n",arr_down,len(arr_down))
	// fmt.Printf("\n")

	// var arr8 []int
	// for i:=0;i<8;i++{
	// 	arr8 =append(arr8,rand.Intn(10)+1)
	// }
	// fmt.Printf("the scores are  %v\n",arr8)
	// arr8_max,arr8_min :=calcSwimScore(arr8)
	// fmt.Printf("the highest score is %v\n",arr8[arr8_max[0]])
	// fmt.Printf("the highest judge are %v\n",arr8_max)
	// fmt.Printf("the lowest score is %v\n",arr8[arr8_min[0]])
	// fmt.Printf("the lowest judge are %v\n",arr8_min)
	// calcSwimJudge(arr8,arr8_max[0],arr8_min[0])
	// var stu Student
	// stu.Name="Jack"
	// stu.Gender='1'
	// stu.Age=20
	// stu.Id=20105824
	// stu.Score=99.8
	// stu.Say()
	// fmt.Println("Name is ",stu.Name)
	// var sen Secnic =Secnic{16,"Jack",-1}
	// sen.CalucSecnicPrice()
	// fmt.Printf("%v,your age is %v,the price is %v\n",sen.Name,sen.Age,sen.Price)
	// a:= model.Newstudent("jack",98.2)
	// fmt.Println("the result is ",(*a).Newscore())
	// var p = model.NewPerson("jack")
	// p.Setage(30)
	// p.Setsal(30000)
	// fmt.Println(*p)
	// fmt.Println(p.Name)
	// fmt.Println(p.Getage())
	// fmt.Println(p.Getsal())
	// var p = model.NewAccount()
	// p.Setno("1234567")
	// p.Setbal(100)ß
	// p.Setpwd("123456")
	// fmt.Println(*p)
	// var a Books
	// a.Goods.Name = "白夜行"
	// a.Goods.Price = 49.9
	// a.Writer ="东野奎吾"
	// a.m.No="1"
	// a.m.Setbal(20)
	// fmt.Println(a)
	// var a Usb2 = &Computer{}
	// a.Start1(1)
	// a.Start1(2)
	// var hs 	Heroslice
	// for i:=0;i<10;i++{
	// 	hc := Hero{
	// 		Name : fmt.Sprintf("hero%d",rand.Intn(100)),
	// 		Age : rand.Intn(100),
	// 	}
	// 	hs = append(hs,hc)
	// }
	// fmt.Printf("before: %v\n",hs)
	// sort.Sort(hs)
	// fmt.Printf("after: %v\n",hs)
	// var a interface{} = Point{1,2}
	// var b Point
	// var flag bool
	// b,flag=a.(Point)
	// fmt.Println(b,flag)
	// var n1 =1.1
	// var n2 float32=2.2
	// var n3 =10
	// var n4 ="s"
	// n5:=Student{Name:"TOM"}
	// n6:=&Student{Name:"TOMmy"}
	// TypeJudge(n1,n2,n3,n4,n5,n6)
	// file,err :=os.Open("/Users/qiankun04/go/src/test/main/test.txt")
	// if err != nil {
	// 	fmt.Println("open file err=",err)
	// }

	// defer file.Close()
	// reader := bufio.NewReader(file)
	// for{
	// 	str,error :=reader.ReadString('\n')
	// 	fmt.Printf("%v,%v\n",str,error)
	// 	if error == io.EOF{
	// 		break
	// 	}
	// }
	// file :="/Users/qiankun04/go/src/test/main/test.txt"
	// content,err :=ioutil.ReadFile(file)
	// if err != nil{
	// 	fmt.Printf("read file error :%v",err)
	// }
	// fmt.Printf("%s",content)
	// filePath := "/Users/qiankun04/go/src/test/main/write_test.txt"
	// file ,err := os.OpenFile(filePath,os.O_RDWR | os.O_TRUNC,0777)
	// if err != nil{
	// 		fmt.Printf("open file error :%v",err)
	// 		return
	// }

	// defer file.Close()

	// reader :=bufio.NewReader(file)
	// for{
	// 	str,error :=reader.ReadString('\n')
	// 	if error == io.EOF{
	// 		break
	// 	}
	// 	fmt.Printf("%v",str)
	// }
	// str :="hello,Garden\n"
	// writer := bufio.NewWriter(file)
	// for i:=0;i<5;i++{
	// 	writer.WriteString(str)
	// }
	// writer.Flush()
	//file1 :="/Users/qiankun04/go/src/test/main/test.txt"
	// file2 :="/Users/qiankun04/go/src/test/main/copy_test.txt"
	// content,err :=ioutil.ReadFile(file1)
	// if err != nil{
	// 	fmt.Printf("read file error :%v",err)
	// }
	// err = ioutil.WriteFile(file2,content,0777)
	// if err != nil{
	// 	fmt.Printf("read file error :%v",err)
	// }
	// srcfile,err :=os.Open(file1)
	// if err != nil{
	// 	fmt.Printf("open file error :%v",err)
	// }

	// defer srcfile.Close()
	// reader :=bufio.NewReader(srcfile)
	// dstfile,err2 :=os.OpenFile(file2,os.O_WRONLY | os.O_CREATE, 0666 )
	// if err2 != nil{
	// 	fmt.Printf("open file error :%v",err2)
	// 	return
	// }
	// defer dstfile.Close()
	// writer := bufio.NewWriter(dstfile)
	// written , err3 :=io.Copy(writer,reader)
	// if err2 != nil{
	// 	fmt.Printf("copy file error :%v",err3)
	// 	return
	// }else{
	// 	fmt.Printf("copy file succeed :%v\n",written)
	// }
	// writer.Flush()
	// file,err :=os.Open(file1)
	// if err != nil{
	// 	fmt.Printf("open file error :%v",err)
	// }
	// defer file.Close()
	// var count CharCount
	// var n int 
	// reader :=bufio.NewReader(file)
	// for{
	// 	n++
	// 	str,err1 :=reader.ReadBytes('\n')
	// 	if err1 == io.EOF{
	// 		break
	// 	}
	// 	for _,v := range str{
	// 		switch {
	// 		case v >='a' && v <='z' :
	// 				count.ChCount++
	// 		case v>='A' && v<='Z':
	// 				count.ChCount++
	// 		case v==' '||v=='\t':
	// 				count.SpaceCount++
	// 		case v>='0'&& v<='9':
	// 				count.NumCount++
	// 		default :
	// 				count.OtherCount++
	// 		}
	// 	}
	// }
	// fmt.Printf("%v,%v\n",count,n)
	// for i,v :=range os.Args{
	// 	fmt.Printf("Args[%v]:%v\n",i,v)
	// }
	// var user,pwd,host string
	// var port int
	// flag.StringVar(&user,"u","root","username,default root")
	// flag.StringVar(&pwd,"p","root","password,default root")
	// flag.StringVar(&host,"h","localhost","host,default localhost")
	// flag.IntVar(&port,"P",3306,"port,default root")
	// flag.Parse()
	// fmt.Printf("%v %v %v %v\n",user,pwd,host,port)
	// monster := Monster{
	// 		Name : "Jack" ,
	// 		Age : 18,
	// 		Birth : "1993/11/27",
	// 		Salary : 300.01,
	// 		Skill : "kill",
	// }
	// a := make(map[string]interface{})
	// a["name"]="tom"
	// a["age"]=20
	// a["gender"]="male"
	// c := make(map[string]interface{})
	// c["name"]="jack"
	// c["age"]=30
	// c["gender"]= []string{"female","male"}
	// var b []map[string]interface{}
	// b =append(b,a)
	// b =append(b,c)
	// d := float64(200.34)
	// content,err :=json.Marshal(&monster)
	// if err != nil{
	// 	fmt.Printf("monster xuliehua error: %v\n",err)
	// }
	// fmt.Printf("monster xuliehua : %s \n%T\n",content,content)
	// content,err =json.Marshal(a)
	// if err != nil{
	// 	fmt.Printf("monster xuliehua error: %v\n",err)
	// }
	// fmt.Printf("a map[string] xuliehua : %s \n%T\n",content,content)
	// content,err =json.Marshal(b)
	// if err != nil{
	// 	fmt.Printf("monster xuliehua error: %v\n",err)
	// }
	// fmt.Printf("b slice of map[string] xuliehua : %s \n%T\n",content,content)
	// content,err =json.Marshal(&d)
	// if err != nil{
	// 	fmt.Printf("monster xuliehua error: %v\n",err)
	// }
	// fmt.Printf("b slice of map[string] xuliehua : %s \n%T\n",content,content)
	// go test()
	// for i:=0;i<10;i++{
	// 	fmt.Printf("hello golang! %v\n",i)
	// 	time.Sleep(time.Second)
	// }
	// cpuNum :=runtime.NumCPU()
	// fmt.Printf("%v\n",cpuNum)
	// runtime.GOMAXPROCS(cpuNum -1)
	// for i:=1;i<=20;i++{
	// 	go fact(i)
	// }
	// time.Sleep(5 * time.Second)
	// lock.Lock()
	// for a,v := range myMap{
	// 	fmt.Printf("map[%v] is %v\n",a,v)
	// }
	// lock.Unlock()
	/* var intChan chan int
	num :=20
	intChan=make(chan int,3)
	fmt.Printf("%v %p\n",intChan,&intChan)
	intChan <- 10
	intChan <- num
	intChan <- 10
	num =<-intChan
	fmt.Printf("%v %v %v\n",len(intChan),cap(intChan),num) */
	// var per  Person
	// var perChan chan Person
	// perChan =make(chan Person,10)
	// for i:=1;i<=10;i++{
	// 	per.Name = "tom"
	// 	per.Age =rand.Intn(80)
	// 	per.Address ="beijing"
	// 	fmt.Printf("%v\n",per)
	// 	perChan <- per
	// }
	// close(perChan)
	// for v :=range perChan{
	// 	fmt.Printf("%v\n",v)
	// }
	// initChan :=make(chan int,50)
	// exitChan :=make(chan bool,1)
	// go writeData(initChan)
	// go readData(initChan,exitChan)
	// _,ok := <-exitChan
	// fmt.Printf("%v\n",ok)
	// numChan :=make(chan int,20)
	// resChan :=make(chan map[int]int,20)
	// exitChan :=make(chan bool,1)
	// flagChan :=make(chan string,8)
	// go writeData(numChan)
	// for i:=1;i<=8;i++{
	// 	go calcData(numChan,resChan,flagChan)
	// }
	// go readData(resChan,exitChan)
	// for {
	// 	if len(flagChan)==8{
	// 		close (resChan)
	// 		break
	// 	}
	// }
	// _,ok := <-exitChan
	// fmt.Printf("%v\n",ok)
	// exitChan := make(chan bool,1)
	// //exitChan1 := make(chan bool,1)
	// filePath1 :="/Users/qiankun04/go/src/learningGo/main/random_write.txt"
	// filePath2 :="/Users/qiankun04/go/src/learningGo/main/sorted_write.txt"
	// //go writeDataToFile(filePath1,20,exitChan)
	// _,ok:=<-exitChan
	// _,ok1:=<-exitChan
	// fmt.Printf("wtrie flag:%v\n",ok)
	// go sortDataInFile(filePath1,filePath2,exitChan)
	// //_,ok1:=<-exitChan
	// fmt.Printf("sort flag:%v\n",ok1)
	// intChan :=make(chan int,2000)
	// primeChan :=make(chan int,3000)
	// exitChan :=make(chan bool,4)
	// filePath :="/Users/qiankun04/go/src/learningGo/main/prime_write.txt"
	// go putNum(2000,intChan)
	// for i:=0;i<4;i++{
	// 	go checkPrime(intChan,primeChan,exitChan)
	// }
	// for{
	// 	if len(exitChan)==4{
	// 		close(primeChan)
	// 		break
	// 	}
	// }
	// file,err :=os.OpenFile(filePath,os.O_RDWR |os.O_CREATE |os.O_TRUNC,0777)
	// if err !=nil {
	// 	fmt.Printf("open file error : %v\n",err)
	// }
	// defer file.Close()
	// writer :=bufio.NewWriter(file)
	// len :=len(primeChan)
	// for i:=0;i<len;i++{
	// 	num := <-primeChan
	// 	fmt.Printf("%v\n",num)
	// 	writer.WriteString(strconv.Itoa(num))
	// 	writer.WriteByte('\n')
	// }
	// writer.Flush()
	// var num int =100
	// reflectTest01(&num)	
	// fmt.Printf("%v\n",num)
	// var stu = Students{"tom",20}
	// reflectTest02(stu)
	// var f1 = 2.3
	// reflectTest03(f1)
	// var mon Monsters = Monsters{
	// 	"cat",
	// 	400,
	// 	88.8,
	// }
	// testStruct(mon)
	// fmt.Printf("%v\n",mon)
   var a Cal = Cal{
	   5,
	   3,
   }
   testCal(a)
} 

type Cal struct{
	Num1 int
	Num2 int
}

func (c Cal)GetSub(name string){
	fmt.Printf("%v has done the minus: %v - %v = %v\n",name,c.Num1,c.Num2,c.Num1-c.Num2)
}

func testCal(a interface{}){
	rtype :=reflect.TypeOf(a)
	rvalue :=reflect.ValueOf(a)
	rnum :=rvalue.NumField()
	for i:=0;i<rnum;i++{
		fmt.Printf(" the Cal struct's %v is %v\n",i,rtype.Field(i).Name)
	}
	var param2 []reflect.Value
	param2=append(param2,reflect.ValueOf("TOM"))
	rvalue.Method(0).Call(param2)
}
type Students struct{
	Name string
	Age int
}

type Monsters struct{
	Name string `json:"name"`
	Age int `json:"age"`
	Score float64
}
func (s Monsters) Print(){
	fmt.Println("---start---")
	fmt.Println(s)
	fmt.Println("---end---")
}
func (s Monsters) Sum(n1,n2 int)(int){
	return n1+n2
}
func (s Monsters) Set(name string,age int,score float64){
	s.Name=name
	s.Age=age
	s.Score=score
	fmt.Println(s)
}

func testStruct(a interface{}){
	rtype :=reflect.TypeOf(a)
	rval :=reflect.ValueOf(a)
	//rkind :=rval.Kind()
	rnum :=rval.NumField()
	fmt.Println("the struct's field count is",rnum)
	for i:=0 ;i<rnum;i++{
		fmt.Printf("field %v 's value is %v,tag is %v\n",i,rval.Field(i),rtype.Field(i).Tag.Get("json"))
	}
	rmethodnum := rval.NumMethod()
	fmt.Println("the struct's method count is",rmethodnum)
	rval.Method(0).Call(nil)
	var param  []reflect.Value
	var param1  []reflect.Value
	param =append(param,reflect.ValueOf(10))
	param =append(param,reflect.ValueOf(40))
	res :=rval.Method(2).Call(param)
	fmt.Printf("res = %v\n",res[0].Int())
	param1=append(param1,reflect.ValueOf("robot"))
	param1=append(param1,reflect.ValueOf(500))
	param1=append(param1,reflect.ValueOf(10.2))
	rval.Method(1).Call(param1)
}
func reflectTest03(b interface{}){
	rtype := reflect.TypeOf(b)
	rvalue := reflect.ValueOf(b)
	rkind := rvalue.Kind()
	iv := rvalue.Interface()
	fv :=iv.(float64)
	fmt.Printf("%v %v %v %v %v\n",rtype,rvalue,rkind,iv,fv)
}
func reflectTest02 (b interface{}){
	rtype :=reflect.TypeOf(b)
	rvalue :=reflect.ValueOf(b)
	rkind :=reflect.ValueOf(b).Kind()
	iv :=rvalue.Interface()
	stu := iv.(Students)
	
	fmt.Printf("%v %v %T %v %T  %v  %v\n",rtype,rvalue,rvalue,iv,iv,stu.Name,rkind)
}
func reflectTest01 (b interface{}){
	rtype :=reflect.TypeOf(b)
	rvalue :=reflect.ValueOf(b)
	iv :=rvalue.Interface()
	rvalue.Elem().SetInt(20)
	fmt.Printf("%v %v %T %v\n",rtype,rvalue,rvalue,iv)
}
func putNum (n int ,intChan chan int){
	for  i:=1;i<=n;i++{
		intChan <- i
	}
	close (intChan)
}
func checkPrime (intChan chan int,primeChan chan int,exitChan chan bool){
	for{
		num,ok:=<-intChan
		//fmt.Printf("%v %v\n",num,ok)
		if !ok{
			exitChan <-true
			break
		}
		flag := true
	for i:=2;i<num;i++{
		if num%i==0{
			flag =false
			break
		}
		}
		if flag==true{
			primeChan <- num
		}
	}	
}
func writeDataToFile(filePath string,n int,exitChan chan bool){
	file,err :=os.OpenFile(filePath,os.O_RDWR |os.O_CREATE |os.O_TRUNC,0777)
	if err !=nil {
		fmt.Printf("open file error : %v\n",err)
	}
	defer file.Close()
	writer :=bufio.NewWriter(file)
	for i:=0;i<n;i++{
		writer.WriteString(strconv.Itoa(rand.Intn(10000)))
		writer.WriteByte('\n')
	}
	writer.Flush()	
	exitChan <- true
	close (exitChan)
}

func sortDataInFile(filePath1 string,filePath2 string,exitChan chan bool){
	file,err :=os.Open(filePath1)
	if err !=nil {
		fmt.Printf("open file error :%v",err)
	}
	defer file.Close()
	reader :=bufio.NewReader(file)
	var slice []int
	for{
		str,err :=reader.ReadString('\n')
		//fmt.Printf("%v ",str)
		if err ==io.EOF{
			break
		}
		num,_:=strconv.Atoi(str[0:len(str)-1])
		slice = append(slice,num)
		
	}
	//fmt.Printf("slice is here\n")
	for i:=1;i<len(slice);i++{
		for j:=0;j<len(slice)-i;j++{
			if slice[j]<slice[j+1]{
				slice[j],slice[j+1]=slice[j+1],slice[j]
			}
		}
	}
	//fmt.Printf("%v\n",slice)
	wirtefile,err1 :=os.OpenFile(filePath2,os.O_RDWR |os.O_CREATE|os.O_TRUNC,0777)
	if err1!=nil {
		fmt.Printf("open file error : %v\n",err)
	}
	defer wirtefile.Close()	
	writer :=bufio.NewWriter(wirtefile)
	for a:=0;a<=len(slice)-1;a++{
		writer.WriteString(strconv.Itoa(slice[a]))
		writer.WriteByte('\n')
	}
	writer.Flush()	
	exitChan <- true
	close (exitChan)
}
func writeData(initChan chan int){
	for i:=1;i<=20;i++{
		initChan <- i
		fmt.Printf("write : %v\n",i)
	}
	close(initChan)
}

func calcData(initChan chan int,resChan chan map[int]int,flagChan chan string){
	for{
		num ,ok:= <- initChan
	if !ok{
		flagChan <- "finish"
		break
		}
	a :=make(map[int]int,1)
	res := 0
	for i:=1;i<=num;i++{
		res +=i
	}
	a[num]=res
	fmt.Printf("calc : %v\n",num)
	resChan <- a
	}
}
func readData(resChan chan map[int]int,exitChan chan bool){
	for{
		 v ,ok:= <- resChan
		if !ok{
			break
			}
		for i,j :=range v{
			fmt.Printf("res[%v]=%v\n",i,j)
		}
	}
	exitChan <- true
	close(exitChan)
}
type Person struct{
	Name string
	Age int
	Address string
}
var(
	myMap = make(map[int]int,10)
	lock sync.Mutex
)
func fact(n int){
	res :=1
	for i:=1;i<=n;i++{
		res*=i
	}
	lock.Lock()
	myMap[n]=res
	lock.Unlock()
}
func test(){
	for i:=0;i<10;i++{
		fmt.Printf("hello world! %v\n",i)
		time.Sleep(time.Second)
	}
}
type Monster struct{
	Name string `json:"namefor"`
	Age int
	Birth string
	Salary float64
	Skill string
}
type CharCount struct{
	ChCount int
	NumCount int
	SpaceCount int
	OtherCount int
}
func PathExists(path string)(bool,error){
	_,err := os.Stat(path)
	if err ==nil{
		return true,nil
	}
	if os.IsNotExist(err){
		return false,nil
	}
	return false,err
}
func TypeJudge(item ...interface{}){
	for i,j :=range item{
		switch j.(type){
	case bool :
		fmt.Printf("the %v's parameter's type is bool,value is %v\n",i,j)
	case float32 :
		fmt.Printf("the %v's parameter's type is float32,value is %v\n",i,j)
	case float64 :
		fmt.Printf("the %v's parameter's type is float64,value is %v\n",i,j)
	case int,int32,int64 :
		fmt.Printf("the %v's parameter's type is int,value is %v\n",i,j)
	case string :
		fmt.Printf("the %v's parameter's type is string,value is %v\n",i,j)
	case Student :
		fmt.Printf("the %v's parameter's type is Student,value is %v\n",i,j)
	case *Student :
		fmt.Printf("the %v's parameter's type is *Student,value is %v\n",i,*(j.(*Student)))
	default :
		fmt.Printf("the %v's parameter's type is unknown,value is %v\n",i,j)
	}
	}
}
type Point struct {
	x int
	y int
}
type Hero struct{
	Name string
	Age int
}
type Heroslice []Hero

func (hs Heroslice) Len()int{
	return len(hs)
}
 
func (hs Heroslice) Less(i,j int)bool{
		return  hs[i].Name < hs[j].Name
}

func (hs Heroslice) Swap(i,j int){
	   hs[i],hs[j] =hs[j],hs[i]
}
type Usb interface{
	Start()
	Stop()
}
type Usb2 interface{
	Start1(a int)
	Stop()
}
type Phone struct{

}
type Camera struct{

}
type Computer struct{

}
func (p *Phone) Start(){
	fmt.Printf("Phone start working.\n")
}
func (p *Phone) Start1(a int){
	fmt.Printf("Phone start working %v too.\n",a)
}
func (p *Phone) Stop(){
	fmt.Printf("Phone stop working.\n")
}
func (c *Camera) Start(){
	fmt.Printf("Camera start working.\n")
}
func (c *Camera) Stop(){
	fmt.Printf("Camera stop working.\n")
}
func (co *Computer) Start1(a int){
	fmt.Printf("Computer start working %v.\n",a)
}
func (c *Computer) Stop(){
	fmt.Printf("Computer stop working.\n")
}
type Goods struct{
	Name string
	Price float64
}
type Books struct{
	Goods
	m model.Account
	Writer string
}
type Bankaccount struct{
	Name string
	Account float64
	Password string
}

func (acc *Bankaccount)QueryAccount(a string,b string)float64{
	if (a==acc.Name && b==acc.Password){
		return  acc.Account
	}
	return -1 
}
type Secnic struct{
	Age int
	Name string
	Price int 
}
func(sen *Secnic)CalucSecnicPrice(){
	if sen.Age<18&&sen.Age>0{
		sen.Price=0
	}else{
		sen.Price=20
	}
}
type Student struct{
	Name string
	Gender byte
	Age int
	Id int
	Score float64
}
func (stu Student)Say(){
	var a string
	if stu.Gender=='1' {
		a="male"
	}else{
		a="female"
	}
	stu.Name="Rose"
	fmt.Println("Name is ",stu.Name)
	fmt.Println("Gender is ",a)
	fmt.Println("Age is ",stu.Age)
	fmt.Println("Id is ",stu.Id)
	fmt.Println("Score is ",stu.Score)
}
func calcSwimScore(arr []int)([]int,[]int){
	var max ,min int = arr[0],arr[0]
	var arr_max ,arr_min = make([]int,1),make([]int,1)
	for i:=0;i<len(arr);i++{
		if (arr[i]>max){
			max = arr[i]
		} else if arr[i]<min{
			min = arr[i]
		}
	}
	for i:=0;i<len(arr);i++{
		if arr[i]==max{
			arr_max =append(arr_max,i)
		}else if arr[i]==min{
			arr_min =append(arr_min,i)
		}
	}
	return arr_max[1:],arr_min[1:]
}

func calcSwimJudge(arr []int,a int ,b int){
	var sum ,avg int 
	//fmt.Printf("%v,%v\n",a,b)
	for i:=0;i<len(arr);i++{
		if(i!=a&&i!=b){
			sum+=arr[i]
		}
	}
	avg = sum/(len(arr)-2)
	fmt.Printf("the avg score is %v,sum is %v\n",avg,sum)
	var max =	int(math.Abs(float64(arr[0]-avg)))
	var min = 	int(math.Abs(float64(arr[0]-avg)))
	var arr_max ,arr_min = make([]int,1),make([]int,1)
	for i:=0;i<len(arr);i++{
		if (int(math.Abs(float64(arr[i]-avg)))>max){
			max = int(math.Abs(float64(arr[i]-avg)))
		} else if int(math.Abs(float64(arr[i]-avg)))<min{
			min = int(math.Abs(float64(arr[i]-avg)))
		}
	}
	for i:=0;i<len(arr);i++{	
	 if(int(math.Abs(float64(arr[i]-avg)))==max){
			arr_max = append (arr_max,i)
		}else if (int(math.Abs(float64(arr[i]-avg)))==min){
			arr_min = append (arr_min,i)
		}
	}
	fmt.Printf("the best judges are %v\n",arr_min[1:])
	fmt.Printf("the worst judges are %v\n",arr_max[1:])
}
func binaryFind(a int,arr []int,bgn int,end int){
	var mid = (bgn+end)/2
	if mid<bgn||mid>end{
		fmt.Printf("not found %v in array\n",a)
		return
	}
	if arr[mid]==a {
		fmt.Printf("found %v,index is %v\n",a,mid)
	}else if arr[mid]<a{
		bgn = mid+1
		binaryFind(a,arr,bgn,end)
	}else{
		end = mid -1
		binaryFind(a,arr,bgn,end)
	}
}
func findString(a string,arr []string){
	var idx []int
	for i,j :=range arr{
		if j==a{
			idx = append(idx,i)
		}
	}
	if (len(arr)>0){
		fmt.Printf("found %v in array,index is %v\n",a,idx)
	}else{
		fmt.Printf("not found %v in array\n",a)
	}	
}
func findMax(arr []int)(int ,int){
	var max int =arr[0]
	var idx int 
	for i:=0;i<len(arr);i++{
		if arr[i]>max{
			max = arr[i]
			idx = i 
		}
	}
	return idx,max
}
func findMin(arr []int)(int ,int){
	var min int =arr[0]
	var idx int 
	for i:=0;i<len(arr);i++{
		if arr[i]<min{
			min = arr[i]
			idx = i 
		}
	}
	return idx,min
}
func sortArrayDesc(arr *[]int){
	for i:=0;i<len(*arr)-1;i++{
		for j:=0;j<len(*arr)-i-1;j++{
			if((*arr)[j]<(*arr)[j+1]){
				(*arr)[j],(*arr)[j+1]=(*arr)[j+1],(*arr)[j]
			}
		}
	}
}

func sortArrayAsc(arr *[]int){
	for i:=0;i<len(*arr)-1;i++{
		for j:=0;j<len(*arr)-i-1;j++{
			if((*arr)[j]>(*arr)[j+1]){
				(*arr)[j],(*arr)[j+1]=(*arr)[j+1],(*arr)[j]
			}
		}
	}
}
func findInArray(a int,arr []int){
	for i,j :=range arr{
		if j==a{
			fmt.Printf("found %v,index is %v\n",a,i)
			return
		}
	}
	fmt.Printf("not found %v in array\n",a)
}
func testArray(arr *[3]int) {
	(*arr)[0] = 0
}

// func checkPrime(a int) bool{
// 	for j:=2;j<=a/2;j++{
// 		if (a%j ==0){
// 			return false
// 		}
// 	}
// 	return true
// }
func checkDay( year,month,day int)bool{
	if(day>=1&&day<=31){
		switch month{
		case 1,3,5,7,8,10,12:
			return true
		case 4,6,9,11:
			if(day==31){
				return false
			}else{
				return true
			}
		case 2:
			if((year%4 == 0 && year%100 != 0) || year%400 == 0){
				if(day<=29){
					return true
				}else{
					return false
				}
			}else{
				if(day<=28){
					return true
				}else{
					return false
				}
			}
		default:
			return false
		}
	}else{
		return false
	}
}
func sum(n1,n2 int) int{
	defer fmt.Printf("ok1 n1=%d\n",n1)
	defer fmt.Printf("ok2 n2=%d\n",n2)
	res:=n1+n2
	fmt.Printf("ok3 res=%d\n",res)
	n1 = 20
	fmt.Printf("ok4 n1=%d\n",n1)
	return res
}

func makeSuffix(suf string) func(string) string{
	return func (str string) string{
		if (strings.HasSuffix(str,suf)== true){
			return str
		}else{
			return str+suf
		}
	}
	
}


func AddUpper() func(int) int{
	var n int =10
	return func(x int) int{
		n=n+x
		return n
	}
}
func swap(n1 int,n2 int)(int,int){
	return n2,n1
}
func init(){
	//fmt.Printf("init here\n")
}
func sumargs(n1 int,args... int) int{
	sum:=n1
	for i:=0;i<len(args);i++{
		sum +=args[i]
	}
	return sum 
}

func calucate(oper string,a int,b int)(res int){
	switch oper{
	case "+":
		res = a+b 
	case "-":
		res = a-b
	case "*":
		res = a*b
	case "/":
		res = a/b
	}
	return res
}