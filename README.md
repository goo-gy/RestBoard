## DB 구성

- Postgre DB Container 구성

```
docker run -d -p 5432:5432 --name practice-db -e POSTGRES_PASSWORD=qwer1234 -e POSTGRES_DB=practice postgres
```

- docker 실행 & DB 접속

```
docker exec -it practice-db /bin/bash
 
psql -U postgres
```