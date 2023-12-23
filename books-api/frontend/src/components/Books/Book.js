import {
  Button,
  Card,
  CardBody,
  CardFooter,
  CardText,
  CardTitle,
} from "reactstrap";

const Book = ({
  isbn,
  title,
  author,
  quantity,
  total,
  handleDelete,
  handleUpdate,
  doCheckout,
  onAdd,
}) => {
  return (
    <Card>
      <CardBody>
        <CardTitle>Title: {title}</CardTitle>
        <CardText>
          ISBN: {isbn}
          <br />
          Author: {author}
          <br />
          Copies in Library: {quantity}
          <br />
          Total Copies: {total}
          <br />
        </CardText>
        <Button name={isbn} onClick={handleUpdate}>
          Update
        </Button>
        {"   "}
        <Button name={isbn} onClick={doCheckout}>
          Checkout
        </Button>
        {"   "}
        <Button name={isbn} onClick={onAdd}>
          Add
        </Button>
      </CardBody>
      <CardFooter>
        <Button name={isbn} onClick={handleDelete}>
          Delete
        </Button>
      </CardFooter>
    </Card>
  );
};

export default Book;

/*
<Card
  style={{
    width: '18rem'
  }}
>
  <img
    alt="Sample"
    src="https://picsum.photos/300/200"
  />
  <CardBody>
    <CardTitle tag="h5">
      Card title
    </CardTitle>
    <CardSubtitle
      className="mb-2 text-muted"
      tag="h6"
    >
      Card subtitle
    </CardSubtitle>
    <CardText>
      Some quick example text to build on the card title and make up the bulk of the card‘s content.
    </CardText>
    <Button>
      Button
    </Button>
  </CardBody>
</Card>
*/
