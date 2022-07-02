package snowflake_test

import (
	"github.com/CarsonSlovoka/go-pkg/v2/crypto/snowflake"
	"testing"
	"time"
)

var BaseT time.Time

func init() {
	BaseT = time.Date(2022, 7, 1, 16, 10, 54, 0, time.UTC)
}

func TestNewNode(t *testing.T) {
	_, err := snowflake.NewNode(0, BaseT, 10, 12)
	if err != nil {
		t.Fatalf("error creating NewNode, %s", err)
	}

	_, err = snowflake.NewNode(5000, BaseT, 10, 12)
	if err == nil {
		t.Fatalf("no error creating NewNode, %s", err)
	}

}

// lazy check if Generate will create duplicate IDs
func TestGenerateDuplicateID(t *testing.T) {
	node, _ := snowflake.NewNode(1, BaseT, 10, 12)
	var x, y snowflake.ID
	for i := 0; i < 1000000; i++ {
		y = node.Generate()
		if x == y {
			t.Errorf("x(%d) & y(%d) are the same", x, y)
		}
		x = y
	}
}

func TestRace(t *testing.T) {
	node, _ := snowflake.NewNode(1, BaseT, 10, 12)

	go func() {
		for i := 0; i < 1000000000; i++ {
			_, _ = snowflake.NewNode(1, BaseT, 10, 12)
		}
	}()

	for i := 0; i < 40000; i++ {
		node.Generate()
	}

	ch := make(chan any)
	go func() {
		for i := 0; i < 40000; i++ {
			node.Generate()
		}
		close(ch)
	}()

	select {
	case <-ch:
	case <-time.After(time.Second * 5):
		close(ch)
	}
}

func TestParse(t *testing.T) {
	n, err := snowflake.NewNode(123, BaseT, 10, 12)
	if err != nil {
		t.Fatalf("error creating NewNode, %s", err)
	}

	id := n.Generate()
	/*
		t.Logf("org    : %#v", id)            // 398977984344064
		t.Logf("Int64    : %#v", id.Int64())  // 398977984344064
		t.Logf("String   : %#v", id.String()) // "398977984344064"
		t.Logf("Base2    : %#v", id.Base2())  // "1011010101101111001001101000001111011000000000000"
	*/

	_ = id.Time(n.BaseTime(), n.ShiftTime())
	if id.Node(n.MaskNode(), n.ShiftNode()) != 123 {
		t.Fatal()
	}

	if snowflake.ParseInt64(int64(id)) != id {
		t.Fatal()
	}

	if v, _ := snowflake.ParseString(id.String()); v != id {
		t.Fatal()
	}

	if v, _ := snowflake.ParseBase2(id.Base2()); v != id {
		t.Fatal()
	}

	n2, err := snowflake.NewNode(12345, BaseT, 14, 12)
	if err != nil {
		t.Fatal(err)
	}
	id2 := n2.Generate()
	// t.Logf("Base2    : %#v", id2.Base2())
	if v := id2.Node(n2.MaskNode(), n2.ShiftNode()); v != 12345 {
		t.Fatal(v)
	}
}

func TestOtherSize(t *testing.T) {
	for _, d := range []struct {
		psw     int64
		numNode uint8
		numStep uint8
	}{
		{0, 0, 0}, // Base2: 101101111001010010100011111
		{2, 3, 0},
		{0, 0, 5},
		{67108862, 26, 10}, // String: 6536855424677705728 Base2: "101101010110111100100110101111111111111111111111111100000000000"
	} {
		n, err := snowflake.NewNode(d.psw, BaseT, d.numNode, d.numStep)
		if err != nil {
			t.Fatal(n, err)
		}
		/*
			id := n.Generate()
			t.Logf("%#v", id.Base2())
			t.Logf("String   : %#v", id.String())
		*/
	}

	n1, _ := snowflake.NewNode(0, time.Date(2022, 7, 1, 16, 10, 54, 0, time.UTC), 0, 0)
	id2022 := n1.Generate()
	// t.Logf("%#v", id2022.Base2()) // 101110001100011100110100010
	n2, _ := snowflake.NewNode(0, time.Date(2008, 7, 1, 16, 10, 54, 0, time.UTC), 0, 0)
	id2008 := n2.Generate()
	// t.Logf("%#v", id2008.Base2()) // 110011011100000111010011011010110100010

	if len(id2008.Base2()) < len(id2022.Base2()) { // 時間越早，其生成的長度也越大
		t.Fatal()
	}
}
