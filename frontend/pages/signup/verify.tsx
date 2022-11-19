import React, { useEffect } from "react";
import Router from "next/router";
import { Alert, Button } from "@mantine/core";
import { IconAlertCircle } from "@tabler/icons";
import OTP from "../../components/OTP";
import { useAppDispatch } from "../../components/hooks/useAppDispatch";
import { verifyUser } from "../../redux/actions/UserAuth";
import { useSelector } from "react-redux";

function verify() {
  const [email, setEmail] = React.useState("");
  const [showError, setShowError] = React.useState(false);
  const [error, setError] = React.useState("");
  const [otp, setOTP] = React.useState("");
  const [loading, setLoading] = React.useState(false);

  const authToken = useSelector((state: any) => state.auth.user?.token);

  useEffect(() => {
    if (localStorage.getItem("email") === null) {
      Router.push("/signup");
    } else {
      setEmail(localStorage.getItem("email"));
    }

    if (localStorage.getItem("token") !== null) {
      Router.push("/dashboard");
    }
  }, []);

  useEffect(() => {
    if (localStorage.getItem("token") !== null) {
      Router.push("/dashboard");
    }
  }, [authToken]);

  const dispatch = useAppDispatch();

  console.log(otp);

  const handleSubmit = (e: any) => {
    e.preventDefault();
    setLoading(true);
    dispatch(verifyUser(email, otp, setLoading, setError, setShowError));
  };

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
      <div className="flex flex-col justify-center items-center w-full h-full space-y-8 lg:space-y-10 xl:space-y-14 2xl:space-y-16">
        <div className="xl:text-6xl 2xl:text-7xl md:text-5xl sm:text-3xl text-2xl  text-white font-bold">
          Verify your email
        </div>
        <p className="font-light xl:text-lg 2xl:text-2xl md:text-base text-sm text-gray-300">
          Enter the OTP sent to{" "}
          <span className="underline underline-offset-4 font-bold">
            {email}
          </span>
        </p>
        <form onSubmit={handleSubmit} className="flex flex-col space-y-10">
          <OTP setOTP={setOTP} />
          <Button
            variant="filled"
            type="submit"
            onSubmit={handleSubmit}
            loading={loading}
            id="verifybtn"
            size="lg"
            className="bg-green-600 w-fit mx-auto font-bold outline-none focus:outline-none"
            color="green"
          >
            Verify
          </Button>
        </form>
      </div>
    </div>
  );
}

export default verify;
