
## Run in Ubuntu

#### Install dieharder
```bash
sudo apt-get install -y dieharder
```
#### Measure

> `pv` used for speed measurement

```bash
go run ./container/rnd_stream.go | pv --bytes --timer | dieharder -a -g 200 -s 1
```

## Docker compose

`docker-compose --file=./container/docker-compose.yml up -d dieharder`

With Colima

https://github.com/abiosoft/colima/discussions/874

```bash
mkdir -p ~/.docker/cli-plugins
ln -sfn $HOMEBREW_PREFIX/opt/docker-compose/bin/docker-compose ~/.docker/cli-plugins/docker-compose
```
