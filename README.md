
# shmetroDB

> 上海地铁车厢号数据库 – 查询车厢号对应列车信息

## 简介

shmetroDB 是一个专注于 **上海地铁车厢号 → 列车信息** 唯一映射查询数据库服务。用户只需输入车厢号（或部分），即可快速定位所属列车编号、车型、车厢编组详情等。
该项目由 来降 开发，本项目前后端分离：

* 后端：采用 Gin (Go) 提供 API 接口。
* 前端：采用 Vite + Vue3 构建用户查询交互界面。



## 支持的线路

*  上海地铁1至18号线，上海地铁浦江线，市域机场线，机场捷运线，松江有轨电车。


## 快速开始

### 克隆项目

```bash
git clone https://github.com/Martyr404/shmetroDB.git
cd shmetroDB
```

### 后端api运行

```bash
cd ./api
go mod tidy
# （可选）设置运行模式为 release
# export GIN_MODE=release
go run main.go
```

### 使用docker运行前端界面（开发模式）

```bash
cd ./gui/gui
docker run -it -p 5173:5173 -v "$(pwd)":/gui -w /gui node:24-alpine sh -c "npm install -g pnpm && sh"
pnpm install
pnpm dev --host
```


## 接口说明

* 请求路径：`POST /api/calculate`
* 请求体示例：

  ```json
  { "carriage_number": "02011" }
  ```
* 成功响应示例：

  ```json
  {
    "Data": [
      {
        "Pk": 1,
        "IsEmpty": false,
        "TrainId": "03001",
        "Train_type": "03A01",
        "Carriage_number": ["02011","02022","02033","02043","02052","02061"],
        "Carriage_index": "0",
        "TrainDetail": "法国进口列车",
        "IsInputCarriageCorrect": true
      }
    ],
    "Msg": "Query carriage number:02011 successfully",
    "StateCode": "2000"
  }
  ```
* 错误或无结果示例：

  ```json
  { "Data": [], "Msg": "Query carriage number:189991 not found", "StateCode": "2000" }
  ```

  或

  ```json
  {
    "Data": { "Code": "0001", "Msg": "Unknown type carriage number." },
    "Msg": "parse carriage number goes wrong",
    "StateCode": "5000"
  }
  ```

## 贡献指南

欢迎提出 issue 或提交 pull request。请先 fork 本仓库、创建分支，修改完成后提交相应 PR。建议开启 `release` 模式、在 `main` 分支之外开发。

## 联系方式

如发现 Bug 或建议功能，请前往项目的 GitHub Issues 页面：
[GitHub Issues](https://github.com/Martyr404/shmetroDB/issues)


