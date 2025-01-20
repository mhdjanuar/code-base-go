# Code Base with go

Proyek ini mengimplementasikan **Clean Architecture** yang diperkenalkan oleh Robert C. Martin (Uncle Bob). Arsitektur ini membantu memisahkan kode berdasarkan tanggung jawab dan menciptakan aplikasi yang mudah dipelihara serta diubah.

## Apa itu Clean Architecture?

[Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html) adalah sebuah pendekatan arsitektur perangkat lunak yang bertujuan untuk mengorganisir kode aplikasi dengan cara yang memisahkan berbagai lapisan tanggung jawab. Dengan memisahkan kode dalam lapisan yang terdefinisi dengan jelas, aplikasi lebih mudah dipelihara, diuji, dan diubah seiring waktu.

## Struktur Proyek

### 1. `cmd`
Folder ini digunakan untuk aplikasi atau skrip yang menjalankan proyek.

### 2. `internal`
Folder ini berisi kode yang bersifat **private** untuk aplikasi Anda dan diorganisir untuk mendukung Clean Architecture.

- **`internal/bootstrap`**
  - Folder ini digunakan untuk inisialisasi dan konfigurasi aplikasi.
  
- **`internal/delivery/http`**
  - Lapisan ini adalah **interface adapters** dalam Clean Architecture. Folder ini bertanggung jawab untuk menerima dan menangani request HTTP, memanggil use case, serta mengembalikan response ke klien.

- **`internal/domain/entities`**
  - Berisi model atau struktur data utama yang merepresentasikan domain aplikasi Anda. Entitas di sini bebas dari ketergantungan pada lapisan lain karena ini merupakan inti domain bisnis Anda.

- **`internal/repository`**
  - Folder ini adalah implementasi dari **data source layer**. Berisi kode untuk berinteraksi dengan penyimpanan data, seperti database. Biasanya diimplementasikan dengan interface agar bisa diuji dan diganti tanpa memengaruhi use case.

- **`internal/usecase`**
  - Lapisan ini adalah **application business rules** dalam Clean Architecture. Berisi logika aplikasi, seperti aturan dan proses bisnis utama. Use case memanggil repository untuk mengambil data atau mengubah data, tetapi tidak memiliki ketergantungan langsung pada teknologi tertentu.

### 3. `pkg`
Folder ini biasanya digunakan untuk kode yang bersifat **reusable** dan tidak spesifik terhadap domain aplikasi tertentu.

#### Fungsi Utama Folder `pkg`

Folder `pkg` berisi utilitas, helper, library, atau abstraksi yang dapat digunakan kembali di berbagai bagian aplikasi. File di sini dirancang secara generik agar tidak bergantung pada domain bisnis aplikasi.

#### Contoh Isi Folder `pkg`

- **Logger:** Abstraksi untuk logging yang dapat digunakan di seluruh aplikasi.
- **Config Loader:** Fungsi untuk membaca file konfigurasi, misalnya dari `.env` atau file YAML.
- **Middleware:** Middleware generik untuk aplikasi HTTP seperti logging, recovery, atau CORS.
- **Error Handling:** Library atau helper untuk menangani error dengan cara yang konsisten di seluruh aplikasi.
- **Utility:** Helper generik seperti fungsi untuk hashing, validasi input, atau formatting tanggal.
