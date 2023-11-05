package snake

type TailMove func(length *[][2]float32, headIndex int) (status TailStatus)

type TailStatus bool

const (
	Dead TailStatus = false
	Live TailStatus = true
)

type tail struct {
	length    [][2]float32
	headIndex int
	status    TailStatus
	tailMove  []TailMove
}

func newTail(tailMove []TailMove, pos, dir [2]float32, len int) *tail {
	t := tail{
		length:   make([][2]float32, len),
		status:   Live,
		tailMove: tailMove,
	}

	for i := 0; i < len; i++ {
		t.length[i][0], t.length[i][1] = pos[0]-dir[0]*float32(i), pos[1]-dir[1]*float32(i)
	}

	return &t
}

func (t *tail) move(dir [2]float32) TailStatus {
	targetHeadIndex := t.headIndex - 1
	if targetHeadIndex < 0 {
		targetHeadIndex = len(t.length) - 1
	}

	t.length[targetHeadIndex][0], t.length[targetHeadIndex][1] = t.length[t.headIndex][0]+dir[0], t.length[t.headIndex][1]+dir[1]
	t.headIndex = targetHeadIndex

	for _, f := range t.tailMove {
		if result := f(&t.length, t.headIndex); result == Dead {
			return Dead
		}
	}

	return Live
}
