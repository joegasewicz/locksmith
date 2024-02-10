# Locksmith 🔐
Identity server

### Usages
To run locksmith locally, first run the docker-compose make cmd
```
make docker-compose
```

Now you can run locksmith
```
go run locksmith.go
```

### Config
Create a yaml file called `locksmith.yaml`

To set default users & role types. *The `user.role` must 
match the `roles` values*.
```yaml
version: 1
roles:
  - admin
  - teachers
  - students
users:
  - name: admin
    role: admin
    email: test1@gmail.com
    password: admin
... etc.
```

### Docker Image
Pull down locksmith from the Docker registry 
```
docker pull bandnoticeboard/locksmith:v1.0.2
```

Or select a version here - [bandnoticeboard/locksmith](https://hub.docker.com/r/bandnoticeboard/locksmith)

### Docker Compose
Example of running locksmith with Docker Compose.
This example also include the required postgres database configuration.
```
  postgres_locksmith:
    image: "postgres:latest"
    ports:
      - "5431:5432"
    environment:
      - POSTGRES_USER=admin
      - POSTGRES_PASSWORD=admin
      - POSTGRES_DB=identity_db
    volumes:
      - ./db/identitydb_vol/:/var/lib/postgresql/data

  locksmith:
    image: "bandnoticeboard/locksmith:v1.0.2"
    ports:
      - "7001:7001"
    env_file:
      - ".env-dev"
    depends_on:
      - postgres_locksmith
    restart: on-failure
```

### Environment Variables
- `TOKEN_SECRET`
- `PGDATABASE`
- `PGUSER`
- `PGPASSWORD` 
- `PGPORT`
- `PGHOST`

Example `.env` file
```text
TOKEN_SECRET=wizard
PGDATABASE=identity_db
PGUSER=admin
PGPASSWORD=admin
PGPORT=5431
PGHOST=host.docker.internal
```

### Endpoints
- GET `http://127.0.0.1:7001/health`
Response:
```
{"Health":"OK"}
```
## Authors

* **joegasewicz** - *Initial work* - [@joegasewicz](https://twitter.com/joegasewicz)

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.
## License
[MIT](https://choosealicense.com/licenses/mit/)