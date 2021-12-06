# API Spec

## Member Service

For member

Request :
- Header :
    - X-Api-Key : "your secret api key"

## All Address find by field address

Request :
- Method : GET
- Endpoint : `api/member/all-address/:filter`
- Header :
    - Content-Type: application/json
    - Accept: application/json
- Body :

Response :

```json 
{
    "code" : "number",
    "status" : "string",
    "data" : [
        {
            "_id": "string",
            "_id_provinsi": "string",
            "provinsi": "string",
            "_id_kota_kabupaten": "string",
            "_id_kelurahan": "string",
            "kota_kabupaten": "string",
            "_id_kecamatan": "string",
            "kecamatan": "string",
            "kelurahan": "string",
            "kodepos": "string",
            "address": "string"
        }
    ]
}
```

## All Address find kodepos by id kecamatan

Request :
- Method : GET
- Endpoint : `api/member/kodepos/:id`
- Header :
    - Content-Type: application/json
    - Accept: application/json
- Body :

Response :

```json 
{
    "code" : "number",
    "status" : "string",
    "data" : [
        {
            "_id": "string",
            "_id_provinsi": "string",
            "provinsi": "string",
            "_id_kota_kabupaten": "string",
            "_id_kelurahan": "string",
            "kota_kabupaten": "string",
            "_id_kecamatan": "string",
            "kecamatan": "string",
            "kelurahan": "string",
            "kodepos": "string",
            "address": "string"
        }
    ]
}
```

## All Address find Kota Kabupaten by id dan nama kota / kabupaten

Request :
- Method : GET
- Endpoint : `api/member/col-kota-kabupaten/1/namakotakabupaten/bandung`
- Header :
    - Content-Type: application/json
    - Accept: application/json
- Body :

Response :

```json 
{
    "code" : "number",
    "status" : "string",
    "data" : [
        {
            "_id": "string",
            "_id_provinsi": "string",
            "provinsi": "string",
            "_id_kelurahan": "string",
            "kota_kabupaten": "string",            
        }
    ]
}
```

## All Address find Kelurahan by id dan nama Kelurahan

Request :
- Method : GET
- Endpoint : `api/member/col-kelurahan/2/namakelurahan/coblong`
- Header :
    - Content-Type: application/json
    - Accept: application/json
- Body :

Response :

```json 
{
    "code" : "number",
    "status" : "string",
    "data" : [
        {
            "_id": "string",
            "_id_provinsi": "string",
            "provinsi": "string",
            "_id_kota_kabupaten": "string",            
            "kota_kabupaten": "string",    
            "_id_kecamatan": "string",
            "kecamatan": "string",
            "kelurahan": "string",        
        }
    ]
}
```

## All Address find Kecamatan by id dan nama Kecamatan

Request :
- Method : GET
- Endpoint : `api/member/col-kecamatan/2/namakecamatan/`
- Header :
    - Content-Type: application/json
    - Accept: application/json
- Body :

Response :

```json 
{
    "code" : "number",
    "status" : "string",
    "data" : [
        {
            "_id": "string",
            "_id_provinsi": "string",
            "provinsi": "string",
            "_id_kota_kabupaten": "string",
            "kota_kabupaten": "string",
            "_id_kecamatan": "string",
            "kecamatan": "string",            
        }
    ]
}
```

## All Address find Provinsi by id dan nama Provinsi

Request :
- Method : GET
- Endpoint : `api/member/provinsi/2/namaprovinsi/jawa barat`
- Header :
    - Content-Type: application/json
    - Accept: application/json
- Body :

Response :

```json 
{
    "code" : "number",
    "status" : "string",
    "data" : [
        {
            "_id": "string",
            "_id_provinsi": "string",
            "provinsi": "string",
            "_id_kota_kabupaten": "string",
            "_id_kelurahan": "string",
            "kota_kabupaten": "string",
            "_id_kecamatan": "string",
            "kecamatan": "string",
            "kelurahan": "string",
            "kodepos": "string",
            "address": "string"
        }
    ]
}
```

## All Member find Member by Gender, fullname dan user id

Request :
- Method : POST
- Endpoint : `api/member/find`
- Header :
    - Content-Type: application/json
    - Accept: application/json
- Body :

```json 
{
    "filter" : "string",
}
```

Response :

```json 
{
    "code" : "number",
    "status" : "string",
    "data" : [
        {
            "_id": "string",
            "full_name": "string",
            "birth_date": "string",
            "profile_picture": "string",
            "gender": "string"
        }
    ]
}
```

## All Member Address find by _id, user_id, name

Request :
- Method : POST
- Endpoint : `api/member/address-find`
- Header :
    - Content-Type: application/json
    - Accept: application/json
- Body :

```json 
{
    "filter" : "string",
}
```

Response :

```json 
{
    "code" : "number",
    "status" : "string",
    "data" : [
        {
            "_id": "string",
            "user_id": "string",
            "name": "string",
            "is_default": "bool",
            "receiver_name": "string",
            "phone": "string",
            "_id_provinsi": "string",
            "provinsi": "string",
            "_id_kota_kabupaten": "string",
            "kota_kabupaten": "string",
            "_id_kecamatan": "string",
            "kecamatan": "string",
            "_id_kelurahan": "string",
            "kelurahan": "string",
            "_id_kode_pos": "string",
            "kode_pos": "string",
            "latitude": "string",
            "longitude": "string",
            "address": "string"
        }
    ]
}
```

## Delete Foto Profile

Request :
- Method : DELETE
- Endpoint : `/api/member/deleted-picture`
- Header :
    - Content-Type: application/json
    - Accept: application/json
- Body :

Response :

```json 
{
    "code" : "number",
    "status" : "string",
    "data" : "string"
}
```
## Delete Foto Profile

Request :
- Method : DELETE
- Endpoint : `/api/member/address/:_id`
- Header :
    - Content-Type: application/json
    - Accept: application/json
- Body :

Response :

```json 
{
    "code" : "number",
    "status" : "string",
    "data" : "string"
}
```

## Member Address set default

Request :
- Method : PUT
- Endpoint : `/api/member/address-is-default`
- Header :
    - Content-Type: application/json
    - Accept: application/json
- Body :

```json 
{    
    "_id":"string",
}
```

Response :

```json 
{
    "code" : "number",
    "status" : "string",
    "data" : "string"
}
```

## Update Member Address

Request :
- Method : PUT
- Endpoint : `/api/member/address`
- Header :
    - Content-Type: application/json
    - Accept: application/json
- Body :

```json 
{    
    "_id": "string",
    "name": "string",
    "is_default": "bool",
    "receiver_name": "string",
    "phone": "string",
    "_id_provinsi": "string",
    "provinsi": "string",
    "_id_kota_kabupaten": "string",
    "kota_kabupaten": "string",
    "_id_kecamatan": "string",
    "kecamatan": "string",
    "_id_kelurahan": "string",
    "kelurahan": "string",
    "_id_kode_pos": "string",
    "kode_pos": "string",
    "latitude": "string",
    "longitude": "string",
    "address": "string"
}
```

Response :

```json 
{
    "code" : "number",
    "status" : "string",
    "data" : {
        "_id": "string",
        "name": "string",
        "is_default": "bool",
        "receiver_name": "string",
        "phone": "string",
        "_id_provinsi": "string",
        "provinsi": "string",
        "_id_kota_kabupaten": "string",
        "kota_kabupaten": "string",
        "_id_kecamatan": "string",
        "kecamatan": "string",
        "_id_kelurahan": "string",
        "kelurahan": "string",
        "_id_kode_pos": "string",
        "kode_pos": "string",
        "latitude": "string",
        "longitude": "string",
        "address": "string"
    }
}
```

## Create Member Address

Request :
- Method : POST
- Endpoint : `/api/member/address`
- Header :
    - Content-Type: application/json
    - Accept: application/json
- Body :

```json 
{    
    "_id": "string",
    "name": "string",
    "is_default": "bool",
    "receiver_name": "string",
    "phone": "string",
    "_id_provinsi": "string",
    "provinsi": "string",
    "_id_kota_kabupaten": "string",
    "kota_kabupaten": "string",
    "_id_kecamatan": "string",
    "kecamatan": "string",
    "_id_kelurahan": "string",
    "kelurahan": "string",
    "_id_kode_pos": "string",
    "kode_pos": "string",
    "latitude": "string",
    "longitude": "string",
    "address": "string"
}
```

Response :

```json 
{
    "code" : "number",
    "status" : "string",
    "data" : {
        "_id": "string",
        "name": "string",
        "is_default": "bool",
        "receiver_name": "string",
        "phone": "string",
        "_id_provinsi": "string",
        "provinsi": "string",
        "_id_kota_kabupaten": "string",
        "kota_kabupaten": "string",
        "_id_kecamatan": "string",
        "kecamatan": "string",
        "_id_kelurahan": "string",
        "kelurahan": "string",
        "_id_kode_pos": "string",
        "kode_pos": "string",
        "latitude": "string",
        "longitude": "string",
        "address": "string"
    }
}
```

## Create Member

Request :
- Method : POST
- Endpoint : `/api/member/create`
- Header :
    - Content-Type: application/json
    - Accept: application/json
- Body :

```json 
{    
    "full_name":"string",
    "birth_date":"string",
    "gender":"string",
    "profile_picture":"string"
}
```

Response :

```json 
{
    "code" : "number",
    "status" : "string",
    "data" : {
        "_id":"string",
        "full_name":"string",
        "birth_date":"string",
        "gender":"string",
        "profile_picture":"string"
    }
}
```