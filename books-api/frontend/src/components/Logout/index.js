import { useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { useAuth } from '../../AuthContext';

const Logout = () => {
  const navigate = useNavigate();
  const { setAuthInfo } = useAuth();
  useEffect(() => {
    localStorage.removeItem("authToken");
    //   localStorage.setItem('isloggedin', false);
    //   localStorage.removeItem('username');
    //   localStorage.removeItem('access');
    setAuthInfo({ isloggedin: "false", username: null, access: null });
    navigate("/login");
  }, []);

  return <h3>Logging Out...</h3>;
};

export default Logout;
