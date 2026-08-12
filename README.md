# Komputasi Paralel: Benchmark Sorting Algorithm di Golang

Proyek ini membandingkan kinerja algoritma **Bubble Sort** dan **Selection Sort** dalam dua pendekatan: **sequential** dan **parallel** (menggunakan Goroutine + WaitGroup) pada bahasa Go. Pengujian dilakukan pada berbagai ukuran data (10.000–500.000 elemen) untuk menganalisis dampak paralelisasi terhadap waktu eksekusi, dilengkapi REST API dan antarmuka web sederhana untuk menjalankan benchmark secara interaktif.

📄 Laporan lengkap (metode pengujian, hasil, analisis, dan penjelasan kode) tersedia di file [`Laporan Dan Panduan Komputasi Paralel - Heru Khoerudin 231351059.pdf`](./Laporan%20Tugas%20Besar%20Heru%20Khoerudin/Laporan%20Dan%20Panduan%20Komputasi%20Paralel%20-%20Heru%20Khoerudin%20231351059.pdf).

## Tujuan

Mengetahui sejauh mana pemanfaatan CPU/logical processor melalui Goroutine dapat meningkatkan efisiensi waktu eksekusi algoritma sorting dibanding versi sequential.

## Tech Stack

- **Bahasa:** Go (Golang)
- **Concurrency:** Goroutine, sync.WaitGroup
- **Interface:** REST API + HTML sederhana untuk menjalankan dan membandingkan benchmark

## Metode Pengujian

- Algoritma: Bubble Sort dan Selection Sort (versi sequential & parallel)
- Jumlah worker: menyesuaikan `runtime.NumCPU()`
- Ukuran data uji: 10.000 – 500.000 elemen (integer acak)
- Metrik: waktu eksekusi (ms), diukur dengan `time.Now()` dan `time.Since()`

## Ringkasan Hasil

| Jumlah Elemen | Bubble Sort — Parallel (ms) | Bubble Sort — Sequential (ms) |
|---|---|---|
| 100.000 | 16.847 | 29.253 |
| 300.000 | 114.199 | 236.322 |
| 500.000 | 251.139 | 655.735 |

| Jumlah Elemen | Selection Sort — Parallel (ms) | Selection Sort — Sequential (ms) |
|---|---|---|
| 100.000 | 8.406 | 14.000 |
| 300.000 | 52.718 | 78.340 |
| 500.000 | 126.531 | 223.945 |

**Kesimpulan singkat:**
- Bubble Sort paralel menunjukkan peningkatan performa yang konsisten dan signifikan dibanding sequential.
- Selection Sort tetap berkompleksitas O(n²) meski dipararelkan, karena hanya bagian pencarian minimum yang dibagi ke banyak worker — sehingga peningkatannya tidak sebesar Bubble Sort.
- Untuk data kecil (<10.000 elemen), versi sequential justru lebih cepat karena overhead pembuatan goroutine.

Detail analisis lengkap ada di laporan PDF.

## Cara Menjalankan

1. Install [Go](https://go.dev/dl/) dan [VSCode](https://code.visualstudio.com/)
2. Clone repo ini, buka foldernya di VSCode
3. Install ekstensi **Live Server** (by Ritwick Dey) di VSCode
4. Jalankan API:
```bash
   go run .
```
   API akan aktif di `http://localhost:8080`
5. Buka `bubble.html` atau `selection.html`, klik kanan → **Open with Live Server**
6. Pilih jumlah data dan algoritma yang ingin dibandingkan, klik **Run Comparison**

## Struktur Kode Singkat

- `BubbleSortSequential` / `BubbleSortParallel` — implementasi Bubble Sort (parallel berbasis odd-even transposition sort)
- `SelectionSort` / `SelectionSortParallel` — implementasi Selection Sort (parallel pada pencarian nilai minimum via `findMinParallel`)
- `GenerateRandomArray` — generator data uji acak
- `RunBubbleSort` / `RunSelectionSort` — fungsi benchmarking utama, mengembalikan data before/after dan waktu eksekusi

## Penulis

**Heru Khoerudin**
Program Studi Teknik Informatika, Sekolah Tinggi Teknologi Wastukancana Purwakarta
