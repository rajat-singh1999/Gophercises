import "./index.css";

import ReturnForm from "./ReturnForm";

const Return = () => {
  return (
    <div className="container">
      <div className="form-div">
        <h3>Book Return</h3>
        <div className="myform">
          <ReturnForm />
        </div>
      </div>
    </div>
  );
};

export default Return;
