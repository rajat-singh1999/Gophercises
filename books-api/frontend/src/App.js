import { Provider } from "react-redux";
import store from "./Redux/store";
import { AuthProvider } from './AuthContext';

import "./App.css";

import { BrowserRouter as Router, Routes, Route } from "react-router-dom";

import Home from "./components/Home";
import Books from "./components/Books";
import Users from "./components/Users";
import AddBook from "./components/AddBook";
import NavBar from "./components/NavBar";
import Return from "./components/Return";
import Login from "./components/LogIn";
import Signup from "./components/Signup";
import Logout from "./components/Logout";

function App() {
  return (
    <Provider store={store}>
      <AuthProvider>
      <NavBar
        color={"dark"}
        localStore={localStorage}
        dark={true}
        fixed={"top"}
        full={"true"}
        container={true}
        expand={true}
      />
      <Router>
        <div className="App">
          <Routes>
            <Route exact path="/login" element={<Login />} />
            <Route exact path="/signup" element={<Signup />} />

            <Route exact path="/" element={<Home />} />
            <Route exact path="/books" element={<Books />} />
            <Route exact path="/addBook" element={<AddBook />} />
            <Route exact path="/users" element={<Users />} />
            <Route exact path="/return" element={<Return />} />
            <Route
              exact
              path="/logout"
              element={<Logout />}
            />
          </Routes>
        </div>
      </Router>
      </AuthProvider>
    </Provider>
  );
}

export default App;
