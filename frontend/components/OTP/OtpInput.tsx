import React from "react";

function OtpInput({
  otpId,
  setOTP,
  OTPVal,
}: {
  otpId: number;
  setOTP: any;
  OTPVal: any;
}) {
  return (
    <input
      type="text"
      id={`otp${otpId}`}
      maxLength={1}
      value={OTPVal}
      onChange={(e) => {
        setOTP(e.target.value);
      }}
      onKeyUp={(e) => {
        if (e.keyCode === 8) {
          if (otpId !== 1) {
            document.getElementById(`otp${otpId - 1}`)?.focus();
          }
        } else {
          if (otpId !== 6 && e.keyCode >= 47 && e.keyCode <= 58) {
            document.getElementById(`otp${otpId + 1}`)?.focus();
          } else {
            document.getElementById(`otp${otpId}`)?.blur();
            document.getElementById(`verifybtn`)?.focus();
          }
        }
      }}
      className="xl:w-16 xl:h-16 md:w-12 md:h-12 h-12 w-12 md:text-xl xl:text-2xl 2xl:text-3xl bg-gray-400 bg-opacity-60 text-center rounded-md font-bold text-white outline-none p-3"
    />
  );
}

export default OtpInput;
