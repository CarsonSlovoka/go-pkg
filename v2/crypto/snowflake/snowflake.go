// Package snowflake provides a very simple Twitter snowflake generator and parser.
package snowflake

import (
	"errors"
	"strconv"
	"sync"
	"time"
)

// A Node struct holds the basic information needed for a snowflake generator
// [time][node][step]
type Node struct {
	baseTime *time.Time // 基準日
	time     int64      // 生成的時間 與 基準日 相減的差(微秒)

	node        int64 // 機器碼 (隨便您設定，有點像Token的意思)
	numNodeBits uint8 // 機器碼有幾碼
	maskNode    int64 // 機器碼的遮罩

	step        int64 // 流水號
	numStepBits uint8 // 流水號有幾碼
	maskStep    int64 // 流水號遮罩

	shiftTime uint8 // 時間戳記的位移碼位
	shiftNode uint8 // 機器碼的位移碼位

	mutex sync.Mutex // 為了不讓外部調用，用小寫命名
}

func (n *Node) BaseTime() time.Time {
	return *n.baseTime
}

func (n *Node) ShiftTime() uint8 {
	return n.shiftTime
}

func (n *Node) ShiftNode() uint8 {
	return n.shiftNode
}

func (n *Node) MaskNode() int64 {
	return n.maskNode
}

func (n *Node) MaskStep() int64 {
	return n.maskStep
}

// NewNode returns a new snowflake node that can be used to generate snowflake IDs
func NewNode(psw int64, baseTime time.Time, numNodeBits, numStepBits uint8) (*Node, error) {
	var nodeMax int64
	nodeMax = -1 ^ (-1 << numNodeBits) // 機器碼的最大可生成數值 // nodeNodeBits如果為5，表示需要二進制 11111
	n := Node{
		baseTime:  &baseTime,
		node:      psw,
		maskNode:  nodeMax << numStepBits,
		maskStep:  -1 ^ (-1 << numStepBits),
		shiftTime: numNodeBits + numStepBits,
		shiftNode: numStepBits,
	}

	if n.node < 0 || n.node > nodeMax {
		return nil, errors.New("Node number must be between 0 and " + strconv.FormatInt(nodeMax, 10))
	}

	return &n, nil
}

func (n *Node) Generate() ID {
	n.mutex.Lock()
	defer n.mutex.Unlock()

	baseTime := *n.baseTime
	now := time.Since(baseTime).Milliseconds()

	if now == n.time { // 如果當前時間與結點時間相同(微秒)，用流水號來區別
		n.step = (n.step + 1) & n.maskStep

		if n.step == 0 { // +1之後如果又循環回來，我們就讓時間設定到下一微秒
			now += 1
			if time.Since(baseTime.Add(time.Duration(now)*time.Millisecond)) < 0 {
				time.Sleep(time.Millisecond)
			}

			/* 寫面這種寫法會多跑很多次
			for i := 0; now <= n.time; i++ {
				fmt.Println(i)
				now = time.Since(baseTime).Milliseconds()
			}
			*/
		}
	} else {
		n.step = 0
	}

	n.time = now

	return ID((now)<<n.shiftTime |
		(n.node << n.shiftNode) |
		(n.step),
	)
}

// An ID is a custom type used for a snowflake ID.  This is used, so we can
// attach methods onto the ID.
type ID int64

// Int64 returns an int64 of the snowflake ID
func (id *ID) Int64() int64 {
	return int64(*id)
}

// String returns a string of the snowflake ID
func (id *ID) String() string {
	return strconv.FormatInt(int64(*id), 10)
}

// Base2 returns a string base2 of the snowflake ID
func (id *ID) Base2() string {
	return strconv.FormatInt(int64(*id), 2)
}

// Time returns an int64 unix timestamp in milliseconds of the snowflake ID time
func (id *ID) Time(baseTime time.Time, shiftTime uint8) time.Time {
	return baseTime.Add((time.Duration(*id) >> shiftTime) * time.Millisecond)
}

// Node returns an int64 of the snowflake ID node number
func (id *ID) Node(maskNode int64, shiftNode uint8) int64 {
	return int64(*id) & maskNode >> shiftNode
}

// Step returns an int64 of the snowflake step (or sequence) number
func (id *ID) Step(stepMask int64) int64 {
	return int64(*id) & stepMask
}

// ParseInt64 converts an int64 into a snowflake ID
func ParseInt64(id int64) ID {
	return ID(id)
}

// ParseString converts a string into a snowflake ID
func ParseString(id string) (ID, error) {
	i, err := strconv.ParseInt(id, 10, 64)
	return ID(i), err
}

// ParseBase2 converts a Base2 string into a snowflake ID
func ParseBase2(id string) (ID, error) {
	i, err := strconv.ParseInt(id, 2, 64)
	return ID(i), err
}
