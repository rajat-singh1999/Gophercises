import {createStore} from "redux";

// authToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3MiOiJhZG1pbiIsImVtYWlsIjoicnNpbmdoMTczNEBnbWFpbC5jb21AZ21haWwuY29tIiwiaWQiOiI1MjI2In0.bV6_4dwsyUzv8nAhcRzasAGuK1n94s8uxdypbdSo9Gs",
const initialState = {
    authToken: "",
  };
function rootReducer(state = initialState, action) {
    switch (action.type) {
      case 'SET_AUTH_TOKEN':
        return { ...state, authToken: action.payload };
      default:
        return state;
    }
  }
  
  const store = createStore(rootReducer);
  
  export default store;