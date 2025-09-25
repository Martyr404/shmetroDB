package orm

import (
	"fmt"
	"strconv"
)

type TrainInfo struct {
	TrainId         string
	Carriage_number []string
	Carriage_type   []string
	TrainDetail     string
}
type Error struct {
	Code    string
	Msg     string
	Verbose error
}

func ParseCarriageNumber(number string) (TrainInfo, *Error) {
	if len(number) == 6 {
		//_ represent the user input carriage type
		line_num, carriage_num, _ := number[:2], number[2:5], number[5:]
		switch line_num {
		case "01":
			{
				//line 1 logic
				trainInfo := TrainInfo{}
				e := &Error{Msg: "to be realized"}
				return trainInfo, e
			}
		case "02":
			{
				//line 2 logic
				trainInfo := TrainInfo{}
				e := &Error{Msg: "to be realized"}
				return trainInfo, e
			}
		case "03":
			{
				//line 3 logic
				trainInfo := TrainInfo{}
				e := &Error{Msg: "to be realized"}
				return trainInfo, e
			}
		case "04":
			{
				//line 4 logic
				trainInfo := TrainInfo{}
				e := &Error{Msg: "to be realized"}
				return trainInfo, e
			}
		case "05":
			{
				//line 5 logic
				trainInfo := TrainInfo{}
				e := &Error{Msg: "to be realized"}
				return trainInfo, e
			}
		case "06":
			{
				//line 6 logic
				trainInfo := TrainInfo{}
				e := &Error{Msg: "to be realized"}
				return trainInfo, e
			}
		case "07":
			{
				//calculate logic
				carriage_num_int, err := strconv.Atoi(carriage_num)
				if err != nil {
					return TrainInfo{}, &Error{Code: "0003", Msg: "Invalid carriage number."}
				}
				calculated_id := carriage_num_int/6 + 1
				carriage_nums, _ := FormatCarriageNumbers(calculated_id, "07", false)
				trainInfo := TrainInfo{
					//若四位数车号700改7000
					TrainId:         strconv.Itoa(700 + calculated_id),
					Carriage_number: carriage_nums,
					Carriage_type:   []string{"1", "2", "3", "3", "2", "1"},
				}
				//judge carriage type is correct
				for _, value := range carriage_nums {
					if number == value {
						return trainInfo, nil
					}
				}
				return trainInfo, &Error{Code: "0006", Msg: "Incorrect carriage type."}
			}
		case "08":
			{
				//line 8 logic
				trainInfo := TrainInfo{}
				e := &Error{Msg: "to be realized"}
				return trainInfo, e
			}
		case "09":
			{
				//line 9 logic
				trainInfo := TrainInfo{}
				e := &Error{Msg: "to be realized"}
				return trainInfo, e
			}
		case "10":
			{
				//line 10 logic
				trainInfo := TrainInfo{}
				e := &Error{Msg: "to be realized"}
				return trainInfo, e
			}
		case "11":
			{
				//line 11 logic
				trainInfo := TrainInfo{}
				e := &Error{Msg: "to be realized"}
				return trainInfo, e
			}
		case "12":
			{
				//line 12 logic
				trainInfo := TrainInfo{}
				e := &Error{Msg: "to be realized"}
				return trainInfo, e
			}
		case "13":
			{
				//line 13 logic
				trainInfo := TrainInfo{}
				e := &Error{Msg: "to be realized"}
				return trainInfo, e
			}
		case "14":
			{
				//line 14 logic
				trainInfo := TrainInfo{}
				e := &Error{Msg: "to be realized"}
				return trainInfo, e
			}
		case "15":
			{
				//line 15 logic
				trainInfo := TrainInfo{}
				e := &Error{Msg: "to be realized"}
				return trainInfo, e
			}
		case "16":
			{
				//line 16 logic
				trainInfo := TrainInfo{}
				e := &Error{Msg: "to be realized"}
				return trainInfo, e
			}
		case "17":
			{
				//line 17 logic
				trainInfo := TrainInfo{}
				e := &Error{Msg: "to be realized"}
				return trainInfo, e
			}
		case "18":
			{
				//line 18 logic
				trainInfo := TrainInfo{}
				e := &Error{Msg: "to be realized"}
				return trainInfo, e
			}
		default:
			{
				return TrainInfo{}, &Error{Code: "0001", Msg: "Unknown type carriage number."}
			}
		}
	} else if len(number) == 5 {
		line_num, _ /*carriage_num*/, _ := number[:2], number[2:4], number[4:]
		switch line_num {
		case "92":
			{
				//line 1 logic
				trainInfo := TrainInfo{}
				e := &Error{Msg: "to be realized"}
				return trainInfo, e
			}
		case "93":
			{
				//line 1 logic
				trainInfo := TrainInfo{}
				e := &Error{Msg: "to be realized"}
				return trainInfo, e
			}
		case "94":
			{
				//line 1 logic
				trainInfo := TrainInfo{}
				e := &Error{Msg: "to be realized"}
				return trainInfo, e
			}
		case "98":
			{
				//line 1 logic
				trainInfo := TrainInfo{}
				e := &Error{Msg: "to be realized"}
				return trainInfo, e
			}
		case "99":
			{
				//line 1 logic
				trainInfo := TrainInfo{}
				e := &Error{Msg: "to be realized"}
				return trainInfo, e
			}
		case "00":
			{
				//line 1/2 logic
				trainInfo := TrainInfo{}
				e := &Error{Msg: "to be realized"}
				return trainInfo, e
			}
		case "01":
			{
				//line 1/2 logic
				trainInfo := TrainInfo{}
				e := &Error{Msg: "to be realized"}
				return trainInfo, e
			}
		case "02":
			{
				//line 3/5 logic
				trainInfo := TrainInfo{}
				e := &Error{Msg: "to be realized"}
				return trainInfo, e
			}
		case "03":
			{
				//line 3/5 logic
				trainInfo := TrainInfo{}
				e := &Error{Msg: "to be realized"}
				return trainInfo, e
			}
		case "04":
			{
				//line 3/5 logic
				trainInfo := TrainInfo{}
				e := &Error{Msg: "to be realized"}
				return trainInfo, e
			}
		default:
			{
				return TrainInfo{}, &Error{Code: "0002", Msg: "Unknown type carriage number"}
			}
		}
	} else {
		return TrainInfo{}, &Error{Code: "0005", Msg: "Invalid carriage number"}
	}
}

// 根据车号生成车厢号
func FormatCarriageNumbers(id int, line string, isAnda bool) ([]string, error) {
	if id > 0 {
		num := []int{6*id - 5, 6*id - 4, 6*id - 3, 6*id - 2, 6*id - 1, 6 * id}
		switch isAnda {
		case true:
			{
				template := []string{"1", "2", "3", "2", "3", "1"}
				res := []string{}
				for i := 0; i < 6; i++ {
					num_str := strconv.Itoa(num[i])
					for len(num_str) < 3 {
						num_str = "0" + num_str
					}
					res = append(res, line+num_str+template[i])
				}
				return res, nil
			}
		case false:
			{
				template := []string{"1", "2", "3", "3", "2", "1"}
				res := []string{}
				for i := 0; i < 6; i++ {
					num_str := strconv.Itoa(num[i])
					for len(num_str) < 3 {
						num_str = "0" + num_str
					}
					res = append(res, line+num_str+template[i])
				}
				return res, nil
			}
		default:
			{
				res := []string{}
				return res, fmt.Errorf("invald isAnda flag , error code : 0007")
			}
		}
	} else {
		return nil, fmt.Errorf("invalid train id , error code : 0004")
	}
}
