docker volume create weact-file &&
docker run \
-d \
-p 8000:8000 \
-v weact-file:/weact-backend/public/file \
weact-backend:v1.10