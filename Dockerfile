# 1. Ambil base image Golang (ibarat ambil OS yang udah ada Go-nya)
FROM golang:alpine

# 2. Bikin folder kerja di dalam container
WORKDIR /app

# 3. Copy file go.mod dan go.sum
COPY go.mod go.sum ./

# 4. Download semua library yang dibutuhkan (Gin, Gorm, dll)
RUN go mod download

# 5. Copy semua kodingan kamu ke dalam container
COPY . .

# 6. Build aplikasinya jadi binary
RUN go build -o main .

# 7. Expose port 8080 (biar bisa diakses dari luar)
EXPOSE 8080

# 8. Perintah buat ngejalanin aplikasi pas container dinyalain
CMD ["./main"]