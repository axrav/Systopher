import React, { useEffect } from "react";
import OtpInput from "./OtpInput";

function index({ setOTP }: { setOTP: any }) {
  const [otp1, setOtp1] = React.useState("");
  const [otp2, setOtp2] = React.useState("");
  const [otp3, setOtp3] = React.useState("");
  const [otp4, setOtp4] = React.useState("");
  const [otp5, setOtp5] = React.useState("");
  const [otp6, setOtp6] = React.useState("");

  useEffect(() => {
    document.getElementById("otp1")?.focus();
  }, []);

  useEffect(() => {
    setOTP(`${otp1}${otp2}${otp3}${otp4}${otp5}${otp6}`);
  }, [otp1, otp2, otp3, otp4, otp5, otp6]);

  return (
    <div className="flex space-x-4 focus-within:cursor-none justify-center items-center w-full">
      <OtpInput otpId={1} setOTP={setOtp1} OTPVal={otp1} />
      <OtpInput otpId={2} setOTP={setOtp2} OTPVal={otp2} />
      <OtpInput otpId={3} setOTP={setOtp3} OTPVal={otp3} />
      <OtpInput otpId={4} setOTP={setOtp4} OTPVal={otp4} />
      <OtpInput otpId={5} setOTP={setOtp5} OTPVal={otp5} />
      <OtpInput otpId={6} setOTP={setOtp6} OTPVal={otp6} />
    </div>
  );
}

export default index;
