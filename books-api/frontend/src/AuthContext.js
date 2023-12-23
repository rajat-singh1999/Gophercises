import React, { createContext, useContext, useState } from 'react';

export const AuthContext = createContext();

export const AuthProvider = ({ children }) => {
  const [authState, setAuthState] = useState({
    isloggedin: localStorage.getItem('isloggedin'),
    username: localStorage.getItem('username'),
    access: localStorage.getItem('access'),
  });

  const setAuthInfo = ({ isloggedin, username, access }) => {
    localStorage.setItem('isloggedin', isloggedin);
    localStorage.setItem('username', username);
    localStorage.setItem('access', access);
    setAuthState({ isloggedin, username, access });
  };

  return (
    <AuthContext.Provider value={{ authState, setAuthInfo }}>
      {children}
    </AuthContext.Provider>
  );
};

export const useAuth = () => useContext(AuthContext);
