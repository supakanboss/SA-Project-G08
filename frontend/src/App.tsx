import { Router, Routes, Route } from "react-router-dom";
import React, { useState, useEffect } from "react";

import CreateSportEquipment from './component/CreateSportEquipment';
import SportEquipmentData from './component/SportEquipmentData';
import Home from './component/Home';
import HomeStaff from "./component/HomeStaff";
import StaffLogin from './component/StaffLogin';


function App() {

  return (
    <Routes>
      <Route path="/" element={<Home />}> </Route>
      <Route path="/HomeStaff" element={<HomeStaff/>}></Route>
      <Route path="/StaffLogin" element={<StaffLogin/>}></Route>
      <Route path="/sport_equipment_data" element={<SportEquipmentData />} />
      <Route path="/create_sport_equipment" element={<CreateSportEquipment />} />      
    </Routes>


  )

}
export default App;