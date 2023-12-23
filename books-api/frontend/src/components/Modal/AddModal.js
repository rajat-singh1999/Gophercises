import { useReducer, useEffect } from "react";
import {
  Modal,
  ModalHeader,
  ModalBody,
  ModalFooter,
  Button,
  Form,
  FormGroup,
  Label,
  Input,
} from "reactstrap";

const initialBookState = {
  id: "",
  title: "",
  author: "",
  addQuantity: "",
};

const reducer = (state, action) => {
  switch (action.type) {
    case "SET_BOOK":
      return {
        ...state,
        [action.field]: action.value,
      };
    case "RESET":
      return initialBookState;
    case "SET_INITIAL_STATE":
      return action.book;
    default:
      return state;
  }
};

const AddModal = ({
  toggle,
  addModal,
  handleAdd,
  book,
}) => {
  const [thisBook, dispatch] = useReducer(reducer, initialBookState);

  useEffect(() => {
    dispatch({ type: "SET_INITIAL_STATE", book: book });
  }, [book]);

  const handleOnChange = (e) => {
    e.preventDefault();
    const name = e.target.name;
    const val = e.target.value;
    dispatch({ type: "SET_BOOK", field: name, value: val });
  };

  return (
    <div>
      <Modal
        isOpen={addModal}
        toggle={toggle}
        className="addModal"
        backdrop="static"
      >
        <ModalHeader toggle={handleAdd}>Adding more copies of {book.title}</ModalHeader>
        <ModalBody>
          <Form>
            <FormGroup>
              <Label for="ISBN">ISBN</Label>
              <Input
                id="id"
                name="id"
                placeholder={thisBook.id}
                type="text"
                value={thisBook.id}
                readOnly
              />
            </FormGroup>

            <FormGroup>
              <Label for="addQuantity">Add Quantity</Label>
              <Input
                id="addQuantity"
                name="addQuantity"
                placeholder={thisBook.addQuantity}
                type="text"
                value={thisBook.addQuantity}
                onChange={handleOnChange}
              />
            </FormGroup>
            
          </Form>
        </ModalBody>
        <ModalFooter>
          <Button
            onClick={() => {
              handleAdd(book.id, thisBook.addQuantity);
              dispatch({ type: "RESET" });
            }}
          >
            Update
          </Button>{" "}
          <Button
            color="secondary"
            onClick={() => {
              dispatch({ type: "RESET" });
              toggle();
            }}
          >
            Cancel
          </Button>
        </ModalFooter>
      </Modal>
    </div>
  );
};

export default AddModal;
