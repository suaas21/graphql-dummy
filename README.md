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
$ graphql-dummy serve
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
> cmd -> api -> schema -> repo -> db infra


## Book Author API:

**Mutation :**

- create Book:
```json
http://localhost:8080/api/v1/public/graphql?query=mutation+_{book(id:"1",name:"Sagor",description:"childhood",author_ids:["1"]){id,name,description}}
```

- Create Author:
```json
http://localhost:8080/api/v1/public/graphql?query=mutation+_{author(id:"1",name:"Sagors childhood",book_ids:["1"]){id,name}}
```
**Query :**
- Get Book by ID
```json
http://localhost:8080/api/v1/public/graphql?query={book(id:"1"){id,name,description,authors{id,name}}}
```

- Get Author by ID
```json
http://localhost:8080/api/v1/public/graphql?query={author(id:"1"){id,name,books{id,name}}}
```
