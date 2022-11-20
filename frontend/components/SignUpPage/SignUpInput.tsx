import React from "react";

function SignUpInput({
  inputType,
  inputPlaceholder,
  inputLabel,
  value,
  setValue,
  error,
  success,
}: {
  inputType: string;
  inputPlaceholder: string;
  inputLabel: string;
  value: string;
  setValue: React.Dispatch<React.SetStateAction<string>>;
  error: Boolean;
  success?: string;
}) {
  const [changed, setChanged] = React.useState(false);
  return (
    <div className="flex flex-col space-y-1">
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
      <input
        placeholder={inputPlaceholder}
        type={inputType}
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
        } bg-gray-900 relative px-1 h-12 outline-none duration-150 border-b font-light md:text-2xl text-lg mb-4`}
      />
    </div>
  );
}

export default SignUpInput;
