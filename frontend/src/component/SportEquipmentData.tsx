import React, { useState, useEffect } from "react";
import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import Container from "@mui/material/Container";
import Paper from "@mui/material/Paper";
import Grid from "@mui/material/Grid";
import { SportEquipmentInterface } from "../models/ISport_Equipment";
import { DataGrid, GridColDef } from "@mui/x-data-grid";
import { Link as RouterLink } from "react-router-dom";
import { ThemeProvider, createTheme } from '@mui/material/styles';
import Home from "../component/Home";
import MemberBar from "../component/MemberBar";
const Theme = createTheme({
  palette: {
    primary: {
      main: '#323232',
    },
    secondary: {
      main: '#FF8B8B',
    },
  },
});


function SportEquipmentData() {
  const [sportquipments, setSportEquipments] = React.useState<SportEquipmentInterface[]>([]);

  const apiUrl = "http://localhost:8080";
  const requestOptionsGet = {
    method: "GET",
    headers: { "Content-Type": "application/json" },};

  const getSportEquipments = async () => {
    const apiUrl = "http://localhost:8080";
    const requestOptions = {
      method: "GET",
      headers: { "Content-Type": "application/json" },
    };

    // fetch ผิด เพราะไม่ใส่ path
    fetch(`${apiUrl}/sport_equipment_data`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        console.log(res.data);
        if (res.data) {
          setSportEquipments(res.data);
        }
      });


  };

  /////////////////// จัดคอลัมน์รายการบันทึก //////////////////////

  const columns: GridColDef[] = [
    { field: "ID", headerName: "No.", width: 20 },
    { field: "StaffID", headerName: "Staff ID", width: 220 },
    { field: "SportTypeID", headerName: "SportType ID", width: 220 },
    {
      field: "SportEquipmentTypeID",
      headerName: "Sport Equipment TypeID",
      width: 220,
    },
    {
      field: "Sport_Equipment_Name",
      headerName: "Sport Equipment Name",
      width: 220,
    },
    {
      field: "Sport_Equipment_Amount",
      headerName: "Sport Equipment Amount",
      width: 220,
    },
  ];

  const fetchSportEquipments = async () => {
    fetch(`${apiUrl}/sport_equipment_data`, requestOptionsGet)
      .then((response) => response.json())
      .then((result) => {
        console.log(result.data);
        setSportEquipments(result.data);
      });
  };

  useEffect(() => {
    getSportEquipments();
    fetchSportEquipments();
  }, []);

  const [token, setToken] = useState<String>("");
  useEffect(() => {
    const token = localStorage.getItem("token");
    if (token) {
      setToken(token);
    }
  }, []);

  if (!token) {

    return <Home />;
  }

  return (
    <div>
      <ThemeProvider theme={Theme}>
      <MemberBar/>
    <div />

      <Container maxWidth="lg">
        <Box display="flex" sx={{marginRight:0 ,margin:2, marginX:126.3}}>  
          <Box>
            <Button  
              sx={{ float: 'right' }}
              component={RouterLink}
              to="/create_sport_equipment"
              variant="contained"
              color="secondary"    
            >
            CreateData
            </Button>  
          </Box> 
        </Box>

        <Paper>
          <Grid item xs={12}>
            <Box
              sx={{
                display: "flex",
                flexWrap: "wrap",
                "& > :not(style)": {
                  m: 1,
                  width: 1,
                  height: 50,
                },
              }}
            >
              <h3>รายการบันทึกข้อมูลอุปกรณ์กีฬา</h3>
            </Box>

            <hr />

            <Box>
              <div style={{ height: 400, width: "100%", marginTop: "20px" }}>
                <DataGrid
                  rows={sportquipments}
                  getRowId={(row) => row.ID}
                  columns={columns}
                  pageSize={5}
                  rowsPerPageOptions={[5]}
                />
              </div>
            </Box>
          </Grid>
        </Paper>
      </Container>
      </ThemeProvider>
    </div>
  );
}

export default SportEquipmentData;
