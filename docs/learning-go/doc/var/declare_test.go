package main

import "testing"

func TestIota(t *testing.T) {
	t.Run("Month", func(t *testing.T) {
		t.Logf("January %d\n", January)
		t.Logf("February %d\n", February)
		t.Logf("March %d\n", March)
		t.Logf("April %d\n", April)
		t.Logf("May %d\n", May)
		t.Logf("June %d\n", June)
		t.Logf("July %d\n", July)
		t.Logf("August %d\n", August)
		t.Logf("September %d\n", September)
		t.Logf("October %d\n", October)
		t.Logf("November %d\n", November)
		t.Logf("December %d\n", December)
	})
	t.Run("Week", func(t *testing.T) {
		t.Logf("Sunday %d\n", Sunday)
		t.Logf("Monday %d\n", Monday)
		t.Logf("Tuesday %d\n", Tuesday)
		t.Logf("Wednesday %d\n", Wednesday)
		t.Logf("Thursday %d\n", Thursday)
		t.Logf("Friday %d\n", Friday)
		t.Logf("Saturday %d\n", Saturday)
	})
}
