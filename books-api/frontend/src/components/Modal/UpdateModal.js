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

const UpdateModal = ({
  toggle,
  updateModel,
  handleUpdate,
  book,
  handleSubmit,
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
        isOpen={updateModel}
        toggle={toggle}
        className="updateModal"
        backdrop="static"
      >
        <ModalHeader toggle={handleUpdate}>Update {book.title}</ModalHeader>
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
              <Label for="Title">Title</Label>
              <Input
                id="title"
                name="title"
                placeholder="Title"
                type="text"
                value={thisBook.title}
                onChange={handleOnChange}
              />
            </FormGroup>
            <FormGroup>
              <Label for="Author">Author</Label>
              <Input
                id="author"
                name="author"
                placeholder="Author Name"
                type="text"
                value={thisBook.author}
                onChange={handleOnChange}
              />
            </FormGroup>
          </Form>
        </ModalBody>
        <ModalFooter>
          <Button
            onClick={() => {
              console.log(book.id, thisBook.title, thisBook.author);
              handleSubmit(book.id, thisBook.title, thisBook.author);
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

export default UpdateModal;
