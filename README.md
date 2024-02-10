# Locksmith 🔐
Identity server

### Usages
To run the locksmith locally, first run the docker-compose make cmd
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
docker pull bandnoticeboard/locksmith:v1.0.1
```

Or select a version here - [bandnoticeboard/locksmith](https://hub.docker.com/r/bandnoticeboard/locksmith)

## Authors

* **joegasewicz** - *Initial work* - [@joegasewicz](https://twitter.com/joegasewicz)

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.
## License
[MIT](https://choosealicense.com/licenses/mit/)