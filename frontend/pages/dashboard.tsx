import React, { useEffect } from "react";
import Router from "next/router";
import api from "../api";
import { useDispatch, useSelector } from "react-redux";
import { getUser, signOut } from "../redux/actions/UserAuth";

function dashboard() {
  const dispatch = useDispatch();
  const user = useSelector((state: any) => state.auth.user);
  const userError = useSelector((state: any) => state.auth.error);

  console.log(user);
  useEffect(() => {
    if (localStorage.getItem("token") === null) {
      Router.push("/login");
    } else {
      dispatch(getUser());
    }
  }, []);

  if (user === null) {
    Router.push("/login");
  }

  if (userError) {
    dispatch(signOut());
    Router.push("/login");
  }
  return (
    <div>
      <h1>dashboard</h1>
      <p
        onClick={() => {
          dispatch(signOut());
        }}
      >
        {user?.email}
      </p>
    </div>
  );
}

export default dashboard;
