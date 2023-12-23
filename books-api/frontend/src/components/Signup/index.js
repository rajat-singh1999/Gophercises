import FormCus from "./Form";
import { useNavigate } from "react-router-dom";
import { useEffect, useState } from "react";

import { UncontrolledAlert } from "reactstrap";

const Signup = ()=>{
    const navigate = useNavigate();
    const [alertState, setAlertState] = useState({
      show: false,
      type: "success",
      message: "good news",
    });

    useEffect(()=>{
      if(localStorage.getItem('isloggedin') === 'true'){
        navigate('/');
      }
    }, []);

    const handleSubmit = ({email, username, password})=>{
        const url = `http://localhost:8080/signup`;
        const options = {
          method: 'POST',
          headers: {
            "Content-Type": "application/json",
            Authorization: localStorage.getItem('authToken'),
          },
          body: JSON.stringify({
            email: email,
            username: username,
            password: password,
          }),
        }
        fetch(url, options)
        .then(res=>res.json())
        .then(data=>{
          console.log(data);
          if(data.type === 'success'){
            navigate('/login');
          } else{
            setAlertState({
              type: "danger",
              show: true,
              message: data.message,
            });
          }
        })

    }
    return(
        <div className="container">
      <div className="form-div">
        <h3>Sign Up!</h3>
        <div className="myform">
          <FormCus handleSubmit={handleSubmit} />
          {alertState.show && (
        <UncontrolledAlert color={alertState.type} fade={true}>
          {alertState.message}
        </UncontrolledAlert>
      )}
        </div>
      </div>
    </div>
    );
}

export default Signup;
