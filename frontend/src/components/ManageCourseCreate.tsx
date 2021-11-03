import * as React from 'react';
import { useEffect, useState } from "react";
import { Link as RouterLink } from "react-router-dom";
import {
  makeStyles,
  Theme,
  createStyles,
  alpha,
} from "@material-ui/core/styles";
import Button from "@material-ui/core/Button";
import TextField from '@material-ui/core/TextField';
import FormControl from "@material-ui/core/FormControl";
import Container from "@material-ui/core/Container";
import Paper from "@material-ui/core/Paper";
import Grid from "@material-ui/core/Grid";
import Box from "@material-ui/core/Box";
import Typography from "@material-ui/core/Typography";
import Divider from "@material-ui/core/Divider";
import Snackbar from "@material-ui/core/Snackbar";
import Select from "@material-ui/core/Select";
import MuiAlert, { AlertProps } from "@material-ui/lab/Alert";
import InputLabel from '@mui/material/InputLabel';
import MenuItem from '@mui/material/MenuItem';


import { ProfessorsInterface } from "../models/IProfessor";
import { CoursesInterface } from "../models/ICourse";
import { TAsInterface } from "../models/ITA";
import { RoomsInterface } from "../models/IRoom";
import { ManageCourseInterface } from "../models/IManageCourse";

import {
  MuiPickersUtilsProvider,
  KeyboardDateTimePicker,
} from "@material-ui/pickers";
import DateFnsUtils from "@date-io/date-fns";
import { getgroups } from "process";
import { getThemeProps } from "@material-ui/styles";

const Alert = (props: AlertProps) => {
  return <MuiAlert elevation={6} variant="filled" {...props} />;
};

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      flexGrow: 1,
    },
    container: {
      marginTop: theme.spacing(2),
    },
    paper: {
      padding: theme.spacing(2),
      color: theme.palette.text.secondary,
    },
  })
);
function ManageCourseCreate() {
  const classes = useStyles();

  const [group, setGroup] = React.useState(Number);
  const [term, setTerm] = React.useState(Number);
  const [teachingTime, setTeachingTime] = React.useState(Number);
  const [ungraduated_year, setUngraduated_year] = React.useState(Number);
  const [selectedDate, setSelectedDate] = useState<Date | null>(new Date());

  const [professers, setProfessers] = useState<ProfessorsInterface[]>([]);
  const [courses, setCourses] = useState<CoursesInterface[]>([]);
  const [tas, setTAs] = useState<TAsInterface[]>([]);
  const [rooms, setRooms] = useState<RoomsInterface[]>([]);
  const [manageCourse, setManageCourse] = useState<Partial<ManageCourseInterface>>(
    {}    
  );
  const [show, setShow] = useState(false);
  const handleClose = () => setShow(false);
  const handleShow = () => setShow(true);
  const [success, setSuccess] = useState(false);
  const [error, setError] = useState(false);

  const apiUrl = "http://localhost:8080";
  const requestOptions = {
    method: "GET",
    headers: { "Content-Type": "application/json" },
  };
  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: unknown }>
  ) => {
    const name = event.target.name as keyof typeof manageCourse;
    setManageCourse({
      ...manageCourse,
      [name]: event.target.value,
    });
  };
  const handleDateChange = (date: Date | null) => {
    console.log(date);
    setSelectedDate(date);
  };
  const handleInputChange = (

    event: React.ChangeEvent<{ id?: string; value: any }>
 
  ) => {
 
    const id = event.target.id as keyof typeof ManageCourseCreate;
 
    const { value } = event.target;
 
    setManageCourse({ ...manageCourse, [id]: value });
 
  };
  const handleGroupChange = (event:any) => {
    setGroup(event.target.value as number);
  };
  const handleTermChange = (event: any) => {
    setTerm(event.target.value as number);
  };
  const handleTeachingTimeChange = (event: any) => {
    setTeachingTime(event.target.value as number);
  };
  const handleUngraduated_yearChange = (event: any) => {
    setUngraduated_year(event.target.value as number);
  };


  const getProfessers = async () => {
    fetch(`${apiUrl}/professers`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setProfessers(res.data);
        } else {
          console.log("else");
        }
      });
  };

  const getCourses = async () => {
    fetch(`${apiUrl}/courses`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setCourses(res.data);
        } else {
          console.log("else");
        }
      });
  };

  const getTAs = async () => {
    fetch(`${apiUrl}/tas`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setTAs(res.data);
        } else {
          console.log("else");
        }
      });
  };

  const getRooms = async () => {
    fetch(`${apiUrl}/rooms`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setRooms(res.data);
        } else {
          console.log("else");
        }
      });
  };
  useEffect(() => {
    getProfessers();
    getCourses();
    getTAs();
    getRooms();
    
  }, []);

  const convertType = (data: string | number | undefined) => {
    let val = typeof data === "string" ? parseInt(data) : data;
    return val;
  };

  function submit() {
    let data = {
      CourseID: convertType(manageCourse.CourseID),
      TAID: convertType(manageCourse.TAID),
      RoomID: convertType(manageCourse.RoomID),
      Group:  group,
	    Term:		term,
      TeachingTime:   teachingTime, 
      Ungraduated_year:    ungraduated_year,
      ManageCourseTime: selectedDate,
    };

    const requestOptionsPost = {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(data),
    };

    fetch(`${apiUrl}/manageCourses`, requestOptionsPost)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setSuccess(true);
        } else {
          setError(true);
        }
      });
  }
  return (
    <Container className={classes.container} maxWidth="md">
      <Snackbar open={success} autoHideDuration={6000} onClose={handleClose}>
        <Alert onClose={handleClose} severity="success">
          บันทึกข้อมูลสำเร็จ
        </Alert>
      </Snackbar>
      <Snackbar open={error} autoHideDuration={6000} onClose={handleClose}>
        <Alert onClose={handleClose} severity="error">
          บันทึกข้อมูลไม่สำเร็จ
        </Alert>
      </Snackbar>
      <Paper className={classes.paper}>
        <Box display="flex">
          <Box flexGrow={1}>
            <Typography
              component="h2"
              variant="h6"
              color="primary"
              gutterBottom
            >
              บันทึกข้อมูล
            </Typography>
          </Box>
        </Box>
        <Divider />
        <Grid container spacing={2} className={classes.root}>
          <Grid item xs={7}>
            <FormControl fullWidth variant="outlined">
              <p>วิชาเรียน</p>
              <Select
                native
                value={manageCourse.CourseID}
                onChange={handleChange}
                inputProps={{
                  name: "CourseID",
                }}
              >
                <option aria-label="None" value="">
                  กรุณาเลือกวิชาเรียน
                </option>
                {courses.map((item: CoursesInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Name}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>
          <Grid item xs={7}>
            <FormControl fullWidth variant="outlined">
              <p>ผู้ช่วยสอน</p>
              <Select
                native
                value={manageCourse.TAID}
                onChange={handleChange}
                inputProps={{
                  name: "TAID",
                }}
              >
                <option aria-label="None" value="">
                กรุณาเลือกผู้ช่วยสอน
                </option>
                {tas.map((item: TAsInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Name}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>
          <Grid item xs={7}>
            <FormControl fullWidth variant="outlined">
              <p>ห้องเรียน</p>
              <Select
                native
                value={manageCourse.TAID}
                onChange={handleChange}
                inputProps={{
                  name: "RoomID",
                }}
              >
                <option aria-label="None" value="">
                กรุณาเลือกห้องเรียน
                </option>
                {rooms.map((item: RoomsInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Number}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>
          

          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>วันที่และเวลาสอน</p>
              <MuiPickersUtilsProvider utils={DateFnsUtils}>
                <KeyboardDateTimePicker
                  name="ManageCoursesTime"
                  value={selectedDate}
                  onChange={handleDateChange}
                  label="กรุณาเลือกวันที่และเวลาสอน"
                  minDate={new Date("2018-01-01T00:00")}
                  format="yyyy/MM/dd hh:mm a"
                />
              </MuiPickersUtilsProvider>
            </FormControl>
          </Grid>
          <Grid item xs={12}>
            <Button
              component={RouterLink}
              to="/manageCourses"
              variant="contained"
            >
              กลับ
            </Button>
            <Button
              style={{ float: "right" }}
              variant="contained"
              onClick={submit}
              color="primary"
            >
              บันทึก
            </Button>
          </Grid>
        </Grid>
      </Paper>
    </Container>
  );
}

export default ManageCourseCreate;
