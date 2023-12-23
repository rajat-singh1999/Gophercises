import { Button, Card, CardBody, CardFooter, CardText, CardTitle } from "reactstrap";

const User = ({ username, id, access, email, issued })=>{
  const issuedString = issued.length>0 ? `Issued Books isbn(s): ${issued}`:`Issued Books isbn(s): N/A`;
    return (
        <Card>
          <CardBody>
            <CardTitle>Username: {username}</CardTitle>
            <CardText>
            ID: {id}
            <br />
            Access: {access}
            <br />
            Email: {email}
            <br />
            {issuedString}
            <br />
            </CardText>
            <Button name={id}>
              Update
            </Button>{'   '}
            <Button name={id}>
              Checkout
            </Button>
          </CardBody>
          <CardFooter>
          <Button name={id}>
              Delete
            </Button>
          </CardFooter>
        </Card>
      );
}

export default User;