# futanaicha静态文件模块
这是futanaicha论坛的一个服务模块

```bash
docker build -t futanaicha/static-service .

mkdir -p /var/www/futanaicha/public

docker run -d \
-p 3000:3000 \
--name futanaicha-static-service \
-v /var/www/futanaicha/public: /var/www/futanaicha/public
futanaicha/static-service

```
## api
+ 获取数据`GET /static/2019-8-29/gsgiwbik.txt`
+ 储存数据`POST /api/static/upload`