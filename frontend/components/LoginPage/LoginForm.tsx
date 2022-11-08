import React from "react";
import LoginInput from "./LoginInput";
import api from "../../api";
import { Button } from "@mantine/core";

function LoginForm({
  error,
  setError,
  setShowError,
}: {
  error: string;
  setError: React.Dispatch<React.SetStateAction<string>>;
  setShowError: React.Dispatch<React.SetStateAction<boolean>>;
}) {
  const [email, setEmail] = React.useState("");
  const [password, setPassword] = React.useState("");
  const [loading, setLoading] = React.useState(false);

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    setLoading(true);
    try {
      const res = await api.post("/auth/login", {
        email: email,
        password: password,
      });
      localStorage.setItem("token", res.data.token);
    } catch (error: any) {
      setLoading(false);
      setShowError(true);
      error.response.status == 400 ? setError("Invalid Credentials") : null;
    }
  };

  return (
    <div className="w-10/12 bg-gray-900 h-fit flex flex-col space-y-12">
      <div className="flex flex-col space-y-3">
        <h1 className="text-white font-semibold md:text-5xl text-2xl">
          Welcome to Systopher
        </h1>
        <p className="text-gray-400 font-light md:text-2xl text-lg">
          Login to your account
        </p>
      </div>
      <form
        onSubmit={handleSubmit}
        className="flex flex-col space-y-8"
        action=""
      >
        <LoginInput
          inputType={"email"}
          inputPlaceholder={"someone@someone.com"}
          inputLabel={"Email"}
          value={email}
          setValue={setEmail}
          error={error}
        />
        <LoginInput
          inputType={"password"}
          inputPlaceholder={"Enter Your Password"}
          inputLabel={"Password"}
          value={password}
          setValue={setPassword}
          error={error}
        />
        <Button
          className="w-full px-4 py-4 bg-gray-800 hover:bg-gray-700 duration-150 rounded-lg text-xl font-semibold text-yellow-500"
          type="submit"
          size="xl"
          loading={loading}
        >
          Login
        </Button>
      </form>
      <div className="">
        <p className="text-gray-500 font-light md:text-xl">
          Don't have an account?{" "}
          <a className="text-yellow-500" href="/register">
            Sign Up
          </a>
        </p>
      </div>
    </div>
  );
}

export default LoginForm;
