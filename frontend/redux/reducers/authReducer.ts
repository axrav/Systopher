const initialState = {};
import {
  AUTH_ERROR,
  GET_USER,
  SIGN_IN,
  SIGN_UP,
  VERIFY_USER,
} from "../types/auth";

export default function authReducer(state = initialState, action: any) {
  switch (action.type) {
    case SIGN_IN:
      localStorage.setItem("token", action.payload.user.token);
      return {
        ...state,
        ...action.payload,
      };
    case SIGN_UP:
      localStorage.setItem("email", action.email);
      return {
        ...state,
        ...action.payload,
      };
    case VERIFY_USER:
      localStorage.removeItem("email");
      localStorage.setItem("token", action.payload.user.token);
      return {
        ...state,
        ...action.payload,
      };

    case AUTH_ERROR:
      return {
        ...state,
        ...action.payload,
      };
    case GET_USER:
      return {
        ...state,
        ...action.payload,
      };
    default:
      return state;
  }
}
