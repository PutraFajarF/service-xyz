# Command Dockerize App
make build

# Evidence Result in folder ./result

# OWASP Implementation
- Validation Input
- Prepared Statement Querry untuk mencegah SQL injection
- Menggunakan parameterized query untuk mencegah SQL injection
- Implemenet go routine & sync.Mutex untuk mencegah race condition
- Implement Detail Log Information for observability
- Implement setMaxIdleCons, setMaxOpenConnsm setConnMaxLifeTime in mysql connection
- Implement timeout in http server connection