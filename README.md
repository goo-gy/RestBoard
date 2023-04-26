## env 작성

```
DB_HOST=[db url]
DB_USER=[username]
DB_PASSWORD=[password]
DB_NAME=[database name]
DB_PORT=[port]
```

Go, docker-compose에서 사용

개발 모드에서 DB_HOST=localhost


## Develop DB Container

``` shell
docker-compose up -d
```

작성된 .env를 기반으로 Database Container 실행


### DB Container 접속

``` shell
docker exec -it practice-db /bin/bash
 
psql -U postgres
```
