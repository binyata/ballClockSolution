package main
import (
	"fmt"
	"strconv"
)

type BallStatus struct {
	Id   int
	BallName string
	CurrentPosition string
	OriginalPosition string
}

type BallList struct {
	Id         int
	Name       string
	BallStatuses []BallStatus
}

type Results struct {
	Id         int
	Days       int
	Main       []BallStatus
	Min        []BallStatus
	FiveMin    []BallStatus
	Hour       []BallStatus
}

func ballCycleIsRestored(generatedBalls BallList, numBalls int) bool {
	if len(generatedBalls.BallStatuses) == numBalls {
		for i := 0; i < len(generatedBalls.BallStatuses); i++ {
			if generatedBalls.BallStatuses[i].CurrentPosition != generatedBalls.BallStatuses[i].OriginalPosition {
				return false
			}
		}
	} else {
		return false
	}
	return true
}

func findLeastUsedBall(generatedBalls BallList) BallStatus {
	return generatedBalls.BallStatuses[0]
}

func UpdateCurrentPositionBallList(generatedBalls BallList) BallList {
	for i := 0; i < len(generatedBalls.BallStatuses); i++{
		generatedBalls.BallStatuses[i].CurrentPosition = "mainRod" + strconv.Itoa(i)
	}
	return generatedBalls
}

func processUntilBallCycleIsRestored(numBalls int, listToProcess BallList, mode int, minuteAmount int) Results {
	var minuteRod BallList
	minuteRod.Id = 2
	minuteRod.Name = "minuteRodCurrentList"
	var fiveMinuteIncrementRod BallList
	fiveMinuteIncrementRod.Id = 3
	fiveMinuteIncrementRod.Name = "fiveMinuteIncrementRodCurrentList"
	var twelveHourRod BallList
	twelveHourRod.Id = 5
	twelveHourRod.Name = "twelveHourRodCurrentList"
	var mainRod BallList = listToProcess
	mainRod.Id = 6
	mainRod.Name = "mainRodCurrentList"
	var minute_tracker int = 0
	var five_minute_tracker int = 0
	var hour_tracker int = 0
	var half_days_passed int = 0
	var transfer_minute_rod BallStatus
	var iteratorMax int = 0;
	if mode == 1 {
		iteratorMax = 9999999
	} else if mode == 2 {
		iteratorMax = minuteAmount
	}
	
	for z := 0; z < iteratorMax; z++ {
		mainRod = UpdateCurrentPositionBallList(mainRod)
		if z > 10 {
			if ballCycleIsRestored(mainRod, numBalls) == true {
				break
			}
		}

		//Begin Clock Operation -- least used ball or least used ball that is in queue first
		// For now will run as if every iterator is a minute
		// while if iterator is not 0 and current position equals originalPosition then return ball cycles after # of days
		transfer_minute_rod = findLeastUsedBall(mainRod)
		minute_tracker += 1

		// minuteRod counter needed
		// , if minuteRod length contains 4 (5 in this case)
		// , then move last ball over to 5 minute increment while other balls go to mainRod (reverse order)
		if len(minuteRod.BallStatuses) < 4 {
			minuteRod.BallStatuses = append(minuteRod.BallStatuses, BallStatus{
				Id:   transfer_minute_rod.Id,
				BallName: transfer_minute_rod.BallName,
				CurrentPosition: transfer_minute_rod.CurrentPosition,
				OriginalPosition: transfer_minute_rod.OriginalPosition,
			})
			mainRod.BallStatuses = append(mainRod.BallStatuses[:0], mainRod.BallStatuses[1:]...)
			mainRod = UpdateCurrentPositionBallList(mainRod)
		} else {
			// 5 minute increment counter needed
			// , if fiveMinuteIncrementRod length contains 11 (12)
			// , then move last ball over to hour rod, while others go to mainRod (reverse order)
			mainRod.BallStatuses = append(mainRod.BallStatuses[:0], mainRod.BallStatuses[1:]...)
			mainRod = UpdateCurrentPositionBallList(mainRod)
			for i := len(minuteRod.BallStatuses) - 1; i > -1; i-- {
				mainRod.BallStatuses = append(mainRod.BallStatuses, BallStatus{
					Id:   minuteRod.BallStatuses[i].Id,
					BallName: minuteRod.BallStatuses[i].BallName,
					CurrentPosition: minuteRod.BallStatuses[i].CurrentPosition,
					OriginalPosition: minuteRod.BallStatuses[i].OriginalPosition,
				})
				minuteRod.BallStatuses = append(minuteRod.BallStatuses[:i], minuteRod.BallStatuses[i+1:]...)
			}
			mainRod = UpdateCurrentPositionBallList(mainRod)
			
			// hour counter needed and half day counter needed
			// , if mainRod length contains 10 (11)
			// , then move all balls over to MainRod since there is always a ball (reverse order)
			if len(fiveMinuteIncrementRod.BallStatuses) < 11 {
				fiveMinuteIncrementRod.BallStatuses = append(fiveMinuteIncrementRod.BallStatuses, BallStatus{
					Id:   transfer_minute_rod.Id,
					BallName: transfer_minute_rod.BallName,
					CurrentPosition: transfer_minute_rod.CurrentPosition,
					OriginalPosition: transfer_minute_rod.OriginalPosition,
				})
				five_minute_tracker += 1
			} else {
				hour_tracker += 1
				for i := len(fiveMinuteIncrementRod.BallStatuses) - 1; i > -1; i-- {
					mainRod.BallStatuses = append(mainRod.BallStatuses, BallStatus{
						Id:   fiveMinuteIncrementRod.BallStatuses[i].Id,
						BallName: fiveMinuteIncrementRod.BallStatuses[i].BallName,
						CurrentPosition: fiveMinuteIncrementRod.BallStatuses[i].CurrentPosition,
						OriginalPosition: fiveMinuteIncrementRod.BallStatuses[i].OriginalPosition,
					})
					fiveMinuteIncrementRod.BallStatuses = append(fiveMinuteIncrementRod.BallStatuses[:i], fiveMinuteIncrementRod.BallStatuses[i+1:]...)
				}
				mainRod = UpdateCurrentPositionBallList(mainRod)
				
				// 12 hour section
				if len(twelveHourRod.BallStatuses) < 11 {
					twelveHourRod.BallStatuses = append(twelveHourRod.BallStatuses, BallStatus{
						Id:   transfer_minute_rod.Id,
						BallName: transfer_minute_rod.BallName,
						CurrentPosition: transfer_minute_rod.CurrentPosition,
						OriginalPosition: transfer_minute_rod.OriginalPosition,
					})
				} else {
					half_days_passed += 1
					for i := len(twelveHourRod.BallStatuses) - 1; i > -1; i-- {
						mainRod.BallStatuses = append(mainRod.BallStatuses, BallStatus{
							Id:   twelveHourRod.BallStatuses[i].Id,
							BallName: twelveHourRod.BallStatuses[i].BallName,
							CurrentPosition: twelveHourRod.BallStatuses[i].CurrentPosition,
							OriginalPosition: twelveHourRod.BallStatuses[i].OriginalPosition,
						})
						twelveHourRod.BallStatuses = append(twelveHourRod.BallStatuses[:i], twelveHourRod.BallStatuses[i+1:]...)	
					}
					mainRod.BallStatuses = append(mainRod.BallStatuses, BallStatus{
						Id:   transfer_minute_rod.Id,
						BallName: transfer_minute_rod.BallName,
						CurrentPosition: transfer_minute_rod.CurrentPosition,
						OriginalPosition: transfer_minute_rod.OriginalPosition,
					})
					mainRod = UpdateCurrentPositionBallList(mainRod)
				}
			}
		}
	}
	fmt.Println("minutes passed: ", minute_tracker)
	fmt.Println("hours passed: ", hour_tracker)
	var finalReturn Results
	finalReturn.Days = (half_days_passed / 2)
	return finalReturn
}


func mode01(numBalls int, mode int, minuteAmount int) string {

	if numBalls > 26 && numBalls < 128 {

		// Need dictionary to show ballName, currentPosition, originalPosition
		var myBallList BallList

		myBallList.Id = 1
		myBallList.Name = "BallListMode01"

		for i := 0; i < (numBalls); i++ {
			myBallList.BallStatuses = append(myBallList.BallStatuses, BallStatus{
				Id:   i,
				BallName: "ballName" + strconv.Itoa(i),
				CurrentPosition: "mainRod" + strconv.Itoa(i),
				OriginalPosition: "mainRod" + strconv.Itoa(i),
			})
		}
		daysPassed := processUntilBallCycleIsRestored(numBalls, myBallList, mode, minuteAmount)
		return strconv.Itoa(numBalls) + " balls cycle after " + strconv.Itoa(daysPassed.Days) + " days. "

	} else {
		return "Please pass in numbers from 27 to 127."
	}
}

func main() {
	fmt.Println(mode01(45, 1, 0))
}

// Planning to output to JSON
//fmt.Println(myBallList.BallStatuses[0].Id)
/*
json1, err := json.Marshal(myBallList)
if err != nil {
log.Fatal("Cannot encode to JSON ", err)
}
//fmt.Fprintf(os.Stdout, "%s", json1)
var myBallList1 BallList 
bytes := []byte(json1)
json.Unmarshal(bytes, &myBallList1)
fmt.Println(myBallList1)
*/