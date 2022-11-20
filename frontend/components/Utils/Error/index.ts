export default (errorCode: string) => {
  switch (errorCode) {
    case "ERR-699":
      return {
        msg: "The credentials you provided are incorrect. Please try again.",
        type: ["email", "password"],
      };
    case "ERR-698":
      return {
        msg: "Invalid Data",
        type: ["email", "password"],
      };
    case "ERR-697":
      return { msg: "Unauthorized Access", type: [] };
    case "ERR-696":
      return {
        msg: "The entered user is not registered on Systopher.",
        type: ["email"],
      };
    case "ERR-695":
      return {
        msg: "The entered username is already taken.",
        type: ["userName"],
      };
    case "ERR-694":
      return { msg: "The entered email is badly formatted.", type: ["email"] };
    case "ERR-693":
      return {
        msg: "The entered email is already in use by another account.",
        type: ["email"],
      };
    case "ERR-692":
      return { msg: "The entered email is badly formatted.", type: ["email"] };
    case "ERR-691":
      return { msg: "The entered username is invalid.", type: ["userName"] };
    case "ERR-690":
      return { msg: "Invalid Server Response", type: [] };
    case "ERR-689":
      return { msg: "Internal Server Error Occurred", type: ["none"] };
    case "ERR-688":
      return { msg: "The entered OTP is Incorrect.", type: ["otp"] };
    case "ERR-687":
      return { msg: "Details not found.", type: ["all"] };
    case "ERR-686":
      return {
        msg: "The Password must be at least 8 characters long with one uppercase letter, one lowercase letter, one number and one special character.",
        type: ["password"],
      };
    case "ERR-685":
      return {
        msg: "The user is already verified. Please login.",
        type: ["otp"],
      };
    case "ERR-684":
      return { msg: "Already exists.", type: ["all"] };
    case "ERR-683":
      return {
        msg: "No response from the server. Please try again.",
        type: ["none"],
      };
  }
};
