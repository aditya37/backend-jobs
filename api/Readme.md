# Coding Conventions

Terimakasih sudah menyempatkan waktu untuk membaca ini. Dalam proyek ini memiliki standarisasi untuk penulisan code dan nama file. Hal ini bertujuan untuk mempermudah dalam pengolahan code dan agar terlihat rapi.

## Penamaan file

Dalam penulisan nama file coding menggunakan standar sebagai berikut :

### Struktur file di project ini dibagai menjadi 6 bagian,yaitu:

* Folder repository
* Folder Controller
* Folder Error
* Folder Middleware
* Folder Network
* Folder Service
* Folder Model

Dari ke-7 folder tersebut (kecuali folder model) akan di isi oleh dua file yang memiliki jenis berbeda, yaitu: 

* Interface 
* Function

format penulisnya adalah :

_I(NamaFunction)(NamaFolder)_

_(NamaFunction)Impl(NamaFolder)_

untuk aturan atau standarisasi penulisan nama file harus disertakan jenis file dan nama foldernya. untuk lebih jelasnya lihat contoh berikut:

```
Membuat file untuk folder repository yang berfungsi untuk menampilkan data, maka untuk penulisannya sebagai berikut:

ICrudRepo.go
CrudImplRepo.go
```

## Penamaan Function
untuk penamaan function menggunakan CamleCase. contoh:

```func GetAllData() {}```

## Import 
jika mengimport sebuah module atau class lebih dari satu gunakan standar penulisan seperti ini 

```
import (
    "nama_module1",
    "nama_module2"
)
```
## Formating
Untuk kurung kurawal ```{ }``` , letakkan kurung kurawal buka ( ```{``` ) di ujung garis atau di akhir  ```function()``` dengan satu spasi contoh 

``` 
func main() {

}
```
dan untuk kurung kurawal tutup ( ```}``` ) harus sejajar dengan ```function()```

oke, sekian tentang penjelasan standart penulisan code.
suatu saat dokumentasi ini akan berubah 

Terima kasih, matur nuwun, arigatou, xiexie.