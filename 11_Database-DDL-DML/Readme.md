## Summary Database DDL DML

- Database adalah kumpulan data yang terorganisir secara sistematis sehingga dapat dengan mudah diakses, dikelola, dan diperbarui.
- Tipe-Tipe Relasi yang terdapat pada database antara lain:
    - One-to-One (1:1): Hubungan satu-satu mengindikasikan bahwa satu baris di tabel A hanya dapat berhubungan dengan satu baris di tabel B, dan sebaliknya.
    - One-to-Many (1:N): Hubungan satu-ke-banyak mengindikasikan bahwa satu baris di tabel A dapat berhubungan dengan banyak baris di tabel B, tetapi satu baris di tabel B hanya dapat berhubungan dengan satu baris di tabel A.
    - Many-to-One (N:1): Hubungan banyak-ke-satu adalah kebalikan dari hubungan satu-ke-banyak. Artinya, banyak baris di tabel A dapat berhubungan dengan satu baris di tabel B.
    - Many-to-Many (N:N): Hubungan banyak-ke-banyak mengindikasikan bahwa banyak baris di tabel A dapat berhubungan dengan banyak baris di tabel B.
- Jenis - Jenis perintah yang ada pada Database SQL berserta dengan lingkupnya:
    - DDL (Database Definiton Language) : CREATE DATABASE, USE (Database name), CREATE TABEL, DROP TABEL, RENAME TABEL, ALTER TABEL, ADD CLOUMN
    - DML (Database Manipulation Language) : INSERT, SELECT, UPDATE, DELETE
    - DCL (Database Control Language) : LIKE/BETWEEN, AND/OR, ORDER BY, LIMIT, JOIN
- jenis - jenis JOIN antara lain :
    - INNER JOIN akan mengembalikan baris-baris dari dua tabel atau lebih yang memenuhi syarat
    - LEFT JOIN akan mengembalikan seluruh baris dari tabel disebelah kiri yang dikenai kondisi ON dan hanya baris dari tabel disebelah kanan yang memenuhi kondisi join
    - RIGHT JOIN akan mengembalikan semua baris dari tabel sebelah kanan yang dikenai kondisi ON dengan data dari tabel sebelah kiri yang memenuhi kondisi join atau kebalikan dari LEFT JOIN
