import { CoursesInterface } from "./ICourse";
import { ProfessorsInterface } from "./IProfessor";
import { TAsInterface } from "./ITA";
import { RoomsInterface } from "./IRoom";

export interface ManageCourseInterface {
    ID: string,
    Group:			string,
    Term:			number,
    TeachingTime:    number,
    Ungraduated_year:    number,
    ManageCourseTime:    Date,

    ProfessorID:    number,
    Professor:      ProfessorsInterface,

    CourseID:   number,
    Course:   CoursesInterface,

    TAID:   number,
    TA:     TAsInterface,

    RoomID:     number,
    Room:       RoomsInterface,

}