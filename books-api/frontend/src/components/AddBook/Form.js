import "bootstrap/dist/css/bootstrap.min.css";

import { useState } from "react";
import {
  Form,
  FormGroup,
  Label,
  Input,
  Button,
  UncontrolledAlert,
} from "reactstrap";

const FormCus = () => {
  const [validationErr, setValidationErr] = useState({
    type: "",
    message: "",
  });
  const [inp, setInp] = useState({
    ISBN: "",
    Title: "",
    Author: "",
    Units: "",
  });

  const handleOnChange = (e) => {
    let name = e.target.name;
    let value = e.target.value;
    setInp((prevInp) => {
      return {
        ...prevInp,
        [name]: value,
      };
    });
  };

  const handleClick = (e) => {
    e.preventDefault();
    if (
      inp.ISBN === "" ||
      inp.Title === "" ||
      inp.Author === "" ||
      inp.Units === ""
    ) {
      setValidationErr((p) => {
        return {
          ...p,
          ["type"]: "danger",
          ["message"]: "Fields cannot be empty!",
        };
      });
    } else if (isNaN(inp.Units)) {
      setValidationErr((p) => {
        return {
          ...p,
          ["type"]: "danger",
          ["message"]: "units field can only be a digit",
        };
      });
    } else {
      // calling some passed function or making an api call to post new data
      let data = {
        id: inp.ISBN,
        title: inp.Title,
        author: inp.Author,
        total: parseInt(inp.Units),
      };
      data = JSON.stringify(data);
      fetch("http://localhost:8080/addbook", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization:
            "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3MiOiJhZG1pbiIsImVtYWlsIjoicnNpbmdoMTczNEBnbWFpbC5jb21AZ21haWwuY29tIiwiaWQiOiI1MjI2In0.bV6_4dwsyUzv8nAhcRzasAGuK1n94s8uxdypbdSo9Gs",
        },
        body: data,
      })
        .then((response) => response.json())
        .then((data) =>{
            if (data.type === "success"){
                setValidationErr((p) => {
                  return {
                    ...p,
                    ["type"]: "success",
                    ["message"]: data.message,
                  };
                })
            } else{
                setValidationErr((p)=>{
                    return{
                        ...p,
                        ["type"]: "danger",
                        ["message"]: data.error.message,
                    }
                })
            }
        })
        .catch((err) => {
          console.error("in catch");
          setValidationErr((p)=>{
            return {
                ...p,
                ["type"]: "danger",
                ["message"]: "something went wrong, check internet connection and try later!!"
            }
          })
        });
    }
  };

  return (
    <div className="addFrom">
      <Form>
        <FormGroup>
          <Label for="ISBN">ISBN</Label>
          <Input
            id="ISBN"
            name="ISBN"
            placeholder="isbn"
            type="text"
            value={inp.ISBN}
            onChange={handleOnChange}
          />
        </FormGroup>
        <FormGroup>
          <Label for="Title">Title</Label>
          <Input
            id="Title"
            name="Title"
            placeholder="Title"
            type="text"
            value={inp.Title}
            onChange={handleOnChange}
          />
        </FormGroup>
        <FormGroup>
          <Label for="Author">Author</Label>
          <Input
            id="Author"
            name="Author"
            placeholder="Author Name"
            type="text"
            value={inp.Author}
            onChange={handleOnChange}
          />
          <Label for="Units">Number of units</Label>
          <Input
            id="Units"
            name="Units"
            placeholder="Number of units"
            type="text"
            value={inp.Units}
            onChange={handleOnChange}
          />
        </FormGroup>
        <Button onClick={handleClick}>Add</Button>
      </Form>
      {validationErr.type === "danger" && (
        <>
          <br />
          <UncontrolledAlert color="danger" fade={true}>
            {validationErr.message}
          </UncontrolledAlert>
        </>
      )}
      {validationErr.type === "success" && (
        <>
          <br />
          <UncontrolledAlert color="success" fade={true}>
            {validationErr.message}
          </UncontrolledAlert>
        </>
      )}
    </div>
  );
};

export default FormCus;
