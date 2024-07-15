To push to the docker registry:

```sh
docker build -t hello-server:v1 .
docker tag hello-server:v1 drio/hello-server:latest
docker push drio/hello-server:latest
```
