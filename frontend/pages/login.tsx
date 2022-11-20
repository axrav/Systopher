import React, { useEffect } from "react";
import Header from "../components/Header";
import LoginForm from "../components/LoginPage/LoginForm";
import { Alert } from "@mantine/core";
import { IconAlertCircle } from "@tabler/icons";
import Router from "next/router";
import Error from "../components/Utils/Notifications/Error";

export default function login() {
  const [error, setError] = React.useState("");
  const [showError, setShowError] = React.useState(false);

  useEffect(() => {
    if (localStorage.getItem("token") !== null) {
      Router.push("/dashboard");
    }
  });

  return (
    <div className="bg-gray-900 h-screen w-screen scrollbar-hide flex justify-between">
      <Error error={error} setShowError={setShowError} showError={showError} />
      <div className="md:w-1/2 w-full h-full scrollbar-hide flex items-center justify-center">
        <LoginForm
          setShowError={setShowError}
          showError={showError}
          error={error}
          setError={setError}
        />
      </div>
      <div className="w-1/2 md:flex hidden h-full scrollbar-hide items-center justify-center">
        <img src="login.svg" alt="" />
      </div>
    </div>
  );
}
