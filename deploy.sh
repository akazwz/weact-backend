docker volume create weact-file &&
docker run \
-d \
-p 8000:8000 \
-v weact-file:/weact-backend/public/file \
registry-vpc.cn-shenzhen.aliyuncs.com/akazwz/weact-backend:v1.10