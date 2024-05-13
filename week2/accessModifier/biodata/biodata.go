package biodata

// dapat di akses di luar package
var NamaLengkap = "Dimas Ramdani"
var Umur = 19

func SayHello(name string) string {
	return "Hallo Nama Saya " + name
}

// tidak dapat di akses di luar package
var jurusan = "Rekayasa Perangkat Lunak"

func sayGoodBye(name string) string {
	return "Selamat tinggal " + name
}
