package soal1

import "fmt"

type SchoolInfo struct {
	Name  string
	Major string
}

type Person struct {
	name       string
	photo      string
	email      string
	age        int
	phone      string
	isMarried  bool
	listSchool []SchoolInfo
}

func Main() {
	profile := Person{
		name:      "Habib Muhammad Rizki",
		photo:     "https://www.google.com/url?sa=i&url=https%3A%2F%2Fid.wikipedia.org%2Fwiki%2FHabib_Burquibah&psig=AOvVaw0UnEmoU0GsUuaKCg5G31UE&ust=1755683324861000&source=images&cd=vfe&opi=89978449&ved=0CBUQjRxqFwoTCJiZw5nMlo8DFQAAAAAdAAAAABAE",
		email:     "habib@mail.com",
		age:       21,
		phone:     "081547827829",
		isMarried: false,
		listSchool: []SchoolInfo{
			{Name: "SMA N 1 Cilacap", Major: "IPS"},
		},
	}
	fmt.Println(profile.name)
	fmt.Println(profile.phone)
	fmt.Println(profile.photo)
	fmt.Println(profile.email)
	fmt.Println(profile.age)
	fmt.Println(profile.phone)
	fmt.Println(profile.isMarried)
	fmt.Println("Daftar Sekolah dan Jurusan:")
	for _, school := range profile.listSchool {
		fmt.Printf("Nama Sekolah: %s, Jurusan: %s\n", school.Name, school.Major)
	}

}
