import { Form, FormGroup, Label, Input, Button } from 'reactstrap';
import { useState } from 'react';
const FormCus = ({ handleSubmit }) => {
    const [inp, setInp] = useState({
        email: "",
        password: "",
    });
    const [showPass, setShowPass] = useState(false);

    const handleInputChange = (e)=>{
        setInp((prev)=>{
            let name = e.target.name;
            let val = e.target.value;

            return {...prev, [name]: val};
        })
    }
  return (
    <div>
      <Form onSubmit={handleSubmit}>
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
          <Label for="password">Password</Label>
          <Input
            id="password"
            name="password"
            placeholder="Enter Password..."
            value={inp.password}
            onChange={handleInputChange}
            type={showPass && "password"}
          />
          <Button color="link" size="sm" onClick={()=>{
            setShowPass(!showPass);
          }}>show</Button>
        </FormGroup>
        <Button type="submit">Log In</Button>
      </Form>
    </div>
  );
};

export default FormCus;
