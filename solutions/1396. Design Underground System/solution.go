type check struct {
	startStation string
	startTime    int
}

type allTrips struct {
	count     int
	sumOfTime int
}

type UndergroundSystem struct {
	currentChecks map[int]check
	data          map[string]allTrips
}

func Constructor() UndergroundSystem {
	return UndergroundSystem{
		currentChecks: make(map[int]check, 0),
		data:          make(map[string]allTrips, 0),
	}
}

func (this *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	this.currentChecks[id] = check{
		startStation: stationName,
		startTime:    t,
	}
}

func (this *UndergroundSystem) CheckOut(id int, stationName string, t int) {
	check := this.currentChecks[id]
	delete(this.currentChecks, id)

	key := this.calcStationKey(check.startStation, stationName)

	if _, ok := this.data[key]; !ok {
		this.data[key] = allTrips{}
	}

	trip := this.data[key]
	trip.count++
	trip.sumOfTime += t - check.startTime
	this.data[key] = trip
}

func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	trips := this.data[this.calcStationKey(startStation, endStation)]
	return float64(trips.sumOfTime) / float64(trips.count)
}

func (this *UndergroundSystem) calcStationKey(startStation string, endStation string) string {
	return startStation + "~" + endStation
}

/**
 * Your UndergroundSystem object will be instantiated and called as such:
 * obj := Constructor();
 * obj.CheckIn(id,stationName,t);
 * obj.CheckOut(id,stationName,t);
 * param_3 := obj.GetAverageTime(startStation,endStation);
 */
