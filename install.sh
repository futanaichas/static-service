docker build -t futanaicha/static-service .
docker network create futanaicha
mkdir -p ~/devopt/wwwroot/static
docker run -d \
--net futanaicha \
--name static-service \
-v ~/devopt/wwwroot/:/var/www/futanaicha/public \
futanaicha/static-service