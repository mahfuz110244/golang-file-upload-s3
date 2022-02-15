## How to upload multiple files in AWS S3 Bucket in Golang

# SETUP

Copy `sample.env` to `.env`.

# RUN

`go run main.go`

# API Test
```
curl --location --request POST 'http://localhost:1323/upload' \
--form 'files=@"/home/mahfuz/Downloads/Fuz.jpg"' \
--form 'files=@"/home/mahfuz/Downloads/Fuz 2x.jpg"' \
--form 'name="Mahfuz"' \
--form 'email="mahfuzku11@gmail.com"'
```