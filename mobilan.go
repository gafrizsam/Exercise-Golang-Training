package main

import "fmt"

type Roda interface {
    Jenis() string
}

type BanKaret struct{}
func (b BanKaret) Jenis() string { return "karet" }

type BanKayu struct{}
func (b BanKayu) Jenis() string { return "kayu" }

type BanBesi struct{}
func (b BanBesi) Jenis() string { return "besi" }

type Pintu struct {
    SuaraKetuk string
    SuaraBuka  string
}

func (p Pintu) Ketuk() string {
    return p.SuaraKetuk
}

func (p Pintu) Buka() string {
    return p.SuaraBuka
}

type Mobil struct {
    Roda       [4]Roda
    PintuKanan Pintu
    PintuKiri  Pintu
}

func (m *Mobil) GantiRoda(posisi int, roda Roda) {
    if posisi < 0 || posisi >= len(m.Roda) {
        fmt.Println("Posisi tidak valid. Posisi harus antara 0 dan 3.")
        return
    }
    m.Roda[posisi] = roda
    fmt.Printf("Roda %s telah dipasang di posisi %d.\n", roda.Jenis(), posisi)
}

func main() {
    mobil := Mobil{
        PintuKanan: Pintu{SuaraKetuk: "Knock", SuaraBuka: "beep"},
        PintuKiri:  Pintu{SuaraKetuk: "beep", SuaraBuka: "Knock"},
    }

    // Ganti roda
    mobil.GantiRoda(0, BanKaret{})
    mobil.GantiRoda(1, BanKayu{})
    mobil.GantiRoda(2, BanBesi{})
    mobil.GantiRoda(3, BanKaret{})

    // Tes pintu
    fmt.Println("Pintu kanan:", mobil.PintuKanan.Ketuk(), "-", mobil.PintuKanan.Buka())
    fmt.Println("Pintu kiri:", mobil.PintuKiri.Ketuk(), "-", mobil.PintuKiri.Buka())
}

// Perubahan di sini: Mengubah nama fungsi untuk mengekspornya
func CreateMobil() Mobil {
    // Inisialisasi roda sebagai array untuk konsistensi
    var roda [4]Roda = [4]Roda{BanKaret{}, BanKaret{}, BanKayu{}, BanBesi{}}
    pintuKanan := Pintu{SuaraKetuk: "Knock", SuaraBuka: "beep"}
    pintuKiri := Pintu{SuaraKetuk: "beep", SuaraBuka: "Knock"}

    mobil := Mobil{
        Roda:       roda,
        PintuKanan: pintuKanan,
        PintuKiri:  pintuKiri,
    }
    return mobil
}
