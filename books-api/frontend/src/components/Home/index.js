import "./index.css";
import { useEffect, useState } from "react";
import { jwtDecode } from "jwt-decode";

const Home = () => {
  const [user, setUser] = useState({});
  const [issuedBooks, setIssuedBooks] = useState([]);

  const authToken = localStorage.getItem("authToken");
  const decodedToken = jwtDecode(authToken);
  const id = decodedToken.id;
  const url = `http://localhost:8080/user/${id}`;

  useEffect(() => {
    fetch(url, {
      method: "GET",
      headers: {
        "content-type": "application/json",
        Authorization: authToken,
      },
    })
      .then((res) => res.json())
      .then((data) => {
        if (data.type === "success") {
          setUser(data["user-details"]);
        }
      })
      .catch((err) => console.log(err));
  }, []);

  useEffect(() => {
    const fetchBooks = async () => {
      if (user.issued && user.issued.length) {
        const books = await Promise.all(
          user.issued.map((id) => {
            return fetch(`http://localhost:8080/book/${id}`, {
              method: "GET",
              headers: {
                Authorization: localStorage.getItem("authToken"),
              },
            })
              .then((res) => res.json())
              .then((data) => {
                if (data.type === "success") {
                  return `${data["book-details"].title}[${data["book-details"].id}]`;
                } else {
                  return "N/A";
                }
              })
              .catch((err) => console.log(err));
          })
        );
        setIssuedBooks(books);
      } else {
        setIssuedBooks(["N/A"]);
      }
    };

    fetchBooks();
  }, [user]);

  return (
    <div className="Home container">
      <div className="user-info">
        <h3>
          Welcome to your dashboard {user.username} you have {user.access}{" "}
          access!
        </h3>
        <p>User ID: {user.id}</p>
        <p>Email: {user.email}</p>
        <p>Issued Books to your name: {issuedBooks.join(", ")}</p>
      </div>
    </div>
  );
};

export default Home;
