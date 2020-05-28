export DEBUG=${DEBUG:="mediasoup:INFO* *WARN* *ERROR*"}

docker run --name ssr-server \
  --restart always \
  -p 8085:8085 \
  -v ${PWD}/.logs:/tmp/ssr/.logs \
  -e DEBUG \
  -d \
  ssr-server:0.1.0