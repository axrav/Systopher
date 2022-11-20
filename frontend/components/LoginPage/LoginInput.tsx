import React from "react";

function LoginInput({
  inputType,
  inputPlaceholder,
  inputLabel,
  value,
  setValue,
  error,
}: {
  inputType: string;
  inputPlaceholder: string;
  inputLabel: string;
  value: string;
  setValue: React.Dispatch<React.SetStateAction<string>>;
  error: Boolean;
}) {
  return (
    <div className="flex flex-col space-y-1">
      <label htmlFor="">
        <p
          className={`${
            error ? "text-red-600 animate-alert-in" : "text-white animate-none"
          } font-light w-fit md:text-lg text-base`}
        >
          {inputLabel}
        </p>
      </label>
      <input
        placeholder={inputPlaceholder}
        type={inputType}
        value={value}
        onChange={(e) => setValue(e.target.value)}
        className={`w-full ${
          error
            ? "text-red-600 placeholder:text-red-900 border-b-red-600"
            : "text-white placeholder:text-gray-700 border-b-gray-600"
        } bg-gray-900 px-1 h-12 outline-none focus:border-b-gray-300 duration-150 border-b font-light md:text-2xl text-lg mb-4`}
      />
    </div>
  );
}

export default LoginInput;
