To push to the docker registry:

```sh
docker build -t echo-server:v1 .
docker tag echo-server:v1 drio/echo-server:latest
docker push drio/echo-server:latest
```
