# A2OS Gin Boilerplate

The example of how A2OS use gin framework.

## Dependencies

- gin
- gin-csrf
- gin-cors
- sessions
- gorm
- viper
- gin-swagger
- swag
- raven-go

## Database

Currently we are using the `MySQL / MariaDB` driver. Redis and MongoDB support is on the way.

| Database | Status |
| :---: | :---: |
| MySQL / MariaDB | ✔️ |
| Redis | ❌ |
| MongoDB | ❌ |

## API Documentation

Update API Documentation:

```bash
$ swag init
```

## Gin-swagger

Annotation documentation: https://swaggo.github.io/swaggo.io/ .
