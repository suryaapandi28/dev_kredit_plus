# KREDIT PLUS
---
SELAMAT DATANG
Di Aplikasi Kredit Plus PT. XYZ Multifinance! Nikmati kemudahan akses pembiayaan untuk White Goods, Motor dan Mobil terbesar di indonesia. Kami Berkomitmen meningkatkan kualiatas hidup anda melalui solusi pembiayaan inovatif dan teknologi terkini.
Kami hadir untuk mempermudah hidup anda! dengan aplikasi Kredit Plus, anda bisa mengajukan berbagai jenis pembiayaan secara cepat dan mudah.

Bersama PT. XYZ Multifinance memujudkan mimpi anda lebih mudah. kami hadir untuk membantu Anda mencapai potensi terbaik dengan solusi pembiayaan yang fleksibel dan terjangkau. Nikmati proses pengajuan yang cepat, persyaratan yang mudah, dan berbagai pilihan produk pembiayaan yang sesuai dengan kebutuhan Anda. Bersama kami, wujudkan gaya hidup yang lebih baik.

KREDIT PLUS !!!

"Solusi Pembiayaan Terpercaya Untuk Semua Kebutuhan Anda"

---
## Installation 👨🏻‍💻

Kreditplus build by [Go](https://go.dev/dl/) Go 1.13+ to run.

1. Clone Repository
   By use terminal/cmd

```sh
git clone https://github.com/suryaapandi28/dev_kredit_plus.git
```

2. Open Repository
   By use terminal/cmd

```sh
cd kreditplus
```

2. Check the .env file and configure with your device

```sh
Example file .env
ENV="yourEnv"
PORT="your using port ex("8080")"
POSTGRES_HOST= "your psql host ex("localhost")"
POSTGRES_PORT="your psql port ex("5432")"
POSTGRES_USER="kreditplus"
POSTGRES_PASSWORD="kreditplus"
POSTGRES_DATABASE="kreditplus"
REDIS_HOST="localhost"
REDIS_PORT="6379"
JWT_SECRET_KEY="erwhdhsajdahjkdhqwiuou"
```

3. Enable the PostgreSQL database
   Option you can use :

   - [pgAdmin](https://www.pgadmin.org/)
   - [NaviCat](https://www.navicat.com/en/download/navicat-premium?gad_source=1&gclid=CjwKCAjwmYCzBhA6EiwAxFwfgFWv6YNc_nwrdL5BByjvaEmUNbzD0vvg-tHgv7x6rFyIx-zSdWYQWhoCRP0QAvD_BwE)
   - Or anything else you usualy use

4. Run the command to create the database and migrate it.
   Make sure you have install migrate cli.
   If you dont, install first by
**If you windows user**🪟

- Open PowerShell. You can do this by searching for PowerShell in the start menu.
- Inside PowerShell, type the code below

```sh
iwr -useb get.scoop.sh | iex
```

Then Install use scoop

```sh
scoop install migrate
```

After all, migrate it by

```sh
migrate -database "postgres://kreditplus:kreditplus@localhost:5432/kreditplus?sslmode=disable" -path db/migrations up
```

5. Configure Docker
   **First Install Docker**
   - Windows User[Docker](https://docs.docker.com/desktop/install/windows-install/)
   - Mac User [Docker](https://docs.docker.com/desktop/install/mac-install/)
   - Then compose it by

```sh
docker-compose up -d
```

6. Run the application

```sh
go mod tidy
go run cmd/app/main.go
```

---

## Tech 💻

Depublic uses a number of open source projects to work properly :

1. **Golang** - High-performance language for scalable apps.
2. **Echo Framework** - Web framework for Go.
3. **PostgreSQL** - Reliable open-source database.
4. **Docker** - Container platform for consistent deployment.

## Architecture 𐄷

- Using Model-View-Controller (MVC) basic architecture with two layers (backend and frontend).
- Using project layout by (https://github.com/golang-standards/project-layout)
- Implementation of search, filter, and sorting features.


## API Documentation 🔗

Documentation for API can be get by :

```sh
https://documenter.getpostman.com/view/24409024/2sA3s9DoK8
```
## Detail Image ⚒️

![WhatsApp Image 2024-08-22 at 16 09 37_c9b525f7](https://github.com/user-attachments/assets/311a0f51-ba03-494b-8b77-e51f66e67427)
