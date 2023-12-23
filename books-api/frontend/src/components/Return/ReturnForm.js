import { useEffect, useState } from "react";
import {
  Form,
  FormGroup,
  Input,
  Label,
  Button,
  UncontrolledAlert,
} from "reactstrap";

const ReturnForm = () => {
  const [inp, setInp] = useState({ ISBN: "", userid: "" });
  const [alertState, setAlertState] = useState({
    show: false,
    type: "success",
    message: "good news",
  });

  useEffect(() => {
    console.log(alertState); // Log the state here
  }, [alertState]); // Add alertState as a dependency

  const handleInputChange = (e) => {
    let name = e.target.name;
    let val = e.target.value;

    setInp({
      ...inp,
      [name]: val,
    });
  };

  const validateInput = () => {
    if (
      isNaN(inp.ISBN) ||
      isNaN(inp.userid) ||
      inp.ISBN === "" ||
      inp.userid === ""
    )
      return false;
    else return true;
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    setAlertState({
      show: false,
      type: "success",
      message: "good news",
    });
    let ok = validateInput();
    if (!ok) {
      setAlertState({
        show: true,
        type: "danger",
        message: "Inputs must be a number.",
      });
    } else {
      const url = `http://localhost:8080/return/${inp.userid}?bid=${inp.ISBN}`;
      const options = {
        method: "GET",
        headers: {
          Authorization: localStorage.getItem("authToken"),
        },
      };
      fetch(url, options)
        .then((res) => res.json())
        .then((data) => {
          console.log(data);
          if (data.type === "error") {
            setAlertState({
              show: true,
              type: "danger",
              message: data.Error,
            });
          } else {
            setAlertState({
              show: true,
              type: "success",
              message: data.message,
            });
          }
        })
        .catch((err) => {
          console.log(err);
          if (err.message === null) {
            setAlertState({
              show: true,
              type: "danger",
              message: err.Error,
            });
          } else {
            setAlertState({
              show: true,
              type: "danger",
              message: "Sorry, something went wrong!",
            });
          }
        });
    }
  };

  return (
    <div>
      <div className="returnForm">
        <Form onSubmit={handleSubmit}>
          <FormGroup>
            <Label for="ISBN">ISBN</Label>
            <Input
              id="ISBN"
              name="ISBN"
              placeholder="isbn"
              value={inp.ISBN}
              onChange={handleInputChange}
            />
            <Label for="userid">UserID</Label>
            <Input
              id="userid"
              name="userid"
              placeholder="user id"
              value={inp.userid}
              onChange={handleInputChange}
            />
          </FormGroup>
          <Button type="submit">Return</Button>
        </Form>
        <br />
      </div>
      {alertState.show && (
        <UncontrolledAlert color={alertState.type} fade={true}>
          {alertState.message}
        </UncontrolledAlert>
      )}
    </div>
  );
};

export default ReturnForm;
