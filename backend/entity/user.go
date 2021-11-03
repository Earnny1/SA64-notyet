package entity
import(
    "gorm.io/gorm"
    "time"
)

type Professor struct{
    gorm.Model
    Name            string
    Professer_id    string  `gorm: "uniqueIndex"`
    Password        string  
    Major           string
    // 1 teacher เป็นเจ้าของได้หลาย ManageCourse
    ManageCourses    []ManageCourse  `gorm: "foreignKey:ProfessorID"`
}
type Course struct{
    gorm.Model
    Name        string
    Code        string  
    Credit      uint
    ManageCourses    []ManageCourse  `gorm: "foreignKey:CourseID"`
    
}

type TA struct{
    gorm.Model
    Name        string
    TA_id       string  
    ManageCourses    []ManageCourse  `gorm: "foreignKey:TAID"`
    
    
}

type Room struct{
    gorm.Model
    Number          uint  
    StudentCount     uint   
    ManageCourses    []ManageCourse  `gorm: "foreignKey:RoomID"`  
    
}
type ManageCourse struct{
    gorm.Model

	Group			uint
	Term			uint
    TeachingTime    uint
    Ungraduated_year       uint
    ManageCourseTime    time.Time

    ProfessorID *uint
    Professor   Professor   `gorm:"references:id"`

    CourseID *uint
    Course  Course  `gorm:"references:id"`

    TAID *uint
    TA   TA     `gorm:"references:id"`

    RoomID *uint
    Room    Room      `gorm:"references:id"`
}
