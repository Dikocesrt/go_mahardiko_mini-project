# Habit Master

## Deskripsi Project
### Tujuan Aplikasi
Habit Master adalah aplikasi inovatif yang dirancang untuk membantu pengguna melacak dan mengelola berbagai kebiasaan hidup mereka secara efektif. Aplikasi ini menyediakan fitur pelacakan yang mudah digunakan, memungkinkan pengguna untuk mencatat aktivitas seperti olahraga, kerja, tidur, dan makan dengan detail. Salah satu fitur unggulan dari Habit Master adalah kemampuannya untuk menghubungkan pengguna dengan para ahli atau expert dalam bidang kesehatan dan kebiasaan hidup. Dengan membayar biaya bulanan, pengguna dapat mengakses layanan ahli yang dapat memberikan dukungan dan bimbingan secara langsung. Selain itu, Habit Master juga menyediakan fitur chatbot sebagai alternatif bagi pengguna yang tidak mampu membayar layanan ahli. Fitur chatbot ini memberikan bantuan dan saran bagi pengguna dalam perjalanan mereka menuju gaya hidup yang lebih sehat dan teratur.

### Mengapa Produk Ini Dipilih untuk Dikembangkan
Aplikasi "Habit Master" dipilih untuk dikembangkan karena beberapa alasan yang sangat penting:

- Meningkatkan Kualitas Hidup: Aplikasi ini dirancang untuk meningkatkan kualitas hidup pengguna dengan membantu mereka mengidentifikasi, melacak, dan memperbaiki kebiasaan-kebiasaan mereka.
- Menyediakan Akses ke Para Ahli: Dengan adanya layanan konsultasi, pengguna dapat dengan mudah mengakses para ahli di berbagai bidang, seperti manajemen stres, kesehatan mental, olahraga, dan gizi. Hal ini memungkinkan pengguna untuk mendapatkan bimbingan dan saran yang tepat sesuai dengan kebutuhan mereka.
- Personalisasi dan Detail yang Lebih Baik: Aplikasi ini memungkinkan pengguna untuk memantau kebiasaan mereka dengan sangat detail. Mereka dapat mencatat aktivitas sehari-hari, termasuk detail olahraga dan makanan, sehingga ahli dapat memberikan saran yang lebih personal dan efektif.
- Integrasi dengan Chatbot: Dengan adanya integrasi dengan chatbot, pengguna yang tidak mampu untuk menggunakan layanan konsultasi ahli masih dapat mendapatkan bantuan dan saran melalui interaksi dengan kecerdasan buatan.
- Konten Informasi yang Berkualitas: Aplikasi ini juga menyediakan artikel-artikel informatif yang dikelola oleh ahli, sehingga pengguna dapat memperluas pengetahuan mereka tentang topik-topik terkait kesehatan dan produktivitas.

Dengan kombinasi fitur-fitur ini, "Habit Master" bertujuan untuk menjadi mitra sehari-hari yang dapat membantu pengguna mencapai potensi terbaik mereka dalam hidup. Dengan fokus pada perbaikan kebiasaan dan kesejahteraan secara keseluruhan, aplikasi ini diharapkan dapat membantu pengguna mencapai gaya hidup yang lebih sehat dan lebih produktif.

## Spesifikasi Fitur
### User
- Registrasi
- Login
- Update Profile
- Manage Aktivitasnya
- Hire Expert
- Chatbot

### Ahli
- Registrasi
- Login
- Update Profile
- Melihat Aktivitas Customer
- Verifikasi pembayaran


### Admin
- Registrasi
- Login
- Manage Bank Account Type
- Manage Expertise
- Manage Activity Type

## Tech Stack
- App Framework -> ECHO
- ORM -> GORM
- DB -> MySQL
- Deployment -> AWS
- Code Structure -> Clean Architecture
- Authentication -> JWT
- Image Storage -> Cloudinary
- CI/CD -> Github Actions
- Containerization -> Docker
- Version Control -> Git & Github
- Other Tools -> Open AI

## ERD
![ERD](https://res.cloudinary.com/dy2fwknbn/image/upload/v1715777838/cctiodjorf1xljnav5mn.jpg)

## Getting Started
1. Clone repository ini
2. Navigasi ke dalam folder project
3. Jalankan perintah go mod tidy
4. Copy file .env.example dan rename menjadi .env
5. Sesuaikan konfigurasi yang ada di file .env
6. Jalankan perintah go run main.go