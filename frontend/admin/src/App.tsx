import { BrowserRouter } from "react-router-dom";
import './App.css';
import Dashboard from './pages/dashboard/Dashboard';
import { useEffect } from "react";

function App() {
  useEffect(() => {
    var currentTheme = document.documentElement.getAttribute("data-theme");
    var targetTheme = "light";

    if (currentTheme === "light") {
      targetTheme = "dark";
    }
    document.documentElement.setAttribute('data-theme', targetTheme)
    localStorage.setItem('theme', targetTheme);
  }, [])
  return (<BrowserRouter>
    <Dashboard />
  </BrowserRouter>)
}

export default App
