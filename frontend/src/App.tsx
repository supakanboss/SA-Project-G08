import CreateBorrow_Sport_Equipment from './component/CreateBorrow_Sport_Equipment';
import Login from './component/Login'
import Home from './component/Home'
import HomeMember from './component/HomeMember'


import { Router, Routes, Route } from "react-router-dom"
import React, { useState, useEffect } from "react"

function App() {

  return (
    <Routes>
      
      <Route path="/" element={<Home />}> </Route>
      <Route path="/Login" element={<Login />}> </Route>
      <Route path="/HomeMember" element={<HomeMember />}> </Route>
      <Route path="/CreateBorrow_Sport_Equipment" element={<CreateBorrow_Sport_Equipment />} /> 
      


    </Routes>
  )
}
export default App;