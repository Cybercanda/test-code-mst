# Test Code MST
## Cara Menjalankan
1. Pastikan Go sudah terinstal: [https://golang.org/dl/](https://golang.org/dl/)
2. Clone repositori ini: `git clone https://github.com/Cybercanda/test-code-mst.git`
3. Masuk ke direktori aplikasi: `cd nama-aplikasi`
4. Jalankan aplikasi: `go run main.go`

## Beberapa paket yang digunakan:
1. gin: `github.com/gin-gonic/gin`
2. viper: `github.com/spf13/viper`

## Dokumentasi
### Base URL
```http
http://localhost:8080
```
### Endpoints
#### Get Players (Get)
```http
http://localhost:8080/api/players
```
Response:
![getall.png](image%2Fgetall.png)

#### Filter Players (Get)
```http
http://localhost:8080/api/players?BirthPlace=Europe
```
Response:
![filter.png](image%2Ffilter.png)

#### Get Players by ID (Get)
```http
http://localhost:8080/api/players/3
```
Response:
![getbyid.png](image%2Fgetbyid.png)

#### Create Players (Post)
```http
http://localhost:8080/api/players
```
file json
```json
{
  "Name": "Bily Hakim",
  "Age": 23,
  "BirthPlace": "Indonesia"
}
```
Response:
![createPlayers.png](image%2FcreatePlayers.png)

#### Api Settings
```http
http://localhost:8080/api/settings
```
Response:
![settings.png](image%2Fsettings.png)