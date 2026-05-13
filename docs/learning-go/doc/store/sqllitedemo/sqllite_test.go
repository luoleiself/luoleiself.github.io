package sqllitedemo

import "testing"

func TestSqlLitePing(t *testing.T) {
	err := lite3.Ping()
	if err != nil {
		t.Fatalf("Error lite3.Ping() %s\n", err)
	}
	t.Log("lite3.Ping() OK")
}
