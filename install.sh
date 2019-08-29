git pull
docker build -t futanaicha/static-service .
docker network create futanaicha
mkdir -p /var/www/futanaicha/public
docker run -d \
-p 3000:3000 \
--net futanaicha
--name static-service \
-v /var/www/futanaicha/public:/var/www/futanaicha/public \
futanaicha/static-service