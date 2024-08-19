package util

import (
	model "attendance-is/models"

	"golang.org/x/crypto/bcrypt"
)

func Seed() {
	pass, _ := bcrypt.GenerateFromPassword([]byte("secret"), bcrypt.DefaultCost)
	users := []model.User{
		{Email: "a@gmail.com", Password: string(pass), Type: 0},
		{Email: "b@gmail.com", Password: string(pass), Type: 0},
		{Email: "c@gmail.com", Password: string(pass), Type: 0},
		{Email: "kurnianingsih@gmail.com", Password: string(pass), Type: 1},
		{Email: "kuwatsantoso@gmail.com", Password: string(pass), Type: 1},
		{Email: "rizky.w.dewantoro@gmail.com", Password: string(pass), Type: 2},
	}
	DB.Create(&users)

	classes := []model.Class{
		{Name: "TIA2021"},
		{Name: "TIB2021"},
		{Name: "TIC2021"},
	}
	DB.Save(&classes)

	students := []model.Student{
		{ID: 1, Nim: "43321001", Name: "ADI RIFTA DWI KURNIAWAN", UserID: users[0].ID},
		{ID: 2, Nim: "43321002", Name: "ADRIAN REYHAN PRANATA"},
		{ID: 3, Nim: "43321003", Name: "AHMAD ADI IRFANSYAH"},
		{ID: 4, Nim: "43321004", Name: "ALFINA NAILAL MUNA"},
		{ID: 5, Nim: "43321005", Name: "ANATASYA OKTA BERLIANA"},
		{ID: 6, Nim: "43321006", Name: "CINTANIA SALSABELLA BERLIANT"},
		{ID: 7, Nim: "43321007", Name: "DENY WISNU SAPUTRO SUKISNO"},
		{ID: 8, Nim: "43321008", Name: "DZAKIY FAIZAL ABDUH"},
		{ID: 9, Nim: "43321009", Name: "FAJAR FATONI PRATAMA"},
		{ID: 10, Nim: "43321010", Name: "FIKRI MAJID"},
		{ID: 11, Nim: "43321011", Name: "FIRMAN SYAHRUL NIZAM"},
		{ID: 12, Nim: "43321012", Name: "HAFID RIZKY ALIYUDO"},
		{ID: 13, Nim: "43321013", Name: "ILHAM YANUAR PUTRA"},
		{ID: 14, Nim: "43321014", Name: "KANNAZ 'IZZUL FATWA"},
		{ID: 15, Nim: "43321015", Name: "LISGIYANTO SOFIYAN"},
		{ID: 16, Nim: "43321016", Name: "M. HILMY IRFAN MAULANA"},
		{ID: 17, Nim: "43321017", Name: "MIRZA SYAHIR NUR RAMADHAN"},
		{ID: 18, Nim: "43321018", Name: "MUHAMMAD DANI AL FAHMI"},
		{ID: 19, Nim: "43321019", Name: "MUHAMMAD WAHYU ANGGORO"},
		{ID: 20, Nim: "43321020", Name: "MUHAMMAD ZAKA TAUFIQUL HAKIM"},
		{ID: 21, Nim: "43321021", Name: "NAUFAL ADLI SANTOSA"},
		{ID: 22, Nim: "43321022", Name: "RIZKY DWI PURNAMA"},
		{ID: 23, Nim: "43321023", Name: "SAKA WIJAYA"},
		{ID: 24, Nim: "43321024", Name: "SINDI"},
		{ID: 25, Nim: "43321025", Name: "SUFYAN HANIF ARIYANA"},
		{ID: 26, Nim: "43321026", Name: "SYAUQI NUR MUHAMMAD"},

		{ID: 1 + 26, Nim: "43321101", Name: "AFIF IHZA DARMAWAN", UserID: users[1].ID},
		{ID: 1 + 27, Nim: "43321102", Name: "AHMAD SYAHRUL FAZA"},
		{ID: 1 + 28, Nim: "43321103", Name: "AINNUR HANIF NUGRAHA"},
		{ID: 1 + 29, Nim: "43321104", Name: "ARDHILLA EKA WINDIARTI"},
		{ID: 1 + 30, Nim: "43321105", Name: "ARYANTO ANDRI SOBIRIN"},
		{ID: 1 + 31, Nim: "43321106", Name: "BHATINDEN SEJIARGA ERGUN GIORTAGMA"},
		{ID: 1 + 32, Nim: "43321107", Name: "DEDY SETIAWAN EKA ARIANTO"},
		{ID: 1 + 33, Nim: "43321108", Name: "DWI RATNA PUSPITA SARI"},
		{ID: 1 + 34, Nim: "43321109", Name: "FEBRI ADI SETIYAWAN"},
		{ID: 1 + 35, Nim: "43321110", Name: "FIDO JAHFAL PRAYOGA"},
		{ID: 1 + 36, Nim: "43321111", Name: "GLADISA WIDADINING CAHYA"},
		{ID: 1 + 37, Nim: "43321112", Name: "HAYA YUMNA ZHARIFAH"},
		{ID: 1 + 38, Nim: "43321113", Name: "INDRA HERDIANA"},
		{ID: 1 + 39, Nim: "43321114", Name: "IZZAL IHZA MAULANA"},
		{ID: 1 + 40, Nim: "43321115", Name: "LUCKY NURFITRIANA"},
		{ID: 1 + 41, Nim: "43321116", Name: "M DANIAR RIFQI"},
		{ID: 1 + 42, Nim: "43321117", Name: "MUHAMAD ALIF JAUZI"},
		{ID: 1 + 43, Nim: "43321118", Name: "MUHAMMAD DAVID AKBAR"},
		{ID: 1 + 44, Nim: "43321119", Name: "MUHAMMAD WISNU AINUN NAJIB"},
		{ID: 1 + 45, Nim: "43321120", Name: "NAWAF ABDUL AZIZ"},
		{ID: 1 + 46, Nim: "43321121", Name: "RAVI NAUFAL SHEVANDAPUTRA"},
		{ID: 1 + 47, Nim: "43321122", Name: "RAYHAN SULTAN AKBAR"},
		{ID: 1 + 48, Nim: "43321123", Name: "RIZKY WAHYU DEWANTORO"},
		{ID: 1 + 49, Nim: "43321124", Name: "SITI AINUN RODHIYAH"},
		{ID: 1 + 50, Nim: "43321125", Name: "TEGAR RADITYA PERMANA PUTRA"},

		{ID: 1 + 51, Nim: "43321201", Name: "AFIF RAMZY BADRANI", UserID: users[2].ID},
		{ID: 1 + 52, Nim: "43321202", Name: "AJI DWI PRAKOSO"},
		{ID: 1 + 53, Nim: "43321203", Name: "ASHABUL KAHFI"},
		{ID: 1 + 54, Nim: "43321204", Name: "BINA SATRIA FAUZI"},
		{ID: 1 + 55, Nim: "43321205", Name: "BUKHARY AZRIELLOREZQA YUFAR"},
		{ID: 1 + 56, Nim: "43321206", Name: "EKA YULIANTO"},
		{ID: 1 + 57, Nim: "43321208", Name: "FAUQA JAHFAL ISKANDAR"},
		{ID: 1 + 58, Nim: "43321209", Name: "GELORAWAN SUSATYO JATIPAMUNGKAS"},
		{ID: 1 + 59, Nim: "43321210", Name: "HENDI AHMAD SHOLEHUDIN"},
		{ID: 1 + 60, Nim: "43321211", Name: "HERLAMBANG ARIEF SAPUTRA"},
		{ID: 1 + 61, Nim: "43321212", Name: "ILHAM RIZKY HARIJANTO"},
		{ID: 1 + 62, Nim: "43321213", Name: "JOVAN ADHIK GUNAWAN"},
		{ID: 1 + 63, Nim: "43321214", Name: "JOYVA ALIANS TAMAM"},
		{ID: 1 + 64, Nim: "43321215", Name: "MIKO BAYU ANGGORO"},
		{ID: 1 + 65, Nim: "43321216", Name: "MUHAMMAD MIZZY"},
		{ID: 1 + 66, Nim: "43321217", Name: "MUHAMMAD NAUFAL SYARIFUDIN"},
		{ID: 1 + 67, Nim: "43321218", Name: "NATYA FAKHRI NUR PRASETYA"},
		{ID: 1 + 68, Nim: "43321219", Name: "RAKHMAT SATRIADI"},
		{ID: 1 + 69, Nim: "43321220", Name: "RIZKI PUTRI FITRIYANI"},
		{ID: 1 + 70, Nim: "43321221", Name: "RIZKI SHAUMI SABIQ"},
		{ID: 1 + 71, Nim: "43321222", Name: "SUSSILO WATI NOVITA"},
		{ID: 1 + 72, Nim: "43321223", Name: "TRI KUSUMA FARADILA"},
		{ID: 1 + 73, Nim: "43321224", Name: "YOEL RAGA ARI YORGA"},
		{ID: 1 + 74, Nim: "43321225", Name: "ZULFA NURMA NOVITA SARI"},
	}

	DB.Save(&students)
	DB.Model(&classes[0]).Association("Students").Replace(students[0:26])
	DB.Model(&classes[1]).Association("Students").Replace(students[26:51])
	DB.Model(&classes[2]).Association("Students").Replace(students[51:75])

	lecturer := []model.Lecturer{
		{Nip: "197904262003122002", Name: "KURNIANINGSIH, S.T., M.T., Dr.", UserID: users[3].ID},
		{Nip: "198407192019031008", Name: "KUWAT SANTOSO, M. KOM", UserID: users[4].ID},
		{Nip: "197403112000121001", Name: "MARDIYONO, S.Kom., M.Sc."},
		{Nip: "199001072019031020", Name: "MUHAMMAD IRWAN YANWARI. S.Kom., M.Eng."},
		{Nip: "199107302019031010", Name: "NURSENO BAYU AJI, S. Kom, M. Kom."},
		{Nip: "198504102014041002", Name: "PRAYITNO, S.ST., M.T."},
		{Nip: "196810252000121001", Name: "TRI RAHARJO YUDANTORO, S. Kom., M.Kom."},
		{Nip: "198703272019022012", Name: "WIKTASARI S.T., M.Kom"},
		{Nip: "199004112019031014", Name: "AFANDI NUR AZIZ THOHARI, S.T., M.Cs"},
		{Nip: "198605292019032009", Name: "AISYATUL KARIMA, S.Kom., MCS"},
		{Nip: "198810142019031007", Name: "AMRAN YOBIOKTABERA, M. KOM"},
		{Nip: "199202052019031009", Name: "ANGGA WAHYU WIBOWO S.Kom., M.Eng"},
		{Nip: "197711192008012013", Name: "IDHAWATI H., S.Kom., M.Kom."},
		{Nip: "198404202015041003", Name: "LILIEK TRIYONO, S.T., M.Kom"},
		{Nip: "196008221988031001", Name: "PARSUMO RAHARDJO, Drs., M.Kom."},
		{Nip: "199401272019032036", Name: "SIRLI FAHRIAH, S.Kom, M.Kom."},
		{Nip: "197501302001121001", Name: "SLAMET HANDOKO, S.Kom., M.Kom."},
		{Nip: "197101172003121001", Name: "SUKAMTO, S.Kom., M.T."},
		{Nip: "197704012005011001", Name: "WAHYU SULISTIYO, S.T., M.Kom."},
	}

	DB.Save(&lecturer)

	courses := []model.Course{
		{Code: "432-191-502", Name: "Desain dan Perancangan Sistem Informasi", Hours: 5},
		{Code: "432-191-503", Name: "Sistem Multimedia", Hours: 5},
		{Code: "432-191-504", Name: "Internet of Things (IoT)", Hours: 5},
		{Code: "432-191-505", Name: "Data Mining", Hours: 5},
		{Code: "432-191-506", Name: "Komputasi Awan", Hours: 5},
		{Code: "432-191-508", Name: "Keamanan, Kesehatan, dan Keselamatan Kerja", Hours: 5},
		{Code: "432-191-509", Name: "Interaksi Manusia dan Komputer", Hours: 2},
	}

	DB.Save(&courses)

	enrollments := []model.Enrollment{
		{ClassID: 1, CourseID: 1, Lecturers: &[]model.Lecturer{{ID: 1}}},
		{ClassID: 2, CourseID: 1, Lecturers: &[]model.Lecturer{{ID: 1}}},
		{ClassID: 3, CourseID: 1, Lecturers: &[]model.Lecturer{{ID: 1}}},
		{ClassID: 1, CourseID: 2, Lecturers: &[]model.Lecturer{{ID: 2}}},
		{ClassID: 2, CourseID: 2, Lecturers: &[]model.Lecturer{{ID: 2}}},
		{ClassID: 3, CourseID: 2, Lecturers: &[]model.Lecturer{{ID: 2}}},
		{ClassID: 1, CourseID: 3, Lecturers: &[]model.Lecturer{{ID: 3}}},
		{ClassID: 2, CourseID: 3, Lecturers: &[]model.Lecturer{{ID: 3}}},
		{ClassID: 3, CourseID: 3, Lecturers: &[]model.Lecturer{{ID: 3}}},
		{ClassID: 1, CourseID: 4, Lecturers: &[]model.Lecturer{{ID: 4}}},
		{ClassID: 2, CourseID: 4, Lecturers: &[]model.Lecturer{{ID: 4}}},
		{ClassID: 3, CourseID: 4, Lecturers: &[]model.Lecturer{{ID: 4}}},
	}

	DB.Create(&enrollments)
}
