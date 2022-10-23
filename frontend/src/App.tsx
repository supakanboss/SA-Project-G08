import CreateReport from './component/CreateReport'
import Home from './component/Home'
import { Router, Routes, Route } from "react-router-dom"
import React, { useState, useEffect } from "react"

function App() {

  return (
    <Routes>
      
      <Route path="/" element={<Home />}> </Route>
      <Route path="/CreateReport" element={<CreateReport />}> </Route>

    </Routes>
  )
}
export default App;