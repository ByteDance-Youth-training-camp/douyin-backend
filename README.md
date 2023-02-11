# Douyin backend
第五届字节跳动青训营后端组项目

### Prerequisites
- docker, docker-compose v2
- golang 1.19
- thumbnailer dependencies:
  - pkg-config
  - ffmpeg
  - c11 compiler

### 启动方式
启动前需要修改minio的host为客户端可访问的外部地址
``` bash
make run
```