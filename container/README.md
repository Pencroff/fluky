
Run in Ubuntu

#### Install dieharder
```bash
sudo apt-get install -y dieharder
```
#### Measure

> `pv` used for speed measurement

```bash
go run ./container/rnd_stream.go | pv --bytes --timer | dieharder -a -g 200 -s 1
```
