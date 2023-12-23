import { useEffect, useState } from "react";
import { Row, Col } from "reactstrap";

import User from "./User";

const Users = () => {
  const [data, setData] = useState([]); // Use a state variable here

  useEffect(() => {
    const url = `http://localhost:8080/users`;
    const options = {
      method: "GET",
      headers: {
        Authorization: localStorage.getItem("authToken"),
      },
    };
    fetch(url, options)
      .then((res) => res.json())
      .then((d) => {
        if (d.type === "success") {
          setData(d.users); // Update the state here
          console.log(d.users);
        } else {
          setData([]); // Update the state here
          console.log(d);
        }
      })
      .catch((err) => console.log(err));
  }, []);

  const listAllUsers = data.map((d) => {
    return (
      <Col sm="4" xs="12" key={d.id}>
        <li style={{ listStyle: "none" }}>
          <User
            username={d.username}
            id={d.id}
            access={d.access}
            email={d.email}
            issued={d.issued}
          />
        </li>
      </Col>
    );
  });

  return (
    <div className="users container">
      <h3>Users</h3>
      <Row>{listAllUsers}</Row>
    </div>
  );
};

export default Users;
