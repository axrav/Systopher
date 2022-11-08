import api from "../../api";
import { SIGN_IN, AUTH_ERROR, GET_USER } from "../types/auth";

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
      setError("");
      setShowError(true);
    }
  };
};

export const getUser = () => {
  return async (dispatch: any) => {
    try {
      const res = await api.get("/server/user", {
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
