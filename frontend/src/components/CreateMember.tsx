import React, { useEffect, useState } from 'react';
import AppBar from '@mui/material/AppBar';
import Box from '@mui/material/Box';
import Toolbar from '@mui/material/Toolbar';
import Typography from '@mui/material/Typography';
import Button from '@mui/material/Button';
import IconButton from '@mui/material/IconButton';
import MenuIcon from '@mui/icons-material/Menu';
import CssBaseline from '@mui/material/CssBaseline';
import Container from '@mui/material/Container';
import Paper from '@mui/material/Paper';
import Grid from '@mui/material/Grid';
import { styled } from '@mui/material/styles';
import TextField from '@mui/material/TextField';
import { LocalizationProvider } from '@mui/x-date-pickers/LocalizationProvider';
import { DatePicker } from '@mui/x-date-pickers/DatePicker';
import { AdapterDateFns } from '@mui/x-date-pickers/AdapterDateFns';
import { FormControl } from '@mui/material';
import Select, { SelectChangeEvent } from "@mui/material/Select";
import { PackageInterface } from '../Models/IPackage';
import { ProvinceInterface } from '../Models/IProvince';
import { GenderInterface } from '../Models/IGender';
import { CreateMemberInterface } from '../Models/ICreateMember';
import Snackbar from "@mui/material/Snackbar";
import MuiAlert, { AlertProps } from "@mui/material/Alert";

const Alert = React.forwardRef<HTMLDivElement, AlertProps>(function Alert(
  props,
  ref
) {
  return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
});


function App() {
  const Item = styled(Paper)(({ theme }) => ({
    backgroundColor: theme.palette.mode === 'dark' ? '#1A2027' : '#fff',
    ...theme.typography.body2,
    padding: theme.spacing(1),
    textAlign: 'center',
    color: theme.palette.text.secondary,
  }));

  const [birthday, setBirthday] = React.useState<Date | null>(null);
  // combo box
  const [gender, setGender] = useState<GenderInterface[]>([]);
  const [province, setProvince] = useState<ProvinceInterface[]>([]);
  const [packages, setPackage] = useState<PackageInterface[]>([]);
  const [CreateMember, setCreateMember] = useState<Partial<CreateMemberInterface>>({});
  
  const apiUrl = "http://localhost:8080";

  const [success, setSuccess] = useState(false);
  const [error, setError] = useState(false);

  const requestOpionsGet = {
    method: "GET",
    headers: { "Content-Type": "application/json" },
  };

  const feachGender = async () => {
    fetch(`${apiUrl}/gender`, requestOpionsGet)
      .then((response) => response.json())
      .then((result) => {
        console.log(result.data);
        setGender(result.data);
      });
  };

  const feachPackage = async () => {
    fetch(`${apiUrl}/package`, requestOpionsGet)
      .then((response) => response.json())
      .then((result) => {
        console.log(result.data);
        setPackage(result.data);
      });
  };

  const feachProvince = async () => {
    fetch(`${apiUrl}/province`, requestOpionsGet)
      .then((response) => response.json())
      .then((result) => {
        console.log(result.data);
        setProvince(result.data);
      });
  };

  const handleInputChange = (
    event: React.ChangeEvent<{ id?: string; value: any }>
  ) => {
    const id = event.target.id as keyof typeof CreateMember;
    const { value } = event.target;
    setCreateMember({ ...CreateMember, [id]: value });
  };

  const handleChange = (event: SelectChangeEvent) => {
    const name = event.target.name as keyof typeof CreateMember
    setCreateMember({
      ...CreateMember,
      [name]: event.target.value,
    });
  };

  const handleClose = (
    event?: React.SyntheticEvent | Event,
    reason?: string
  ) => {
    if (reason === "clickaway") {
      return;
    }
    setSuccess(false);
    setError(false);
  };

  //เรียกใช้ feach//
  useEffect(() => {
    feachGender();
    feachPackage();
    feachProvince();
  }, []);

  console.log(CreateMember)

  const convertType = (data: string | number | undefined) => {
    let val = typeof data === "string" ? parseInt(data) : data;
    return val;
  };

  function submit() {
    let data = {
      Email: CreateMember.Email,
      Password: CreateMember.Password,
      Member_NAME: CreateMember.Member_NAME,
      Age: convertType (CreateMember.Age),
      Weight: convertType (CreateMember.Weight),
      Height: convertType (CreateMember.Height),
      Tel: CreateMember.Tel,
      DOB: birthday,
      GenderID: convertType (CreateMember.GenderID),
      ProvinceID: convertType (CreateMember.ProvinceID),
      PackageID: convertType (CreateMember.PackageID),
    }

    console.log("data",data);

    const requestOptions = {
      method: "POST",
      headers: {
        Authorization: `Bearer ${localStorage.getItem("token")}`,
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    };

    fetch(`${apiUrl}/createmember`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        console.log(res);
        if (res.data) {
          setSuccess(true);
        } else {
          setError(true);
        }
      });   
  }

  return (
    <div>
      <Box sx={{ flexGrow: 1 }}>
        <AppBar position="static">
          <Toolbar>
            <IconButton
              size="large"
              edge="start"
              color="inherit"
              aria-label="menu"
              sx={{ mr: 2 }}
            >
              <MenuIcon />
            </IconButton>

            <Typography variant="h6" component="div" sx={{ flexGrow: 1 }}>
              SA G08 สถานกีฬา
            </Typography>

            <Button color="inherit">
              Login
            </Button>

          </Toolbar>
        </AppBar>
      </Box>

      <Container maxWidth="lg">
        <Box
          sx={{
            mt: 2,
          }}
          >
          <Snackbar
            open={success}
            autoHideDuration={3000}
            onClose={handleClose}
            anchorOrigin={{ vertical: "top", horizontal: "center" }}
          >
            <Alert onClose={handleClose} severity="success">
              บันทึกข้อมูลสำเร็จ
            </Alert>
          </Snackbar>
          <Snackbar
            open={error}
            autoHideDuration={6000}
            onClose={handleClose}
            anchorOrigin={{ vertical: "top", horizontal: "center" }}
          >
            <Alert onClose={handleClose} severity="error">
              บันทึกข้อมูลไม่สำเร็จ
            </Alert>
          </Snackbar>
        </Box>

        <Paper>
          <Box
            display={"flex"}
            sx={{
              marginTop: 2,
              padding: 2,
            }}
          >
            <h2>สมัครสมาชิก</h2>
          </Box>
          <hr />

          <Grid container spacing={2} sx={{ marginX: 5, marginY: 1, padding: 2, }}>

            {/* e-mail */}
            <Grid item xs={6}>
              <p>E-mail</p>
              <TextField
                fullWidth
                id="Email"
                type="string"
                label="E-mail"
                variant="outlined"
                name="Email"
                value={CreateMember.Email}
                onChange={handleInputChange}
              />
            </Grid>

            <Grid item xs={6}>
              {/*<Item>xs=4</Item>*/}
            </Grid>

            {/* password */}
            <Grid item xs={6}>
              <p>Password</p>
              <TextField
                fullWidth
                id="Password"
                type="password"
                label="Password"
                variant="outlined" 
                name="Password"
                value={CreateMember.Password}
                onChange={handleInputChange}
                />
            </Grid>

            <Grid item xs={6}>
              {/*<Item>xs=8</Item>*/}
            </Grid>

            {/* ชื่อ-สกุล */}
            <Grid item xs={6}>
              <p>ชื่อ-สกุล</p>
              <TextField
                fullWidth
                id="Member_NAME"
                type="string"
                label="ชื่อ-สกุล"
                variant="outlined" 
                name="Member_NAME"
                value={CreateMember.Member_NAME}
                onChange={handleInputChange}
                />
            </Grid>

            <Grid item xs={6}>
              {/*<Item>xs=8</Item>*/}
            </Grid>

            {/* อายุ */}
            <Grid item xs={3}>
              <p>อายุ</p>
              <TextField
                fullWidth
                id="Age"
                type="number"
                InputProps={{ inputProps: { min: 1 } }}
                label="อายุ"
                variant="outlined" 
                name="Age"
                value={CreateMember.Age }
                onChange={handleInputChange}
                />
            </Grid>

            {/* วันเกิด */}
            <Grid item xs={3}>
              <FormControl fullWidth variant="outlined">
                <p>วันเกิด</p>
                <LocalizationProvider dateAdapter={AdapterDateFns}>
                  <DatePicker
                    label="วันเกิด"
                    value={birthday}
                    onChange={(newValue) => {
                      console.log(newValue)
                      setBirthday(newValue);
                    }}
                    renderInput={(params) => <TextField {...params} />}
                  />
                </LocalizationProvider>
              </FormControl>
            </Grid>

            <Grid item xs={6}>
              {/*<Item>xs=8</Item>*/}
            </Grid>

            {/* เพศ */}
            <Grid item xs={3}>
              <p>เพศ</p>
              <Select fullWidth id="Gender"
                onChange={handleChange}
                native
                value={CreateMember.GenderID + ""}
                inputProps={{ name : "GenderID"}}>
                <option aria-label="None" value=""
                >
                  กรุณาเลือกเพศ
                </option>
                {gender.map((item) => (
                  <option key={item.ID} value={item.ID} >
                    {item.Gender_Type}
                  </option>
                ))}
              </Select>
            </Grid>

            <Grid item xs={9}>
              {/*<Item>xs=8</Item>*/}
            </Grid>

            {/* น้ำหนัก */}
            <Grid item xs={3}>
              <p>น้ำหนัก</p>
              <TextField
                fullWidth
                id="Weight"
                type="number"
                InputProps={{ inputProps: { min: 1 } }}
                label="น้ำหนัก"
                variant="outlined" 
                name="Weight"
                value={CreateMember.Weight}
                onChange={handleInputChange}
                />
            </Grid>

            {/* ส่วนสูง */}
            <Grid item xs={3}>
              <p>ส่วนสูง</p>
              <TextField
                fullWidth
                id="Height"
                type="number"
                InputProps={{ inputProps: { min: 1 } }}
                label="ส่วนสูง"
                variant="outlined" 
                name="Height"
                value={CreateMember.Height}
                onChange={handleInputChange}
                />
            </Grid>

            <Grid item xs={6}>
              {/*<Item>xs=8</Item>*/}
            </Grid>

            {/* จังหวัด */}
            <Grid item xs={4}>
              <p>จังหวัด</p>
              <Select fullWidth id="Province"
                onChange={handleChange}
                native
                value={CreateMember.ProvinceID + ""}
                inputProps={{ name : "ProvinceID"}}
                >
                <option aria-label="None" value="">
                  กรุณาเลือกจังหวัด
                </option>
                {province.map((item) => (
                  <option key={item.ID} value={item.ID} >
                    {item.Province_Type}
                  </option>
                ))}
              </Select>
            </Grid>

            <Grid item xs={8}>
              {/*<Item>xs=8</Item>*/}
            </Grid>

            {/* เบอร์โทรศัพท์ */}
            <Grid item xs={4}>
              <p>เบอร์โทรศัพท์</p>
              <TextField
                fullWidth
                id="Tel"
                type="string"
                label="เบอร์โทรศัพท์"
                variant="outlined" 
                name="Tel"
                value={CreateMember.Tel}
                onChange={handleInputChange}
                />
            </Grid>

            <Grid item xs={8}>
              {/*<Item>xs=8</Item>*/}
            </Grid>

            {/* package */}
            <Grid item xs={4}>
              <p>Package</p>
              <Select fullWidth id="Package"
                onChange={handleChange}
                native
                value={CreateMember.PackageID + ""}
                inputProps={{ name : "PackageID"}}
                >
                <option aria-label="None" value="">
                  กรุณาเลือกPackage
                </option>
                {packages.map((item) => (
                  <option key={item.ID} value={item.ID} >
                    {item.Package_Type}
                  </option>
                ))}
              </Select>
            </Grid>

            <Grid item xs={8}>
              {/*<Item>xs=8</Item>*/}
            </Grid>

            <Grid item xs={12} sx={{ marginY: 1, }}>
              <Button 
              variant="contained" 
              color="info">
                BACK
              </Button>

              <Button 
              onClick={submit}
              variant="contained" 
              color="success" 
              sx={{ marginX: 23, }}>
                SUBMIT
              </Button>
            </Grid>

          </Grid>

        </Paper>
      </Container>

    </div>
  );
}

export default App;
