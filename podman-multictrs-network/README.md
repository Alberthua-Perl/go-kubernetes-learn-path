## Demo for Podman Multi-Containers Network

In this demo we will create two podman containers called `times-app` exposed 8080/tcp port
and `cities-app` exposed 8090/tcp port which are all in the `cities` container network based
`cni` network backend.

In `rootless` network namespace, we create container network and containers:

```bash
$ podman network create cities
# create podman container network in rootless-netns
```
These two apps connect with each other by container name in cities network enabled `dns`.
So we can use following commands to run times-app:

```bash
$ podman run --name times-app \
  --network cities -p 8080:8080 -d \
  <registry_uri>/<user_or_org>/podman-info-times:v0.1
# run times-app container specified cities network

$ podman inspect times-app \
  -f '{{.NetworkSettings.Network.cities.IPAddress}}'
# verify container IP address

$ podman run --rm --network cities \
  <registry_uri>/ubi8/ubi-minimal:8.5 \
  curl -s http://<times-app_ip_address>:8080/times/BKK && echo
# verify times-app response

$ podman run --rm --network cities \
  <registry_uri>/ubi8/ubi-minimal:8.5 \
  curl -s http://times-app:8080/times/BKK && echo
# use container name to verify response
```

As cities-app use `TIMES_APP_URL` environment variable to connect times-app, so use following
command to run:

```bash
$ podman run --name cities-app \
  --network cities -p 8090:8090 -d \
  -e TIMES_APP_URL=http://times-app:8080/times \
  <registry_uri>/<user_or_org>/podman-info-cities:v0.1
# view app code to verify TIMES_APP_URL function

$  curl -s http://localhost:8090/cities/MAD | jq
{
  "name": "Madrid",
  "time": "2022-06-01T12:34:56.123Z",
  "population": 3223000,
  "country":"Spain"
}
# get associated city info
```
