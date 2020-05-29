package main
import(
	"fmt"
	"net"
	"io"
)

func main(){
	fmt.Printf("start Listening..\n")
	listen,err := net.Listen("tcp","0.0.0.0:45673")
	if err != nil{
		fmt.Printf("Listen error %v.\n",err)
		return
	}
	defer listen.Close()

	for{
		fmt.Printf("wait Client Listen.\n")
		conn,err1 :=listen.Accept()
		if err != nil{
			fmt.Printf("Accept error %v.\n",err1)
		} else{
			fmt.Printf("Accept succeed  conn is %v, ip is %v\n",conn,conn.RemoteAddr().String())
		}
		go process(conn)
	}
	//fmt.Printf("Listen is %v.\n",listen)
}

func process(conn net.Conn){
	defer conn.Close()
	fmt.Printf("Server wait to read %v\n",conn.RemoteAddr().String())
	for{
		buf := make([]byte,1024)
		n,err2 := conn.Read(buf)
		if err2 == io.EOF{
			fmt.Printf("Server read compelet, exiting. %v\n",err2)
			return 
		}
		fmt.Printf("%v",string(buf[:n]))
	}
}