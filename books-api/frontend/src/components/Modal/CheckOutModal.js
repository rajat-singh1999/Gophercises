import { useEffect, useState } from "react";
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

const CheckOutModal = ({
  toggle,
  checkoutModal,
  handleCheckout,
  id
}) => {
    const [userid, setUserid] = useState("")
    const [bookid, setBookId] = useState(999999);
    useEffect(()=>{
        setBookId(id)
    }, [id]);

  const handleOnChange = (e) => {
    setUserid(e.target.value);
  };

  return (
    <div>
      <Modal
        isOpen={checkoutModal}
        toggle={toggle}
        className="checkOutModal"
        backdrop="static"
      >
        <ModalHeader toggle={handleCheckout}>Update {bookid}</ModalHeader>
        <ModalBody>
          <Form>
            <FormGroup>
              <Label for="ISBN">Book ISBN</Label>
              <Input
                id="id"
                name="id"
                placeholder={bookid}
                type="text"
                value={bookid}
                readOnly
              />
            </FormGroup>
            
            <FormGroup>
              <Label for="userid">User ID</Label>
              <Input
                id="userid"
                name="userid"
                placeholder="Enter user id here"
                type="text"
                value={userid}
                onChange={handleOnChange}
              />
            </FormGroup>
          </Form>
        </ModalBody>
        <ModalFooter>
          <Button
            onClick={() => {
                handleCheckout(bookid, userid);
                toggle();
            }}
          >
            Confirm Checkout
          </Button>{" "}
          <Button
            color="secondary"
            onClick={() => {
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

export default CheckOutModal;
