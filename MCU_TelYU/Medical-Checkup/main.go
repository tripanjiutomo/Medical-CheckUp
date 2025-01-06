package main

import (
    "fmt"
    "os"
    "os/exec"
    "strings"
)

type PaketMCU struct {
    IDPaket   int
    NamaPaket string
    Harga     float64
}

type PemeriksaanFisik struct {
    TinggiBadan   int
    BeratBadan    int
    TekananDarah  string
    DenyutNadi    int
    SuhuTubuh     float64
    Mata          string
    Hemoglobin    string
    Trombosit     string
    GulaDarah     string
    Kolestrol     string
    RontgenDada   string
    FungsiHatiSGOT string
    FungsiHatiSGPT string
    EKG           string
}

type Pasien struct {
    Nama          string
    JenisKelamin  string
    Umur          int
    Alamat        string
    PaketMCU      PaketMCU
    TanggalMasuk  string
    Fisik         PemeriksaanFisik
}

var (
    dataPasien    []Pasien
    daftarPaketMCU = []PaketMCU{
        {IDPaket: 1, NamaPaket: "Paket A", Harga: 250000},
        {IDPaket: 2, NamaPaket: "Paket B", Harga: 300000},
        {IDPaket: 3, NamaPaket: "Paket C", Harga: 350000},
    }
)

func clearScreen() {
    cmd := exec.Command("cmd", "/c", "cls")
    cmd.Stdout = os.Stdout
    cmd.Run()
}

func writeToFile(filename string, content string) {
    // Membuat direktori output jika belum ada
    if err := os.MkdirAll("output", os.ModePerm); err != nil {
        fmt.Println("Error creating directory:", err)
        return
    }

    file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    if _, err := file.WriteString(content + "\n"); err != nil {
        fmt.Println("Error writing to file:", err)
    }
}

func tambahPasien() {
    clearScreen()
    var pasien Pasien
    fmt.Println("Masukkan data pasien:")
    fmt.Print("Nama          : ")
    fmt.Scanln(&pasien.Nama)
    fmt.Print("Jenis Kelamin : ")
    fmt.Scanln(&pasien.JenisKelamin)
    fmt.Print("Umur          : ")
    fmt.Scanln(&pasien.Umur)
    fmt.Print("Alamat        : ")
    fmt.Scanln(&pasien.Alamat)
    fmt.Print("(DD-MM-YYYY)  : ")
    fmt.Scanln(&pasien.TanggalMasuk)

    // Pilih Paket MCU
    listPaketMCU()
    var pilihanPaket int
    fmt.Print("Pilih ID Paket MCU : ")
    fmt.Scanln(&pilihanPaket)

    // Cari paket MCU yang sesuai
    paketDitemukan := false
    for _, paket := range daftarPaketMCU {
        if paket.IDPaket == pilihanPaket {
            pasien.PaketMCU = paket
            paketDitemukan = true
            break
        }
    }

    if !paketDitemukan {
        fmt.Println("ID Paket MCU tidak valid.")
        return
    }

    dataPasien = append(dataPasien, pasien)
    fmt.Println("Data pasien berhasil ditambahkan.")
    writeToFile("output/data_pasien.txt", fmt.Sprintf("Nama: %s, Jenis Kelamin: %s, Umur: %d, Alamat: %s, Paket: %s", pasien.Nama, pasien.JenisKelamin, pasien.Umur, pasien.Alamat, pasien.PaketMCU.NamaPaket))

    fmt.Println("\nTekan Enter untuk melanjutkan...")
    fmt.Scanln()
    clearScreen()
}

func manageDataPasien() {
    clearScreen()
	if len(dataPasien) == 0 {
        fmt.Println("Tidak ada data pasien saat ini.")
        fmt.Println("\nTekan Enter untuk melanjutkan...")
        fmt.Scanln()
        return
    }

    fmt.Println("Data Pasien MCU:")
    var output strings.Builder
    for i, pasien := range dataPasien {
        output.WriteString(fmt.Sprintf("%d. Nama         : %s\n", i+1, pasien.Nama))
        output.WriteString(fmt.Sprintf("   Jenis Kelamin: %s\n", pasien.JenisKelamin))
        output.WriteString(fmt.Sprintf("   Umur         : %d\n", pasien.Umur))
        output.WriteString(fmt.Sprintf("   Alamat       : %s\n", pasien.Alamat))
        output.WriteString(fmt.Sprintf("   Paket MCU    : %s (Harga: %.2f)\n", pasien.PaketMCU.NamaPaket, pasien.PaketMCU.Harga))
        output.WriteString(fmt.Sprintf("   Tanggal Masuk: %s\n", pasien.TanggalMasuk))
        output.WriteString("  --------------------\n")
    }

    // Simpan ke file
    writeToFile("output/data_pasien.txt", output.String())

    fmt.Println("\nPilihan:")
    fmt.Println("1. Cari data pasien")
    fmt.Println("0. Kembali ke menu utama")
    fmt.Print("Pilih: ")
    var pilihan int
    fmt.Scanln(&pilihan)

    if pilihan == 0 {
        return
    } else if pilihan == 1 {
        clearScreen()
        var keyword string
        fmt.Print("Masukkan keyword pencarian (nama/usia/jenis kelamin/paket/tanggal): ")
        fmt.Scanln(&keyword)
        clearScreen()

        var hasil []Pasien
        for _, pasien := range dataPasien {
            if strings.Contains(strings.ToLower(pasien.Nama), strings.ToLower(keyword)) ||
                strings.Contains(strings.ToLower(pasien.JenisKelamin), strings.ToLower(keyword)) ||
                strings.Contains(strings.ToLower(fmt.Sprint(pasien.Umur)), strings.ToLower(keyword)) ||
                strings.Contains(strings.ToLower(pasien.PaketMCU.NamaPaket), strings.ToLower(keyword)) ||
                strings.Contains(strings.ToLower(pasien.TanggalMasuk), strings.ToLower(keyword)) ||
                strings.Contains(strings.ToLower(pasien.Alamat), strings.ToLower(keyword)) {
                hasil = append(hasil, pasien)
            }
        }

        if len(hasil) == 0 {
            fmt.Println("Data pasien tidak ditemukan.")
            fmt.Println("\nTekan Enter untuk melanjutkan...")
            fmt.Scanln()
            return
        }

        fmt.Println("Data pasien ditemukan:")
        for i, pasien := range hasil {
            fmt.Printf("%d. Nama          : %s\n", i+1, pasien.Nama)
            fmt.Printf("   Jenis Kelamin : %s\n", pasien.JenisKelamin)
            fmt.Printf("   Umur          : %d\n", pasien.Umur)
            fmt.Printf("   Alamat        : %s\n", pasien.Alamat)
            fmt.Printf("   Paket MCU     : %s (Harga: %.2f)\n", pasien.PaketMCU.NamaPaket, pasien.PaketMCU.Harga)
            fmt.Printf("   Tanggal Masuk : %s\n", pasien.TanggalMasuk)
            fmt.Println("  --------------------")
        }

        var pilihanPasien int
        fmt.Print("Pilih nomor pasien: ")
        fmt.Scanln(&pilihanPasien)

        if pilihanPasien > 0 && pilihanPasien <= len(hasil) {
            // Tampilkan submenu untuk pasien yang dipilih
            for {
                fmt.Println("\nPilihan:")
                fmt.Println("1. Perbarui data pasien")
                fmt.Println("2. Hapus data pasien")
                fmt.Println("0. Kembali ke menu pencarian")
                fmt.Print("Pilih: ")
                var subPilihan int
                fmt.Scanln(&subPilihan)

                switch subPilihan {
                case 1:
                    updatePasien(hasil[pilihanPasien-1])
                case 2:
                    deletePasien(pilihanPasien - 1)
                    return
                case 0:
                    return
                default:
                    fmt.Println("Pilihan tidak valid.")
                }
            }
        }
    } else {
        fmt.Println("Pilihan tidak valid.")
        fmt.Println("\nTekan Enter untuk melanjutkan...")
        fmt.Scanln()
    }
}

func updatePasien(pasienLama Pasien) {
    clearScreen()
    var pasienBaru Pasien
	fmt.Println("Masukkan data baru untuk pasien:")
    fmt.Print("Nama          : ")
    fmt.Scanln(&pasienBaru.Nama)
    fmt.Print("Jenis Kelamin : ")
    fmt.Scanln(&pasienBaru.JenisKelamin)
    fmt.Print("Umur          : ")
    fmt.Scanln(&pasienBaru.Umur)
    fmt.Print("Alamat        : ")
    fmt.Scanln(&pasienBaru.Alamat)
    fmt.Print("(DD-MM-YYYY)  : ")
    fmt.Scanln(&pasienBaru.TanggalMasuk)

    // Pilih Paket MCU
    listPaketMCU()
    var pilihanPaket int
    fmt.Print("Pilih ID Paket MCU: ")
    fmt.Scanln(&pilihanPaket)

    // Cari paket MCU yang sesuai
    paketDitemukan := false
    for _, paket := range daftarPaketMCU {
        if paket.IDPaket == pilihanPaket {
            pasienBaru.PaketMCU = paket
            paketDitemukan = true
            break
        }
    }

    if !paketDitemukan {
        fmt.Println("ID Paket MCU tidak valid.")
        return
    }

    // Perbarui data pasien di slice
    for i, pasien := range dataPasien {
        if pasien == pasienLama {
            dataPasien[i] = pasienBaru
            fmt.Println("Data pasien berhasil diperbarui.")
            writeToFile("output/data_pasien.txt", fmt.Sprintf("Data pasien diperbarui: %s", pasienBaru.Nama))
            return
        }
    }

    fmt.Println("Data pasien tidak ditemukan.")
}

func deletePasien(index int) {
    clearScreen()
    if index < 0 || index >= len(dataPasien) {
        fmt.Println("Nomor pasien tidak valid.")
        fmt.Println("\nTekan Enter untuk melanjutkan...")
        fmt.Scanln()
        return
    } 
    writeToFile("output/data_pasien.txt", fmt.Sprintf("Data pasien dihapus: %s", dataPasien[index].Nama))
    dataPasien = append(dataPasien[:index], dataPasien[index+1:]...)
    fmt.Println("Data pasien berhasil dihapus.")
    fmt.Println("\nTekan Enter untuk melanjutkan...")
    fmt.Scanln()
    clearScreen()
}

func masukkanMCUPasien() {
    clearScreen()
    var keyword string
    fmt.Print("Masukkan nama pasien untuk memasukkan data MCU: ")
    fmt.Scanln(&keyword)

    pasienDitemukan := false
    var pasien *Pasien // Pointer untuk menyimpan data pasien yang ditemukan
    for i := range dataPasien {
        if dataPasien[i].Nama == keyword {
            pasien = &dataPasien[i]
            pasienDitemukan = true
            break
        }
    }

    if !pasienDitemukan {
        fmt.Println("Data pasien tidak ditemukan.")
        fmt.Println("\nTekan Enter untuk melanjutkan...")
        fmt.Scanln()
        return
    }

    fmt.Println("\nMasukkan data MCU untuk pasien", pasien.Nama)

    switch pasien.PaketMCU.IDPaket {
    case 1:
        fmt.Print("Tinggi Badan (cm)    : ")
        fmt.Scanln(&pasien.Fisik.TinggiBadan)
        fmt.Print("Berat Badan (kg)     : ")
        fmt.Scanln(&pasien.Fisik.BeratBadan)
        fmt.Print("Tekanan Darah (mmHg) : ")
        fmt.Scanln(&pasien.Fisik.TekananDarah)
        fmt.Print("Denyut Nadi (bpm)    : ")
        fmt.Scanln(&pasien.Fisik.DenyutNadi)
        fmt.Print("Suhu Tubuh (°C)      : ")
        fmt.Scanln(&pasien.Fisik.SuhuTubuh)
        fmt.Print("Mata                 : ")
        fmt.Scanln(&pasien.Fisik.Mata)
        fmt.Print("Hemoglobin           : ")
        fmt.Scanln(&pasien.Fisik.Hemoglobin)
        fmt.Print("Trombosit            : ")
        fmt.Scanln(&pasien.Fisik.Trombosit)
    case 2:
        fmt.Print("Tinggi Badan (cm)    : ")
        fmt.Scanln(&pasien.Fisik.TinggiBadan)
        fmt.Print("Berat Badan (kg)     : ")
        fmt.Scanln(&pasien.Fisik.BeratBadan)
        fmt.Print("Tekanan Darah (mmHg) : ")
        fmt.Scanln(&pasien.Fisik.TekananDarah)
        fmt.Print("Denyut Nadi (bpm)    : ")
        fmt.Scanln(&pasien.Fisik.DenyutNadi)
        fmt.Print("Suhu Tubuh (°C)      : ")
        fmt.Scanln(&pasien.Fisik.SuhuTubuh)
        fmt.Print("Mata                 : ")
        fmt.Scanln(&pasien.Fisik.Mata)
        fmt.Print("Hemoglobin           : ")
        fmt.Scanln(&pasien.Fisik.Hemoglobin)
        fmt.Print("Trombosit            : ")
        fmt.Scanln(&pasien.Fisik.Trombosit)
        fmt.Print("Gula Darah           : ")
        fmt.Scanln(&pasien.Fisik.GulaDarah)
        fmt.Print("Kolestrol            : ")
        fmt.Scanln(&pasien.Fisik.Kolestrol)
        fmt.Print("Rontgen Dada         : ")
        fmt.Scanln(&pasien.Fisik.RontgenDada)
    case 3:
        fmt.Print("Tinggi Badan (cm)    : ")
        fmt.Scanln(&pasien.Fisik.TinggiBadan)
        fmt.Print("Berat Badan (kg)     : ")
        fmt.Scanln(&pasien.Fisik.BeratBadan)
        fmt.Print("Tekanan Darah (mmHg) : ")
        fmt.Scanln(&pasien.Fisik.TekananDarah)
        fmt.Print("Denyut Nadi (bpm)    : ")
        fmt.Scanln(&pasien.Fisik.DenyutNadi)
        fmt.Print("Suhu Tubuh (°C)      : ")
        fmt.Scanln(&pasien.Fisik.SuhuTubuh)
        fmt.Print("Mata                 : ")
        fmt.Scanln(&pasien.Fisik.Mata)
        fmt.Print("Hemoglobin           : ")
        fmt.Scanln(&pasien.Fisik.Hemoglobin)
        fmt.Print("Trombosit            : ")
        fmt.Scanln(&pasien.Fisik.Trombosit)
        fmt.Print("Gula Darah           : ")
        fmt.Scanln(&pasien.Fisik.GulaDarah)
        fmt.Print("Kolestrol            : ")
        fmt.Scanln(&pasien.Fisik.Kolestrol)
        fmt.Print("Rontgen Dada         : ")
        fmt.Scanln(&pasien.Fisik.RontgenDada)
        fmt.Print("Fungsi Hati SGOT     : ")
        fmt.Scanln(&pasien.Fisik.FungsiHatiSGOT)
        fmt.Print("Fungsi Hati SGPT     : ")
        fmt.Scanln(&pasien.Fisik.FungsiHatiSGPT)
        fmt.Print("EKG                  : ")
        fmt.Scanln(&pasien.Fisik.EKG)
    default:
        fmt.Println("Paket MCU tidak valid.")
        fmt.Println("\nTekan Enter untuk melanjutkan...")
        fmt.Scanln()
        return
    }

    fmt.Println("Data MCU pasien berhasil ditambahkan.")
    writeToFile("output/hasil_mcu.txt", fmt.Sprintf("Data MCU untuk pasien %s berhasil ditambahkan.", pasien.Nama))
    fmt.Println("\nTekan Enter untuk melanjutkan...")
    fmt.Scanln()
}

func tampilkanHasilMCUPasien() {
    clearScreen()
    fmt.Println("\nData Hasil MCU Pasien:")
    if len(dataPasien) == 0 {
        fmt.Println("Tidak ada data pasien.")
        fmt.Println("\nTekan Enter untuk melanjutkan...")
        fmt.Scanln()
        return
    }

    var output strings.Builder
    for i, pasien := range dataPasien {
        output.WriteString(fmt.Sprintf("%d. Nama         : %s\n", i+1, pasien.Nama))
        switch pasien.PaketMCU.IDPaket {
        case 1:
            output.WriteString(fmt.Sprintf("   Tinggi Badan : %d cm\n", pasien.Fisik.TinggiBadan))
            output.WriteString(fmt.Sprintf("   Berat Badan  : %d kg\n", pasien.Fisik.BeratBadan))
            output.WriteString(fmt.Sprintf("   Tekanan Darah: %s mmHg\n", pasien.Fisik.TekananDarah))
            output.WriteString(fmt.Sprintf("   Denyut Nadi  : %d bpm\n", pasien.Fisik.DenyutNadi))
            output.WriteString(fmt.Sprintf("   Suhu Tubuh   : %.2f °C\n", pasien.Fisik.SuhuTubuh))
            output.WriteString(fmt.Sprintf("   Mata         : %s\n", pasien.Fisik.Mata))
            output.WriteString(fmt.Sprintf("   Hemoglobin   : %s\n", pasien.Fisik.Hemoglobin))
			output.WriteString(fmt.Sprintf("   Trombosit    : %s\n", pasien.Fisik.Trombosit))
        case 2:
            output.WriteString(fmt.Sprintf("   Tinggi Badan : %d cm\n", pasien.Fisik.TinggiBadan))
            output.WriteString(fmt.Sprintf("   Berat Badan  : %d kg\n", pasien.Fisik.BeratBadan))
            output.WriteString(fmt.Sprintf("   Tekanan Darah: %s mmHg\n", pasien.Fisik.TekananDarah))
            output.WriteString(fmt.Sprintf("   Denyut Nadi  : %d bpm\n", pasien.Fisik.DenyutNadi))
            output.WriteString(fmt.Sprintf("   Suhu Tubuh   : %.2f °C\n", pasien.Fisik.SuhuTubuh))
            output.WriteString(fmt.Sprintf("   Mata         : %s\n", pasien.Fisik.Mata))
            output.WriteString(fmt.Sprintf("   Hemoglobin   : %s\n", pasien.Fisik.Hemoglobin))
            output.WriteString(fmt.Sprintf("   Trombosit    : %s\n", pasien.Fisik.Trombosit))
            output.WriteString(fmt.Sprintf("   Gula Darah   : %s\n", pasien.Fisik.GulaDarah))
            output.WriteString(fmt.Sprintf("   Kolestrol    : %s\n", pasien.Fisik.Kolestrol))
            output.WriteString(fmt.Sprintf("   Rontgen Dada : %s\n", pasien.Fisik.RontgenDada))
        case 3:
            output.WriteString(fmt.Sprintf("   Tinggi Badan : %d cm\n", pasien.Fisik.TinggiBadan))
            output.WriteString(fmt.Sprintf("   Berat Badan  : %d kg\n", pasien.Fisik.BeratBadan))
            output.WriteString(fmt.Sprintf("   Tekanan Darah: %s mmHg\n", pasien.Fisik.TekananDarah))
            output.WriteString(fmt.Sprintf("   Denyut Nadi  : %d bpm\n", pasien.Fisik.DenyutNadi))
            output.WriteString(fmt.Sprintf("   Suhu Tubuh   : %.2f °C\n", pasien.Fisik.SuhuTubuh))
            output.WriteString(fmt.Sprintf("   Mata         : %s\n", pasien.Fisik.Mata))
            output.WriteString(fmt.Sprintf("   Hemoglobin   : %s\n", pasien.Fisik.Hemoglobin))
            output.WriteString(fmt.Sprintf("   Trombosit    : %s\n", pasien.Fisik.Trombosit))
            output.WriteString(fmt.Sprintf("   Gula Darah   : %s\n", pasien.Fisik.GulaDarah))
            output.WriteString(fmt.Sprintf("   Kolestrol    : %s\n", pasien.Fisik.Kolestrol))
            output.WriteString(fmt.Sprintf("   Rontgen Dada : %s\n", pasien.Fisik.RontgenDada))
            output.WriteString(fmt.Sprintf("   F-Hati SGOT  : %s\n", pasien.Fisik.FungsiHatiSGOT))
            output.WriteString(fmt.Sprintf("   F-Hati SGPT  : %s\n", pasien.Fisik.FungsiHatiSGPT))
            output.WriteString(fmt.Sprintf("   EKG          : %s\n", pasien.Fisik.EKG))
        }
        output.WriteString("   --------------------\n")
    }

    // Simpan hasil ke file
    writeToFile("output/hasil_mcu.txt", output.String())

    var keyword string
    fmt.Print("\nMasukkan keyword untuk mencari hasil MCU pasien: ")
    fmt.Scanln(&keyword)

    var hasilPencarian []Pasien
    for _, pasien := range dataPasien {
        if strings.Contains(strings.ToLower(pasien.Nama), strings.ToLower(keyword)) ||
            strings.Contains(strings.ToLower(pasien.JenisKelamin), strings.ToLower(keyword)) ||
            strings.Contains(strings.ToLower(fmt.Sprint(pasien.Umur)), strings.ToLower(keyword)) ||
            strings.Contains(strings.ToLower(pasien.Alamat), strings.ToLower(keyword)) ||
            strings.Contains(strings.ToLower(pasien.TanggalMasuk), strings.ToLower(keyword)) {
            hasilPencarian = append(hasilPencarian, pasien)
        }
    }

    if len(hasilPencarian) == 0 {
        fmt.Println("Data pasien tidak ditemukan.")
        fmt.Println("\nTekan Enter untuk melanjutkan...")
        fmt.Scanln()
        return
    }

    fmt.Println("\nHasil pencarian:")
    for i, pasien := range hasilPencarian {
        if pasien.PaketMCU.IDPaket == 1 {
			fmt.Printf("%d. Nama        : %s\n", i+1, pasien.Nama)
            fmt.Printf("   Tinggi Badan : %d cm\n", pasien.Fisik.TinggiBadan)
            fmt.Printf("   Berat Badan  : %d kg\n", pasien.Fisik.BeratBadan)
            fmt.Printf("   Tekanan Darah: %s mmHg\n", pasien.Fisik.TekananDarah)
            fmt.Printf("   Denyut Nadi  : %d bpm\n", pasien.Fisik.DenyutNadi)
            fmt.Printf("   Suhu Tubuh   : %.2f °C\n", pasien.Fisik.SuhuTubuh)
            fmt.Printf("   Mata         : %s\n", pasien.Fisik.Mata)
            fmt.Printf("   Hemoglobin   : %s\n", pasien.Fisik.Hemoglobin)
            fmt.Printf("   Trombosit    : %s\n", pasien.Fisik.Trombosit)
        } else if pasien.PaketMCU.IDPaket == 2 {
            fmt.Printf("%d. Nama        : %s\n", i+1, pasien.Nama)
            fmt.Printf("   Tinggi Badan : %d cm\n", pasien.Fisik.TinggiBadan)
            fmt.Printf("   Berat Badan  : %d kg\n", pasien.Fisik.BeratBadan)
            fmt.Printf("   Tekanan Darah: %s mmHg\n", pasien.Fisik.TekananDarah)
            fmt.Printf("   Denyut Nadi  : %d bpm\n", pasien.Fisik.DenyutNadi)
            fmt.Printf("   Suhu Tubuh   : %.2f °C\n", pasien.Fisik.SuhuTubuh)
            fmt.Printf("   Mata         : %s\n", pasien.Fisik.Mata)
            fmt.Printf("   Hemoglobin   : %s\n", pasien.Fisik.Hemoglobin)
            fmt.Printf("   Trombosit    : %s\n", pasien.Fisik.Trombosit)
            fmt.Printf("   Gula Darah   : %s\n", pasien.Fisik.GulaDarah)
            fmt.Printf("   Kolestrol    : %s\n", pasien.Fisik.Kolestrol)
            fmt.Printf("   Rontgen Dada : %s\n", pasien.Fisik.RontgenDada)
        } else if pasien.PaketMCU.IDPaket == 3 {
            fmt.Printf("%d. Nama        : %s\n", i+1, pasien.Nama)
            fmt.Printf("   Tinggi Badan : %d cm\n", pasien.Fisik.TinggiBadan)
            fmt.Printf("   Berat Badan  : %d kg\n", pasien.Fisik.BeratBadan)
            fmt.Printf("   Tekanan Darah: %s mmHg\n", pasien.Fisik.TekananDarah)
            fmt.Printf("   Denyut Nadi  : %d bpm\n", pasien.Fisik.DenyutNadi)
            fmt.Printf("   Suhu Tubuh   : %.2f °C\n", pasien.Fisik.SuhuTubuh)
            fmt.Printf("   Mata         : %s\n", pasien.Fisik.Mata)
            fmt.Printf("   Hemoglobin   : %s\n", pasien.Fisik.Hemoglobin)
            fmt.Printf("   Trombosit    : %s\n", pasien.Fisik.Trombosit)
            fmt.Printf("   Gula Darah   : %s\n", pasien.Fisik.GulaDarah)
            fmt.Printf("   Kolestrol    : %s\n", pasien.Fisik.Kolestrol)
            fmt.Printf("   Rontgen Dada : %s\n", pasien.Fisik.RontgenDada)
            fmt.Printf("   F-Hati SGOT  : %s\n", pasien.Fisik.FungsiHatiSGOT)
            fmt.Printf("   F-Hati SGPT  : %s\n", pasien.Fisik.FungsiHatiSGPT)
            fmt.Printf("   EKG          : %s\n", pasien.Fisik.EKG)
        }
        fmt.Println("   --------------------")
    }

    // Simpan hasil pencarian ke file
    writeToFile("output/hasil_mcu.txt", output.String())
}

func laporanPemasukan() {
    clearScreen()
    totalPemasukan := 0.0
    for _, pasien := range dataPasien {
        totalPemasukan += pasien.PaketMCU.Harga
    }
    laporan := fmt.Sprintf("Total Pemasukan: Rp %.2f\n", totalPemasukan)
    fmt.Println(laporan)
	writeToFile("output/laporan_pemasukan.txt", laporan)
    fmt.Println("\nTekan Enter untuk kembali ke menu...")
    fmt.Scanln()
}

func listPaketMCU() {
    clearScreen()
    fmt.Println("\nDaftar Paket MCU:")
    var output strings.Builder
    for _, paket := range daftarPaketMCU {
        output.WriteString(fmt.Sprintf("%d. %s (Harga: %.2f)\n", paket.IDPaket, paket.NamaPaket, paket.Harga))
        switch paket.IDPaket {
        case 1:
            output.WriteString("   - Tinggi Badan\n")
            output.WriteString("   - Berat Badan\n")
            output.WriteString("   - Tekanan Darah\n")
            output.WriteString("   - Denyut Nadi\n")
            output.WriteString("   - Suhu Tubuh\n")
            output.WriteString("   - Mata\n")
            output.WriteString("   - Hemoglobin\n")
            output.WriteString("   - Trombosit\n")
        case 2:
            output.WriteString("   - Tinggi Badan\n")
            output.WriteString("   - Berat Badan\n")
            output.WriteString("   - Tekanan Darah\n")
            output.WriteString("   - Denyut Nadi\n")
            output.WriteString("   - Suhu Tubuh\n")
            output.WriteString("   - Mata\n")
            output.WriteString("   - Hemoglobin\n")
            output.WriteString("   - Trombosit\n")
            output.WriteString("   - Gula Darah\n")
            output.WriteString("   - Kolestrol\n")
            output.WriteString("   - Rontgen Dada\n")
        case 3:
            output.WriteString("   - Tinggi Badan\n")
            output.WriteString("   - Berat Badan\n")
            output.WriteString("   - Tekanan Darah\n")
            output.WriteString("   - Denyut Nadi\n")
            output.WriteString("   - Suhu Tubuh\n")
            output.WriteString("   - Mata\n")
            output.WriteString("   - Hemoglobin\n")
            output.WriteString("   - Trombosit\n")
            output.WriteString("   - Gula Darah\n")
            output.WriteString("   - Kolestrol\n")
            output.WriteString("   - Rontgen Dada\n")
            output.WriteString("   - Fungsi Hati SGOT\n")
            output.WriteString("   - Fungsi Hati SGPT\n")
            output.WriteString("   - EKG\n")
        }
    }
    // Simpan daftar paket ke file
    writeToFile("output/paket_mcu.txt", output.String())
    fmt.Println("\nTekan Enter untuk melanjutkan...")
    fmt.Scanln()
    clearScreen()
}

func menu() {
    for {
        clearScreen()
        fmt.Println("\n=== Aplikasi Medical Check-Up ===")
        fmt.Println("1. Paket MCU")
        fmt.Println("2. Tambah Data Pasien MCU")
        fmt.Println("3. Data Pasien MCU") // Menggunakan manageDataPasien
        fmt.Println("4. Masukkan MCU Pasien") 
        fmt.Println("5. Hasil MCU Pasien")    
        fmt.Println("6. Laporan Pemasukan")  
        fmt.Println("0. Keluar")
        fmt.Print("Pilih menu: ")

        var pilihan int
        fmt.Scanln(&pilihan)

        switch pilihan {
        case 1:
            listPaketMCU()
        case 2:
            tambahPasien()
        case 3:
            manageDataPasien() // Gabungan ListDataPasien dan cariPasien
        case 4:
            masukkanMCUPasien() // Fungsi baru untuk memasukkan hasil MCU
        case 5:
            tampilkanHasilMCUPasien() // Fungsi baru untuk menampilkan hasil MCU
        case 6:
            laporanPemasukan()
        case 0:
            fmt.Println("Terima kasih! Sampai jumpa.")
            return
        default:
            fmt.Println("Pilihan tidak valid, silakan coba lagi.")
        }
    }
}

func main() {
    menu()
}