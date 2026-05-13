package externaltesting_test

import (
	"fmt"
	"testing"

	"github.com/luoleiself/learning-go/testing/externaltesting"
)

func TestExportTest(t *testing.T) {
	fmt.Println("此处位于 externaltesting_test 包下的 TestExportTest 方法中")
	fmt.Printf("externaltesting.Jan %d\n", externaltesting.Jan)
	fmt.Printf("externaltesting.Feb %d\n", externaltesting.Feb)
	fmt.Printf("externaltesting.Mar %d \n", externaltesting.Mar)
	fmt.Printf("externaltesting.Apr %d \n", externaltesting.Apr)
	fmt.Printf("externaltesting.May %d \n", externaltesting.May)
	fmt.Printf("externaltesting.Jun %d \n", externaltesting.Jun)
	fmt.Printf("externaltesting.Jul %d \n", externaltesting.Jul)
	fmt.Printf("externaltesting.Aug %d \n", externaltesting.Aug)
	fmt.Printf("externaltesting.Sep %d \n", externaltesting.Sep)
	fmt.Printf("externaltesting.Oct %d \n", externaltesting.Oct)
	fmt.Printf("externaltesting.Nov %d \n", externaltesting.Nov)
	fmt.Printf("externaltesting.Dec %d \n", externaltesting.Dec)
	t.Log("success!!!")
}
