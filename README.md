# Habit Tracker and Consultation

## Deskripsi Project
### Tujuan Aplikasi
"Habit Tracker and Consultation" adalah sebuah aplikasi yang dirancang untuk membantu pengguna melacak kebiasaan mereka dan meningkatkan produktivitas sehari-hari. Aplikasi ini juga menyediakan layanan konsultasi dengan para ahli di berbagai bidang, yang dapat membantu pengguna dalam mencapai tujuan mereka, baik itu terkait dengan kesehatan mental, kesehatan fisik, atau bidang lainnya.

### Mengapa Produk Ini Dipilih untuk Dikembangkan
Aplikasi "Habit Tracker and Consultation" dipilih untuk dikembangkan karena beberapa alasan yang sangat penting:

- Meningkatkan Kualitas Hidup: Aplikasi ini dirancang untuk meningkatkan kualitas hidup pengguna dengan membantu mereka mengidentifikasi, melacak, dan memperbaiki kebiasaan-kebiasaan mereka.
- Menyediakan Akses ke Para Ahli: Dengan adanya layanan konsultasi, pengguna dapat dengan mudah mengakses para ahli di berbagai bidang, seperti manajemen stres, kesehatan mental, olahraga, dan gizi. Hal ini memungkinkan pengguna untuk mendapatkan bimbingan dan saran yang tepat sesuai dengan kebutuhan mereka.
- Personalisasi dan Detail yang Lebih Baik: Aplikasi ini memungkinkan pengguna untuk memantau kebiasaan mereka dengan sangat detail. Mereka dapat mencatat aktivitas sehari-hari, termasuk detail olahraga dan makanan, sehingga ahli dapat memberikan saran yang lebih personal dan efektif.
- Integrasi dengan Chatbot: Dengan adanya integrasi dengan chatbot, pengguna yang tidak mampu untuk menggunakan layanan konsultasi ahli masih dapat mendapatkan bantuan dan saran melalui interaksi dengan kecerdasan buatan.
- Konten Informasi yang Berkualitas: Aplikasi ini juga menyediakan artikel-artikel informatif yang dikelola oleh ahli, sehingga pengguna dapat memperluas pengetahuan mereka tentang topik-topik terkait kesehatan dan produktivitas.

Dengan kombinasi fitur-fitur ini, "Habit Tracker and Consultation" bertujuan untuk menjadi mitra sehari-hari yang dapat membantu pengguna mencapai potensi terbaik mereka dalam hidup. Dengan fokus pada perbaikan kebiasaan dan kesejahteraan secara keseluruhan, aplikasi ini diharapkan dapat membantu pengguna mencapai gaya hidup yang lebih sehat dan lebih produktif.

## Spesifikasi Fitur
### User
- User dapat melakukan registrasi dan login
- User dapat mengubah profilnya
- User dapat mencatat berbagai aktivitas sehari-hari dengan detail aktivitasnya
- User dapat mengubah aktivitas yang sudah tercatat
- User dapat menghapus aktivitas yang sudah tercatat
- User dapat melihat aktivitas yang sudah tercatat
- User dapat melihat ahli 
- User dapat menghire ahli di berbagai bidang dalam waktu 1 bulan
- User dapat melakukan pembayaran ke ahli yang dihire melalui transfer manual
- User dapat chat dengan chatbot

### Ahli
- Ahli dapat melakukan registrasi dan login
- Ahli dapat mengubah profilnya
- Ahli dapat melihat daftar customernya
- Ahli dapat melihat aktivitas customernya
- Ahli dapat memverifikasi pembayaran user yang ingin menghire dirinya


### Admin
- Admin dapat mengelola tipe bank account yang tersedia
- Admin dapat mengelola jenis expertise yang tersedia
- Admin dapat mengelola jenis aktivitas yang tersedia

## Tech Stack
- App Framework => Echo
- ORM => Gorm
- DB => MySQL
- Deployment => GCP Cloud Run
- Code Structure => Clean Architecture
- Authentication => JWT
- Other Tools => Integrasi OpenAI untuk Chat Bot dan Integrasi Cloudinary untuk menyimpan gambar