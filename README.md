## How to upload multiple files in AWS S3 Bucket in Golang

# SETUP

Copy `sample.env` to `.env` and update config file according to your AWS configuration

# RUN

`go run main.go`

# API Test
For Single File Upload Curl URL:
```
curl --location --request POST 'http://localhost:1323/upload' \
--form 'file=@"/home/mahfuz/Downloads/Fuz.jpg"' \
--form 'name="Mahfuz"' \
--form 'email="mahfuzku11@gmail.com"'
```

Response
```
{
    "status_code": 200,
    "success": true,
    "data": {
        "name": "Fuz",
        "url": "https://gorillamove.s3.ap-southeast-1.amazonaws.com/files/16-02-2022-11-30-57AM-85C8D8-Fuz.jpg",
        "size": 41064,
        "extension": ".jpg",
        "upload_status": true,
        "message": "upload file successfully to s3 bucket"
    }
}
```

For Bulk File Upload Curl URL:
```
curl --location --request POST 'http://localhost:1323/upload/bulk' \
--form 'files=@"/home/mahfuz/Downloads/Fuz.jpg"' \
--form 'files=@"/home/mahfuz/Downloads/Fuz 2x.jpg"' \
--form 'name="Mahfuz"' \
--form 'email="mahfuzku11@gmail.com"'
```

Response
```
{
    "status_code": 200,
    "success": true,
    "data": [
        {
            "name": "Fuz",
            "url": "https://gorillamove.s3.ap-southeast-1.amazonaws.com/files/16-02-2022-11-29-07AM-0ED7EA-Fuz.jpg",
            "size": 41064,
            "extension": ".jpg",
            "upload_status": true,
            "message": "upload file successfully to s3 bucket"
        },
        {
            "name": "Fuz 2x",
            "url": "https://gorillamove.s3.ap-southeast-1.amazonaws.com/files/16-02-2022-11-29-07AM-512B25-Fuz%202x.jpg",
            "size": 82008,
            "extension": ".jpg",
            "upload_status": true,
            "message": "upload file successfully to s3 bucket"
        }
    ]
}
```