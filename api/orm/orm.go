package orm

import (
	"fmt"
	"regexp"
	"shmetroDB/psql"
	"strconv"
	"strings"
)

type TrainInfo struct {
	Pk                     int
	IsEmpty                bool
	TrainId                string
	Train_type             string
	Carriage_number        []string
	Carriage_index         string
	TrainDetail            string
	IsInputCarriageCorrect bool
}
type Error struct {
	Code         string
	Msg          string
	Verbose      error
	VerboseError *Error
}

func ParseCarriageNumber(number string) ([]TrainInfo, *Error) {
	if len(number) == 6 {
		//_ represent the user input carriage type
		line_num, carriage_num, _ := number[:2], number[2:5], number[5:]
		switch line_num {
		case "01":
			{
				//line 1 logic
				carriage_num_int, err := strconv.Atoi(carriage_num)
				if err != nil {
					return []TrainInfo{{IsEmpty: true}}, &Error{Code: "0020", Msg: "Invalid carriage number."}
				}
				switch {
				case (carriage_num_int >= 235 && carriage_num_int <= 362):
					{
						if (carriage_num_int-234)%8 == 0 {
							calculated_id := (carriage_num_int-234)/8 + 39
							carriage_nums, _ := FormatCarriageNumbers(calculated_id, "01", false)
							trainInfo := TrainInfo{
								//若三位数车号7000改700
								IsEmpty:         false,
								TrainId:         "0" + strconv.Itoa(1000+calculated_id),
								Carriage_number: carriage_nums,
								Carriage_index:  "7",
							}
							//judge carriage type is correct
							for _, value := range carriage_nums {
								if number == value {
									trainInfo.IsInputCarriageCorrect = true
									return []TrainInfo{trainInfo}, nil
								}
							}
							trainInfo.IsInputCarriageCorrect = false
							return []TrainInfo{trainInfo}, nil
						} else {
							calculated_id := (carriage_num_int-234)/8 + 40
							carriage_nums, _ := FormatCarriageNumbers(calculated_id, "01", false)
							trainInfo := TrainInfo{
								//若三位数车号7000改700
								IsEmpty:         false,
								TrainId:         "0" + strconv.Itoa(1000+calculated_id),
								Carriage_number: carriage_nums,
								Carriage_index:  strconv.Itoa((carriage_num_int-234)%8 - 1),
							}
							//judge carriage type is correct
							for _, value := range carriage_nums {
								if number == value {
									trainInfo.IsInputCarriageCorrect = true
									return []TrainInfo{trainInfo}, nil
								}
							}
							trainInfo.IsInputCarriageCorrect = false
							return []TrainInfo{trainInfo}, nil
						}
					}
				case (carriage_num_int >= 468 && carriage_num_int <= 715):
					{
						if (carriage_num_int-467)%8 == 0 {
							calculated_id := (carriage_num_int-467)/8 + 55
							carriage_nums, _ := FormatCarriageNumbers(calculated_id, "01", false)
							trainInfo := TrainInfo{
								//若三位数车号7000改700
								IsEmpty:         false,
								TrainId:         "0" + strconv.Itoa(1000+calculated_id),
								Carriage_number: carriage_nums,
								Carriage_index:  "7",
							}
							//judge carriage type is correct
							for _, value := range carriage_nums {
								if number == value {
									trainInfo.IsInputCarriageCorrect = true
									return []TrainInfo{trainInfo}, nil
								}
							}
							trainInfo.IsInputCarriageCorrect = false
							return []TrainInfo{trainInfo}, nil
						} else {
							calculated_id := (carriage_num_int-467)/8 + 55
							carriage_nums, _ := FormatCarriageNumbers(calculated_id, "01", false)
							trainInfo := TrainInfo{
								//若三位数车号7000改700
								IsEmpty:         false,
								TrainId:         "0" + strconv.Itoa(1000+calculated_id),
								Carriage_number: carriage_nums,
								Carriage_index:  strconv.Itoa((carriage_num_int-467)%8 - 1),
							}
							//judge carriage type is correct
							for _, value := range carriage_nums {
								if number == value {
									trainInfo.IsInputCarriageCorrect = true
									return []TrainInfo{trainInfo}, nil
								}
							}
							trainInfo.IsInputCarriageCorrect = false
							return []TrainInfo{trainInfo}, nil
						}
					}
				default:
					//针对01A04的六位数车号车厢
					t := TrainInfo{}
					Err := QueryCarriage("1", number, &t)
					if Err != nil {
						e := &Error{Code: "0023", Msg: "Line1 Search 01A04 Carriages Goes Wrong", VerboseError: Err}
						t.IsEmpty = true
						return []TrainInfo{t}, e
					}
					t.IsEmpty = false
					return []TrainInfo{t}, nil
				}
			}
		case "02":
			{
				//line 2 logic
				carriage_num_int, err := strconv.Atoi(carriage_num)
				if err != nil {
					return []TrainInfo{{IsEmpty: true}}, &Error{Code: "0003", Msg: "Invalid carriage number."}
				}
				switch {
				case (carriage_num_int >= 129 && carriage_num_int <= 424):
					{
						//line 2 logic
						if carriage_num_int%8 == 0 {
							calculated_id := carriage_num_int/8 + 16
							carriage_nums, _ := FormatCarriageNumbers(calculated_id, "02", false)
							trainInfo := TrainInfo{
								//若三位数车号7000改700
								IsEmpty:         false,
								TrainId:         "0" + strconv.Itoa(2000+calculated_id),
								Carriage_number: carriage_nums,
								Carriage_index:  "7",
							}
							//judge carriage type is correct
							for _, value := range carriage_nums {
								if number == value {
									trainInfo.IsInputCarriageCorrect = true
									return []TrainInfo{trainInfo}, nil
								}
							}
							trainInfo.IsInputCarriageCorrect = false
							return []TrainInfo{trainInfo}, nil
						} else {
							calculated_id := carriage_num_int/8 + 17
							carriage_nums, _ := FormatCarriageNumbers(calculated_id, "02", false)
							trainInfo := TrainInfo{
								//若三位数车号7000改700
								IsEmpty:         false,
								TrainId:         "0" + strconv.Itoa(2000+calculated_id),
								Carriage_number: carriage_nums,
								Carriage_index:  strconv.Itoa(carriage_num_int%8 - 1),
							}
							//judge carriage type is correct
							for _, value := range carriage_nums {
								if number == value {
									trainInfo.IsInputCarriageCorrect = true
									return []TrainInfo{trainInfo}, nil
								}
							}
							trainInfo.IsInputCarriageCorrect = false
							return []TrainInfo{trainInfo}, nil
						}

					}
				case (carriage_num_int >= 425 && carriage_num_int <= 488):
					{
						//line 2 logic
						if carriage_num_int%4 == 0 {
							calculated_id := carriage_num_int/4 - 37
							carriage_nums, _ := FormatCarriageNumbers(calculated_id, "02", false)
							trainInfo := TrainInfo{
								//若三位数车号7000改700
								IsEmpty:         false,
								TrainId:         "0" + strconv.Itoa(2000+calculated_id),
								Carriage_number: carriage_nums,
								Carriage_index:  "7",
							}
							//judge carriage type is correct
							for _, value := range carriage_nums {
								if number == value {
									trainInfo.IsInputCarriageCorrect = true
									return []TrainInfo{trainInfo}, nil
								}
							}
							trainInfo.IsInputCarriageCorrect = false
							return []TrainInfo{trainInfo}, nil
						} else {
							calculated_id := carriage_num_int/4 - 36
							carriage_nums, _ := FormatCarriageNumbers(calculated_id, "02", false)
							trainInfo := TrainInfo{
								//若三位数车号7000改700
								IsEmpty:         false,
								TrainId:         "0" + strconv.Itoa(2000+calculated_id),
								Carriage_number: carriage_nums,
								Carriage_index:  strconv.Itoa(carriage_num_int%4 - 1),
							}
							//judge carriage type is correct
							for _, value := range carriage_nums {
								if number == value {
									trainInfo.IsInputCarriageCorrect = true
									return []TrainInfo{trainInfo}, nil
								}
							}
							trainInfo.IsInputCarriageCorrect = false
							return []TrainInfo{trainInfo}, nil
						}
					}
				case (carriage_num_int >= 489 && carriage_num_int <= 552):
					{
						//line 2 logic
						if carriage_num_int%4 == 0 {
							calculated_id := carriage_num_int/4 - 53
							carriage_nums, _ := FormatCarriageNumbers(calculated_id, "02", false)
							trainInfo := TrainInfo{
								//若三位数车号7000改700
								IsEmpty:         false,
								TrainId:         "0" + strconv.Itoa(2000+calculated_id),
								Carriage_number: carriage_nums,
								Carriage_index:  "6",
							}
							//judge carriage type is correct
							for _, value := range carriage_nums {
								if number == value {
									trainInfo.IsInputCarriageCorrect = true
									return []TrainInfo{trainInfo}, nil
								}
							}
							trainInfo.IsInputCarriageCorrect = false
							return []TrainInfo{trainInfo}, nil
						} else {
							calculated_id := carriage_num_int/4 - 52
							carriage_nums, _ := FormatCarriageNumbers(calculated_id, "02", false)
							trainInfo := TrainInfo{
								//若三位数车号7000改700
								IsEmpty:         false,
								TrainId:         "0" + strconv.Itoa(2000+calculated_id),
								Carriage_number: carriage_nums,
								Carriage_index:  strconv.Itoa(carriage_num_int%4 + 3),
							}
							//judge carriage type is correct
							for _, value := range carriage_nums {
								if number == value {
									trainInfo.IsInputCarriageCorrect = true
									return []TrainInfo{trainInfo}, nil
								}
							}
							trainInfo.IsInputCarriageCorrect = false
							return []TrainInfo{trainInfo}, nil
						}
					}
				case (carriage_num_int <= 553 && carriage_num_int <= 800):
					{
						//line 2 logic
						if carriage_num_int%8 == 0 {
							calculated_id := carriage_num_int/8 + 16
							carriage_nums, _ := FormatCarriageNumbers(calculated_id, "02", false)
							trainInfo := TrainInfo{
								//若三位数车号7000改700
								IsEmpty:         false,
								TrainId:         "0" + strconv.Itoa(2000+calculated_id),
								Carriage_number: carriage_nums,
								Carriage_index:  "7",
							}
							//judge carriage type is correct
							for _, value := range carriage_nums {
								if number == value {
									trainInfo.IsInputCarriageCorrect = true
									return []TrainInfo{trainInfo}, nil
								}
							}
							trainInfo.IsInputCarriageCorrect = false
							return []TrainInfo{trainInfo}, nil
						} else {
							calculated_id := carriage_num_int/8 + 17
							carriage_nums, _ := FormatCarriageNumbers(calculated_id, "02", false)
							trainInfo := TrainInfo{
								//若三位数车号7000改700
								IsEmpty:         false,
								TrainId:         "0" + strconv.Itoa(2000+calculated_id),
								Carriage_number: carriage_nums,
								Carriage_index:  strconv.Itoa(carriage_num_int%8 - 1),
							}
							//judge carriage type is correct
							for _, value := range carriage_nums {
								if number == value {
									trainInfo.IsInputCarriageCorrect = true
									return []TrainInfo{trainInfo}, nil
								}
							}
							trainInfo.IsInputCarriageCorrect = false
							return []TrainInfo{trainInfo}, nil
						}
					}
				default:
					{
						return []TrainInfo{{IsEmpty: true}}, &Error{Code: "0018", Msg: "Invalid carriage number."}
					}
				}
			}
		case "03":
			{
				//line3 logic
				carriage_num_int, err := strconv.Atoi(carriage_num)
				if err != nil {
					return []TrainInfo{{IsEmpty: true}}, &Error{Code: "0003", Msg: "Invalid carriage number."}
				}
				if carriage_num_int%6 == 0 {
					calculated_id := carriage_num_int / 6
					carriage_nums, _ := FormatCarriageNumbers(calculated_id, "03", false)
					trainInfo := TrainInfo{
						//若三位数车号7000改700
						IsEmpty:         false,
						TrainId:         "0" + strconv.Itoa(3000+calculated_id),
						Carriage_number: carriage_nums,
						Carriage_index:  "5",
					}
					//judge carriage type is correct
					for _, value := range carriage_nums {
						if number == value {
							trainInfo.IsInputCarriageCorrect = true
							return []TrainInfo{trainInfo}, nil
						}
					}
					trainInfo.IsInputCarriageCorrect = false
					return []TrainInfo{trainInfo}, nil
				} else {
					calculated_id := carriage_num_int/6 + 1
					carriage_nums, _ := FormatCarriageNumbers(calculated_id, "03", false)
					trainInfo := TrainInfo{
						//若三位数车号9000改900
						IsEmpty:         false,
						TrainId:         "0" + strconv.Itoa(3000+calculated_id),
						Carriage_number: carriage_nums,
						Carriage_index:  strconv.Itoa(carriage_num_int%6 - 1),
					}
					//judge carriage type is correct
					for _, value := range carriage_nums {
						if number == value {
							trainInfo.IsInputCarriageCorrect = true
							return []TrainInfo{trainInfo}, nil
						}
					}
					trainInfo.IsInputCarriageCorrect = false
					return []TrainInfo{trainInfo}, nil
				}
			}
		case "04":
			{
				//line4 logic
				carriage_num_int, err := strconv.Atoi(carriage_num)
				if err != nil {
					return []TrainInfo{{IsEmpty: true}}, &Error{Code: "0003", Msg: "Invalid carriage number."}
				}
				if carriage_num_int%6 == 0 {
					calculated_id := carriage_num_int / 6
					carriage_nums, _ := FormatCarriageNumbers(calculated_id, "04", false)
					tid := ""
					if calculated_id >= 37 && calculated_id <= 49 {
						tid = "0" + strconv.Itoa(3000+calculated_id)
					} else {
						tid = "0" + strconv.Itoa(4000+calculated_id)
					}
					trainInfo := TrainInfo{
						//若三位数车号7000改700
						IsEmpty:         false,
						TrainId:         tid,
						Carriage_number: carriage_nums,
						Carriage_index:  "5",
					}
					//judge carriage type is correct
					for _, value := range carriage_nums {
						if number == value {
							trainInfo.IsInputCarriageCorrect = true
							return []TrainInfo{trainInfo}, nil
						}
					}
					trainInfo.IsInputCarriageCorrect = false
					return []TrainInfo{trainInfo}, nil
				} else {
					calculated_id := carriage_num_int/6 + 1
					carriage_nums, _ := FormatCarriageNumbers(calculated_id, "04", false)
					tid := ""
					if calculated_id >= 37 && calculated_id <= 49 {
						tid = "0" + strconv.Itoa(3000+calculated_id)
					} else {
						tid = "0" + strconv.Itoa(4000+calculated_id)
					}
					trainInfo := TrainInfo{
						//若三位数车号7000改700
						IsEmpty:         false,
						TrainId:         tid,
						Carriage_number: carriage_nums,
						Carriage_index:  strconv.Itoa(carriage_num_int%6 - 1),
					}
					//judge carriage type is correct
					for _, value := range carriage_nums {
						if number == value {
							trainInfo.IsInputCarriageCorrect = true
							return []TrainInfo{trainInfo}, nil
						}
					}
					trainInfo.IsInputCarriageCorrect = false
					return []TrainInfo{trainInfo}, nil
				}
			}
		case "05":
			{
				//line5 logic
				//6位数情况
				carriage_num_int, err := strconv.Atoi(carriage_num)
				if err != nil {
					return []TrainInfo{{IsEmpty: true}}, &Error{Code: "0003", Msg: "Invalid carriage number."}
				}
				if carriage_num_int >= 69 && carriage_num_int <= 266 {
					//6节车逻辑
					if (carriage_num_int-68)%6 == 0 {
						calculated_id := (carriage_num_int-68)/6 + 18
						carriage_nums, _ := FormatCarriageNumbers(calculated_id, "05", false)
						trainInfo := TrainInfo{
							//若三位数车号7000改700
							IsEmpty:         false,
							TrainId:         "05" + strconv.Itoa(5000+calculated_id),
							Carriage_number: carriage_nums,
							Carriage_index:  "5",
						}
						//judge carriage type is correct
						for _, value := range carriage_nums {
							if number == value {
								trainInfo.IsInputCarriageCorrect = true
								return []TrainInfo{trainInfo}, nil
							}
						}
						trainInfo.IsInputCarriageCorrect = false
						return []TrainInfo{trainInfo}, nil
					} else {
						calculated_id := (carriage_num_int-68)/6 + 19
						carriage_nums, _ := FormatCarriageNumbers(calculated_id, "05", false)
						trainInfo := TrainInfo{
							//若三位数车号9000改900
							IsEmpty:         false,
							TrainId:         "0" + strconv.Itoa(5000+calculated_id),
							Carriage_number: carriage_nums,
							Carriage_index:  strconv.Itoa((carriage_num_int-68)%6 - 1),
						}
						//judge carriage type is correct
						for _, value := range carriage_nums {
							if number == value {
								trainInfo.IsInputCarriageCorrect = true
								return []TrainInfo{trainInfo}, nil
							}
						}
						trainInfo.IsInputCarriageCorrect = false
						return []TrainInfo{trainInfo}, nil
					}
				} else {
					//4节车逻辑
					trainInfo := TrainInfo{}
					E := QueryCarriage("5", number, &trainInfo)
					if E != nil {
						return []TrainInfo{{IsEmpty: true}}, E
					}
					if trainInfo.TrainId == "" {
						trainInfo.IsEmpty = true
						E = &Error{Code: "0018", Msg: "Corresponding train not found."}
						return []TrainInfo{trainInfo}, E
					} else {
						//找到列车 逻辑
						trainInfo.IsEmpty = false
						carriage_nums := trainInfo.Carriage_number
						//judge carriage type is correct
						for _, value := range carriage_nums {
							if number == value {
								trainInfo.IsInputCarriageCorrect = true
								return []TrainInfo{trainInfo}, nil
							}
						}
						trainInfo.IsInputCarriageCorrect = false
						return []TrainInfo{trainInfo}, nil
					}
				}
			}
		case "06":
			{
				//line6 logic
				carriage_num_int, err := strconv.Atoi(carriage_num)
				if err != nil {
					return []TrainInfo{{IsEmpty: true}}, &Error{Code: "0003", Msg: "Invalid carriage number."}
				}
				switch {
				case carriage_num_int >= 1 && carriage_num_int <= 12:
					{
						//能够整除逻辑
						if carriage_num_int%4 == 0 {
							calculated_id := carriage_num_int / 4
							carriage_nums, _ := FormatCarriageNumbers(calculated_id, "06", false)
							trainInfo := TrainInfo{
								IsEmpty:         false,
								TrainId:         "0" + strconv.Itoa(6000+calculated_id),
								Carriage_number: carriage_nums,
								Carriage_index:  "3",
							}
							//judge carriage type is correct
							for _, value := range carriage_nums {
								if number == value {
									trainInfo.IsInputCarriageCorrect = true
									return []TrainInfo{trainInfo}, nil
								}
							}
							trainInfo.IsInputCarriageCorrect = false
							return []TrainInfo{trainInfo}, nil
						} else {
							//不能整除逻辑
							calculated_id := carriage_num_int/4 + 1
							carriage_nums, _ := FormatCarriageNumbers(calculated_id, "06", false)
							trainInfo := TrainInfo{
								IsEmpty:         false,
								TrainId:         "0" + strconv.Itoa(6000+calculated_id),
								Carriage_number: carriage_nums,
								Carriage_index:  strconv.Itoa(carriage_num_int%4 - 1),
							}
							//judge carriage type is correct
							for _, value := range carriage_nums {
								if number == value {
									trainInfo.IsInputCarriageCorrect = true
									return []TrainInfo{trainInfo}, nil
								}
							}
							trainInfo.IsInputCarriageCorrect = false
							return []TrainInfo{trainInfo}, nil
						}
					}
				case carriage_num_int >= 13 && carriage_num_int <= 48:
					{
						//能够整除逻辑
						if carriage_num_int%4 == 0 {
							calculated_id := carriage_num_int/4 + 1
							carriage_nums, _ := FormatCarriageNumbers(calculated_id, "06", false)
							trainInfo := TrainInfo{
								IsEmpty:         false,
								TrainId:         "0" + strconv.Itoa(6000+calculated_id),
								Carriage_number: carriage_nums,
								Carriage_index:  "3",
							}
							//judge carriage type is correct
							for _, value := range carriage_nums {
								if number == value {
									trainInfo.IsInputCarriageCorrect = true
									return []TrainInfo{trainInfo}, nil
								}
							}
							trainInfo.IsInputCarriageCorrect = false
							return []TrainInfo{trainInfo}, nil
						} else {
							//不能整除逻辑
							calculated_id := carriage_num_int/4 + 2
							carriage_nums, _ := FormatCarriageNumbers(calculated_id, "06", false)
							trainInfo := TrainInfo{
								IsEmpty:         false,
								TrainId:         "0" + strconv.Itoa(6000+calculated_id),
								Carriage_number: carriage_nums,
								Carriage_index:  strconv.Itoa(carriage_num_int%4 - 1),
							}
							//judge carriage type is correct
							for _, value := range carriage_nums {
								if number == value {
									trainInfo.IsInputCarriageCorrect = true
									return []TrainInfo{trainInfo}, nil
								}
							}
							trainInfo.IsInputCarriageCorrect = false
							return []TrainInfo{trainInfo}, nil
						}
					}
				case carriage_num_int >= 49 && carriage_num_int <= 84:
					{
						//能够整除逻辑
						if carriage_num_int%4 == 0 {
							calculated_id := carriage_num_int/4 + 2
							carriage_nums, _ := FormatCarriageNumbers(calculated_id, "06", false)
							trainInfo := TrainInfo{
								IsEmpty:         false,
								TrainId:         "0" + strconv.Itoa(6000+calculated_id),
								Carriage_number: carriage_nums,
								Carriage_index:  "3",
							}
							//judge carriage type is correct
							for _, value := range carriage_nums {
								if number == value {
									trainInfo.IsInputCarriageCorrect = true
									return []TrainInfo{trainInfo}, nil
								}
							}
							trainInfo.IsInputCarriageCorrect = false
							return []TrainInfo{trainInfo}, nil
						} else {
							//不能整除逻辑
							calculated_id := carriage_num_int/4 + 3
							carriage_nums, _ := FormatCarriageNumbers(calculated_id, "06", false)
							trainInfo := TrainInfo{
								IsEmpty:         false,
								TrainId:         "0" + strconv.Itoa(6000+calculated_id),
								Carriage_number: carriage_nums,
								Carriage_index:  strconv.Itoa(carriage_num_int%4 - 1),
							}
							//judge carriage type is correct
							for _, value := range carriage_nums {
								if number == value {
									trainInfo.IsInputCarriageCorrect = true
									return []TrainInfo{trainInfo}, nil
								}
							}
							trainInfo.IsInputCarriageCorrect = false
							return []TrainInfo{trainInfo}, nil
						}
					}
				case carriage_num_int >= 85 && carriage_num_int <= 120:
					{
						//能够整除逻辑
						if carriage_num_int%4 == 0 {
							calculated_id := carriage_num_int/4 + 3
							carriage_nums, _ := FormatCarriageNumbers(calculated_id, "06", false)
							trainInfo := TrainInfo{
								IsEmpty:         false,
								TrainId:         "0" + strconv.Itoa(6000+calculated_id),
								Carriage_number: carriage_nums,
								Carriage_index:  "3",
							}
							//judge carriage type is correct
							for _, value := range carriage_nums {
								if number == value {
									trainInfo.IsInputCarriageCorrect = true
									return []TrainInfo{trainInfo}, nil
								}
							}
							trainInfo.IsInputCarriageCorrect = false
							return []TrainInfo{trainInfo}, nil
						} else {
							//不能整除逻辑
							calculated_id := carriage_num_int/4 + 4
							carriage_nums, _ := FormatCarriageNumbers(calculated_id, "06", false)
							trainInfo := TrainInfo{
								IsEmpty:         false,
								TrainId:         "0" + strconv.Itoa(6000+calculated_id),
								Carriage_number: carriage_nums,
								Carriage_index:  strconv.Itoa(carriage_num_int%4 - 1),
							}
							//judge carriage type is correct
							for _, value := range carriage_nums {
								if number == value {
									trainInfo.IsInputCarriageCorrect = true
									return []TrainInfo{trainInfo}, nil
								}
							}
							trainInfo.IsInputCarriageCorrect = false
							return []TrainInfo{trainInfo}, nil
						}
					}
				case carriage_num_int >= 121 && carriage_num_int <= 156:
					{
						//能够整除逻辑
						if carriage_num_int%4 == 0 {
							calculated_id := carriage_num_int/4 + 4
							carriage_nums, _ := FormatCarriageNumbers(calculated_id, "06", false)
							trainInfo := TrainInfo{
								IsEmpty:         false,
								TrainId:         "0" + strconv.Itoa(6000+calculated_id),
								Carriage_number: carriage_nums,
								Carriage_index:  "3",
							}
							//judge carriage type is correct
							for _, value := range carriage_nums {
								if number == value {
									trainInfo.IsInputCarriageCorrect = true
									return []TrainInfo{trainInfo}, nil
								}
							}
							trainInfo.IsInputCarriageCorrect = false
							return []TrainInfo{trainInfo}, nil
						} else {
							//不能整除逻辑
							calculated_id := carriage_num_int/4 + 5
							carriage_nums, _ := FormatCarriageNumbers(calculated_id, "06", false)
							trainInfo := TrainInfo{
								IsEmpty:         false,
								TrainId:         "0" + strconv.Itoa(6000+calculated_id),
								Carriage_number: carriage_nums,
								Carriage_index:  strconv.Itoa(carriage_num_int%4 - 1),
							}
							//judge carriage type is correct
							for _, value := range carriage_nums {
								if number == value {
									trainInfo.IsInputCarriageCorrect = true
									return []TrainInfo{trainInfo}, nil
								}
							}
							trainInfo.IsInputCarriageCorrect = false
							return []TrainInfo{trainInfo}, nil
						}
					}
				case carriage_num_int >= 157 && carriage_num_int <= 192:
					{
						//能够整除逻辑
						if carriage_num_int%4 == 0 {
							calculated_id := carriage_num_int/4 + 5
							carriage_nums, _ := FormatCarriageNumbers(calculated_id, "06", false)
							trainInfo := TrainInfo{
								IsEmpty:         false,
								TrainId:         "0" + strconv.Itoa(6000+calculated_id),
								Carriage_number: carriage_nums,
								Carriage_index:  "3",
							}
							//judge carriage type is correct
							for _, value := range carriage_nums {
								if number == value {
									trainInfo.IsInputCarriageCorrect = true
									return []TrainInfo{trainInfo}, nil
								}
							}
							trainInfo.IsInputCarriageCorrect = false
							return []TrainInfo{trainInfo}, nil
						} else {
							//不能整除逻辑
							calculated_id := carriage_num_int/4 + 6
							carriage_nums, _ := FormatCarriageNumbers(calculated_id, "06", false)
							trainInfo := TrainInfo{
								IsEmpty:         false,
								TrainId:         "0" + strconv.Itoa(6000+calculated_id),
								Carriage_number: carriage_nums,
								Carriage_index:  strconv.Itoa(carriage_num_int%4 - 1),
							}
							//judge carriage type is correct
							for _, value := range carriage_nums {
								if number == value {
									trainInfo.IsInputCarriageCorrect = true
									return []TrainInfo{trainInfo}, nil
								}
							}
							trainInfo.IsInputCarriageCorrect = false
							return []TrainInfo{trainInfo}, nil
						}
					}
				case carriage_num_int >= 193:
					{
						//能够整除逻辑
						if carriage_num_int%4 == 0 {
							calculated_id := carriage_num_int/4 + 6
							carriage_nums, _ := FormatCarriageNumbers(calculated_id, "06", false)
							trainInfo := TrainInfo{
								IsEmpty:         false,
								TrainId:         "0" + strconv.Itoa(6000+calculated_id),
								Carriage_number: carriage_nums,
								Carriage_index:  "3",
							}
							//judge carriage type is correct
							for _, value := range carriage_nums {
								if number == value {
									trainInfo.IsInputCarriageCorrect = true
									return []TrainInfo{trainInfo}, nil
								}
							}
							trainInfo.IsInputCarriageCorrect = false
							return []TrainInfo{trainInfo}, nil
						} else {
							//不能整除逻辑
							calculated_id := carriage_num_int/4 + 7
							carriage_nums, _ := FormatCarriageNumbers(calculated_id, "06", false)
							trainInfo := TrainInfo{
								IsEmpty:         false,
								TrainId:         "0" + strconv.Itoa(6000+calculated_id),
								Carriage_number: carriage_nums,
								Carriage_index:  strconv.Itoa(carriage_num_int%4 - 1),
							}
							//judge carriage type is correct
							for _, value := range carriage_nums {
								if number == value {
									trainInfo.IsInputCarriageCorrect = true
									return []TrainInfo{trainInfo}, nil
								}
							}
							trainInfo.IsInputCarriageCorrect = false
							return []TrainInfo{trainInfo}, nil
						}
					}
				}
				return []TrainInfo{{IsEmpty: true}}, &Error{Code: "0012", Msg: "Unexpected field reached"}
			}
		case "07":
			{
				//line7 logic
				carriage_num_int, err := strconv.Atoi(carriage_num)
				if err != nil {
					return []TrainInfo{{IsEmpty: true}}, &Error{Code: "0003", Msg: "Invalid carriage number."}
				}
				if carriage_num_int%6 == 0 {
					calculated_id := carriage_num_int / 6
					carriage_nums, _ := FormatCarriageNumbers(calculated_id, "07", false)
					trainInfo := TrainInfo{
						//若三位数车号7000改700
						IsEmpty:         false,
						TrainId:         "0" + strconv.Itoa(7000+calculated_id),
						Carriage_number: carriage_nums,
						Carriage_index:  "5",
					}
					//judge carriage type is correct
					for _, value := range carriage_nums {
						if number == value {
							trainInfo.IsInputCarriageCorrect = true
							return []TrainInfo{trainInfo}, nil
						}
					}
					trainInfo.IsInputCarriageCorrect = false
					return []TrainInfo{trainInfo}, nil
				} else {
					calculated_id := carriage_num_int/6 + 1
					carriage_nums, _ := FormatCarriageNumbers(calculated_id, "07", false)
					trainInfo := TrainInfo{
						//若三位数车号9000改900
						IsEmpty:         false,
						TrainId:         "0" + strconv.Itoa(7000+calculated_id),
						Carriage_number: carriage_nums,
						Carriage_index:  strconv.Itoa(carriage_num_int%6 - 1),
					}
					//judge carriage type is correct
					for _, value := range carriage_nums {
						if number == value {
							trainInfo.IsInputCarriageCorrect = true
							return []TrainInfo{trainInfo}, nil
						}
					}
					trainInfo.IsInputCarriageCorrect = false
					return []TrainInfo{trainInfo}, nil
				}
			}
		case "08":
			{
				//line8 logic
				carriage_num_int, err := strconv.Atoi(carriage_num)
				if err != nil {
					return []TrainInfo{{IsEmpty: true}}, &Error{Code: "0003", Msg: "Invalid carriage number."}
				}
				if carriage_num_int <= 168 {
					//6节车逻辑
					if carriage_num_int%6 == 0 {
						calculated_id := carriage_num_int / 6
						carriage_nums, _ := FormatCarriageNumbers(calculated_id, "08", false)
						trainInfo := TrainInfo{
							//若三位数车号7000改700
							IsEmpty:         false,
							TrainId:         "0" + strconv.Itoa(8000+calculated_id),
							Carriage_number: carriage_nums,
							Carriage_index:  "5",
						}
						//judge carriage type is correct
						for _, value := range carriage_nums {
							if number == value {
								trainInfo.IsInputCarriageCorrect = true
								return []TrainInfo{trainInfo}, nil
							}
						}
						trainInfo.IsInputCarriageCorrect = false
						return []TrainInfo{trainInfo}, nil
					} else {
						calculated_id := carriage_num_int/6 + 1
						carriage_nums, _ := FormatCarriageNumbers(calculated_id, "08", false)
						trainInfo := TrainInfo{
							//若三位数车号9000改900
							IsEmpty:         false,
							TrainId:         "0" + strconv.Itoa(8000+calculated_id),
							Carriage_number: carriage_nums,
							Carriage_index:  strconv.Itoa(carriage_num_int%6 - 1),
						}
						//judge carriage type is correct
						for _, value := range carriage_nums {
							if number == value {
								trainInfo.IsInputCarriageCorrect = true
								return []TrainInfo{trainInfo}, nil
							}
						}
						trainInfo.IsInputCarriageCorrect = false
						return []TrainInfo{trainInfo}, nil
					}
				} else {
					//7节车逻辑
					if (carriage_num_int-168)%7 == 0 {
						calculated_id := (carriage_num_int-168)/7 + 28
						carriage_nums, _ := FormatCarriageNumbers(calculated_id, "08", false)
						trainInfo := TrainInfo{
							//若三位数车号7000改700
							IsEmpty:         false,
							TrainId:         "0" + strconv.Itoa(8000+calculated_id),
							Carriage_number: carriage_nums,
							Carriage_index:  "6",
						}
						//judge carriage type is correct
						for _, value := range carriage_nums {
							if number == value {
								trainInfo.IsInputCarriageCorrect = true
								return []TrainInfo{trainInfo}, nil
							}
						}
						trainInfo.IsInputCarriageCorrect = false
						return []TrainInfo{trainInfo}, nil
					} else {
						calculated_id := (carriage_num_int-168)/7 + 29
						carriage_nums, _ := FormatCarriageNumbers(calculated_id, "08", false)
						trainInfo := TrainInfo{
							//若三位数车号9000改900
							IsEmpty:         false,
							TrainId:         "0" + strconv.Itoa(8000+calculated_id),
							Carriage_number: carriage_nums,
							Carriage_index:  strconv.Itoa((carriage_num_int-168)%7 - 1),
						}
						//judge carriage type is correct
						for _, value := range carriage_nums {
							if number == value {
								trainInfo.IsInputCarriageCorrect = true
								return []TrainInfo{trainInfo}, nil
							}
						}
						trainInfo.IsInputCarriageCorrect = false
						return []TrainInfo{trainInfo}, nil
					}
				}

			}
		case "09":
			{
				//line9 logic
				carriage_num_int, err := strconv.Atoi(carriage_num)
				if err != nil {
					return []TrainInfo{{IsEmpty: true}}, &Error{Code: "0003", Msg: "Invalid carriage number."}
				}
				if carriage_num_int%6 == 0 {
					calculated_id := carriage_num_int / 6
					carriage_nums, _ := FormatCarriageNumbers(calculated_id, "09", false)
					trainInfo := TrainInfo{
						//若三位数车号9000改900
						IsEmpty:         false,
						TrainId:         "0" + strconv.Itoa(9000+calculated_id),
						Carriage_number: carriage_nums,
						Carriage_index:  "5",
					}
					//judge carriage type is correct
					for _, value := range carriage_nums {
						if number == value {
							trainInfo.IsInputCarriageCorrect = true
							return []TrainInfo{trainInfo}, nil
						}
					}
					trainInfo.IsInputCarriageCorrect = false
					return []TrainInfo{trainInfo}, nil
				} else {
					calculated_id := carriage_num_int/6 + 1
					carriage_nums, _ := FormatCarriageNumbers(calculated_id, "09", false)
					trainInfo := TrainInfo{
						//若三位数车号9000改900
						IsEmpty:         false,
						TrainId:         "0" + strconv.Itoa(9000+calculated_id),
						Carriage_number: carriage_nums,
						Carriage_index:  strconv.Itoa(carriage_num_int%6 - 1),
					}
					//judge carriage type is correct
					for _, value := range carriage_nums {
						if number == value {
							trainInfo.IsInputCarriageCorrect = true
							return []TrainInfo{trainInfo}, nil
						}
					}
					trainInfo.IsInputCarriageCorrect = false
					return []TrainInfo{trainInfo}, nil
				}
			}
		case "10":
			{
				//line10 logic
				carriage_num_int, err := strconv.Atoi(carriage_num)
				if err != nil {
					return []TrainInfo{{IsEmpty: true}}, &Error{Code: "0003", Msg: "Invalid carriage number."}
				}
				if carriage_num_int%6 == 0 {
					calculated_id := carriage_num_int / 6
					carriage_nums, _ := FormatCarriageNumbers(calculated_id, "10", false)
					trainInfo := TrainInfo{
						//若三位数车号12000改1200
						IsEmpty:         false,
						TrainId:         strconv.Itoa(10000 + calculated_id),
						Carriage_number: carriage_nums,
						Carriage_index:  "5",
					}
					//judge carriage type is correct
					for _, value := range carriage_nums {
						if number == value {
							trainInfo.IsInputCarriageCorrect = true
							return []TrainInfo{trainInfo}, nil
						}
					}
					trainInfo.IsInputCarriageCorrect = false
					return []TrainInfo{trainInfo}, nil
				} else {
					calculated_id := carriage_num_int/6 + 1
					carriage_nums, _ := FormatCarriageNumbers(calculated_id, "10", false)
					trainInfo := TrainInfo{
						IsEmpty:         false,
						TrainId:         strconv.Itoa(10000 + calculated_id),
						Carriage_number: carriage_nums,
						Carriage_index:  strconv.Itoa(carriage_num_int%6 - 1),
					}
					//judge carriage type is correct
					for _, value := range carriage_nums {
						if number == value {
							trainInfo.IsInputCarriageCorrect = true
							return []TrainInfo{trainInfo}, nil
						}
					}
					trainInfo.IsInputCarriageCorrect = false
					return []TrainInfo{trainInfo}, nil
				}
			}
		case "11":
			{
				//line11 logic
				carriage_num_int, err := strconv.Atoi(carriage_num)
				if err != nil {
					return []TrainInfo{{IsEmpty: true}}, &Error{Code: "0003", Msg: "Invalid carriage number."}
				}
				if carriage_num_int%6 == 0 {
					calculated_id := carriage_num_int / 6
					carriage_nums, _ := FormatCarriageNumbers(calculated_id, "11", false)
					trainInfo := TrainInfo{
						//若三位数车号12000改1200
						IsEmpty:         false,
						TrainId:         strconv.Itoa(11000 + calculated_id),
						Carriage_number: carriage_nums,
						Carriage_index:  "5",
					}
					//judge carriage type is correct
					for _, value := range carriage_nums {
						if number == value {
							trainInfo.IsInputCarriageCorrect = true
							return []TrainInfo{trainInfo}, nil
						}
					}
					trainInfo.IsInputCarriageCorrect = false
					return []TrainInfo{trainInfo}, nil
				} else {
					calculated_id := carriage_num_int/6 + 1
					carriage_nums, _ := FormatCarriageNumbers(calculated_id, "11", false)
					trainInfo := TrainInfo{
						IsEmpty:         false,
						TrainId:         strconv.Itoa(11000 + calculated_id),
						Carriage_number: carriage_nums,
						Carriage_index:  strconv.Itoa(carriage_num_int%6 - 1),
					}
					//judge carriage type is correct
					for _, value := range carriage_nums {
						if number == value {
							trainInfo.IsInputCarriageCorrect = true
							return []TrainInfo{trainInfo}, nil
						}
					}
					trainInfo.IsInputCarriageCorrect = false
					return []TrainInfo{trainInfo}, nil
				}
			}
		case "12":
			{
				//line12 logic
				carriage_num_int, err := strconv.Atoi(carriage_num)
				if err != nil {
					return []TrainInfo{{IsEmpty: true}}, &Error{Code: "0003", Msg: "Invalid carriage number."}
				}
				if carriage_num_int%6 == 0 {
					calculated_id := carriage_num_int / 6
					carriage_nums, _ := FormatCarriageNumbers(calculated_id, "12", false)
					trainInfo := TrainInfo{
						//若三位数车号12000改1200
						IsEmpty:         false,
						TrainId:         strconv.Itoa(12000 + calculated_id),
						Carriage_number: carriage_nums,
						Carriage_index:  "5",
					}
					//judge carriage type is correct
					for _, value := range carriage_nums {
						if number == value {
							trainInfo.IsInputCarriageCorrect = true
							return []TrainInfo{trainInfo}, nil
						}
					}
					trainInfo.IsInputCarriageCorrect = false
					return []TrainInfo{trainInfo}, nil
				} else {
					calculated_id := carriage_num_int/6 + 1
					carriage_nums, _ := FormatCarriageNumbers(calculated_id, "12", false)
					trainInfo := TrainInfo{
						IsEmpty:         false,
						TrainId:         strconv.Itoa(12000 + calculated_id),
						Carriage_number: carriage_nums,
						Carriage_index:  strconv.Itoa(carriage_num_int%6 - 1),
					}
					//judge carriage type is correct
					for _, value := range carriage_nums {
						if number == value {
							trainInfo.IsInputCarriageCorrect = true
							return []TrainInfo{trainInfo}, nil
						}
					}
					trainInfo.IsInputCarriageCorrect = false
					return []TrainInfo{trainInfo}, nil
				}
			}
		case "13":
			{
				//line13 logic
				carriage_num_int, err := strconv.Atoi(carriage_num)
				if err != nil {
					return []TrainInfo{{IsEmpty: true}}, &Error{Code: "0003", Msg: "Invalid carriage number."}
				}
				if carriage_num_int%6 == 0 {
					calculated_id := carriage_num_int / 6
					carriage_nums, _ := FormatCarriageNumbers(calculated_id, "13", false)
					trainInfo := TrainInfo{
						//若三位数车号12000改1200
						IsEmpty:         false,
						TrainId:         strconv.Itoa(13000 + calculated_id),
						Carriage_number: carriage_nums,
						Carriage_index:  "5",
					}
					//judge carriage type is correct
					for _, value := range carriage_nums {
						if number == value {
							trainInfo.IsInputCarriageCorrect = true
							return []TrainInfo{trainInfo}, nil
						}
					}
					trainInfo.IsInputCarriageCorrect = false
					return []TrainInfo{trainInfo}, nil
				} else {
					calculated_id := carriage_num_int/6 + 1
					carriage_nums, _ := FormatCarriageNumbers(calculated_id, "13", false)
					trainInfo := TrainInfo{
						IsEmpty:         false,
						TrainId:         strconv.Itoa(13000 + calculated_id),
						Carriage_number: carriage_nums,
						Carriage_index:  strconv.Itoa(carriage_num_int%6 - 1),
					}
					//judge carriage type is correct
					for _, value := range carriage_nums {
						if number == value {
							trainInfo.IsInputCarriageCorrect = true
							return []TrainInfo{trainInfo}, nil
						}
					}
					trainInfo.IsInputCarriageCorrect = false
					return []TrainInfo{trainInfo}, nil
				}
			}
		case "14":
			{
				//line14 logic
				carriage_num_int, err := strconv.Atoi(carriage_num)
				if err != nil {
					return []TrainInfo{{IsEmpty: true}}, &Error{Code: "0003", Msg: "Invalid carriage number."}
				}
				if carriage_num_int%6 == 0 {
					calculated_id := carriage_num_int / 6
					carriage_nums, _ := FormatCarriageNumbers(calculated_id, "14", false)
					trainInfo := TrainInfo{
						//若三位数车号12000改1200
						IsEmpty:         false,
						TrainId:         strconv.Itoa(14000 + calculated_id),
						Carriage_number: carriage_nums,
						Carriage_index:  "5",
					}
					//judge carriage type is correct
					for _, value := range carriage_nums {
						if number == value {
							trainInfo.IsInputCarriageCorrect = true
							return []TrainInfo{trainInfo}, nil
						}
					}
					trainInfo.IsInputCarriageCorrect = false
					return []TrainInfo{trainInfo}, nil
				} else {
					calculated_id := carriage_num_int/6 + 1
					carriage_nums, _ := FormatCarriageNumbers(calculated_id, "14", false)
					trainInfo := TrainInfo{
						IsEmpty:         false,
						TrainId:         strconv.Itoa(14000 + calculated_id),
						Carriage_number: carriage_nums,
						Carriage_index:  strconv.Itoa(carriage_num_int%6 - 1),
					}
					//judge carriage type is correct
					for _, value := range carriage_nums {
						if number == value {
							trainInfo.IsInputCarriageCorrect = true
							return []TrainInfo{trainInfo}, nil
						}
					}
					trainInfo.IsInputCarriageCorrect = false
					return []TrainInfo{trainInfo}, nil
				}
			}
		case "15":
			{
				//line15 logic
				carriage_num_int, err := strconv.Atoi(carriage_num)
				if err != nil {
					return []TrainInfo{{IsEmpty: true}}, &Error{Code: "0003", Msg: "Invalid carriage number."}
				}
				if carriage_num_int%6 == 0 {
					calculated_id := carriage_num_int / 6
					carriage_nums, _ := FormatCarriageNumbers(calculated_id, "15", false)
					trainInfo := TrainInfo{
						//若三位数车号12000改1200
						IsEmpty:         false,
						TrainId:         strconv.Itoa(15000 + calculated_id),
						Carriage_number: carriage_nums,
						Carriage_index:  "5",
					}
					//judge carriage type is correct
					for _, value := range carriage_nums {
						if number == value {
							trainInfo.IsInputCarriageCorrect = true
							return []TrainInfo{trainInfo}, nil
						}
					}
					trainInfo.IsInputCarriageCorrect = false
					return []TrainInfo{trainInfo}, nil
				} else {
					calculated_id := carriage_num_int/6 + 1
					carriage_nums, _ := FormatCarriageNumbers(calculated_id, "15", false)
					trainInfo := TrainInfo{
						IsEmpty:         false,
						TrainId:         strconv.Itoa(15000 + calculated_id),
						Carriage_number: carriage_nums,
						Carriage_index:  strconv.Itoa(carriage_num_int%6 - 1),
					}
					//judge carriage type is correct
					for _, value := range carriage_nums {
						if number == value {
							trainInfo.IsInputCarriageCorrect = true
							return []TrainInfo{trainInfo}, nil
						}
					}
					trainInfo.IsInputCarriageCorrect = false
					return []TrainInfo{trainInfo}, nil
				}
			}
		case "16":
			{
				//line16 logic
				carriage_num_int, err := strconv.Atoi(carriage_num)
				if err != nil {
					return []TrainInfo{{IsEmpty: true}}, &Error{Code: "0003", Msg: "Invalid carriage number."}
				}
				if carriage_num_int <= 138 {
					//3节车逻辑
					if carriage_num_int%3 == 0 {
						calculated_id := carriage_num_int / 3
						carriage_nums, _ := FormatCarriageNumbers(calculated_id, "16", false)
						trainInfo := TrainInfo{
							//若三位数车号7000改700
							IsEmpty:         false,
							TrainId:         strconv.Itoa(16000 + calculated_id),
							Carriage_number: carriage_nums,
							Carriage_index:  "2",
						}
						//judge carriage type is correct
						for _, value := range carriage_nums {
							if number == value {
								trainInfo.IsInputCarriageCorrect = true
								return []TrainInfo{trainInfo}, nil
							}
						}
						trainInfo.IsInputCarriageCorrect = false
						return []TrainInfo{trainInfo}, nil
					} else {
						calculated_id := carriage_num_int/3 + 1
						carriage_nums, _ := FormatCarriageNumbers(calculated_id, "16", false)
						trainInfo := TrainInfo{
							//若三位数车号9000改900
							IsEmpty:         false,
							TrainId:         strconv.Itoa(16000 + calculated_id),
							Carriage_number: carriage_nums,
							Carriage_index:  strconv.Itoa(carriage_num_int%3 - 1),
						}
						//judge carriage type is correct
						for _, value := range carriage_nums {
							if number == value {
								trainInfo.IsInputCarriageCorrect = true
								return []TrainInfo{trainInfo}, nil
							}
						}
						trainInfo.IsInputCarriageCorrect = false
						return []TrainInfo{trainInfo}, nil
					}
				} else {
					//6节车逻辑
					if (carriage_num_int-138)%6 == 0 {
						calculated_id := (carriage_num_int-138)/6 + 46
						carriage_nums, _ := FormatCarriageNumbers(calculated_id, "16", false)
						trainInfo := TrainInfo{
							//若三位数车号7000改700
							IsEmpty:         false,
							TrainId:         strconv.Itoa(16000 + calculated_id),
							Carriage_number: carriage_nums,
							Carriage_index:  "5",
						}
						//judge carriage type is correct
						for _, value := range carriage_nums {
							if number == value {
								trainInfo.IsInputCarriageCorrect = true
								return []TrainInfo{trainInfo}, nil
							}
						}
						trainInfo.IsInputCarriageCorrect = false
						return []TrainInfo{trainInfo}, nil
					} else {
						calculated_id := (carriage_num_int-138)/6 + 46
						carriage_nums, _ := FormatCarriageNumbers(calculated_id, "16", false)
						trainInfo := TrainInfo{
							//若三位数车号9000改900
							IsEmpty:         false,
							TrainId:         strconv.Itoa(16000 + calculated_id),
							Carriage_number: carriage_nums,
							Carriage_index:  strconv.Itoa((carriage_num_int-138)%6 - 1),
						}
						//judge carriage type is correct
						for _, value := range carriage_nums {
							if number == value {
								trainInfo.IsInputCarriageCorrect = true
								return []TrainInfo{trainInfo}, nil
							}
						}
						trainInfo.IsInputCarriageCorrect = false
						return []TrainInfo{trainInfo}, nil
					}
				}

			}
		case "17":
			{
				//line17 logic
				carriage_num_int, err := strconv.Atoi(carriage_num)
				if err != nil {
					return []TrainInfo{{IsEmpty: true}}, &Error{Code: "0003", Msg: "Invalid carriage number."}
				}
				if carriage_num_int%6 == 0 {
					calculated_id := carriage_num_int / 6
					carriage_nums, _ := FormatCarriageNumbers(calculated_id, "17", false)
					trainInfo := TrainInfo{
						//若三位数车号12000改1200
						IsEmpty:         false,
						TrainId:         strconv.Itoa(17000 + calculated_id),
						Carriage_number: carriage_nums,
						Carriage_index:  "5",
					}
					//judge carriage type is correct
					for _, value := range carriage_nums {
						if number == value {
							trainInfo.IsInputCarriageCorrect = true
							return []TrainInfo{trainInfo}, nil
						}
					}
					trainInfo.IsInputCarriageCorrect = false
					return []TrainInfo{trainInfo}, nil
				} else {
					calculated_id := carriage_num_int/6 + 1
					carriage_nums, _ := FormatCarriageNumbers(calculated_id, "17", false)
					trainInfo := TrainInfo{
						IsEmpty:         false,
						TrainId:         strconv.Itoa(17000 + calculated_id),
						Carriage_number: carriage_nums,
						Carriage_index:  strconv.Itoa(carriage_num_int%6 - 1),
					}
					//judge carriage type is correct
					for _, value := range carriage_nums {
						if number == value {
							trainInfo.IsInputCarriageCorrect = true
							return []TrainInfo{trainInfo}, nil
						}
					}
					trainInfo.IsInputCarriageCorrect = false
					return []TrainInfo{trainInfo}, nil
				}
			}
		case "18":
			{
				//line18 logic
				carriage_num_int, err := strconv.Atoi(carriage_num)
				if err != nil {
					return []TrainInfo{{IsEmpty: true}}, &Error{Code: "0003", Msg: "Invalid carriage number."}
				}
				if carriage_num_int%6 == 0 {
					calculated_id := carriage_num_int / 6
					carriage_nums, _ := FormatCarriageNumbers(calculated_id, "18", false)
					trainInfo := TrainInfo{
						//若三位数车号12000改1200
						IsEmpty:         false,
						TrainId:         strconv.Itoa(18000 + calculated_id),
						Carriage_number: carriage_nums,
						Carriage_index:  "5",
					}
					//judge carriage type is correct
					for _, value := range carriage_nums {
						if number == value {
							trainInfo.IsInputCarriageCorrect = true
							return []TrainInfo{trainInfo}, nil
						}
					}
					trainInfo.IsInputCarriageCorrect = false
					return []TrainInfo{trainInfo}, nil
				} else {
					calculated_id := carriage_num_int/6 + 1
					carriage_nums, _ := FormatCarriageNumbers(calculated_id, "18", false)
					trainInfo := TrainInfo{
						IsEmpty:         false,
						TrainId:         strconv.Itoa(18000 + calculated_id),
						Carriage_number: carriage_nums,
						Carriage_index:  strconv.Itoa(carriage_num_int%6 - 1),
					}
					//judge carriage type is correct
					for _, value := range carriage_nums {
						if number == value {
							trainInfo.IsInputCarriageCorrect = true
							return []TrainInfo{trainInfo}, nil
						}
					}
					trainInfo.IsInputCarriageCorrect = false
					return []TrainInfo{trainInfo}, nil
				}
			}
		default:
			{
				return []TrainInfo{{IsEmpty: true}}, &Error{Code: "0001", Msg: "Unknown type carriage number."}
			}
		}
	} else if len(number) == 5 {
		line_num, _ /*carriage_num*/, _ := number[:2], number[2:4], number[4:]
		switch line_num {
		case "92":
			{
				//line 1 logic
				t := TrainInfo{}
				Err := QueryCarriage("1", number, &t)
				if Err != nil {
					e := &Error{Code: "0024", Msg: "Line1 Search 01A01 Carriages goes wrong", VerboseError: Err}
					t.IsEmpty = true
					return []TrainInfo{t}, e
				}
				t.IsEmpty = false
				return []TrainInfo{t}, nil
			}
		case "93":
			{
				//line 1 logic
				t := TrainInfo{}
				Err := QueryCarriage("1", number, &t)
				if Err != nil {
					e := &Error{Code: "0025", Msg: "Line1 Search 01A01/01A02 Carriages goes wrong", VerboseError: Err}
					t.IsEmpty = true
					return []TrainInfo{t}, e
				}
				t.IsEmpty = false
				return []TrainInfo{t}, nil
			}
		case "94":
			{
				//line 1 logic
				t := TrainInfo{}
				Err := QueryCarriage("1", number, &t)
				if Err != nil {
					e := &Error{Code: "0026", Msg: "Line1 Search 01A01/01A02 Carriages goes wrong", VerboseError: Err}
					t.IsEmpty = true
					return []TrainInfo{t}, e
				}
				t.IsEmpty = false
				return []TrainInfo{t}, nil
			}
		case "98":
			{
				//line 1 logic
				t := TrainInfo{}
				Err := QueryCarriage("1", number, &t)
				if Err != nil {
					e := &Error{Code: "0027", Msg: "Line1 Search 01A03 Carriages goes wrong", VerboseError: Err}
					t.IsEmpty = true
					return []TrainInfo{t}, e
				}
				t.IsEmpty = false
				return []TrainInfo{t}, nil
			}
		case "99":
			{
				//line 1 logic
				t := TrainInfo{}
				Err := QueryCarriage("1", number, &t)
				if Err != nil {
					e := &Error{Code: "0028", Msg: "Line1 Search Carriages goes wrong", VerboseError: Err}
					t.IsEmpty = true
					return []TrainInfo{t}, e
				}
				t.IsEmpty = false
				return []TrainInfo{t}, nil
			}
		case "14":
			{
				t := TrainInfo{}
				Err := QueryCarriage("1", number, &t)
				if Err != nil {
					e := &Error{Code: "0029", Msg: "Line1 Search Carriages goes wrong", VerboseError: Err}
					t.IsEmpty = true
					return []TrainInfo{t}, e
				}
				t.IsEmpty = false
				return []TrainInfo{t}, nil
			}
		case "00":
			{
				//line 1/2 logic
				t1, t2 := TrainInfo{}, TrainInfo{}
				Err1 := QueryCarriage("1", number, &t1)
				Err2 := QueryCarriage("2", number, &t2)
				if Err1 != nil || Err2 != nil {
					Err1.VerboseError = Err2
					return []TrainInfo{{IsEmpty: true}}, Err1
				}
				if t1.TrainId == "" {
					t1.IsEmpty = true
					if t2.TrainId == "" {
						t2.IsEmpty = true
						Err := Error{Code: "0030", Msg: "Corresponding train in line1/2 not found"}
						return []TrainInfo{{IsEmpty: true}}, &Err
					} else {
						t2.IsEmpty = false
						return []TrainInfo{t2}, nil
					}
				} else {
					t1.IsEmpty = false
					return []TrainInfo{t1}, nil
				}
			}
		case "01":
			{
				//line 1/2 logic
				t1, t2 := TrainInfo{}, TrainInfo{}
				Err1 := QueryCarriage("1", number, &t1)
				Err2 := QueryCarriage("2", number, &t2)
				if Err1 != nil || Err2 != nil {
					Err1.VerboseError = Err2
					return []TrainInfo{{IsEmpty: true}}, Err1
				}
				if t1.TrainId == "" {
					t1.IsEmpty = true
					if t2.TrainId == "" {
						t2.IsEmpty = true
						Err := Error{Code: "0031", Msg: "Corresponding train in line1/2 not found"}
						return []TrainInfo{{IsEmpty: true}}, &Err
					} else {
						t2.IsEmpty = false
						return []TrainInfo{t2}, nil
					}
				} else {
					t1.IsEmpty = false
					return []TrainInfo{t1}, nil
				}
			}
		case "02":
			{
				//line 3/5 logic
				t1, t2 := TrainInfo{}, TrainInfo{}
				Err1 := QueryCarriage("3", number, &t1)
				Err2 := QueryCarriage("5", number, &t2)
				if Err1 != nil || Err2 != nil {
					if Err1 == nil {
						return []TrainInfo{{IsEmpty: true}}, Err2
					} else {
						return []TrainInfo{{IsEmpty: true}}, Err1
					}
				}
				if t1.TrainId == "" {
					if t2.TrainId == "" {
						t1.IsEmpty = true
						t2.IsEmpty = true
						Err := Error{Code: "0032", Msg: "Corresponding train in line3/5 not found"}
						return []TrainInfo{{IsEmpty: true}}, &Err
					} else {
						t1.IsEmpty = true
						t2.IsEmpty = false
						return []TrainInfo{t2}, nil
					}
				} else {
					if t2.TrainId == "" {
						t1.IsEmpty = false
						t2.IsEmpty = true
						return []TrainInfo{t1}, nil
					} else {
						t1.IsEmpty = false
						t2.IsEmpty = false
						return []TrainInfo{t1, t2}, nil
					}
				}

			}
		case "03":
			{
				//line 3/5 logic
				t1, t2 := TrainInfo{}, TrainInfo{}
				Err1 := QueryCarriage("3", number, &t1)
				Err2 := QueryCarriage("5", number, &t2)
				if Err1 != nil || Err2 != nil {
					if Err1 == nil {
						return []TrainInfo{{IsEmpty: true}}, Err2
					} else {
						return []TrainInfo{{IsEmpty: true}}, Err1
					}
				}
				if t1.TrainId == "" {
					if t2.TrainId == "" {
						t1.IsEmpty = true
						t2.IsEmpty = true
						Err := Error{Code: "0033", Msg: "Corresponding train in line3/5 not found"}
						return []TrainInfo{{IsEmpty: true}}, &Err
					} else {
						t1.IsEmpty = true
						t2.IsEmpty = false
						return []TrainInfo{t2}, nil
					}
				} else {
					if t2.TrainId == "" {
						t1.IsEmpty = false
						t2.IsEmpty = true
						return []TrainInfo{t1}, nil
					} else {
						t1.IsEmpty = false
						t2.IsEmpty = false
						return []TrainInfo{t1, t2}, nil
					}
				}

			}
		case "04":
			{
				//line 3/5 logic
				t1, t2 := TrainInfo{}, TrainInfo{}
				Err1 := QueryCarriage("3", number, &t1)
				Err2 := QueryCarriage("5", number, &t2)
				if Err1 != nil || Err2 != nil {
					if Err1 == nil {
						return []TrainInfo{{IsEmpty: true}}, Err2
					} else {
						return []TrainInfo{{IsEmpty: true}}, Err1
					}
				}
				if t1.TrainId == "" {
					if t2.TrainId == "" {
						t1.IsEmpty = true
						t2.IsEmpty = true
						Err := Error{Code: "0034", Msg: "Corresponding train in line3/5 not found"}
						return []TrainInfo{{IsEmpty: true}}, &Err
					} else {
						t1.IsEmpty = true
						t2.IsEmpty = false
						return []TrainInfo{t2}, nil
					}
				} else {
					if t2.TrainId == "" {
						t1.IsEmpty = false
						t2.IsEmpty = true
						return []TrainInfo{t1}, nil
					} else {
						t1.IsEmpty = false
						t2.IsEmpty = false
						return []TrainInfo{t1, t2}, nil
					}
				}

			}
		default:
			{
				return []TrainInfo{{IsEmpty: true}}, &Error{Code: "0002", Msg: "Unknown type carriage number"}
			}
		}
	} else if len(number) == 7 {
		numberLower := strings.ToLower(number)
		lineNum, carriageNum := numberLower[:3], numberLower[3:6]
		switch lineNum {
		case "jc4":
			{
				//jc logic
				carriage_num_int, err := strconv.Atoi(carriageNum)
				if err != nil {
					return []TrainInfo{{IsEmpty: true}}, &Error{Code: "0003", Msg: "Invalid carriage number."}
				}
				if carriage_num_int%4 == 0 {
					calculated_id := carriage_num_int / 4
					carriage_nums, _ := FormatCarriageNumbers(calculated_id, "jc4", false)
					trainInfo := TrainInfo{
						//若三位数车号12000改1200
						IsEmpty:         false,
						TrainId:         "JC4" + fmt.Sprintf("%03d", calculated_id),
						Carriage_number: carriage_nums,
						Carriage_index:  "3",
					}
					//judge carriage type is correct
					for _, value := range carriage_nums {
						if number == value {
							trainInfo.IsInputCarriageCorrect = true
							return []TrainInfo{trainInfo}, nil
						}
					}
					trainInfo.IsInputCarriageCorrect = false
					return []TrainInfo{trainInfo}, nil
				} else {
					calculated_id := carriage_num_int/4 + 1
					carriage_nums, _ := FormatCarriageNumbers(calculated_id, "jc4", false)
					trainInfo := TrainInfo{
						IsEmpty:         false,
						TrainId:         "JC4" + fmt.Sprintf("%03d", calculated_id),
						Carriage_number: carriage_nums,
						Carriage_index:  strconv.Itoa(carriage_num_int%4 - 1),
					}
					//judge carriage type is correct
					for _, value := range carriage_nums {
						if number == value {
							trainInfo.IsInputCarriageCorrect = true
							return []TrainInfo{trainInfo}, nil
						}
					}
					trainInfo.IsInputCarriageCorrect = false
					return []TrainInfo{trainInfo}, nil
				}
			}
		case "jc8":
			{
				//jc logic
				carriage_num_int, err := strconv.Atoi(carriageNum)
				if err != nil {
					return []TrainInfo{{IsEmpty: true}}, &Error{Code: "0003", Msg: "Invalid carriage number."}
				}
				if carriage_num_int%8 == 0 {
					calculated_id := carriage_num_int / 8
					carriage_nums, _ := FormatCarriageNumbers(calculated_id, "jc8", false)
					trainInfo := TrainInfo{
						//若三位数车号12000改1200
						IsEmpty:         false,
						TrainId:         "JC8" + fmt.Sprintf("%03d", calculated_id),
						Carriage_number: carriage_nums,
						Carriage_index:  "7",
					}
					//judge carriage type is correct
					for _, value := range carriage_nums {
						if number == value {
							trainInfo.IsInputCarriageCorrect = true
							return []TrainInfo{trainInfo}, nil
						}
					}
					trainInfo.IsInputCarriageCorrect = false
					return []TrainInfo{trainInfo}, nil
				} else {
					calculated_id := carriage_num_int/8 + 1
					carriage_nums, _ := FormatCarriageNumbers(calculated_id, "jc8", false)
					trainInfo := TrainInfo{
						IsEmpty:         false,
						TrainId:         "JC8" + fmt.Sprintf("%03d", calculated_id),
						Carriage_number: carriage_nums,
						Carriage_index:  strconv.Itoa(carriage_num_int%8 - 1),
					}
					//judge carriage type is correct
					for _, value := range carriage_nums {
						if number == value {
							trainInfo.IsInputCarriageCorrect = true
							return []TrainInfo{trainInfo}, nil
						}
					}
					trainInfo.IsInputCarriageCorrect = false
					return []TrainInfo{trainInfo}, nil
				}
			}
		case "t01":
			{
				//jc logic
				carriage_num_int, err := strconv.Atoi(carriageNum)
				if err != nil {
					return []TrainInfo{{IsEmpty: true}}, &Error{Code: "0003", Msg: "Invalid carriage number."}
				}
				if carriage_num_int%4 == 0 {
					calculated_id := carriage_num_int / 4
					carriage_nums, _ := FormatCarriageNumbers(calculated_id, "t01", false)
					trainInfo := TrainInfo{
						//若三位数车号12000改1200
						IsEmpty:         false,
						TrainId:         "T01" + fmt.Sprintf("%02d", calculated_id),
						Carriage_number: carriage_nums,
						Carriage_index:  "3",
					}
					//judge carriage type is correct
					for _, value := range carriage_nums {
						if number == value {
							trainInfo.IsInputCarriageCorrect = true
							return []TrainInfo{trainInfo}, nil
						}
					}
					trainInfo.IsInputCarriageCorrect = false
					return []TrainInfo{trainInfo}, nil
				} else {
					calculated_id := carriage_num_int/4 + 1
					carriage_nums, _ := FormatCarriageNumbers(calculated_id, "t01", false)
					trainInfo := TrainInfo{
						IsEmpty:         false,
						TrainId:         "T01" + fmt.Sprintf("%02d", calculated_id),
						Carriage_number: carriage_nums,
						Carriage_index:  strconv.Itoa(carriage_num_int%4 - 1),
					}
					//judge carriage type is correct
					for _, value := range carriage_nums {
						if number == value {
							trainInfo.IsInputCarriageCorrect = true
							return []TrainInfo{trainInfo}, nil
						}
					}
					trainInfo.IsInputCarriageCorrect = false
					return []TrainInfo{trainInfo}, nil
				}
			}
		default:
			{
				return []TrainInfo{{IsEmpty: true}}, &Error{Code: "0002", Msg: "Unknown type carriage number"}
			}

		}
	} else if len(number) == 8 {
		numberLower := strings.ToLower(number)
		lineNum, carriageNum := numberLower[:4], numberLower[4:7]
		switch lineNum {
		case "jy01":
			{
				//jc logic
				carriage_num_int, err := strconv.Atoi(carriageNum)
				if err != nil {
					return []TrainInfo{{IsEmpty: true}}, &Error{Code: "0003", Msg: "Invalid carriage number."}
				}
				if carriage_num_int%4 == 0 {
					calculated_id := carriage_num_int / 4
					carriage_nums, _ := FormatCarriageNumbers(calculated_id, "jy01", false)
					trainInfo := TrainInfo{
						//若三位数车号12000改1200
						IsEmpty:         false,
						TrainId:         "JY01" + fmt.Sprintf("%02d", calculated_id),
						Carriage_number: carriage_nums,
						Carriage_index:  "3",
					}
					//judge carriage type is correct
					for _, value := range carriage_nums {
						if number == value {
							trainInfo.IsInputCarriageCorrect = true
							return []TrainInfo{trainInfo}, nil
						}
					}
					trainInfo.IsInputCarriageCorrect = false
					return []TrainInfo{trainInfo}, nil
				} else {
					calculated_id := carriage_num_int/4 + 1
					carriage_nums, _ := FormatCarriageNumbers(calculated_id, "jy01", false)
					trainInfo := TrainInfo{
						IsEmpty:         false,
						TrainId:         "JY01" + fmt.Sprintf("%02d", calculated_id),
						Carriage_number: carriage_nums,
						Carriage_index:  strconv.Itoa(carriage_num_int%4 - 1),
					}
					//judge carriage type is correct
					for _, value := range carriage_nums {
						if number == value {
							trainInfo.IsInputCarriageCorrect = true
							return []TrainInfo{trainInfo}, nil
						}
					}
					trainInfo.IsInputCarriageCorrect = false
					return []TrainInfo{trainInfo}, nil
				}
			}
		case "sj01":
			{
				//jc logic
				carriage_num_int, err := strconv.Atoi(carriageNum)
				if err != nil {
					return []TrainInfo{{IsEmpty: true}}, &Error{Code: "0003", Msg: "Invalid carriage number."}
				}
				if carriage_num_int%5 == 0 {
					calculated_id := carriage_num_int / 5
					carriage_nums, _ := FormatCarriageNumbers(calculated_id, "sj01", false)
					trainInfo := TrainInfo{
						//若三位数车号12000改1200
						IsEmpty:         false,
						TrainId:         "SJ01" + fmt.Sprintf("%02d", calculated_id),
						Carriage_number: carriage_nums,
						Carriage_index:  "4",
					}
					//judge carriage type is correct
					for _, value := range carriage_nums {
						if number == value {
							trainInfo.IsInputCarriageCorrect = true
							return []TrainInfo{trainInfo}, nil
						}
					}
					trainInfo.IsInputCarriageCorrect = false
					return []TrainInfo{trainInfo}, nil
				} else {
					calculated_id := carriage_num_int/5 + 1
					carriage_nums, _ := FormatCarriageNumbers(calculated_id, "sj01", false)
					trainInfo := TrainInfo{
						IsEmpty:         false,
						TrainId:         "SJ01" + fmt.Sprintf("%02d", calculated_id),
						Carriage_number: carriage_nums,
						Carriage_index:  strconv.Itoa(carriage_num_int%5 - 1),
					}
					//judge carriage type is correct
					for _, value := range carriage_nums {
						if number == value {
							trainInfo.IsInputCarriageCorrect = true
							return []TrainInfo{trainInfo}, nil
						}
					}
					trainInfo.IsInputCarriageCorrect = false
					return []TrainInfo{trainInfo}, nil
				}
			}
		default:
			{
				return []TrainInfo{{IsEmpty: true}}, &Error{Code: "0002", Msg: "Unknown type carriage number"}
			}

		}
	} else {
		return []TrainInfo{{IsEmpty: true}}, &Error{Code: "0005", Msg: "Invalid carriage number"}
	}

}

// 根据车号生成车厢号
func FormatCarriageNumbers(id int, line string, isAnda bool) ([]string, *Error) {
	if id > 0 {
		switch line {
		case "01":
			{
				switch {
				case (id >= 40 && id <= 55):
					{
						num := []int{8*(id-39) + 227, 8*(id-39) + 228, 8*(id-39) + 229, 8*(id-39) + 230, 8*(id-39) + 231, 8*(id-39) + 232, 8*(id-39) + 233, 8*(id-39) + 234}
						template := []string{"1", "2", "3", "2", "3", "3", "2", "1"}
						res := []string{}
						for i := 0; i < 8; i++ {
							res = append(res, line+strconv.Itoa(num[i])+template[i])
						}
						return res, nil
					}
				case (id >= 56 && id <= 86):
					{
						num := []int{8*(id-55) + 460, 8*(id-55) + 461, 8*(id-55) + 462, 8*(id-55) + 463, 8*(id-55) + 464, 8*(id-55) + 465, 8*(id-55) + 466, 8*(id-55) + 467}
						template := []string{"1", "2", "3", "2", "3", "3", "2", "1"}
						res := []string{}
						for i := 0; i < 8; i++ {
							res = append(res, line+strconv.Itoa(num[i])+template[i])
						}
						return res, nil
					}
				default:
					{
						E := &Error{Code: "0022", Msg: "Unexpected field reached"}
						return []string{}, E
					}
				}
			}
		case "02":
			{
				switch {
				case (id >= 33 && id <= 69):
					{
						num := []int{8*(id-16) - 7, 8*(id-16) - 6, 8*(id-16) - 5, 8*(id-16) - 4, 8*(id-16) - 3, 8*(id-16) - 2, 8*(id-16) - 1, 8 * (id - 16)}
						template := []string{"1", "2", "2", "3", "2", "3", "2", "1"}
						res := []string{}
						for i := 0; i < 8; i++ {
							res = append(res, line+strconv.Itoa(num[i])+template[i])
						}
						return res, nil
					}
				case (id >= 70 && id <= 85):
					{
						num := []int{4*(id+37) - 3, 4*(id+37) - 2, 4*(id+37) - 1, 4*(id+53) - 3, 4*(id+53) - 2, 4*(id+53) - 1, 4 * (id + 53), 4 * (id + 37)}
						template := []string{"1", "2", "2", "3", "2", "3", "2", "1"}
						res := []string{}
						for i := 0; i < 8; i++ {
							res = append(res, line+strconv.Itoa(num[i])+template[i])
						}
						return res, nil
					}
				case (id >= 86 && id <= 116):
					{
						num := []int{8*(id-16) - 7, 8*(id-16) - 6, 8*(id-16) - 5, 8*(id-16) - 4, 8*(id-16) - 3, 8*(id-16) - 2, 8*(id-16) - 1, 8 * (id - 16)}
						template := []string{"1", "2", "3", "2", "3", "3", "2", "1"}
						res := []string{}
						for i := 0; i < 8; i++ {
							res = append(res, line+strconv.Itoa(num[i])+template[i])
						}
						return res, nil
					}
				default:
					{
						E := &Error{Code: "0019", Msg: "Unexpected field reached"}
						return []string{}, E
					}
				}
			}
		case "05":
			{
				switch {
				case (id >= 1 && id <= 18):
					//pass

					return []string{}, nil

				case (id >= 19):
					{
						num := []int{6*(id-18) + 63, 6*(id-18) + 64, 6*(id-18) + 65, 6*(id-18) + 66, 6*(id-18) + 67, 6*(id-18) + 68}
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
						E := &Error{Code: "0013", Msg: "Unexpected field reached"}
						return []string{}, E
					}
				}
			}
		case "06":
			{
				//line6 logic
				switch {
				case (id >= 1 && id <= 3):
					{
						num := []int{4*id - 3, 4*id - 2, 4*id - 1, 4 * id}
						template := []string{"1", "2", "2", "1"}
						res := []string{}
						for i := 0; i < 4; i++ {
							num_str := strconv.Itoa(num[i])
							for len(num_str) < 3 {
								num_str = "0" + num_str
							}
							res = append(res, line+num_str+template[i])
						}
						return res, nil
					}
				case (id >= 5 && id <= 13):
					{
						num := []int{4*id - 7, 4*id - 6, 4*id - 5, 4*id - 4}
						template := []string{"1", "2", "2", "1"}
						res := []string{}
						for i := 0; i < 4; i++ {
							num_str := strconv.Itoa(num[i])
							for len(num_str) < 3 {
								num_str = "0" + num_str
							}
							res = append(res, line+num_str+template[i])
						}
						return res, nil
					}
				case (id >= 15 && id <= 23):
					{
						num := []int{4*id - 11, 4*id - 10, 4*id - 9, 4*id - 8}
						template := []string{"1", "2", "2", "1"}
						res := []string{}
						for i := 0; i < 4; i++ {
							num_str := strconv.Itoa(num[i])
							for len(num_str) < 3 {
								num_str = "0" + num_str
							}
							res = append(res, line+num_str+template[i])
						}
						return res, nil
					}
				case (id >= 25 && id <= 33):
					{
						num := []int{4*id - 15, 4*id - 14, 4*id - 13, 4*id - 12}
						template := []string{"1", "2", "2", "1"}
						res := []string{}
						for i := 0; i < 4; i++ {
							num_str := strconv.Itoa(num[i])
							for len(num_str) < 3 {
								num_str = "0" + num_str
							}
							res = append(res, line+num_str+template[i])
						}
						return res, nil
					}
				case (id >= 35 && id <= 43):
					{
						num := []int{4*id - 19, 4*id - 18, 4*id - 17, 4*id - 16}
						template := []string{"1", "2", "2", "1"}
						res := []string{}
						for i := 0; i < 4; i++ {
							num_str := strconv.Itoa(num[i])
							for len(num_str) < 3 {
								num_str = "0" + num_str
							}
							res = append(res, line+num_str+template[i])
						}
						return res, nil
					}
				case (id >= 45 && id <= 53):
					{
						num := []int{4*id - 23, 4*id - 22, 4*id - 21, 4*id - 20}
						template := []string{"1", "2", "2", "1"}
						res := []string{}
						for i := 0; i < 4; i++ {
							num_str := strconv.Itoa(num[i])
							for len(num_str) < 3 {
								num_str = "0" + num_str
							}
							res = append(res, line+num_str+template[i])
						}
						return res, nil
					}
				case (id >= 55):
					{
						num := []int{4*id - 27, 4*id - 26, 4*id - 25, 4*id - 24}
						template := []string{"1", "2", "2", "1"}
						res := []string{}
						for i := 0; i < 4; i++ {
							num_str := strconv.Itoa(num[i])
							for len(num_str) < 3 {
								num_str = "0" + num_str
							}
							res = append(res, line+num_str+template[i])
						}
						return res, nil
					}
				}
				res := []string{}
				return res, nil
			}
		case "14":
			{
				num := []int{8*id - 7, 8*id - 6, 8*id - 5, 8*id - 4, 8*id - 3, 8*id - 2, 8*id - 1, 8 * id}
				template := []string{"1", "2", "3", "2", "3", "3", "2", "1"}
				res := []string{}
				for i := 0; i < 8; i++ {
					num_str := strconv.Itoa(num[i])
					for len(num_str) < 3 {
						num_str = "0" + num_str
					}
					res = append(res, line+num_str+template[i])
				}
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
		case "jc4":
			{
				num := []int{4*id - 3, 4*id - 2, 4*id - 1, 4 * id}
				template := []string{"1", "2", "3", "4"}
				res := []string{}
				for i := 0; i < 4; i++ {
					num_str := fmt.Sprintf("%03d", num[i])
					res = append(res, "JC4"+num_str+template[i])
				}
				return res, nil
			}
		case "jc8":
			{
				num := []int{8*id - 7, 8*id - 6, 8*id - 5, 8*id - 4, 8*id - 3, 8*id - 2, 8*id - 1, 8 * id}
				template := []string{"1", "2", "3", "4", "5", "6", "7", "8"}
				res := []string{}
				for i := 0; i < 8; i++ {
					num_str := fmt.Sprintf("%03d", num[i])
					res = append(res, "JC8"+num_str+template[i])
				}
				return res, nil

			}
		case "t01":
			{
				num := []int{4*id - 3, 4*id - 2, 4*id - 1, 4 * id}
				template := []string{"1", "2", "2", "1"}
				res := []string{}
				for i := 0; i < 4; i++ {
					num_str := fmt.Sprintf("%03d", num[i])
					res = append(res, "T01"+num_str+template[i])
				}
				return res, nil
			}
		case "jy01":
			{
				num := []int{4*id - 3, 4*id - 2, 4*id - 1, 4 * id}
				template := []string{"1", "2", "2", "1"}
				res := []string{}
				for i := 0; i < 4; i++ {
					num_str := fmt.Sprintf("%03d", num[i])
					res = append(res, "JY01"+num_str+template[i])
				}
				return res, nil
			}
		case "sj01":
			{
				num := []int{5*id - 4, 5*id - 3, 5*id - 2, 5*id - 1, 5 * id}
				template := []string{"1", "2", "3", "2", "1"}
				res := []string{}
				for i := 0; i < 5; i++ {
					num_str := fmt.Sprintf("%03d", num[i])
					res = append(res, "SJ01"+num_str+template[i])
				}
				return res, nil
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
							num_str := fmt.Sprintf("%03d", num[i])
							res = append(res, line+num_str+template[i])
						}
						return res, nil
					}
				case false:
					{
						template := []string{"1", "2", "3", "3", "2", "1"}
						res := []string{}
						for i := 0; i < 6; i++ {
							num_str := fmt.Sprintf("%03d", num[i])
							res = append(res, line+num_str+template[i])
						}
						return res, nil
					}
				default:
					{
						res := []string{}
						return res, &Error{Code: "0007", Msg: "invald isAnda flag"}
					}
				}
			}
		}
	} else {
		return nil, &Error{Code: "0007", Msg: "invalid train id"}
	}
}

// 验证车厢号的合法性，支持大小写混合
func ValidateCarriageNumbers(number string) bool {
	pattern := `(?i)^(?:\d{5}|\d{6}|T\d{6}|JC\d{5}|SJ\d{6}|JY\d{6})$`
	matched, _ := regexp.MatchString(pattern, number)
	return matched
}
func QueryInfo(train_id string, t *TrainInfo) *Error {
	if len(train_id) == 5 {
		if strings.ToLower(string(train_id[0])) != "t" {
			line, _ := train_id[:2], train_id[2:]
			line_int_tmp, _ := strconv.Atoi(line)
			line_pop_zero := strconv.Itoa(line_int_tmp)
			sql := "select * from line" + line_pop_zero + " where train_id='" + train_id + "'"

			err := psql.Init()
			if err != nil {
				e := &Error{Code: "0008", Msg: "PostgreSQL initialize error", Verbose: err}
				return e
			}
			db := psql.GetDB()
			// 执行查询
			res, err := db.Query(sql)
			if err != nil {
				e := &Error{Code: "0009", Msg: "SQL Query error", Verbose: err}
				return e
			}
			defer res.Close()
			defer db.Close()
			//fmt.Printf("pk\ttrain_id\ttrain_type\ttrain_detail\n")
			info_cpy := TrainInfo{}
			for res.Next() {
				err := res.Scan(&info_cpy.Pk, &info_cpy.TrainId, &info_cpy.Train_type, &info_cpy.TrainDetail)
				if err != nil {
					e := &Error{Code: "0010", Msg: "Scan data from psql goes wrong", Verbose: err}
					return e
				}
			}
			//judge the target info_cpy whether exist
			if info_cpy.TrainId == "" {
				t.Pk = 0
				t.TrainId = ""
				t.Train_type = ""
				t.TrainDetail = ""
				t.Carriage_number = []string{}
				t.IsEmpty = true
			} else {
				t.Pk = info_cpy.Pk
				t.TrainId = info_cpy.TrainId
				t.Train_type = info_cpy.Train_type
				t.TrainDetail = info_cpy.TrainDetail
				t.IsEmpty = false
			}
			return nil

		} else {
			//pj
			sql := "select * from \"linePJ\" where train_id = '" + train_id + "';"
			err := psql.Init()
			if err != nil {
				e := &Error{Code: "0008", Msg: "PostgreSQL initialize error", Verbose: err}
				return e
			}
			db := psql.GetDB()
			// 执行查询
			res, err := db.Query(sql)
			if err != nil {
				e := &Error{Code: "0009", Msg: "SQL Query error", Verbose: err}
				return e
			}
			defer res.Close()
			defer db.Close()
			//fmt.Printf("pk\ttrain_id\ttrain_type\ttrain_detail\n")
			info_cpy := TrainInfo{}
			for res.Next() {
				err := res.Scan(&info_cpy.Pk, &info_cpy.TrainId, &info_cpy.Train_type, &info_cpy.TrainDetail)
				if err != nil {
					e := &Error{Code: "0010", Msg: "Scan data from psql goes wrong", Verbose: err}
					return e
				}
			}
			//judge the target info_cpy whether exist
			if info_cpy.TrainId == "" {
				t.Pk = 0
				t.TrainId = ""
				t.Carriage_number = []string{}
				t.Train_type = ""
				t.TrainDetail = ""
				t.IsEmpty = true
			} else {
				t.Pk = info_cpy.Pk
				t.TrainId = info_cpy.TrainId
				t.Train_type = info_cpy.Train_type
				t.TrainDetail = info_cpy.TrainDetail
				t.IsEmpty = false
			}
			return nil
		}

	} else if len(train_id) == 6 {
		line := strings.ToLower(string(train_id[:2]))
		switch line {
		case "sj":
			{
				sql := "select * from \"lineSJ01\" where train_id = '" + train_id + "';"
				err := psql.Init()
				if err != nil {
					e := &Error{Code: "0008", Msg: "PostgreSQL initialize error", Verbose: err}
					return e
				}
				db := psql.GetDB()
				// 执行查询
				res, err := db.Query(sql)
				if err != nil {
					e := &Error{Code: "0009", Msg: "SQL Query error", Verbose: err}
					return e
				}
				defer res.Close()
				defer db.Close()
				//fmt.Printf("pk\ttrain_id\ttrain_type\ttrain_detail\n")
				info_cpy := TrainInfo{}
				for res.Next() {
					err := res.Scan(&info_cpy.Pk, &info_cpy.TrainId, &info_cpy.Train_type, &info_cpy.TrainDetail)
					if err != nil {
						e := &Error{Code: "0010", Msg: "Scan data from psql goes wrong", Verbose: err}
						return e
					}
				}
				//judge the target info_cpy whether exist
				if info_cpy.TrainId == "" {
					t.Pk = 0
					t.TrainId = ""
					t.Train_type = ""
					t.TrainDetail = ""
					t.Carriage_number = []string{}
					t.IsEmpty = true
				} else {
					t.Pk = info_cpy.Pk
					t.TrainId = info_cpy.TrainId
					t.Train_type = info_cpy.Train_type
					t.TrainDetail = info_cpy.TrainDetail
					t.IsEmpty = false
				}
				return nil
			}
		case "jy":
			{
				sql := "select * from \"lineJY01\" where train_id = '" + train_id + "';"
				err := psql.Init()
				if err != nil {
					e := &Error{Code: "0008", Msg: "PostgreSQL initialize error", Verbose: err}
					return e
				}
				db := psql.GetDB()
				// 执行查询
				res, err := db.Query(sql)
				if err != nil {
					e := &Error{Code: "0009", Msg: "SQL Query error", Verbose: err}
					return e
				}
				defer res.Close()
				defer db.Close()
				//fmt.Printf("pk\ttrain_id\ttrain_type\ttrain_detail\n")
				info_cpy := TrainInfo{}
				for res.Next() {
					err := res.Scan(&info_cpy.Pk, &info_cpy.TrainId, &info_cpy.Train_type, &info_cpy.TrainDetail)
					if err != nil {
						e := &Error{Code: "0010", Msg: "Scan data from psql goes wrong", Verbose: err}
						return e
					}
				}
				//judge the target info_cpy whether exist
				if info_cpy.TrainId == "" {
					t.Pk = 0
					t.TrainId = ""
					t.Train_type = ""
					t.TrainDetail = ""
					t.Carriage_number = []string{}
					t.IsEmpty = true
				} else {
					t.Pk = info_cpy.Pk
					t.TrainId = info_cpy.TrainId
					t.Train_type = info_cpy.Train_type
					t.TrainDetail = info_cpy.TrainDetail
					t.IsEmpty = false
				}
				return nil
			}
		case "jc":
			{
				sql := "select * from \"lineJC\" where train_id = '" + train_id + "';"
				err := psql.Init()
				if err != nil {
					e := &Error{Code: "0008", Msg: "PostgreSQL initialize error", Verbose: err}
					return e
				}
				db := psql.GetDB()
				// 执行查询
				res, err := db.Query(sql)
				if err != nil {
					e := &Error{Code: "0009", Msg: "SQL Query error", Verbose: err}
					return e
				}
				defer res.Close()
				defer db.Close()
				//fmt.Printf("pk\ttrain_id\ttrain_type\ttrain_detail\n")
				info_cpy := TrainInfo{}
				for res.Next() {
					err := res.Scan(&info_cpy.Pk, &info_cpy.TrainId, &info_cpy.Train_type, &info_cpy.TrainDetail)
					if err != nil {
						e := &Error{Code: "0010", Msg: "Scan data from psql goes wrong", Verbose: err}
						return e
					}
				}
				//judge the target info_cpy whether exist
				if info_cpy.TrainId == "" {
					t.Pk = 0
					t.TrainId = ""
					t.Train_type = ""
					t.TrainDetail = ""
					t.Carriage_number = []string{}
					t.IsEmpty = true
				} else {
					t.Pk = info_cpy.Pk
					t.TrainId = info_cpy.TrainId
					t.Train_type = info_cpy.Train_type
					t.TrainDetail = info_cpy.TrainDetail
					t.IsEmpty = false
				}
				return nil
			}
		default:
			{
				t.Pk = 0
				t.TrainId = ""
				t.Train_type = ""
				t.TrainDetail = ""
				t.Carriage_number = []string{}
				t.IsEmpty = true
				return &Error{Code: "0035", Msg: "Unknown type train id"}
			}
		}

	} else {
		e := &Error{Code: "0011", Msg: "Invalid input train_id"}
		return e
	}
}
func QueryCarriage(line string, number string, t *TrainInfo) *Error {
	sql := "select * from \"line" + line + "-carriages<->trains\" where left(carriage_number,length(carriage_number)-1)=left('" + number + "',length('" + number + "')-1)"
	err := psql.Init()
	if err != nil {
		e := &Error{Code: "0015", Msg: "PostgreSQL initialize error", Verbose: err}
		return e
	}
	db := psql.GetDB()
	// 执行查询
	res, err := db.Query(sql)
	if err != nil {
		e := &Error{Code: "0016", Msg: "SQL Query error", Verbose: err}
		return e
	}
	defer res.Close()
	defer db.Close()

	for res.Next() {
		all_carriage := ""
		correct_carriage := ""
		err := res.Scan(&correct_carriage, &t.TrainId, &t.Carriage_index, &all_carriage)

		t.Carriage_number = strings.Split(all_carriage, ",")
		if err != nil {
			e := &Error{Code: "0017", Msg: "Scan data from psql goes wrong", Verbose: err}
			return e
		}
		if correct_carriage == number {
			t.IsInputCarriageCorrect = true
		} else {
			t.IsInputCarriageCorrect = false
		}
	}
	return nil
}
