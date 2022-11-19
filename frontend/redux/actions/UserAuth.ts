import api from "../../api";
import {
  SIGN_IN,
  AUTH_ERROR,
  GET_USER,
  SIGN_UP,
  VERIFY_USER,
} from "../types/auth";
import errorCodeToMessage from "../../components/Utils/Error";

export const signIn = (
  email: string,
  password: string,
  setLoading: any,
  setError: any,
  setShowError: any
) => {
  return async (dispatch: any) => {
    try {
      const res = await api.post("/auth/login", { email, password });
      dispatch({ type: SIGN_IN, payload: { user: res.data, error: null } });
    } catch (err: any) {
      dispatch({
        type: AUTH_ERROR,
        payload: { error: err.response, user: null },
      });
      setLoading(false);
      setError(errorCodeToMessage(err.response.data.code));
      setShowError(true);
    }
  };
};

export const signUp = (
  email: string,
  userName: string,
  password: string,
  setLoading: any,
  setError: any,
  setShowError: any
) => {
  return async (dispatch: any) => {
    try {
      const res = await api.post("/auth/signup", { email, password, userName });
      console.log(res.data);
      dispatch({
        type: SIGN_UP,
        payload: { signedUp: true, error: null },
        email,
      });
    } catch (err: any) {
      dispatch({
        type: AUTH_ERROR,
        payload: { error: err.response, user: null },
      });
      setLoading(false);
      setError(errorCodeToMessage(err.response.data.code));
      setShowError(true);
    }
  };
};

export const verifyUser = (
  email: string,
  otp: string,
  setLoading: any,
  setError: any,
  setShowError: any
) => {
  return async (dispatch: any) => {
    try {
      const res = await api.post("/auth/verify", { email, otp });
      dispatch({ type: VERIFY_USER, payload: { user: res.data, error: null } });
    } catch (err: any) {
      dispatch({
        type: AUTH_ERROR,
        payload: { error: err.response, user: null },
      });
      setLoading(false);
      setError(errorCodeToMessage(err.response.data.code));
      setShowError(true);
    }
  };
};

export const getUser = () => {
  return async (dispatch: any) => {
    try {
      const res = await api.get("/api/user", {
        headers: { Authorization: `Bearer ${localStorage.getItem("token")}` },
      });
      dispatch({ type: GET_USER, payload: { user: res.data } });
    } catch (err: any) {
      dispatch({ type: AUTH_ERROR, payload: { error: err.response } });
    }
  };
};

export const signOut = () => {
  return async (dispatch: any) => {
    try {
      localStorage.removeItem("token");
      dispatch({ type: GET_USER, payload: { user: null } });
    } catch (err: any) {
      dispatch({ type: AUTH_ERROR, payload: { error: err.response } });
    }
  };
};

export const clearErrorMessage = () => {
  return { type: AUTH_ERROR, payload: { error: false } };
};
