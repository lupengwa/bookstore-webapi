# Adding data to database
1. Start database - ./postgres_start.sh
   - note: you can make this persistent by specifying volumes in the script such as adding:
   ```
   -e PGDATA=/var/lib/postgresql/data/pgdata \
   -v /Users/frankmoley/.local/docker/data:/var/lib/postgresql/data \
   ```
2. enter data folder, copy schema and data file to container, then run sql
```
   //copy
   docker cp ./schema.sql local-pg:/docker-entrypoint-initdb.d/schema.sql
   docker cp ./data.sql local-pg:/docker-entrypoint-initdb.d/data.sql
   
   //execute db query
   docker exec -u postgres local-pg psql postgres postgres -f docker-entrypoint-initdb.d/schema.sql
   docker exec -u postgres local-pg psql postgres postgres -f docker-entrypoint-initdb.d/data.sql
```
3. Exec into docker container
   ```
   docker exec -it local-pg /bin/bash
   ```

4. Launch psql from inside the docker container, you can run any postgres sql here
   ```
   psql -U postgres
   ```



