前后端通信请求
request:
 header:
    cookie:${cookies}
 body:
    carriage_number:070072

启动docker的nodejs22容器
docker run -it -p 5173:5173 -v $(pwd):/gui -w /gui node:22-alpine sh -c "npm install -g pnpm && sh"

启动vue项目
pnpm create vue@latest gui

sql初始化表格示例
insert into line9 (pk, train_id, train_type, train_detail)
select 
  generate_series(1, 105) as pk,
  -- 生成类似 '09001'、'09002'...'09105' 的train_id（假设格式）
  '09' || lpad(generate_series(1, 105)::text, 3, '0') as train_id,
  -- 生成相同格式的train_type（根据需求调整）
  '09' || lpad(generate_series(1, 105)::text, 3, '0') as train_type,
  '平平无奇' as train_detail;



c.JSON(http.StatusOK, gin.H{
				"msg": "you are searching number " + req.Carriage_number,
				"result": gin.H{
					"TrainId":                    trainInfo.TrainId,
					"Carriage_num":               trainInfo.Carriage_number,
					"Carriage_index":             trainInfo.Carriage_index,
					"Train_type":                 trainInfo.Train_type,
					"Train_detail":               trainInfo.TrainDetail,
					"isInputCarriageTypeCorrect": true,
				},
			})