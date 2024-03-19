package seed

import (
	"eagle-backend-dashboard/config"
	"eagle-backend-dashboard/entity"
	"log"
)

func MenuSeeders() {
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
			UrlKey:      "laporan",
			Description: "Main menu",
			ParentID:    &mainID,
			Icon:        "material-symbols:report-rounded",
			Url:         "/laporan/360-profil/360-profil",
		},
		{
			ID:          3,
			Name:        "Manajemen Data & Proses",
			UrlKey:      "manajemen-data-proses",
			Description: "Main menu",
			ParentID:    &mainID,
			Icon:        "material-symbols:bookmark-manager-rounded",
			Url:         "/manajemen-data-proses/koneksi-ke-sumber-data/tambah-koneksi",
		},
		{
			ID:          4,
			Name:        "Olah Data",
			UrlKey:      "olah-data",
			Description: "Main menu",
			ParentID:    &mainID,
			Icon:        "material-symbols-light:data-table-rounded",
			Url:         "/olah-data/query-data",
		},
		{
			ID:          5,
			Name:        "Pemantauan Sistem",
			UrlKey:      "pemantauan-sistem",
			Description: "Main menu",
			ParentID:    &mainID,
			Icon:        "ri:scan-2-fill",
			Url:         "/pemantauan-sistem/lineage",
		},
		{
			ID:          6,
			Name:        "360 Profil",
			UrlKey:      "laporan,360-profil",
			Description: "360 Profil menu",
			ParentID:    &laporanID,
			Url:         "/laporan/360-profil/360-profil",
			Icon:        "ic:round-360",
		},
		{
			ID:          7,
			Name:        "Penyediaan",
			UrlKey:      "laporan,penyediaan",
			Description: "Penyediaan menu",
			ParentID:    &laporanID,
			Url:         "/laporan/penyediaan/penyediaan",
			Icon:        "majesticons:box",
		},
		{
			ID:          8,
			Name:        "Pendidikan",
			UrlKey:      "laporan,pendidikan",
			Description: "Pendidikan menu",
			ParentID:    &laporanID,
			Url:         "/laporan/pendidikan/pendidikan",
			Icon:        "material-symbols:history-edu",
		},
		{
			ID:          9,
			Name:        "Penggunaan",
			UrlKey:      "laporan,penggunaan",
			Description: "Penggunaan menu",
			ParentID:    &laporanID,
			Url:         "/laporan/penggunaan/penggunaan",
			Icon:        "material-symbols:note-stack",
		},
		{
			ID:          10,
			Name:        "Perawatan",
			UrlKey:      "laporan,perawatan",
			Description: "Perawatan menu",
			ParentID:    &laporanID,
			Url:         "/laporan/perawatan/perawatan",
			Icon:        "fluent:slide-settings-24-filled",
		},
		{
			ID:          11,
			Name:        "Pemisahan",
			UrlKey:      "laporan,pemisahan",
			Description: "Pemisahan menu",
			ParentID:    &laporanID,
			Url:         "/laporan/pemisahan/pemisahan",
			Icon:        "ant-design:split-cells-outlined",
		},
		{
			ID:          12,
			Name:        "Data DSP",
			UrlKey:      "laporan,data-dsp",
			Description: "Data DSP menu",
			ParentID:    &laporanID,
			Url:         "/laporan/data-dsp",
			Icon:        "tabler:binary-tree",
		},
		{
			ID:          13,
			Name:        "Koneksi ke Sumber Data",
			UrlKey:      "manajemen-data-proses,koneksi-ke-sumber-data",
			Description: "Koneksi ke Sumber Data menu",
			ParentID:    &manajemenDataProsesID,
			Url:         "/manajemen-data-proses/koneksi-ke-sumber-data/tambah-koneksi",
			Icon:        "majesticons:data",
		},
		{
			ID:          14,
			Name:        "Penarikan Data",
			UrlKey:      "manajemen-data-proses,penarikan-data",
			Description: "Penarikan Data menu",
			ParentID:    &manajemenDataProsesID,
			Url:         "/manajemen-data-proses/penarikan-data/penarikan-data",
			Icon:        "fluent:arrow-sync-circle-24-filled",
		},
		{
			ID:          15,
			Name:        "Pemrosesan Data",
			UrlKey:      "manajemen-data-proses,pemrosesan-data",
			Description: "Pemrosesan Data menu",
			ParentID:    &manajemenDataProsesID,
			Url:         "/manajemen-data-proses/pemrosesan-data/pemrosesan-data",
			Icon:        "fluent:slide-settings-24-filled",
		},
		{
			ID:          16,
			Name:        "Query Data",
			Description: "Query Data menu",
			ParentID:    &olahDataID,
			UrlKey:      "olah-data,query-data",
			Url:         "/olah-data/query-data",
			Icon:        "mdi:sql-query",
		},
		{
			ID:          17,
			Name:        "Input Data",
			Description: "Input Data menu",
			ParentID:    &olahDataID,
			UrlKey:      "olah-data,input-data",
			Url:         "/olah-data/input-data",
			Icon:        "lucide:folder-input",
		},
		{
			ID:          18,
			Name:        "Lineage",
			Description: "Lineage menu",
			ParentID:    &pemantauanSistemID,
			UrlKey:      "pemantauan-sistem,lineage",
			Url:         "/pemantauan-sistem/lineage",
			Icon:        "fa6-solid:user-group",
		},
		{
			ID:          19,
			Name:        "Services",
			Description: "Services menu",
			ParentID:    &pemantauanSistemID,
			UrlKey:      "pemantauan-sistem,services",
			Url:         "/pemantauan-sistem/services",
			Icon:        "ri:customer-service-2-fill",
		},
		{
			ID:          20,
			Name:        "360 Profil",
			Description: "360 Profil menu",
			ParentID:    &n360ProfilID,
			UrlKey:      "laporan,360-profil,360-profil",
			Url:         "/laporan/360-profil/360-profil",
			Icon:        "ic:round-360",
		},
		{
			ID:          21,
			Name:        "Admin 360 Profil",
			Description: "Admin 360 Profil menu",
			ParentID:    &n360ProfilID,
			UrlKey:      "laporan,360-profil,admin-360-profil",
			Url:         "/laporan/360-profil/admin-360-profil",
			Icon:        "ic:round-360",
		},
		{
			ID:          22,
			Name:        "Penyediaan",
			Description: "Penyediaan menu",
			ParentID:    &penyediaanID,
			UrlKey:      "laporan,penyediaan,penyediaan",
			Icon:        "majesticons:box",
		},
		{
			ID:          23,
			Name:        "Admin Penyediaan",
			Description: "Admin Penyediaan menu",
			ParentID:    &penyediaanID,
			UrlKey:      "laporan,penyediaan,admin-penyediaan",
			Icon:        "majesticons:box",
		},
		{
			ID:          24,
			Name:        "Pendidikan",
			Description: "Pendidikan menu",
			ParentID:    &pendidikanID,
			UrlKey:      "laporan,pendidikan,pendidikan",
			Icon:        "material-symbols:history-edu",
		},
		{
			ID:          25,
			Name:        "Admin Pendidikan",
			Description: "Admin Pendidikan menu",
			ParentID:    &pendidikanID,
			UrlKey:      "laporan,pendidikan,admin-pendidikan",
			Icon:        "material-symbols:history-edu",
		},
		{
			ID:          26,
			Name:        "Penggunaan",
			Description: "Penggunaan menu",
			ParentID:    &penggunaanID,
			UrlKey:      "laporan,penggunaan,penggunaan",
			Icon:        "material-symbols:note-stack",
		},
		{
			ID:          27,
			Name:        "Admin Penggunaan",
			Description: "Admin Penggunaan menu",
			ParentID:    &penggunaanID,
			UrlKey:      "laporan,penggunaan,admin-penggunaan",
			Icon:        "material-symbols:note-stack",
		},
		{
			ID:          28,
			Name:        "Perawatan",
			Description: "Perawatan menu",
			ParentID:    &perawatanID,
			UrlKey:      "laporan,perawatan,perawatan",
			Icon:        "fluent:slide-settings-24-filled",
		},
		{
			ID:          29,
			Name:        "Admin Perawatan",
			Description: "Admin Perawatan menu",
			ParentID:    &perawatanID,
			UrlKey:      "laporan,perawatan,admin-perawatan",
			Icon:        "fluent:slide-settings-24-filled",
		},
		{
			ID:          30,
			Name:        "Pemisahan",
			Description: "Pemisahan menu",
			ParentID:    &pemisahanID,
			UrlKey:      "laporan,pemisahan,pemisahan",
			Icon:        "ant-design:split-cells-outlined",
		},
		{
			ID:          31,
			Name:        "Admin Pemisahan",
			Description: "Admin Pemisahan menu",
			ParentID:    &pemisahanID,
			UrlKey:      "laporan,pemisahan,admin-pemisahan",
			Icon:        "ant-design:split-cells-outlined",
		},
		{
			ID:          32,
			Name:        "Tambah Koneksi",
			Description: "Tambah Koneksi menu",
			ParentID:    &koneksiKeSumberDataID,
			UrlKey:      "manajemen-data-proses,koneksi-ke-sumber-data,tambah-koneksi",
			Icon:        "majesticons:data",
		},
		{
			ID:          33,
			Name:        "Ubah Koneksi",
			Description: "Ubah Koneksi menu",
			ParentID:    &koneksiKeSumberDataID,
			UrlKey:      "manajemen-data-proses,koneksi-ke-sumber-data,ubah-koneksi",
			Icon:        "majesticons:data",
		},
		{
			ID:          34,
			Name:        "Tambah Penarikan Data",
			Description: "Tambah Penarikan Data menu",
			ParentID:    &penarikanDataID,
			UrlKey:      "manajemen-data-proses,penarikan-data,tambah-penarikan-data",
			Icon:        "fluent:arrow-sync-circle-24-filled",
		},
		{
			ID:          35,
			Name:        "Ubah Penarikan Data",
			Description: "Ubah Penarikan Data menu",
			ParentID:    &penarikanDataID,
			UrlKey:      "manajemen-data-proses,penarikan-data,ubah-penarikan-data",
			Icon:        "fluent:arrow-sync-circle-24-filled",
		},
		{
			ID:          36,
			Name:        "Tambah Pemrosesan Data",
			Description: "Tambah Pemrosesan Data menu",
			ParentID:    &pemrosesanDataID,
			UrlKey:      "manajemen-data-proses,pemrosesan-data,tambah-pemrosesan-data",
			Icon:        "fluent:slide-settings-24-filled",
		},
		{
			ID:          37,
			Name:        "Ubah Pemrosesan Data",
			Description: "Ubah Pemrosesan Data menu",
			ParentID:    &pemrosesanDataID,
			UrlKey:      "manajemen-data-proses,pemrosesan-data,ubah-pemrosesan-data",
			Icon:        "fluent:slide-settings-24-filled",
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
