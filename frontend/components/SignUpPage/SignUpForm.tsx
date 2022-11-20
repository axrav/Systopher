import React, { useEffect } from "react";
import SignUpInput from "./SignUpInput";
import Router from "next/router";
import api from "../../api";
import { Button } from "@mantine/core";
import { useDispatch, useSelector } from "react-redux";
import { clearErrorMessage, signUp } from "../../redux/actions/UserAuth";
import { useAppDispatch } from "../hooks/useAppDispatch";

function SignUpForm({
  error,
  setError,
  showError,
  setShowError,
}: {
  error: string;
  setError: React.Dispatch<React.SetStateAction<string>>;
  showError: boolean;
  setShowError: React.Dispatch<React.SetStateAction<boolean>>;
}) {
  const [email, setEmail] = React.useState("");
  const [password, setPassword] = React.useState("");
  const [rePassword, setRePassword] = React.useState("");
  const [passwordVerify, setPasswordVerify] = React.useState(false);
  const [loading, setLoading] = React.useState(false);
  const [userName, setUserName] = React.useState("");

  const [userNameError, setUserNameError] = React.useState(false);
  const [emailError, setEmailError] = React.useState(false);
  const [passwordError, setPasswordError] = React.useState(false);
  const [rePasswordError, setRePasswordError] = React.useState(false);

  const dispatch = useAppDispatch();

  const user = useSelector((state: any) => state.auth.user);
  const userError = useSelector((state: any) => state.auth.error);

  console.log(user);
  console.log(userError);

  const handleSubmit = async () => {
    if (passwordVerify && userName && email && password && rePassword) {
      setLoading(true);
      dispatch(
        signUp(
          email,
          userName,
          password,
          setLoading,
          setError,
          setShowError,
          setEmailError,
          setPasswordError,
          setUserNameError,
          setRePasswordError
        )
      );
    } else {
      setError("Please fill out all the fields");
      setShowError(true);
    }
  };

  if (user) {
    Router.push("/dashboard");
  }

  useEffect(() => {
    if (!showError) {
      setError("");
      setEmailError(false);
      setPasswordError(false);
      setRePasswordError(false);
      setUserNameError(false);
    }
  }, [showError]);

  useEffect(() => {
    if (password === rePassword) {
      setPasswordVerify(true);
      setPasswordError(false);
      setRePasswordError(false);
    } else {
      setPasswordVerify(false);
    }
  }, [rePassword, password]);

  useEffect(() => {
    if (error !== "" && !passwordVerify) {
      setPasswordError(true);
      setRePasswordError(true);
    }
  }, [error]);

  return (
    <div className="w-10/12 h-fit flex flex-col space-y-12">
      <div className="flex flex-col space-y-3">
        <h1 className="text-white font-semibold md:text-5xl text-2xl">
          Welcome to Systopher
        </h1>
        <p className="text-gray-400 font-light md:text-2xl text-lg">
          Sign Up for a new Account!
        </p>
      </div>
      <form
        onSubmit={(e) => e.preventDefault()}
        className="flex flex-col space-y-8"
        action=""
      >
        <SignUpInput
          inputType={"email"}
          inputPlaceholder={"someone@someone.com"}
          inputLabel={"Email"}
          value={email}
          setValue={setEmail}
          error={emailError}
        />
        <SignUpInput
          inputType={"text"}
          inputPlaceholder={"myuser"}
          inputLabel={"Username"}
          value={userName}
          setValue={setUserName}
          error={userNameError}
        />
        <SignUpInput
          inputType={"password"}
          inputPlaceholder={"Enter Your Password"}
          inputLabel={"Password"}
          value={password}
          setValue={setPassword}
          error={passwordError}
        />
        <SignUpInput
          inputType={"password"}
          inputPlaceholder={"Re-Enter Your Password"}
          inputLabel={"Re Enter Password"}
          value={rePassword}
          setValue={setRePassword}
          error={rePasswordError}
        />
        <Button
          className="w-full px-4 py-4 bg-gray-800 hover:bg-gray-700 duration-150 rounded-lg text-xl font-semibold text-yellow-500"
          type="submit"
          size="xl"
          loading={loading}
          onClick={handleSubmit}
        >
          Sign Up
        </Button>
      </form>
      <div className="">
        <p className="text-gray-500 font-light md:text-xl">
          Already have an account?{" "}
          <button
            className="text-yellow-500"
            onClick={() => {
              Router.push("/login");
            }}
          >
            Login
          </button>
        </p>
      </div>
    </div>
  );
}

export default SignUpForm;
