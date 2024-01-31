# gomemstress

[![Docker Repository on Quay](https://quay.io/repository/maschulz/gomemstress/status "Docker Repository on Quay")](https://quay.io/repository/maschulz/gomemstress)

> `gomemstress` is a dead-simple tool that just allocates memory.


## Usage

```sh
./memstress -h
Usage of ./memstress:
  -max int
    	Maximum memory that should be consumed (default 9223372036854775807)
  -size int
    	Chunk size in MiB to allocate per cycle (default 100)
  -sleep int
    	Time in ms to wait between allocations (default 250)
```
