package entity

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("se-g07.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	database.AutoMigrate(
		&Teacher{},
		&Student{},
		&Semester{},
		&ExamType{},
		&Program{},
		&Course{},
		&ExamSchedule{},
		&AddCourse{},
		&Petition{},
		&Registrar{},
		&IncreaseGrades{},
		&Grades{},
		&RequestStatus{},
		&RequestExam{},
		&RecordPetition{},
		&RegisCourse{},
		&Withdrawal{},
	)

	db = database

	password, err := bcrypt.GenerateFromPassword([]byte("123456"), 14)

	registrar1 := Registrar{
		ID_registrar: "R5678912",
		Prefix:       "Miss",
		Name:         "Fasai Jangloei",
		Email:        "fj@gmail.com",
		Password:     string(password),
	}
	db.Model(&Registrar{}).Create(&registrar1)

	registrar2 := Registrar{
		ID_registrar: "R2034567",
		Prefix:       "Mr.",
		Name:         "Sompong Naja",
		Email:        "sn@gmail.com",
		Password:     string(password),
	}
	db.Model(&Registrar{}).Create(&registrar2)

	registrar3 := Registrar{
		ID_registrar: "R1234567",
		Prefix:       "Mr.",
		Name:         "Meta  Kunwat",
		Email:        "mk@gmail.com",
		Password:     string(password),
	}
	db.Model(&Registrar{}).Create(&registrar3)

	registrar4 := Registrar{
		ID_registrar: "R2345678",
		Prefix:       "Mr.",
		Name:         "Apiwat Sompung",
		Email:        "ap@gmail.com",
		Password:     string(password),
	}
	db.Model(&Registrar{}).Create(&registrar4)

	registrar5 := Registrar{
		ID_registrar: "R5901258",
		Prefix:       "Miss",
		Name:         "Jeeninee Khuptawuttinun",
		Email:        "Jeeninee@gmail.com",
		Password:     string(password),
	}
	db.Model(&Registrar{}).Create(&registrar5)

	student1 := Student{
		ID_student: "B6202385",
		Prefix:     "Miss",
		Name:       "Phatcha Sisuwo",
		Major:      "CPE",
		Year:       3,
		Email:      "phatcha@gmail.com",
		Password:   string(password),
	}
	db.Model(&Student{}).Create(&student1)

	student2 := Student{
		ID_student: "B6202743",
		Prefix:     "Miss",
		Name:       "Narudee Arunno",
		Major:      "CPE",
		Year:       3,
		Email:      "narudee@gmail.com",
		Password:   string(password),
	}
	db.Model(&Student{}).Create(&student2)

	student3 := Student{
		ID_student: "B6214449",
		Prefix:     "Miss",
		Name:       "Suwanan Thamsui",
		Major:      "CPE",
		Year:       3,
		Email:      "suwanan@gmail.com",
		Password:   string(password),
	}
	db.Model(&Student{}).Create(&student3)

	student4 := Student{
		ID_student: "B6230760",
		Prefix:     "Miss",
		Name:       "Patnarin Aiewchoei",
		Major:      "CPE",
		Year:       3,
		Email:      "patnarin@gmail.com",
		Password:   string(password),
	}
	db.Model(&Student{}).Create(&student4)

	student5 := Student{
		ID_student: "B5924615",
		Prefix:     "Mr.",
		Name:       "Pawarit Praneetponkrang",
		Major:      "CPE",
		Year:       6,
		Email:      "pawarit@gmail.com",
		Password:   string(password),
	}
	db.Model(&Student{}).Create(&student5)

	student6 := Student{
		ID_student: "B5901258",
		Prefix:     "Miss",
		Name:       "Khuptawuttinun",
		Major:      "CPE",
		Year:       6,
		Email:      "Khuptawuttinun@gmail.com",
		Password:   string(password),
	}
	db.Model(&Student{}).Create(&student6)

	teacher1 := Teacher{
		ID_teacher: "T000001",
		Name:       "ผู้ช่วยศาสตราจารย์ ดร.ชาญวิทย์ แก้วกสิ",
		Email:      "chanwit@gmail.com",
		Password:   string(password),
		Prefix:     "Mr.",
		Major:      "CPE",
	}
	db.Model(&Teacher{}).Create(&teacher1)

	teacher2 := Teacher{
		ID_teacher: "T0023581",
		Name:       "ผู้ช่วยศาสตราจารย์ ดร.นันทวุฒิ คะอังกุ",
		Email:      "nuntawut@gmail.com",
		Password:   string(password),
		Prefix:     "Mr.",
		Major:      "CPE",
	}
	db.Model(&Teacher{}).Create(&teacher2)

	teacher3 := Teacher{
		ID_teacher: "T0157690",
		Name:       "ผู้ช่วยศาสตราจารย์ ดร.ศรัญญา กาญจนวัฒนา",
		Email:      "sarunya@gmail.com",
		Password:   string(password),
		Prefix:     "Miss",
		Major:      "CPE",
	}
	db.Model(&Teacher{}).Create(&teacher3)

	teacher4 := Teacher{
		ID_teacher: "T1578952",
		Name:       "ผู้ช่วยศาสตราจารย์ ดร.ปรเมศวร์ ห่อแก้ว",
		Email:      "paramate@gmail.com",
		Password:   string(password),
		Prefix:     "Mr.",
		Major:      "CPE",
	}
	db.Model(&Teacher{}).Create(&teacher4)

	//Course Data
	course1 := Course{
		Coursename:   "SOFTWARE ENGINEERING",
		Coursenumber: 523332,
	}
	db.Model(&Course{}).Create(&course1)
	course2 := Course{
		Coursename:   "COMPUTER AND COMMUNICATION",
		Coursenumber: 523352,
	}
	db.Model(&Course{}).Create(&course2)
	course3 := Course{
		Coursename:   "OPERATING SYSTEMS",
		Coursenumber: 523354,
	}
	db.Model(&Course{}).Create(&course3)
	course4 := Course{
		Coursename:   "System Analysis and Design",
		Coursenumber: 523331,
	}
	db.Model(&Course{}).Create(&course4)
	course5 := Course{
		Coursename:   "DATABASE SYSTEMS",
		Coursenumber: 523211,
	}
	db.Model(&Course{}).Create(&course5)
	course6 := Course{
		Coursename:   "COMPUTER STATISTICS",
		Coursenumber: 523301,
	}
	db.Model(&Course{}).Create(&course6)

	// Program data
	program1 := Program{
		Programname: "Thai Program",
	}
	db.Model(&Program{}).Create(&program1)
	program2 := Program{
		Programname: "International Program",
	}
	db.Model(&Program{}).Create(&program2)

	//Semester Data
	Semester1 := Semester{
		Semester: "ภาคการศึกษาที่ 1",
	}
	db.Model(&Semester{}).Create(&Semester1)

	Semester2 := Semester{
		Semester: "ภาคการศึกษาที่ 2",
	}
	db.Model(&Semester{}).Create(&Semester2)

	Semester3 := Semester{
		Semester: "ภาคการศึกษาที่ 3",
	}
	db.Model(&Semester{}).Create(&Semester3)

	//RequestStatus Data
	status1 := RequestStatus{
		Status: "รอดำเนินการ",
	}
	db.Model(&RequestStatus{}).Create(&status1)

	/*status2 := RequestStatus{
		Status: "กำลังดำเนินการ",
	}
	db.Model(&RequestStatus{}).Create(&status2)
	status3 := RequestStatus{
		Status: "ดำเนินการสำเร็จแล้ว",
	}
	db.Model(&RequestStatus{}).Create(&status3)*/

	//
	claim1 := Petition{
		Claim: "เกินกว่าหน่วยกิตกำหนด",
	}
	db.Model(&Petition{}).Create(&claim1)
	claim2 := Petition{
		Claim: "ต่ำกว่าหน่วยกิตกำหนด",
	}
	db.Model(&Petition{}).Create(&claim2)

	// ExamType Data
	type1 := ExamType{
		Type: "กลางภาค",
	}
	db.Model(&ExamType{}).Create(&type1)
	type2 := ExamType{
		Type: "ปลายภาค",
	}
	db.Model(&ExamType{}).Create(&type2)
	grade1 := Grades{
		Grade: "A",
	}
	db.Model(&Grades{}).Create(&grade1)
	grade2 := Grades{
		Grade: "B+",
	}
	db.Model(&Grades{}).Create(&grade2)
	grade3 := Grades{
		Grade: "B",
	}
	db.Model(&Grades{}).Create(&grade3)
	grade4 := Grades{
		Grade: "C+",
	}
	db.Model(&Grades{}).Create(&grade4)
	grade5 := Grades{
		Grade: "C",
	}
	db.Model(&Grades{}).Create(&grade5)
	grade6 := Grades{
		Grade: "D+",
	}
	db.Model(&Grades{}).Create(&grade6)
	grade7 := Grades{
		Grade: "D",
	}
	db.Model(&Grades{}).Create(&grade7)
	grade8 := Grades{
		Grade: "F",
	}
	db.Model(&Grades{}).Create(&grade8)

	regis1 := RegisCourse{
		Student: student2,
		Course:  course4,
	}
	db.Model(&RegisCourse{}).Create(&regis1)

	// ข้อมูล regiscourse
	db.Model(&RegisCourse{}).Create(&RegisCourse{
		Student: student2,
		Course:  course1,
	})

	db.Model(&RegisCourse{}).Create(&RegisCourse{
		Student: student2,
		Course:  course2,
	})

	db.Model(&RegisCourse{}).Create(&RegisCourse{
		Student: student2,
		Course:  course3,
	})

	db.Model(&RegisCourse{}).Create(&RegisCourse{
		Student: student1,
		Course:  course4,
	})

	db.Model(&RegisCourse{}).Create(&RegisCourse{
		Student: student1,
		Course:  course5,
	})

	db.Model(&RegisCourse{}).Create(&RegisCourse{
		Student: student1,
		Course:  course6,
	})

	db.Model(&RegisCourse{}).Create(&RegisCourse{
		Student: student3,
		Course:  course1,
	})

	db.Model(&RegisCourse{}).Create(&RegisCourse{
		Student: student3,
		Course:  course2,
	})

	db.Model(&RegisCourse{}).Create(&RegisCourse{
		Student: student4,
		Course:  course3,
	})

	db.Model(&RegisCourse{}).Create(&RegisCourse{
		Student: student4,
		Course:  course4,
	})

	db.Model(&RegisCourse{}).Create(&RegisCourse{
		Student: student5,
		Course:  course5,
	})

	db.Model(&RegisCourse{}).Create(&RegisCourse{
		Student: student5,
		Course:  course6,
	})

	// // ข้อมูล Withdrawal
	// db.Model(&Withdrawal{}).Create(&Withdrawal{
	// 	Student:        student2,
	// 	RegisCourse:    regis1,
	// 	Teacher:        teacher1,
	// 	Semester:       Semester1,
	// 	YearTime:       2564,
	// 	Reason:         "so hard",
	// 	RemainCredit:   15,
	// 	WithdrawalTime: time.Now(),
	// })

}
