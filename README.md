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
- User dapat mencatat berbagai aktivitas sehari-hari dengan detail aktivitasnya
- User dapat mengubah aktivitas yang sudah tercatat
- User dapat menghapus aktivitas yang sudah tercatat
- User dapat melihat aktivitas yang sudah tercatat
- User dapat mencari dan melihat detail ahli 
- User dapat menghire ahli di berbagai bidang dalam suatu kurun waktu
- User dapat melakukan pembayaran ke ahli yang dihire melalui transfer manual
- User dapat melakukan konsultasi melalui chat kepada ahli yang telah dihire
- User dapat memberikan ulasan kepada ahli yang telah dihire
- User dapat melihat riwayat hire ahli
- User dapat chat dengan chatbot terkait aktivitas dan goalsnya
- User dapat melihat artikel artikel informatif
- User dapat melakukan report terkait ahli dan artikel dengan alasan yang jelas

### Ahli
- Ahli dapat memverifikasi pembayaran user yang ingin menghire dirinya
- Ahli dapat melihat informasi aktivitas customernya
- Ahli dapat melakukan chat dengan customer
- Ahli dapat melihat ulasan terkait dirinya
- Ahli dapat membuat artikel informatif untuk mempromosikan dirinya
- Ahli dapat melihat artikel yang dia buat
- Ahli dapat mengubah artikel yang dia buat
- Ahli dapat menghapus artikel yang dia buat
- Ahli dapat melakukan report terkait user dan artikel dengan alasan yang jelas

### Admin
- Admin dapat mengelola akun dan artikel
- Admin dapat melihat hasil report dari user dan ahli
- Admin dapat melakukan takedown akun dan menghapus artikel apabila report yang dilakukan benar

## Tech Stack
- App Framework => Echo
- ORM => Gorm
- DB => MySQL
- Deployment => GCP Cloud Run
- Code Structure => Clean Architecture
- Authentication => JWT
- Other Tools => Integrasi Open AI untuk Chat Bot