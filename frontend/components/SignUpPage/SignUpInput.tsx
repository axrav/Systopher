import React from "react";
import { IconEye, IconEyeOff } from "@tabler/icons";

function SignUpInput({
  inputType,
  inputPlaceholder,
  inputLabel,
  value,
  setValue,
  error,
  success,
  showPassword,
  setShowPassword,
}: {
  inputType: string;
  inputPlaceholder: string;
  inputLabel: string;
  value: string;
  setValue: React.Dispatch<React.SetStateAction<string>>;
  error: Boolean;
  success?: string;
  showPassword?: Boolean;
  setShowPassword?: React.Dispatch<React.SetStateAction<boolean>>;
}) {
  const [changed, setChanged] = React.useState(false);
  return (
    <div className="flex flex-col space-y-1 select-none">
      <label htmlFor="">
        <p
          className={`${
            error
              ? "text-red-600 animate-alert-in"
              : success === "yes"
              ? "text-green-600"
              : "text-white animate-none"
          } font-light w-fit md:text-lg text-base`}
        >
          {inputLabel}
        </p>
      </label>
      <div className="relative">
        <div
          hidden={inputType !== "password"}
          onClick={() => {
            if (setShowPassword) setShowPassword(!showPassword);
          }}
          className="absolute z-50 hover:bg-gray-800 px-2 py-2 cursor-pointer rounded-xl right-0"
        >
          {showPassword ? <IconEye /> : <IconEyeOff />}
        </div>
        <input
          placeholder={inputPlaceholder}
          type={showPassword ? "text" : inputType}
          value={value}
          onChange={(e) => {
            setValue(e.target.value);
            setChanged(true);
          }}
          className={`w-full ${
            error
              ? "text-red-600 placeholder:text-red-900 border-b-red-600 focus:border-b-red-300"
              : success === "yes"
              ? "text-green-600 placeholder:text-green-900 border-b-green-600 focus:border-b-green-300"
              : "text-white placeholder:text-gray-700 border-b-gray-600 focus:border-b-gray-300"
          } bg-transparent relative px-1 h-12 outline-none duration-150 border-b font-light md:text-2xl text-lg mb-4`}
        />
      </div>
    </div>
  );
}

export default SignUpInput;
