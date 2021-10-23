# GraphQl Dummy

## Start Api

## Build
```bash
$ ./build.sh
or
$ make build
```

## Application binary
```bash
$ book-author serve
or
$ make run
```

## Container dev
```bash
$ docker-compose up --build
or
$ make serve
```

## GuideLine

* api folder contains api code

* infra contains drivers like db, messaging, cache etc
* repo folder contains database code
* model folder contains model
* schema folder contains application graphql schema code

### flow
> cmd -> api -> schema -> repo, models, cache, messaging