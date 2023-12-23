import "./index.css";
import Form from "./Form";

const AddBook = () => {
  return (
    <div className="container">
      <div className="form-div">
        <h3>Component for adding books</h3>
        <div className="myForm">
          <Form />
        </div>
      </div>
    </div>
  );
};

export default AddBook;
