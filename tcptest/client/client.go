package main
import(
	"fmt"
	"net"
	"bufio"
	"os"
)

func main(){
	conn,err := net.Dial("tcp","127.0.0.1:45673")
	if err != nil{
		fmt.Print("Connect err. %v\n",err)
		return
	} 
	fmt.Printf("Connect succeed. %v\n",conn)
	for{
	reader := bufio.NewReader(os.Stdin)
	content,err1 := reader.ReadString('\n') 
	if err1 != nil{
		fmt.Print("Read err. %v\n",err1)
	} 
	fmt.Printf("send : %v",content)
	if content=="exit\n" {
		break	
	}
	_,err2 :=conn.Write([]byte(content))
	if err2 != nil{
		fmt.Print("Write err. %v\n",err2)
	} 
	}
	//Label:
	fmt.Printf("client has sent many bytes and exit.\n")
}