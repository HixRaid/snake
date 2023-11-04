package snake

type TailStatus bool

const (
	Dead TailStatus = false
	Live TailStatus = true
)

type tail struct {
	length    [][2]float32
	fieldSize [2]float32
	headIndex int
	status    TailStatus
}

func newTail(fieldSize, pos, dir [2]float32, len int) *tail {
	t := tail{
		length:    make([][2]float32, len),
		fieldSize: fieldSize,
		status:    Live,
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

	switch {
	case t.length[targetHeadIndex][0] > t.fieldSize[0]-1:
		t.length[targetHeadIndex][0] = 0
	case t.length[targetHeadIndex][0] < 0:
		t.length[targetHeadIndex][0] = t.fieldSize[0] - 1
	case t.length[targetHeadIndex][1] > t.fieldSize[1]-1:
		t.length[targetHeadIndex][1] = 0
	case t.length[targetHeadIndex][1] < 0:
		t.length[targetHeadIndex][1] = t.fieldSize[1] - 1
	}

	for i, v := range t.length {
		if v == t.length[t.headIndex] && i != t.headIndex {
			return Dead
		}
	}

	return Live
}
