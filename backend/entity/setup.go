package entity

import ( 

	"time"
     
	"golang.org/x/crypto/bcrypt"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm" 
    
)

var db *gorm.DB

func DB() *gorm.DB { 
    return db
}

func SetupDatabase() {
    database, err := gorm.Open(sqlite.Open("sa-64.db"), &gorm.Config{}) 
    if err != nil {
        panic("failed to connect database") 
}

    database.AutoMigrate(
        &Professor{}, 
		&Course{}, 
		&TA{}, 
		&Room{}, 
		&ManageCourse{},
)

    db = database 

	password, err := bcrypt.GenerateFromPassword([]byte("123456"), 14)

	//Taecher Data
    db.Model(&Professor{}).Create(&Professor{
		Name: "ดร.ศรัญญา กาญจนวัฒนา",
		Professer_id: "R1234567",
		Password: string(password),
		Major: "วิศวกรรมสิ่งแวดล้อม",
	})
    

    db.Model(&Professor{}).Create(&Professor{
		Name:  "ดร.ชาญวิทย์ แก้วกสิ",
		Professer_id: "R555552",
		Password: string(password),
		Major: "วิศวกรรมคอมพิวเตอร์",
	})
	

    // --- Course Data
	Course1 := Course{
		Name:  "Database Systems",
		Code:   "523211",
		Credit: 4,
	}
	db.Model(&Course{}).Create(&Course1)

	Course2 := Course{
		Name:  "Computer Engineering Projec",
		Code:   "523480",
		Credit: 4,
	}
	db.Model(&Course{}).Create(&Course2)

	Course3 := Course{
		Name:  "Problem Solving with Programming",
		Code:   "523203",
		Credit: 2,
	}
	db.Model(&Course{}).Create(&Course3)

    // --- TA Data
    A := TA{
		Name: "นายA",
        TA_id: "G11111",
	}
	db.Model(&TA{}).Create(&A)

	B := TA{
		Name: "นายB",
        TA_id: "G22222",
	}
	db.Model(&TA{}).Create(&B)

	C := TA{
		Name: "นายC",
        TA_id: "G33333",
	}
	db.Model(&TA{}).Create(&C)

    // --- Room Data
    room1 := Room{
		Number: 1211,
        StudentCount: 120,
	}
	db.Model(&Room{}).Create(&room1)

    room2 := Room{
		Number: 4101,
        StudentCount: 1500,
	}
	db.Model(&Room{}).Create(&room2)

    room3 := Room{
		Number: 5203,
        StudentCount: 600,
	}
	db.Model(&Room{}).Create(&room3)

    //ManageCourse data
	var sarunya Professor
	var chanwit Professor

    db.Raw("SELECT * FROM professors WHERE professer_id = ?", "R1234567").Scan(&sarunya)
	db.Raw("SELECT * FROM professors WHERE professer_id = ?", "R555552").Scan(&chanwit)

	manage1 := ManageCourse{

		Group:			1,
		Term:			2,
    	TeachingTime:    8,
    	Ungraduated_year:       2564,
    	ManageCourseTime:    time.Date(2021, 07, 12, 10, 00, 00, 0000000, time.FixedZone("UTC+7",7*3600 )),
        Course: Course1,
        TA: A,
        Room:  room1,
		Professor: sarunya,	
	}
	db.Model(&ManageCourse{}).Create(&manage1)
	manage2 := ManageCourse{

		Group:			5,
		Term:			3,
    	TeachingTime:    12,
    	Ungraduated_year:       2564,
    	ManageCourseTime:    time.Date(2021, 11, 25, 12, 00, 00, 0000000, time.FixedZone("UTC+7",7*3600 )),
        Course: Course1,
        TA: B,
        Room:  room2,
		Professor: chanwit,	
	}
	db.Model(&ManageCourse{}).Create(&manage2)

}
