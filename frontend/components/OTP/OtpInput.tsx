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
        console.log("keycode ", e.key);
        if (e.key === "Backspace") {
          if (otpId !== 1) {
            document.getElementById(`otp${otpId - 1}`)?.focus();
          }
        } else {
          if (
            otpId !== 6 &&
            ["0", "1", "2", "3", "4", "5", "6", "7", "8", "9"].includes(e.key)
          ) {
            document.getElementById(`otp${otpId + 1}`)?.focus();
          } else {
            document.getElementById(`otp${otpId}`)?.blur();
            document.getElementById(`verifybtn`)?.focus();
          }
        }
      }}
      className="xl:w-16 xl:h-16 md:w-12 md:h-12 h-9 w-9 md:text-xl xl:text-2xl 2xl:text-3xl bg-gray-400 bg-opacity-60 text-center rounded-md font-bold text-white outline-none p-3"
    />
  );
}

export default OtpInput;
