package orm

import (
	"fmt"
	"regexp"
	"strconv"
)

type TrainInfo struct {
	Pk              int
	TrainId         string
	Train_type      string
	Carriage_number []string
	Carriage_index  string
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
				//line4 logic
				carriage_num_int, err := strconv.Atoi(carriage_num)
				if err != nil {
					return TrainInfo{}, &Error{Code: "0003", Msg: "Invalid carriage number."}
				}
				if carriage_num_int%6 == 0 {
					calculated_id := carriage_num_int / 6
					carriage_nums, _ := FormatCarriageNumbers(calculated_id, "04", false)
					trainInfo := TrainInfo{
						//若三位数车号7000改700
						TrainId:         "0" + strconv.Itoa(4000+calculated_id),
						Carriage_number: carriage_nums,
						Carriage_index:  "5",
					}
					//judge carriage type is correct
					for _, value := range carriage_nums {
						if number == value {
							return trainInfo, nil
						}
					}
					return trainInfo, &Error{Code: "0006", Msg: "Incorrect carriage type."}
				} else {
					calculated_id := carriage_num_int/6 + 1
					carriage_nums, _ := FormatCarriageNumbers(calculated_id, "04", false)
					trainInfo := TrainInfo{
						//若三位数车号9000改900
						TrainId:         "0" + strconv.Itoa(4000+calculated_id),
						Carriage_number: carriage_nums,
						Carriage_index:  strconv.Itoa(carriage_num_int%6 - 1),
					}
					//judge carriage type is correct
					for _, value := range carriage_nums {
						if number == value {
							return trainInfo, nil
						}
					}
					return trainInfo, &Error{Code: "0006", Msg: "Incorrect carriage type."}
				}
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
				//line7 logic
				carriage_num_int, err := strconv.Atoi(carriage_num)
				if err != nil {
					return TrainInfo{}, &Error{Code: "0003", Msg: "Invalid carriage number."}
				}
				if carriage_num_int%6 == 0 {
					calculated_id := carriage_num_int / 6
					carriage_nums, _ := FormatCarriageNumbers(calculated_id, "07", false)
					trainInfo := TrainInfo{
						//若三位数车号7000改700
						TrainId:         "0" + strconv.Itoa(7000+calculated_id),
						Carriage_number: carriage_nums,
						Carriage_index:  "5",
					}
					//judge carriage type is correct
					for _, value := range carriage_nums {
						if number == value {
							return trainInfo, nil
						}
					}
					return trainInfo, &Error{Code: "0006", Msg: "Incorrect carriage type."}
				} else {
					calculated_id := carriage_num_int/6 + 1
					carriage_nums, _ := FormatCarriageNumbers(calculated_id, "07", false)
					trainInfo := TrainInfo{
						//若三位数车号9000改900
						TrainId:         "0" + strconv.Itoa(7000+calculated_id),
						Carriage_number: carriage_nums,
						Carriage_index:  strconv.Itoa(carriage_num_int%6 - 1),
					}
					//judge carriage type is correct
					for _, value := range carriage_nums {
						if number == value {
							return trainInfo, nil
						}
					}
					return trainInfo, &Error{Code: "0006", Msg: "Incorrect carriage type."}
				}
			}
		case "08":
			{
				//line8 logic
				carriage_num_int, err := strconv.Atoi(carriage_num)
				if err != nil {
					return TrainInfo{}, &Error{Code: "0003", Msg: "Invalid carriage number."}
				}
				if carriage_num_int <= 168 {
					//6节车逻辑
					if carriage_num_int%6 == 0 {
						calculated_id := carriage_num_int / 6
						carriage_nums, _ := FormatCarriageNumbers(calculated_id, "08", false)
						trainInfo := TrainInfo{
							//若三位数车号7000改700
							TrainId:         "0" + strconv.Itoa(8000+calculated_id),
							Carriage_number: carriage_nums,
							Carriage_index:  "5",
						}
						//judge carriage type is correct
						for _, value := range carriage_nums {
							if number == value {
								return trainInfo, nil
							}
						}
						return trainInfo, &Error{Code: "0006", Msg: "Incorrect carriage type."}
					} else {
						calculated_id := carriage_num_int/6 + 1
						carriage_nums, _ := FormatCarriageNumbers(calculated_id, "08", false)
						trainInfo := TrainInfo{
							//若三位数车号9000改900
							TrainId:         "0" + strconv.Itoa(8000+calculated_id),
							Carriage_number: carriage_nums,
							Carriage_index:  strconv.Itoa(carriage_num_int%6 - 1),
						}
						//judge carriage type is correct
						for _, value := range carriage_nums {
							if number == value {
								return trainInfo, nil
							}
						}
						return trainInfo, &Error{Code: "0006", Msg: "Incorrect carriage type."}
					}
				} else {
					//7节车逻辑
					if (carriage_num_int-168)%7 == 0 {
						calculated_id := (carriage_num_int-168)/7 + 28
						carriage_nums, _ := FormatCarriageNumbers(calculated_id, "08", false)
						trainInfo := TrainInfo{
							//若三位数车号7000改700
							TrainId:         "0" + strconv.Itoa(8000+calculated_id),
							Carriage_number: carriage_nums,
							Carriage_index:  "6",
						}
						//judge carriage type is correct
						for _, value := range carriage_nums {
							if number == value {
								return trainInfo, nil
							}
						}
						return trainInfo, &Error{Code: "0006", Msg: "Incorrect carriage type."}
					} else {
						calculated_id := (carriage_num_int-168)/7 + 29
						carriage_nums, _ := FormatCarriageNumbers(calculated_id, "08", false)
						trainInfo := TrainInfo{
							//若三位数车号9000改900
							TrainId:         "0" + strconv.Itoa(8000+calculated_id),
							Carriage_number: carriage_nums,
							Carriage_index:  strconv.Itoa(carriage_num_int%7 - 1),
						}
						//judge carriage type is correct
						for _, value := range carriage_nums {
							if number == value {
								return trainInfo, nil
							}
						}
						return trainInfo, &Error{Code: "0006", Msg: "Incorrect carriage type."}
					}
				}

			}
		case "09":
			{
				//line9 logic
				carriage_num_int, err := strconv.Atoi(carriage_num)
				if err != nil {
					return TrainInfo{}, &Error{Code: "0003", Msg: "Invalid carriage number."}
				}
				if carriage_num_int%6 == 0 {
					calculated_id := carriage_num_int / 6
					carriage_nums, _ := FormatCarriageNumbers(calculated_id, "09", false)
					trainInfo := TrainInfo{
						//若三位数车号9000改900
						TrainId:         "0" + strconv.Itoa(9000+calculated_id),
						Carriage_number: carriage_nums,
						Carriage_index:  "5",
					}
					//judge carriage type is correct
					for _, value := range carriage_nums {
						if number == value {
							return trainInfo, nil
						}
					}
					return trainInfo, &Error{Code: "0006", Msg: "Incorrect carriage type."}
				} else {
					calculated_id := carriage_num_int/6 + 1
					carriage_nums, _ := FormatCarriageNumbers(calculated_id, "09", false)
					trainInfo := TrainInfo{
						//若三位数车号9000改900
						TrainId:         "0" + strconv.Itoa(9000+calculated_id),
						Carriage_number: carriage_nums,
						Carriage_index:  strconv.Itoa(carriage_num_int%6 - 1),
					}
					//judge carriage type is correct
					for _, value := range carriage_nums {
						if number == value {
							return trainInfo, nil
						}
					}
					return trainInfo, &Error{Code: "0006", Msg: "Incorrect carriage type."}
				}
			}
		case "10":
			{
				//line10 logic
				carriage_num_int, err := strconv.Atoi(carriage_num)
				if err != nil {
					return TrainInfo{}, &Error{Code: "0003", Msg: "Invalid carriage number."}
				}
				if carriage_num_int%6 == 0 {
					calculated_id := carriage_num_int / 6
					carriage_nums, _ := FormatCarriageNumbers(calculated_id, "10", false)
					trainInfo := TrainInfo{
						//若三位数车号12000改1200
						TrainId:         strconv.Itoa(10000 + calculated_id),
						Carriage_number: carriage_nums,
						Carriage_index:  "5",
					}
					//judge carriage type is correct
					for _, value := range carriage_nums {
						if number == value {
							return trainInfo, nil
						}
					}
					return trainInfo, &Error{Code: "0006", Msg: "Incorrect carriage type."}
				} else {
					calculated_id := carriage_num_int/6 + 1
					carriage_nums, _ := FormatCarriageNumbers(calculated_id, "10", false)
					trainInfo := TrainInfo{
						TrainId:         strconv.Itoa(10000 + calculated_id),
						Carriage_number: carriage_nums,
						Carriage_index:  strconv.Itoa(carriage_num_int%6 - 1),
					}
					//judge carriage type is correct
					for _, value := range carriage_nums {
						if number == value {
							return trainInfo, nil
						}
					}
					return trainInfo, &Error{Code: "0006", Msg: "Incorrect carriage type."}
				}
			}
		case "11":
			{
				//line11 logic
				carriage_num_int, err := strconv.Atoi(carriage_num)
				if err != nil {
					return TrainInfo{}, &Error{Code: "0003", Msg: "Invalid carriage number."}
				}
				if carriage_num_int%6 == 0 {
					calculated_id := carriage_num_int / 6
					carriage_nums, _ := FormatCarriageNumbers(calculated_id, "11", false)
					trainInfo := TrainInfo{
						//若三位数车号12000改1200
						TrainId:         strconv.Itoa(11000 + calculated_id),
						Carriage_number: carriage_nums,
						Carriage_index:  "5",
					}
					//judge carriage type is correct
					for _, value := range carriage_nums {
						if number == value {
							return trainInfo, nil
						}
					}
					return trainInfo, &Error{Code: "0006", Msg: "Incorrect carriage type."}
				} else {
					calculated_id := carriage_num_int/6 + 1
					carriage_nums, _ := FormatCarriageNumbers(calculated_id, "11", false)
					trainInfo := TrainInfo{
						TrainId:         strconv.Itoa(11000 + calculated_id),
						Carriage_number: carriage_nums,
						Carriage_index:  strconv.Itoa(carriage_num_int%6 - 1),
					}
					//judge carriage type is correct
					for _, value := range carriage_nums {
						if number == value {
							return trainInfo, nil
						}
					}
					return trainInfo, &Error{Code: "0006", Msg: "Incorrect carriage type."}
				}
			}
		case "12":
			{
				//line12 logic
				carriage_num_int, err := strconv.Atoi(carriage_num)
				if err != nil {
					return TrainInfo{}, &Error{Code: "0003", Msg: "Invalid carriage number."}
				}
				if carriage_num_int%6 == 0 {
					calculated_id := carriage_num_int / 6
					carriage_nums, _ := FormatCarriageNumbers(calculated_id, "12", false)
					trainInfo := TrainInfo{
						//若三位数车号12000改1200
						TrainId:         strconv.Itoa(12000 + calculated_id),
						Carriage_number: carriage_nums,
						Carriage_index:  "5",
					}
					//judge carriage type is correct
					for _, value := range carriage_nums {
						if number == value {
							return trainInfo, nil
						}
					}
					return trainInfo, &Error{Code: "0006", Msg: "Incorrect carriage type."}
				} else {
					calculated_id := carriage_num_int/6 + 1
					carriage_nums, _ := FormatCarriageNumbers(calculated_id, "12", false)
					trainInfo := TrainInfo{
						TrainId:         strconv.Itoa(12000 + calculated_id),
						Carriage_number: carriage_nums,
						Carriage_index:  strconv.Itoa(carriage_num_int%6 - 1),
					}
					//judge carriage type is correct
					for _, value := range carriage_nums {
						if number == value {
							return trainInfo, nil
						}
					}
					return trainInfo, &Error{Code: "0006", Msg: "Incorrect carriage type."}
				}
			}
		case "13":
			{
				//line13 logic
				carriage_num_int, err := strconv.Atoi(carriage_num)
				if err != nil {
					return TrainInfo{}, &Error{Code: "0003", Msg: "Invalid carriage number."}
				}
				if carriage_num_int%6 == 0 {
					calculated_id := carriage_num_int / 6
					carriage_nums, _ := FormatCarriageNumbers(calculated_id, "13", false)
					trainInfo := TrainInfo{
						//若三位数车号12000改1200
						TrainId:         strconv.Itoa(13000 + calculated_id),
						Carriage_number: carriage_nums,
						Carriage_index:  "5",
					}
					//judge carriage type is correct
					for _, value := range carriage_nums {
						if number == value {
							return trainInfo, nil
						}
					}
					return trainInfo, &Error{Code: "0006", Msg: "Incorrect carriage type."}
				} else {
					calculated_id := carriage_num_int/6 + 1
					carriage_nums, _ := FormatCarriageNumbers(calculated_id, "13", false)
					trainInfo := TrainInfo{
						TrainId:         strconv.Itoa(13000 + calculated_id),
						Carriage_number: carriage_nums,
						Carriage_index:  strconv.Itoa(carriage_num_int%6 - 1),
					}
					//judge carriage type is correct
					for _, value := range carriage_nums {
						if number == value {
							return trainInfo, nil
						}
					}
					return trainInfo, &Error{Code: "0006", Msg: "Incorrect carriage type."}
				}
			}
		case "14":
			{
				//line14 logic
				carriage_num_int, err := strconv.Atoi(carriage_num)
				if err != nil {
					return TrainInfo{}, &Error{Code: "0003", Msg: "Invalid carriage number."}
				}
				if carriage_num_int%6 == 0 {
					calculated_id := carriage_num_int / 6
					carriage_nums, _ := FormatCarriageNumbers(calculated_id, "14", false)
					trainInfo := TrainInfo{
						//若三位数车号12000改1200
						TrainId:         strconv.Itoa(14000 + calculated_id),
						Carriage_number: carriage_nums,
						Carriage_index:  "5",
					}
					//judge carriage type is correct
					for _, value := range carriage_nums {
						if number == value {
							return trainInfo, nil
						}
					}
					return trainInfo, &Error{Code: "0006", Msg: "Incorrect carriage type."}
				} else {
					calculated_id := carriage_num_int/6 + 1
					carriage_nums, _ := FormatCarriageNumbers(calculated_id, "14", false)
					trainInfo := TrainInfo{
						TrainId:         strconv.Itoa(14000 + calculated_id),
						Carriage_number: carriage_nums,
						Carriage_index:  strconv.Itoa(carriage_num_int%6 - 1),
					}
					//judge carriage type is correct
					for _, value := range carriage_nums {
						if number == value {
							return trainInfo, nil
						}
					}
					return trainInfo, &Error{Code: "0006", Msg: "Incorrect carriage type."}
				}
			}
		case "15":
			{
				//line15 logic
				carriage_num_int, err := strconv.Atoi(carriage_num)
				if err != nil {
					return TrainInfo{}, &Error{Code: "0003", Msg: "Invalid carriage number."}
				}
				if carriage_num_int%6 == 0 {
					calculated_id := carriage_num_int / 6
					carriage_nums, _ := FormatCarriageNumbers(calculated_id, "15", false)
					trainInfo := TrainInfo{
						//若三位数车号12000改1200
						TrainId:         strconv.Itoa(15000 + calculated_id),
						Carriage_number: carriage_nums,
						Carriage_index:  "5",
					}
					//judge carriage type is correct
					for _, value := range carriage_nums {
						if number == value {
							return trainInfo, nil
						}
					}
					return trainInfo, &Error{Code: "0006", Msg: "Incorrect carriage type."}
				} else {
					calculated_id := carriage_num_int/6 + 1
					carriage_nums, _ := FormatCarriageNumbers(calculated_id, "15", false)
					trainInfo := TrainInfo{
						TrainId:         strconv.Itoa(15000 + calculated_id),
						Carriage_number: carriage_nums,
						Carriage_index:  strconv.Itoa(carriage_num_int%6 - 1),
					}
					//judge carriage type is correct
					for _, value := range carriage_nums {
						if number == value {
							return trainInfo, nil
						}
					}
					return trainInfo, &Error{Code: "0006", Msg: "Incorrect carriage type."}
				}
			}
		case "16":
			{
				//line16 logic
				carriage_num_int, err := strconv.Atoi(carriage_num)
				if err != nil {
					return TrainInfo{}, &Error{Code: "0003", Msg: "Invalid carriage number."}
				}
				if carriage_num_int <= 138 {
					//3节车逻辑
					if carriage_num_int%3 == 0 {
						calculated_id := carriage_num_int / 3
						carriage_nums, _ := FormatCarriageNumbers(calculated_id, "16", false)
						trainInfo := TrainInfo{
							//若三位数车号7000改700
							TrainId:         strconv.Itoa(16000 + calculated_id),
							Carriage_number: carriage_nums,
							Carriage_index:  "2",
						}
						//judge carriage type is correct
						for _, value := range carriage_nums {
							if number == value {
								return trainInfo, nil
							}
						}
						return trainInfo, &Error{Code: "0006", Msg: "Incorrect carriage type."}
					} else {
						calculated_id := carriage_num_int/3 + 1
						carriage_nums, _ := FormatCarriageNumbers(calculated_id, "16", false)
						trainInfo := TrainInfo{
							//若三位数车号9000改900
							TrainId:         strconv.Itoa(16000 + calculated_id),
							Carriage_number: carriage_nums,
							Carriage_index:  strconv.Itoa(carriage_num_int%3 - 1),
						}
						//judge carriage type is correct
						for _, value := range carriage_nums {
							if number == value {
								return trainInfo, nil
							}
						}
						return trainInfo, &Error{Code: "0006", Msg: "Incorrect carriage type."}
					}
				} else {
					//7节车逻辑
					if (carriage_num_int-138)%6 == 0 {
						calculated_id := (carriage_num_int-138)/6 + 46
						carriage_nums, _ := FormatCarriageNumbers(calculated_id, "16", false)
						trainInfo := TrainInfo{
							//若三位数车号7000改700
							TrainId:         strconv.Itoa(16000 + calculated_id),
							Carriage_number: carriage_nums,
							Carriage_index:  "5",
						}
						//judge carriage type is correct
						for _, value := range carriage_nums {
							if number == value {
								return trainInfo, nil
							}
						}
						return trainInfo, &Error{Code: "0006", Msg: "Incorrect carriage type."}
					} else {
						calculated_id := (carriage_num_int-138)/6 + 46
						carriage_nums, _ := FormatCarriageNumbers(calculated_id, "16", false)
						trainInfo := TrainInfo{
							//若三位数车号9000改900
							TrainId:         strconv.Itoa(16000 + calculated_id),
							Carriage_number: carriage_nums,
							Carriage_index:  strconv.Itoa(carriage_num_int%6 - 1),
						}
						//judge carriage type is correct
						for _, value := range carriage_nums {
							if number == value {
								return trainInfo, nil
							}
						}
						return trainInfo, &Error{Code: "0006", Msg: "Incorrect carriage type."}
					}
				}

			}
		case "17":
			{
				//line17 logic
				carriage_num_int, err := strconv.Atoi(carriage_num)
				if err != nil {
					return TrainInfo{}, &Error{Code: "0003", Msg: "Invalid carriage number."}
				}
				if carriage_num_int%6 == 0 {
					calculated_id := carriage_num_int / 6
					carriage_nums, _ := FormatCarriageNumbers(calculated_id, "17", false)
					trainInfo := TrainInfo{
						//若三位数车号12000改1200
						TrainId:         strconv.Itoa(17000 + calculated_id),
						Carriage_number: carriage_nums,
						Carriage_index:  "5",
					}
					//judge carriage type is correct
					for _, value := range carriage_nums {
						if number == value {
							return trainInfo, nil
						}
					}
					return trainInfo, &Error{Code: "0006", Msg: "Incorrect carriage type."}
				} else {
					calculated_id := carriage_num_int/6 + 1
					carriage_nums, _ := FormatCarriageNumbers(calculated_id, "17", false)
					trainInfo := TrainInfo{
						TrainId:         strconv.Itoa(17000 + calculated_id),
						Carriage_number: carriage_nums,
						Carriage_index:  strconv.Itoa(carriage_num_int%6 - 1),
					}
					//judge carriage type is correct
					for _, value := range carriage_nums {
						if number == value {
							return trainInfo, nil
						}
					}
					return trainInfo, &Error{Code: "0006", Msg: "Incorrect carriage type."}
				}
			}
		case "18":
			{
				//line18 logic
				carriage_num_int, err := strconv.Atoi(carriage_num)
				if err != nil {
					return TrainInfo{}, &Error{Code: "0003", Msg: "Invalid carriage number."}
				}
				if carriage_num_int%6 == 0 {
					calculated_id := carriage_num_int / 6
					carriage_nums, _ := FormatCarriageNumbers(calculated_id, "18", false)
					trainInfo := TrainInfo{
						//若三位数车号12000改1200
						TrainId:         strconv.Itoa(18000 + calculated_id),
						Carriage_number: carriage_nums,
						Carriage_index:  "5",
					}
					//judge carriage type is correct
					for _, value := range carriage_nums {
						if number == value {
							return trainInfo, nil
						}
					}
					return trainInfo, &Error{Code: "0006", Msg: "Incorrect carriage type."}
				} else {
					calculated_id := carriage_num_int/6 + 1
					carriage_nums, _ := FormatCarriageNumbers(calculated_id, "18", false)
					trainInfo := TrainInfo{
						TrainId:         strconv.Itoa(18000 + calculated_id),
						Carriage_number: carriage_nums,
						Carriage_index:  strconv.Itoa(carriage_num_int%6 - 1),
					}
					//judge carriage type is correct
					for _, value := range carriage_nums {
						if number == value {
							return trainInfo, nil
						}
					}
					return trainInfo, &Error{Code: "0006", Msg: "Incorrect carriage type."}
				}
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
		switch line {
		case "06":
			{
				//line6
				res := []string{}
				return res, nil
			}
		case "16":
			{
				if id <= 46 {
					num := []int{3*id - 2, 3*id - 1, 3 * id}
					template := []string{"1", "3", "1"}
					res := []string{}
					for i := 0; i < 3; i++ {
						num_str := strconv.Itoa(num[i])
						for len(num_str) < 3 {
							num_str = "0" + num_str
						}
						res = append(res, line+num_str+template[i])
					}
					return res, nil
				} else {
					num := []int{(id-46)*6 + 133, (id-46)*6 + 134, (id-46)*6 + 135, (id-46)*6 + 136, (id-46)*6 + 137, (id-46)*6 + 138}
					template := []string{"1", "2", "3", "3", "2", "1"}
					res := []string{}
					for i := 0; i < 6; i++ {
						num_str := strconv.Itoa(num[i])
						res = append(res, line+num_str+template[i])
					}
					return res, nil
				}
			}
		case "08":
			{
				if id <= 28 {
					num := []int{6*id - 5, 6*id - 4, 6*id - 3, 6*id - 2, 6*id - 1, 6 * id}
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
				} else {
					num := []int{7*(id-28) + 162, 7*(id-28) + 163, 7*(id-28) + 164, 7*(id-28) + 165, 7*(id-28) + 166, 7*(id-28) + 167, 7*(id-28) + 168}
					template := []string{"1", "2", "3", "3", "3", "2", "1"}
					res := []string{}
					for i := 0; i < 7; i++ {
						num_str := strconv.Itoa(num[i])
						res = append(res, line+num_str+template[i])
					}
					return res, nil
				}
			}
		default:
			{
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
			}
		}
	} else {
		return nil, fmt.Errorf("invalid train id , error code : 0004")
	}
}

// 验证车厢号的合法性，支持大小写混合
func ValidateCarriageNumbers(number string) bool {
	pattern := `(?i)^(?:\d{5}|\d{6}|T\d{6}|JC\d{5}|JY\d{6})$`
	matched, _ := regexp.MatchString(pattern, number)
	return matched
}
