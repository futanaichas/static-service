git pull https://github.com/futanaichas/static-service.git && cd static-service
docker build -t futanaicha/static-service .
mkdir -p /var/www/futanaicha/public
docker run -d \
-p 3000:3000 \
--name futanaicha-static-service \
-v /var/www/futanaicha/public: /var/www/futanaicha/public
futanaicha/static-service