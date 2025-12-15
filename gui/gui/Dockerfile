# 1. 拉取 Node 24 官方镜像
FROM node:24

# 2. 设置工作目录, 设置 CI 环境
WORKDIR /gui
ENV CI=true

# 3. 全局安装 pnpm
RUN ["npm","install","-g","pnpm"]

# 4. 拷贝宿主机当前目录到容器
COPY . .

# 5. 安装依赖并构建
RUN pnpm install && pnpm build

# 6. 暴露端口
EXPOSE 10011

# 7. 启动 preview
CMD ["pnpm", "preview", "--host", "0.0.0.0", "--port", "10011"]
