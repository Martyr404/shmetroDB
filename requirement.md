前后端通信请求
request:
 header:
    cookie:${cookies}
 body:
    carriage_number:080071

启动docker的nodejs22容器
docker run -it -p 5173:5173 -v $(pwd):/gui -w /gui node:22-alpine sh -c "npm install -g pnpm && sh"

启动vue项目
pnpm create vue@latest gui