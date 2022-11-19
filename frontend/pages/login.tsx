import React, { useEffect } from "react";
import Header from "../components/Header";
import LoginForm from "../components/LoginPage/LoginForm";
import { Alert } from "@mantine/core";
import { IconAlertCircle } from "@tabler/icons";
import Router from "next/router";

export default function login() {
  const [error, setError] = React.useState("");
  const [showError, setShowError] = React.useState(false);

  //   useEffect(() => {
  //     if (localStorage.getItem("token") !== null) {
  //       Router.push("/dashboard");
  //     }
  //   });

  return (
    <div className="bg-gray-900 h-screen w-screen scrollbar-hide flex justify-between overflow-y-scroll">
      <div className="absolute bottom-5 right-5">
        <Alert
          icon={<IconAlertCircle size={16} />}
          onClose={() => setShowError(false)}
          title="Login Error!"
          color="red"
          withCloseButton
          hidden={!showError}
        >
          {error}
        </Alert>
      </div>
      <div className="md:w-1/2 w-full h-full flex items-center justify-center">
        <LoginForm
          setShowError={setShowError}
          error={error}
          setError={setError}
        />
      </div>
      <div className="w-1/2 md:flex hidden h-full items-center justify-center">
        <img src="login.svg" alt="" />
      </div>
    </div>
  );
}
