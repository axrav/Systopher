import React from "react";

function SignUpInput({
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
  error: string;
}) {
  return (
    <div className="flex flex-col space-y-1">
      <label htmlFor="">
        <p className="text-white font-light md:text-lg text-base">
          {inputLabel}
        </p>
      </label>
      <input
        placeholder={inputPlaceholder}
        type={inputType}
        value={value}
        onChange={(e) => setValue(e.target.value)}
        className={`w-full ${
          error ? "bg-red-400" : "bg-trasnparent"
        } placeholder:text-gray-700 px-1 h-12 outline-none focus:border-b-gray-300 duration-150 border-b border-b-gray-600 bg-transparent text-white font-light md:text-2xl text-lg mb-4`}
      />
    </div>
  );
}

export default SignUpInput;
