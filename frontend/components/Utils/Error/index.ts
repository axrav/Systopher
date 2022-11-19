export default (errorCode: string) => {
  switch (errorCode) {
    case "ERR-699":
      return "The credentials you provided are incorrect. Please try again.";
    case "ERR-698":
      return "Invalid Data";
    case "ERR-697":
      return "Unauthorized Access";
    case "ERR-696":
      return "The entered user does not exist.";
    case "ERR-695":
      return "The entered username is already taken.";
    case "ERR-694":
      return "The entered email is badly formatted.";
    case "ERR-693":
      return "The entered email is already in use by another account.";
    case "ERR-692":
      return "The entered email is badly formatted.";
    case "ERR-691":
      return "The entered username is invalid.";
    case "ERR-690":
      return "Invalid Server Response";
    case "ERR-689":
      return "Internal Server Error Occurred";
    case "ERR-688":
      return "The entered OTP is invalid.";
    case "ERR-687":
      return "Details not found.";
    case "ERR-686":
      return "The Password must be at least 8 characters long with one uppercase letter, one lowercase letter, one number and one special character.";
    case "ERR-685":
      return "The user is already verified. Please login.";
    case "ERR-684":
      return "Already exists.";
    case "ERR-683":
      return "No response from the server. Please try again.";
  }
};
