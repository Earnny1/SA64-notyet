import { useEffect, useState } from "react";
import { Link as RouterLink } from "react-router-dom";
import { createStyles, makeStyles, Theme } from "@material-ui/core/styles";
import Typography from "@material-ui/core/Typography";
import Button from "@material-ui/core/Button";
import Container from "@material-ui/core/Container";
import Paper from "@material-ui/core/Paper";
import Box from "@material-ui/core/Box";
import Table from "@material-ui/core/Table";
import TableBody from "@material-ui/core/TableBody";
import TableCell from "@material-ui/core/TableCell";
import TableContainer from "@material-ui/core/TableContainer";
import TableHead from "@material-ui/core/TableHead";
import TableRow from "@material-ui/core/TableRow";
import { ManageCourseInterface } from "../models/IManageCourse";
import { format } from 'date-fns'

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    container: {
      marginTop: theme.spacing(2),
    },
    table: {
      minWidth: 650,
    },
    tableSpace: {
      marginTop: 20,
    },
  })
);

function ManageCourses() {
  const classes = useStyles();
  const [manageCourses, setWatchVideos] = useState<ManageCourseInterface[]>([]);
  const apiUrl = "http://localhost:8080";
  const requestOptions = {
    method: "GET",
    headers: { "Content-Type": "application/json" },
  };

  const getManageCourses = async () => {
    fetch(`${apiUrl}/manageCourses/create`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        console.log(res.data);
        if (res.data) {
          setWatchVideos(res.data);
        } else {
          console.log("else");
        }
      });
  };

  useEffect(() => {
    getManageCourses();
  }, []);

  return (
    <div>
      <Container className={classes.container} maxWidth="md">
        <Box display="flex">
          <Box flexGrow={1}>
            <Typography
              component="h2"
              variant="h6"
              color="primary"
              gutterBottom
            >
              รายละเอียดวิชาเรียน
            </Typography>
          </Box>
          <Box>
            <Button
              component={RouterLink}
              to="/manageCourses/create"
              variant="contained"
              color="primary"
            >
              สร้างข้อมูล
            </Button>
          </Box>
        </Box>
        <TableContainer component={Paper} className={classes.tableSpace}>
          <Table className={classes.table} aria-label="simple table">
            <TableHead>
              <TableRow>
              <TableCell align="center" width="10%">
                  ลำดับ
                </TableCell>
                <TableCell align="center" width="20%">
                  วิชาเรียน
                </TableCell>
                <TableCell align="center" width="20%">
                  ผู้ช่วยสอน
                </TableCell>
                <TableCell align="center" width="20%">
                  ห้องเรียน
                </TableCell>
                <TableCell align="center" width="20%">
                  กลุ่มเรียน
                </TableCell>
                <TableCell align="center" width="20%">
                  ภาคเรียน
                </TableCell>
                <TableCell align="center" width="20%">
                  ปีการศึกษา
                </TableCell>
                <TableCell align="center" width="20%">
                  ระยะเวลาสอน
                </TableCell>
                <TableCell align="center" width="30%">
                  วันที่และเวลาสอน
                </TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {manageCourses.map((item: ManageCourseInterface) => (
                <TableRow key={item.ID}>
                  <TableCell align="center">{item.ID}</TableCell>
                  <TableCell align="center">{item.Course.Name}</TableCell>
                  <TableCell align="center">{item.TA.Name}</TableCell>
                  <TableCell align="center">{item.Room.Number}</TableCell>
                  
                  <TableCell align="center">{item.Group}</TableCell>
                  <TableCell align="center">{item.Term}</TableCell>
                  <TableCell align="center">{item.Ungraduated_year}</TableCell>
                  <TableCell align="center">{item.TeachingTime}</TableCell>
                  <TableCell align="center">{format((new Date(item.ManageCourseTime)), 'dd MMMM yyyy hh:mm a')}</TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TableContainer>
      </Container>
    </div> 
  );
}
export default ManageCourses;