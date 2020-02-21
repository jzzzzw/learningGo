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