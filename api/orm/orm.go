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

func ParseCarriageNumber(number string) (TrainInfo, error) {
	if len(number) == 6 {
		//_ represent the user input carriage type
		line_num, carriage_num, _ := number[:2], number[2:5], number[5:]
		switch line_num {
		case "01":
			{
				//line 1 logic
			}
		case "02":
			{
				//line 2 logic
			}
		case "03":
			{
				//line 3 logic
			}
		case "04":
			{
				//line 4 logic
			}
		case "05":
			{
				//line 5 logic
			}
		case "06":
			{
				//line 6 logic
			}
		case "07":
			{
				//calculate logic
				carriage_num_int, err := strconv.Atoi(carriage_num)
				if err != nil {
					return TrainInfo{}, fmt.Errorf("invalid carriage number , error code : 0003")
				}
				calculated_id := carriage_num_int/6 + 1
				carriage_nums, _ := FormatCarriageNumbers(calculated_id, "07", false)
				trainInfo := TrainInfo{
					TrainId:         strconv.Itoa(calculated_id),
					Carriage_number: carriage_nums,
					Carriage_type:   []string{"1", "2", "3", "3", "2", "1"},
				}
				//judge carriage type is correct
				for _, value := range carriage_nums {
					if number == value {
						return trainInfo, nil
					}
				}
				return trainInfo, fmt.Errorf("incorrect carriage type , error code : 0006")
			}
		case "08":
			{
				//line 8 logic
			}
		case "09":
			{
				//line 9 logic
			}
		case "10":
			{
				//line 10 logic
			}
		case "11":
			{
				//line 11 logic
			}
		case "12":
			{
				//line 12 logic
			}
		case "13":
			{
				//line 13 logic
			}
		case "14":
			{
				//line 14 logic
			}
		case "15":
			{
				//line 15 logic
			}
		case "16":
			{
				//line 16 logic
			}
		case "17":
			{
				//line 17 logic
			}
		case "18":
			{
				//line 18 logic
			}
		default:
			{
				return TrainInfo{}, fmt.Errorf("unknown type carriage number , error code : 0001")
			}
		}
	} else if len(number) == 5 {
		line_num, carriage_num, _ := number[:2], number[2:4], number[4:]
		switch line_num {
		case "92":
			{
				//line 1 logic
			}
		case "93":
			{
				//line 1 logic
			}
		case "94":
			{
				//line 1 logic
			}
		case "98":
			{
				//line 1 logic
			}
		case "99":
			{
				//line 1 logic
			}
		case "00":
			{
				//line 1/2 logic
			}
		case "01":
			{
				//line 1/2 logic
			}
		case "02":
			{
				//line 3/5 logic
			}
		case "03":
			{
				//line 3/5 logic
			}
		case "04":
			{
				//line 3/5 logic
			}
		default:
			{
				return TrainInfo{}, fmt.Errorf("unknown type carriage number , error code : 0002")
			}
		}
	} else {
		return TrainInfo{}, fmt.Errorf("invalid carriage number , error code : 0005")
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
