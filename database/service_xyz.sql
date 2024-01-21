DROP TABLE IF EXISTS consumer_info;
-- Tabel untuk menyimpan informasi konsumen
CREATE TABLE consumer_info (
    id INT PRIMARY KEY AUTO_INCREMENT,
    nik VARCHAR(20) NOT NULL,
    email VARCHAR(255) NOT NULL,
    gender VARCHAR(10),
    full_name VARCHAR(100) NOT NULL,
    legal_name VARCHAR(100) NOT NULL,
    tempat_lahir VARCHAR(50),
    tanggal_lahir DATE,
    gaji INT,
    foto_ktp VARCHAR(255),
    foto_selfie VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

DROP TABLE IF EXISTS consumer_transaction;
-- Tabel untuk menyimpan transaksi konsumen
CREATE TABLE consumer_transaction (
    id INT PRIMARY KEY AUTO_INCREMENT,
    consumer_id INT,
    otr INT NOT NULL,
    admin_fee INT NOT NULL,
    jumlah_cicilan INT NOT NULL,
    jumlah_bunga FLOAT,
    nama_asset VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (consumer_id) REFERENCES consumer_info(id)
);

DROP TABLE IF EXISTS loan;
-- Tabel untuk menyimpan informasi pinjaman
CREATE TABLE loan (
    id INT PRIMARY KEY AUTO_INCREMENT,
    consumer_id INT,
    amount INT NOT NULL,
    total_loan INT NOT NULL,
    loan_interest FLOAT,
    monthly_payment INT,
    status VARCHAR(20),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (consumer_id) REFERENCES consumer_info(id)
);

DROP TABLE IF EXISTS transaction;
-- Tabel untuk menyimpan transaksi pembayaran pinjaman
CREATE TABLE transaction (
    id INT PRIMARY KEY AUTO_INCREMENT,
    loan_id INT,
    consumer_id INT,
    amount INT NOT NULL,
    payment_date DATE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (loan_id) REFERENCES loan(id),
    FOREIGN KEY (consumer_id) REFERENCES consumer_info(id)
);
