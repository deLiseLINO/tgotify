import React, { useState, useEffect } from "react";
import AuthorizationPage from "../AuthorizationPage/AuthorizationPage";
import { Routes, Route, Link } from "react-router-dom";
import MainPage from "../MainPage/MainPage";
import NotFoundPage from "../NotFoundPage/NotFoundPage";
import { jwtDecode } from "jwt-decode";

function App() {
  const [isAuth, setIsAuth] = useState(false);

  const getAuth = (newPost) => {
    setIsAuth(newPost);
  };

  useEffect(() => {
    let token = localStorage.getItem("token");
    if (token) {
      const decoded = jwtDecode(token);

      const exp = decoded.exp
      const now = new Date().getTime() / 1000

      if (exp - now >= 0){
        setIsAuth(true);
      } else {
        setIsAuth(false);
      }

    }
  }, []);

  return (
    <Routes>
      <Route path="/*" element={<NotFoundPage />} />

      {isAuth ? (
        <Route path="/" element={<MainPage />} />
      ) : (
        <Route path="/" element={<AuthorizationPage isAuth={getAuth} />} />
      )}
    </Routes>
  );
}

export default App;
