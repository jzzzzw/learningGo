package calc

import(
	"testing"

)


func TestAddUpper(t *testing.T){
	res :=addUpper(10)
	if res !=50{
		t.Fatalf("test AddUpper(10) failed ,expected 50 ,get %v",res)
	}

	t.Logf("test AddUpper(10) succeed ")
}

func TestSub(t *testing.T){
	res :=sub(10,10)
	if res !=0{
		t.Fatalf("test sub(10,10) failed ,expected 0 ,get %v",res)
	}

	t.Logf("est sub(10,10)  succeed ")
}

func TestStore(t *testing.T){
	m := Monster{"tom",20,"cooking"}
	res :=m.Store()
	if res == false{
		t.Fatalf("test store failed ,expected true ,get %v",res)
	}
	t.Logf("est store  succeed ")
}

func TestRestore(t *testing.T){
	var m Monster
	filepath :="/Users/qiankun04/go/src/learningGo/test/store_test.txt"
	res:=m.Restore(filepath)
	if res == false{
		t.Fatalf("test restore() failed ,expected true ,get %v",res)
	}
	if m.Name=="tom"{
		t.Logf("test restore()  succeed , m is %v ,name is %v",m,m.Name)
	}else{
		t.Fatalf("test restore() failed ")
	}
	
}