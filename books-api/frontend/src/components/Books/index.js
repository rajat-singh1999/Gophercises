import React, { useEffect, useState } from "react";
import { useNavigate } from 'react-router-dom'
import {Row, Col, UncontrolledAlert} from "reactstrap";

import Book from "./Book";
import UpdateModal from "../Modal/UpdateModal";
import CheckOutModal from "../Modal/CheckOutModal";
import AddModal from "../Modal/AddModal";

const Books = () => {
  const navigate = useNavigate();
  const [data, setData] = useState([]);
  const [updateModel, setUpdateModal] = useState(false);
  const [updateBook, setUpdateBook] = useState({});
  const [cBookID, setCBookID] = useState(null);
  const [checkoutModal, setCheckoutModal] = useState(false);
  const [addModal, setAddModal] = useState(false);

  const [alertState, setAlertState] = useState({
    show: false,
    type: "success",
    message: "",
  })

  const toggle = () => {
    setUpdateModal(!updateModel);
  };

  const toggleCheckout = () => {
    setCheckoutModal(!checkoutModal);
  };

  const toggleAdd = () => {
    setAddModal(!addModal);
  };

  const getBooks = () => {
    if(localStorage.getItem("authToken") === null){
      navigate('/login');
    }
    fetch("http://localhost:8080/books", {
      headers: {
        Authorization: localStorage.getItem('authToken'),
      },
    })
      .then((response) => response.json())
      .then((data) => {
        if(data.type === "tokenerror"){
          console.log("problem with auth token, try to login again!");
          localStorage.removeItem("authToken");
          setAlertState({
            show: true,
            type: "danger",
            message: "problem with auth token, try to login again!",
          });
        }
        else if(data.type === "error"){
          console.log("something went wrong while getting books", data.error.message);
          setAlertState({
            show: true,
            type: "success",
            message: data.error.message,
          });
        }
        else{
          setData(data)
          setAlertState({
            show: false,
            type: "success",
            message: "",
          });
        }
      })
      .catch((error) => console.error(error));
  };

  const handleDelete = (e) => {
    const id = e.target.name;
    console.log(`Button clicked for id=${id}`);
    const url = "http://localhost:8080/delete/" + id;
    fetch(url, {
      method: "DELETE",
      headers: {
        Authorization: localStorage.getItem('authToken'),
      },
    })
      .then((response) => response.json())
      .then((res) => {
        if (res.type === "error") {
          console.log("error while deleting");
          setAlertState({
            show: true,
            type: "danger",
            message: res.error.message,
          });
          setTimeout(getBooks(), 3000);
        } else {
          console.log("successfully deleted");
          getBooks();
          setAlertState({
            show: true,
            type: "success",
            message: "successfully deleted",
          });
        }
      })
      .catch((err) => {
        console.log(err);
        getBooks();
      });
  };

  const getBook = async (bookid) => {
    const url = `http://localhost:8080/book/${bookid}`;
    try {
      const response = await fetch(url, {
        method: "GET",
        headers: {
          Authorization: localStorage.getItem('authToken'),
        },
      });
      const bookData = await response.json();
      return bookData;
    } catch (err) {
      return err;
    }
  };
  
  const handleUpdate = async (e) => {
    console.log(`Button clicked for book id: ${e.target.name}`);
    const bookData = await getBook(e.target.name);
    console.log(bookData);
    if (bookData.type === "success") {
      setUpdateBook(() => {
        return bookData["book-details"];
      });
      toggle();
    } else if (bookData.type === "error") {
      console.log(bookData);
      setAlertState({
        show: true,
        type: "danger",
        message: bookData.error.message,
      });
      return;
    } else {
      console.log(bookData);
      return;
    }
  };
  
  const handleUpdateSubmit = (id, title, author) => {
    const url = `http://localhost:8080/update`;
    console.log(id, title, author);
    fetch(url, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: localStorage.getItem('authToken'),
      },
      body: JSON.stringify({
        id,
        title,
        author,
      }),
    })
      .then((res) => res.json())
      .then((data) => {
        console.log(data);
        if(data.type === "error"){
          setAlertState({
            show: true,
            type: "danger",
            message: data.error.message,
          });
          return;
        }
        toggle();
        getBooks();
      })
      .catch((err) => {
        console.log(err);
        setAlertState({
          show: true,
          type: "danger",
          message: "something went wrong!",
        });
      });
  };

  const doCheckout = (e) => {
    setCBookID(e.target.name);
    toggleCheckout();
  };

  const handleCheckout = (bookid, userid) => {
    console.log(`checkout operation for: ${bookid}, ${userid}`);
    const url = `http://localhost:8080/checkout/${userid}?id=${bookid}`;
    const options = {
      method: "PATCH",
      headers: {
        Authorization: localStorage.getItem('authToken'),
      },
    };

    fetch(url, options)
      .then((res) => res.json())
      .then((d) => {
        console.log(d);
        if (d.type === "success") {
          console.log("SUCCESS");
        } else {
          console.log("FAILED");
        }
        getBooks();
      })
      .catch((err) => console.log(err));
  };

  const onAdd = async (e) => {
    console.log("Add call for: ", e.target.name);
    const bookData = await getBook(e.target.name);
    setUpdateBook(() => {
      return bookData["book-details"];
    });
    toggleAdd()
  };

  const handleAdd = async (bookid, count) => {
    const url = `http://localhost:8080/addexistingbook/${bookid}?c=${count}`;
    const options = {
      method: "GET",
      headers: {
        Authorization: localStorage.getItem('authToken'),
      },
    };
    fetch(url, options)
      .then((res) => res.json())
      .then((d) => {
        if (d.type === "success") {
          toggleAdd();
          getBooks();
        } else {
          console.log(d);
          getBooks();
        }
      })
      .catch((err) => console.log(err));
  };

  useEffect(getBooks, []);
  

  const renderBooks = data.map((d) => {
    return (
      
      <Col sm="4" xs="12" key={d.id}>
        <li style={{ listStyle: "none" }}>
          <Book
            isbn={d.id}
            title={d.title}
            author={d.author}
            quantity={d.quantity}
            total={d.total}
            handleDelete={handleDelete}
            handleUpdate={handleUpdate}
            doCheckout={doCheckout}
            onAdd={onAdd}
          />
        </li>
        <br />
        <br />
      </Col>
    );
  });

  const renderUpdateModal = (
    <UpdateModal
      toggle={toggle}
      updateModel={updateModel}
      book={updateBook}
      handleUpdate={handleUpdate}
      handleSubmit={handleUpdateSubmit}
    />
  );

  const renderCheckOutModal = (
    <CheckOutModal
      toggle={toggleCheckout}
      checkoutModal={checkoutModal}
      handleCheckout={handleCheckout}
      id={cBookID}
    />
  );

  const renderAddModal = (
    <AddModal
      toggle={toggleAdd}
      addModal={addModal}
      handleAdd={handleAdd}
      book={updateBook}
    />
  );

  return (
    <div className="container">
      <h3>List of available books:</h3>
      {alertState.show && (
        <UncontrolledAlert color={alertState.type} fade={true}>
          {alertState.message}
        </UncontrolledAlert>
      )}
      <Row>{renderBooks}</Row>
      {renderUpdateModal}
      {renderCheckOutModal}
      {renderAddModal}
    </div>
  );
};

export default Books;
