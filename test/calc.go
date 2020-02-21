package calc
import(
	_"testing"
	"encoding/json"
	"os"
	"fmt"
	"bufio"
	"io"
)

func addUpper(n int) int{
	res :=0 
	for i:=1 ;i<=10;i++{
		res +=i
	}
	return res
}

func sub(n1,n2 int) int{
	 
	return n1-n2
}

type Monster struct{
	Name string
	Age int
	Skill string
}

func(m *Monster)Store()bool{

	content,err :=json.Marshal(m)
	if err != nil{
		fmt.Printf("Monster serializes error: %v\n",err)
	}
	filePath := "/Users/qiankun04/go/src/learningGo/test/store_test.txt"
	file,error := os.OpenFile(filePath,os.O_RDWR | os.O_CREATE | os.O_TRUNC,0777)
	if error !=nil {
		fmt.Printf("open file error :%v",error)
		return false
	}

	defer file.Close()
	
	writer := bufio.NewWriter(file)
	_,err1 :=writer.WriteString(string(content)+"\n")
	writer.Flush()
	if err1 !=nil {
		fmt.Printf("write file error :%v",err1)
		return false
	}
	return true 
}

func(m *Monster)restore(filepath string)(bool){

	file,error := os.OpenFile(filepath,os.O_RDWR | os.O_CREATE ,0777)
	if error !=nil {
		fmt.Printf("open file error :%v",error)
		return false 
	}

	defer file.Close()
	
	reader :=bufio.NewReader(file)
	for{
		str,error :=reader.ReadString('\n')
		if error == io.EOF{
		 	break
		}
	  err := json.Unmarshal([]byte(str),&m)
	  if err !=nil {
		fmt.Printf("Unserialize error :%v",err)
		return false 
		}
	  fmt.Printf("%v\n",m)
	}
	return true 

}