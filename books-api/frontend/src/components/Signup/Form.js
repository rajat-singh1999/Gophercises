import { Form, FormGroup, Label, Input, Button, UncontrolledAlert } from "reactstrap";
import { useState } from "react";

const FormCus = ({handleSubmit}) => {
  const [inp, setInp] = useState({
    email: "",
    password: "",
    username: "",
    cpassword: "",
  });
  const [alertState, setAlertState] = useState({
    show: false,
    type: "success",
    message: "good news",
  });
  const handleInputChange = (e) => {
    setInp((prev) => {
      let name = e.target.name;
      let val = e.target.value;
      return { ...prev, [name]: val };
    });
  };

  const [showPass, setShowPass] = useState(false);
  const validateInps = () => {
    if(inp.email!=='' && inp.email!==' ' && inp.password!=='' && inp.password!==' ' && inp.cpassword!=='' && inp.cpassword!==' ' && inp.username!=='' && inp.username!==' '){
      if (inp.password !== inp.cpassword) {
        setAlertState({
          show: true,
          type: "danger",
          message: "passwords dont match",
        });
        return false;
      } else if (inp.password.length < 4) {
        setAlertState({
          show: true,
          type: "danger",
          message: "password cannot be less than 4 characters long",
        });
        return false;
      } else {
        setAlertState({
          show: false,
          type: "success",
          message: "",
        });
        return true;
      }
    } else{
      setAlertState({
        show: true,
        type: "danger",
        message: "no field should be empty!",
      });
      return false;
    }
    
  };

  const validateAndSubmit = (e)=>{
    e.preventDefault();
    if(validateInps()){
      handleSubmit(inp);
    }
}


  return (
    <div>
      <Form onSubmit={validateAndSubmit}>
        <FormGroup>
          <Label for="email">Email</Label>
          <Input
            id="email"
            name="email"
            placeholder="Enter your email..."
            value={inp.email}
            onChange={handleInputChange}
            type="email"
          />
          <Label for="username">Username</Label>
          <Input
            id="username"
            name="username"
            placeholder="Enter a username..."
            value={inp.username}
            onChange={handleInputChange}
            type={"text"}
          />
          <Label for="password">Password</Label>
          <Input
            id="password"
            name="password"
            placeholder="Enter Password..."
            value={inp.password}
            onChange={handleInputChange}
            type={showPass && "password"}
          />
          <Label for="cpassword">Confirm Password</Label>
          <Input
            id="cpassword"
            name="cpassword"
            placeholder="confirm password..."
            value={inp.cpassword}
            onChange={handleInputChange}
            type={showPass && "password"}
          />
          <Button
            color="link"
            size="sm"
            onClick={() => {
              setShowPass(!showPass);
            }}
          >
            show
          </Button>
        </FormGroup>
        <Button type="submit">Sign Up</Button>
      </Form>
      {alertState.show && (
        <UncontrolledAlert color={alertState.type} fade={true}>
          {alertState.message}
        </UncontrolledAlert>
      )}
    </div>
  );
};

export default FormCus;
