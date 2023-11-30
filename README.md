# SIAKAD SISWA SERVICE
Ini adalah salah satu demo dalam rancang bangun SISTEM INFORMASI AKADEMIS dengan pendekatan Microservice.
Pada service ini, akan dibangun sample CRUD dari entitas "SISWA" walaupun belum terintegrasi dengan service lainnya.


## Sistem Requirements
1. GO 1.20
2. MariaDB 10.4.27


## Database Preparation
1. Buat database dengan nama yang ingin ditentukan
2. Buka file .env dan sesuaikan isi dari env tersebut dan isi nama database sesuai nama yang dibuat

## Service Preparation
1. Buka terminal
2. Ketikan 
``` bash
 go mod tidy
 go mod vendor
```
3. Perhatikan agar depedency bisa terinstall sepenuhnya

## Eksekusi service
``` bash
go run ../cmd/main.go
```