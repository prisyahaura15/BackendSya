# Befous
Nama	= Ibrohim Mubarok <br />
NPM		= 1214081 <br />
Kelas	= 3C <br />

## Mengambil semua data gejson

Link API-nya

```
https://asia-southeast2-gis3-401509.cloudfunctions.net/BefousAmbilDataGeojson
```

Response

```
[{
      "type": "Feature",
      "properties": {
        "name": "nama property"
      },
      "geometry": {
        "coordinates": [
          x,y
        ],
        "type": "tipe geojson"
      }
    }]
```

## Registrasi Akun

Link API-nya

```
https://asia-southeast2-gis3-401509.cloudfunctions.net/BefousMembuatUser
```

Body

```
{
    "username": "input username di sini",
    "password": "input password di sini",
	"role": "input role di sini"
}
```

Response

```
{"status":true,"message":"Berhasil Input data"}
```

## Login Akun

### Membuat Token

Link API-nya

```
https://asia-southeast2-gis3-401509.cloudfunctions.net/BefousMembuatTokenUser
```

Body

```
{
    "username": "input username di sini",
    "password": "input password di sini"
}
```

Response bila berhasil

```
{"status":true,"token":"token yang didapat","message":"Selamat Datang"}
```

Response bila gagal

```
{"status":false,"message":"Password Salah"}
```

### Menyimpan Token

Link API-nya

```
https://asia-southeast2-gis3-401509.cloudfunctions.net/BefousLoginUser

```

Header

```
Login : masukkan token di sini
```

Response bila berhasil

```
{
    "status": true,
    "message": "data User berhasil diambil",
    "data": [
        {
            "username": "data",
            "password": "data",
            "role": "role"
        },
        {
            "username": "data",
            "password": "data",
            "role": "role"
        }
    ]
}
```

Response bila gagal

```
{"status":false,"message":"Data Username tidak ada di database"}
```

## Delete Akun

Link API-nya

```
https://asia-southeast2-gis3-401509.cloudfunctions.net/BefousHapusUser

```

Body

```
{
    "username": "input username di sini"
}
```

Response bila berhasil

```
{"status":false,"message":"Berhasil Delete data"}
```

Response bila gagal

```
{"status":false,"message":"error parsing application/json: EOF"}
```