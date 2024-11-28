import Auth from "./pages/auth/Auth.tsx"
import Home from "./pages/home/Home.tsx"
import Game from "./pages/game/Game.tsx"
import { ToastContainer } from "react-toastify"
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';

function App() {
  return (
    <>
        <Router>
          <Routes>  
            <Route path="/Login" element={<Auth isLogin={true} />}></Route>
            <Route path="/Register" element={<Auth isLogin={false}/>}></Route>
            <Route>
              <Route path="/Home" element={<Home />}></Route>
              <Route path="/Game/:gameId" element={<Game />}></Route>
            </Route>
          </Routes>
        <ToastContainer />
      </Router>
    </>
  )
}

export default App