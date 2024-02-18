import (
	"container/heap"
	"sort"
)

type bookedRoom struct {
	end   int
	index int
}
type mhBookedRooms []*bookedRoom

func (mh mhBookedRooms) Len() int      { return len(mh) }
func (mh mhBookedRooms) Swap(i, j int) { mh[i], mh[j] = mh[j], mh[i] }
func (mh mhBookedRooms) Less(i, j int) bool {
	return mh[i].end < mh[j].end || (mh[i].end == mh[j].end && mh[i].index < mh[j].index)
}

func (mh *mhBookedRooms) Push(a interface{}) {
	*mh = append(*mh, a.(*bookedRoom))
}

func (mh *mhBookedRooms) Pop() interface{} {
	hLen := len(*mh) - 1
	el := (*mh)[hLen]
	*mh = (*mh)[:hLen]
	return el
}

type mhAvailableRooms []int

func (mh mhAvailableRooms) Len() int      { return len(mh) }
func (mh mhAvailableRooms) Swap(i, j int) { mh[i], mh[j] = mh[j], mh[i] }
func (mh mhAvailableRooms) Less(i, j int) bool {
	return mh[i] < mh[j]
}

func (mh *mhAvailableRooms) Push(a interface{}) {
	*mh = append(*mh, a.(int))
}

func (mh *mhAvailableRooms) Pop() interface{} {
	hLen := len(*mh) - 1
	el := (*mh)[hLen]
	*mh = (*mh)[:hLen]
	return el
}

func mostBooked(n int, meetings [][]int) int {
	bookedRooms := &mhBookedRooms{}
	availableRooms := &mhAvailableRooms{}
	var maxBookedRoomIdx, maxBookedRoomCount int
	bookingCount := make([]int, n)
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0] || (meetings[i][0] == meetings[j][0] && meetings[i][1] < meetings[j][1])
	})
	for i := 0; i < n; i++ {
		heap.Push(availableRooms, i)
	}
	for _, meeting := range meetings {
		// cleanup
		for bookedRooms.Len() > 0 {
			booked := heap.Pop(bookedRooms).(*bookedRoom)
			if booked.end > meeting[0] {
				heap.Push(bookedRooms, booked) // push booked room back
				break
			} else {
				heap.Push(availableRooms, booked.index)
			}
		}
		if availableRooms.Len() > 0 {
			room := heap.Pop(availableRooms).(int)
			heap.Push(bookedRooms, &bookedRoom{end: meeting[1], index: room})
			bookingCount[room]++
			// update index if booking count is greater or a lower room index has same count
			if bookingCount[room] > maxBookedRoomCount || (bookingCount[room] == maxBookedRoomCount && room < maxBookedRoomIdx) {
				maxBookedRoomIdx = room
				maxBookedRoomCount = bookingCount[room]
			}
			continue
		}
		// no available rooms, use the room which is ending first
		booked := heap.Pop(bookedRooms).(*bookedRoom)
		bookingCount[booked.index]++
		// update index if booking count is greater or a lower room index has same count
		if bookingCount[booked.index] > maxBookedRoomCount || (bookingCount[booked.index] == maxBookedRoomCount && booked.index < maxBookedRoomIdx) {
			maxBookedRoomIdx = booked.index
			maxBookedRoomCount = bookingCount[booked.index]
		}
		// push to booked
		newEnd := booked.end + (meeting[1] - meeting[0])
		heap.Push(bookedRooms, &bookedRoom{end: newEnd, index: booked.index})
	}
	return maxBookedRoomIdx
}
