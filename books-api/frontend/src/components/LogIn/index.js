import { useEffect, useState } from "react";
import {UncontrolledAlert} from 'reactstrap';
import Form from "./Form";
import { useNavigate } from "react-router-dom";
import { jwtDecode } from "jwt-decode";

import { useAuth } from "../../AuthContext";

const Login = () => {
  const { setAuthInfo } = useAuth();
  const navigate = useNavigate();
  const [alertState, setAlertState] = useState({
    show: false,
    type: "success",
    message: "",
  })
  useEffect(()=>{
    if(localStorage.getItem('isloggedin')==='true'){
      navigate('/')
    }
  }, []);

  const handleSubmit = (e) => {
    e.preventDefault();
    const email = e.target.email.value;
    const pass = e.target.password.value;
    const url = `http://localhost:8080/login`;
    const options = {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        email: email,
        password: pass,
      }),
    };
    console.log(options);
    fetch(url, options)
      .then((res) => res.json())
      .then((d) => {
        console.log(d);
        if(d.type === "success"){
          const decodedToken = jwtDecode(d.token);
          localStorage.setItem('authToken', d.token);
          // localStorage.setItem('isloggedin', 'true');
          // localStorage.setItem('username', decodedToken.username);
          // localStorage.setItem('access', decodedToken.access);
          console.log(decodedToken);
          setAuthInfo({ isloggedin: 'true', username: decodedToken.username, access: decodedToken.access });
          navigate('/');
        }
        else{
          setAlertState({
            show: true,
            type: "danger",
            message: d.error.message,
          })
        }
      }) // save in local storage or show an error on screen
      .catch((err) => {
        console.log(err)
        setAlertState({
          show: true,
          type: "danger",
          message: "something went wrong, try again later!",
        })
      });
  };
  return (
    <div className="container">
      <div className="form-div">
        <h3>Login!</h3>
        <div className="myform">
          <Form handleSubmit={handleSubmit} />
        </div>
      {alertState.show && (
        <UncontrolledAlert color={alertState.type} fade={true}>
          {alertState.message}
        </UncontrolledAlert>
      )}
      </div>
    </div>
  );
};

export default Login;
