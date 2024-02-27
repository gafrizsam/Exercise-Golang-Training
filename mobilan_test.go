package main

import (
    "testing"
)

// Tes membuat mobil dengan roda dan pintu yang diharapkan
func TestMobilCreation(t *testing.T) {
    mobil := CreateMobil() // Menggunakan fungsi yang telah diperbaiki

    // Memperbaiki akses ke panjang array Roda
    if len(mobil.Roda) != 4 {
        t.Errorf("Jumlah roda yang diharapkan adalah 4, tetapi mendapatkan %d", len(mobil.Roda))
    }

    // Memperbaiki akses ke suara pintu menggunakan metode Ketuk dan Buka
    if mobil.PintuKanan.Ketuk() != "Knock" || mobil.PintuKiri.Buka() != "Knock" {
        t.Errorf("Suara pintu tidak sesuai dengan yang diharapkan")
    }
}

// Tes mengganti roda pada mobil
func TestGantiRoda(t *testing.T) {
    mobil := CreateMobil() // Inisialisasi mobil
    rodaBaru := BanKaret{} // Asumsikan ini adalah objek roda baru yang ingin dipasang

    mobil.GantiRoda(0, rodaBaru) // Memperbaiki nama metode sesuai dengan definisi

    // Memperbaiki pengecekan jenis roda yang diganti
    if mobil.Roda[0].Jenis() != rodaBaru.Jenis() {
        t.Errorf("Gagal mengganti roda, jenis roda yang diharapkan %s, tetapi mendapatkan %s", rodaBaru.Jenis(), mobil.Roda[0].Jenis())
    }
}

// Tes suara ketuk dan buka pint
