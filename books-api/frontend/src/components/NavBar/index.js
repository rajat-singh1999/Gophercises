import "bootstrap/dist/css/bootstrap.min.css";
import { useAuth } from '../../AuthContext';

import { useEffect, useState } from "react";
import {
  Navbar,
  NavbarBrand,
  Nav,
  NavItem,
  NavLink,
  NavbarText,
  Collapse,
  NavbarToggler,
} from "reactstrap";

const NavBar = (args) => {
  const { authState } = useAuth();
  const localStore = args.localStore;
  useEffect(()=>{
    console.log('navbar rendered')
  },[]);
  const [isOpen, setIsOpen] = useState(false);

  const toggle = () => setIsOpen(!isOpen);

  return (
    <div>
      <Navbar {...args}>
        <NavbarBrand href="/">Awesome Library</NavbarBrand>
        <NavbarToggler onClick={toggle} />
        <Collapse isOpen={isOpen} navbar>
          <Nav className="me-auto" navbar>
            
            {authState.isloggedin==='false' && <>
            <NavItem>
              <NavLink href="/login">Login</NavLink>
            </NavItem>
            <NavItem>
              <NavLink href="/signup">SignUp</NavLink>
            </NavItem>            
            </>
            }

            {localStore.getItem('access')==='admin' && <NavItem>
              <NavLink href="/users">Users</NavLink>
            </NavItem>}
            {authState.isloggedin==='true' && <NavItem>
              <NavLink href="/books">Books</NavLink>
            </NavItem>}
            {localStore.getItem('access')==='admin' && <NavItem>
              <NavLink href="/addBook">AddBook</NavLink>
            </NavItem>}
            {localStore.getItem('access')==='admin' && <NavItem>
              <NavLink href="/return">Return</NavLink>
            </NavItem>}
          </Nav>
          <NavbarText>{authState.isloggedin&&localStore.getItem('username')}</NavbarText>
          <Nav>
          {authState.isloggedin==='true' && <NavItem>
              <NavLink href="/logout">Logout</NavLink>
            </NavItem>}
          </Nav>
        </Collapse>
      </Navbar>
    </div>
  );
};

export default NavBar;
