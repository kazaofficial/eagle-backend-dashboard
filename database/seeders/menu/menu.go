package main

import (
	"eagle-backend-dashboard/config"
	"eagle-backend-dashboard/entity"
	"log"
)

func main() {
	// Membuat koneksi dari config/database.go
	dbConfig, err := config.NewDatabaseConfig()
	if err != nil {
		// log with comment
		log.Fatalf(err.Error())
	}

	// Connect ke database
	db, err := config.Connect(dbConfig)
	if err != nil {
		// log with comment
		log.Fatalf(err.Error())
	}

	mainID := 1
	laporanID := 2
	manajemenDataProsesID := 3
	olahDataID := 4
	pemantauanSistemID := 5
	n360ProfilID := 6
	penyediaanID := 7
	pendidikanID := 8
	penggunaanID := 9
	perawatanID := 10
	pemisahanID := 11
	koneksiKeSumberDataID := 13
	penarikanDataID := 14
	pemrosesanDataID := 15

	menus := []entity.Menu{
		{
			ID:          1,
			Name:        "No Parent",
			Description: "First menu, parent for main menu",
		},
		{
			ID:          2,
			Name:        "Laporan",
			Description: "Main menu",
			ParentID:    &mainID,
		},
		{
			ID:          3,
			Name:        "Manajemen Data & Proses",
			Description: "Main menu",
			ParentID:    &mainID,
		},
		{
			ID:          4,
			Name:        "Olah Data",
			Description: "Main menu",
			ParentID:    &mainID,
		},
		{
			ID:          5,
			Name:        "Pemantauan Sistem",
			Description: "Main menu",
			ParentID:    &mainID,
		},
		{
			ID:          6,
			Name:        "360 Profil",
			Description: "360 Profil menu",
			ParentID:    &laporanID,
		},
		{
			ID:          7,
			Name:        "Penyediaan",
			Description: "Penyediaan menu",
			ParentID:    &laporanID,
		},
		{
			ID:          8,
			Name:        "Pendidikan",
			Description: "Penyediaan menu",
			ParentID:    &laporanID,
		},
		{
			ID:          9,
			Name:        "Penggunaan",
			Description: "Penggunaan menu",
			ParentID:    &laporanID,
		},
		{
			ID:          10,
			Name:        "Perawatan",
			Description: "Perawatan menu",
			ParentID:    &laporanID,
		},
		{
			ID:          11,
			Name:        "Pemisahan",
			Description: "Main menu",
			ParentID:    &laporanID,
		},
		{
			ID:          12,
			Name:        "Data DSP",
			Description: "Main menu",
			ParentID:    &laporanID,
		},
		{
			ID:          13,
			Name:        "Koneksi ke Sumber Data",
			Description: "Koneksi ke Sumber Data menu",
			ParentID:    &manajemenDataProsesID,
		},
		{
			ID:          14,
			Name:        "Penarikan Data",
			Description: "Penarikan Data menu",
			ParentID:    &manajemenDataProsesID,
		},
		{
			ID:          15,
			Name:        "Pemrosesan Data",
			Description: "Pemrosesan Data menu",
			ParentID:    &manajemenDataProsesID,
		},
		{
			ID:          16,
			Name:        "Query Data",
			Description: "Query Data menu",
			ParentID:    &olahDataID,
		},
		{
			ID:          17,
			Name:        "Input Data",
			Description: "Input Data menu",
			ParentID:    &olahDataID,
		},
		{
			ID:          18,
			Name:        "Lineage",
			Description: "Lineage menu",
			ParentID:    &pemantauanSistemID,
		},
		{
			ID:          19,
			Name:        "Services",
			Description: "Services menu",
			ParentID:    &pemantauanSistemID,
		},
		{
			ID:          20,
			Name:        "360 Profil",
			Description: "360 Profil menu",
			ParentID:    &n360ProfilID,
		},
		{
			ID:          21,
			Name:        "Admin 360 Profil",
			Description: "Admin 360 Profil menu",
			ParentID:    &n360ProfilID,
		},
		{
			ID:          22,
			Name:        "Penyediaan",
			Description: "Penyediaan menu",
			ParentID:    &penyediaanID,
		},
		{
			ID:          23,
			Name:        "Admin Penyediaan",
			Description: "Admin Penyediaan menu",
			ParentID:    &penyediaanID,
		},
		{
			ID:          24,
			Name:        "Pendidikan",
			Description: "Pendidikan menu",
			ParentID:    &pendidikanID,
		},
		{
			ID:          25,
			Name:        "Admin Pendidikan",
			Description: "Admin Pendidikan menu",
			ParentID:    &pendidikanID,
		},
		{
			ID:          26,
			Name:        "Penggunaan",
			Description: "Penggunaan menu",
			ParentID:    &penggunaanID,
		},
		{
			ID:          27,
			Name:        "Admin Penggunaan",
			Description: "Admin Penggunaan menu",
			ParentID:    &penggunaanID,
		},
		{
			ID:          28,
			Name:        "Perawatan",
			Description: "Perawatan menu",
			ParentID:    &perawatanID,
		},
		{
			ID:          29,
			Name:        "Admin Perawatan",
			Description: "Admin Perawatan menu",
			ParentID:    &perawatanID,
		},
		{
			ID:          30,
			Name:        "Pemisahan",
			Description: "Pemisahan menu",
			ParentID:    &pemisahanID,
		},
		{
			ID:          31,
			Name:        "Admin Pemisahan",
			Description: "Admin Pemisahan menu",
			ParentID:    &pemisahanID,
		},
		{
			ID:          32,
			Name:        "Tambah Koneksi",
			Description: "Tambah Koneksi menu",
			ParentID:    &koneksiKeSumberDataID,
		},
		{
			ID:          33,
			Name:        "Ubah Koneksi",
			Description: "Ubah Koneksi menu",
			ParentID:    &koneksiKeSumberDataID,
		},
		{
			ID:          34,
			Name:        "Tambah Penarikan Data",
			Description: "Tambah Penarikan Data menu",
			ParentID:    &penarikanDataID,
		},
		{
			ID:          35,
			Name:        "Ubah Penarikan Data",
			Description: "Ubah Penarikan Data menu",
			ParentID:    &penarikanDataID,
		},
		{
			ID:          36,
			Name:        "Tambah Pemrosesan Data",
			Description: "Tambah Pemrosesan Data menu",
			ParentID:    &pemrosesanDataID,
		},
		{
			ID:          37,
			Name:        "Ubah Pemrosesan Data",
			Description: "Ubah Pemrosesan Data menu",
			ParentID:    &pemrosesanDataID,
		},
	}

	for _, menu := range menus {
		err := db.FirstOrCreate(&menu).Error
		if err != nil {
			log.Fatalf(err.Error())
		}
	}

	log.Println("Menus seeded successfully")
}
