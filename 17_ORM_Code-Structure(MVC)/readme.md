## Summary ORM_Code-Structure(MVC)

- Object-Relational Mapping, adalah teknik dalam pengembangan perangkat lunak yang memungkinkan pengembang untuk mengakses dan memanipulasi data dalam basis data menggunakan objek dan kelas di dalam bahasa pemrograman, tanpa perlu menulis kueri SQL secara langsung. Dengan kata lain, ORM memungkinkan representasi objek dalam program untuk dihubungkan ke struktur data dalam basis data relasional.
- MVC adalah singkatan dari Model-View-Controller, sebuah desain arsitektur perangkat lunak yang digunakan untuk memisahkan komponen-komponen utama dalam sebuah aplikasi. Dengan memisahkan logika aplikasi ke dalam tiga komponen terpisah, MVC membantu mengorganisir kode dengan lebih baik, membuatnya lebih mudah dipahami, dielola, dan diperbarui. komponen dalam arsitektur MVC antara lain:
   - Model :
      - Model mewakili data dan aturan bisnis aplikasi. Ini mencakup logika untuk mengelola dan memanipulasi data, serta mengirimkan pembaruan ke View atau Controller saat data berubah.
      - Model bertanggung jawab untuk memperbarui tampilan (View) ketika data yang terkait dengan tampilan tersebut berubah.
      - Dalam konteks basis data, model sering kali berhubungan dengan operasi-operasi basis data seperti pengambilan, penyimpanan, dan pembaruan data.
  - View :
      - View adalah bagian dari aplikasi yang bertanggung jawab untuk menampilkan data kepada pengguna. Ini dapat berupa antarmuka pengguna grafis (seperti halaman web, formulir, atau aplikasi mobile), atau output dalam bentuk lain seperti file XML atau JSON.
      - View tidak memiliki logika bisnis. Tugasnya hanyalah menampilkan data yang diberikan kepadanya oleh Controller atau Model.
  - Controller:
      - Controller bertindak sebagai perantara antara Model dan View. Saat pengguna berinteraksi dengan View, Controller menangkap tindakan tersebut dan memprosesnya
      - Controller mengambil input dari pengguna (melalui View) dan mengonversinya menjadi perintah untuk Model atau View.
      - Controller menerima pembaruan dari Model (saat data berubah) dan memperbarui View untuk mencerminkan perubahan tersebut.
