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